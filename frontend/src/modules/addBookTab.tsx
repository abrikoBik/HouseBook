import { useEffect, useState, useMemo } from "react";
import { useNavigate } from "react-router";
import "../styles/addBook.css";

interface ApiBook {
  id: string;
  title: string | null;
  author: string | null;
  image_uuid: string | null;
  isbn?: string;
  genre?: string;
  year?: number | null;
}
interface AddBookTabProps {
  onClose?: () => void;
}

// Обновите сигнатуру компонента
const AddBookTab: React.FC<AddBookTabProps> = ({ onClose }) => {
  const navigate = useNavigate();

  const [allBooks, setAllBooks] = useState<ApiBook[]>([]);
  const [searchQuery, setSearchQuery] = useState("");
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // Загрузка всех книг при монтировании
  useEffect(() => {
    const fetchBooks = async () => {
      try {
        setIsLoading(true);
        setError(null);

        // Используем limit для предотвращения слишком больших ответов
        const response = await fetch("/api/books?limit=100");

        if (!response.ok) {
          throw new Error(`Ошибка загрузки: ${response.status}`);
        }

        const data: ApiBook[] = await response.json();
        setAllBooks(data);
      } catch (err) {
        console.error("Failed to fetch books:", err);
        setError("Не удалось загрузить список книг. Попробуйте позже.");
      } finally {
        setIsLoading(false);
      }
    };

    fetchBooks();
  }, []);

  // Фильтрация книг по поисковому запросу
  const filteredBooks = useMemo(() => {
    if (!searchQuery.trim()) return allBooks.slice(0, 50); // Показываем первые 50 без поиска

    const query = searchQuery.toLowerCase().trim();
    return allBooks
      .filter((book) => {
        const title = book.title?.toLowerCase() || "";
        const author = book.author?.toLowerCase() || "";
        const isbn = book.isbn?.toLowerCase() || "";

        return (
          title.includes(query) ||
          author.includes(query) ||
          isbn.includes(query)
        );
      })
      .slice(0, 50);
  }, [allBooks, searchQuery]);

  const handleBookSelect = async (bookId: string) => {
    try {
      const userId = localStorage.getItem("UserID");
      if (!userId) {
        alert("Пожалуйста, войдите в аккаунт");
        return;
      }

      const libsResponse = await fetch(`/api/users/${userId}/libraries`);
      if (!libsResponse.ok) throw new Error("Не удалось получить библиотеки");

      const libraries = await libsResponse.json();
      if (libraries.length === 0) {
        alert("У вас нет библиотек. Создайте библиотеку сначала.");
        return;
      }

      // Добавляем книгу в первую библиотеку
      const addResponse = await fetch(
        `/api/libraries/${libraries[1].id}/books`,
        {
          method: "PATCH",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ bookid: bookId }),
        },
      );

      if (addResponse.ok) {
        alert("Книга добавлена в вашу библиотеку!");
        navigate("/library"); // Перенаправляем в библиотеку
      } else {
        const errData = await addResponse.json().catch(() => ({}));
        alert(`Ошибка: ${errData.error || "Не удалось добавить книгу"}`);
      }
    } catch (err) {
      console.error("Error adding book to library:", err);
      alert("Произошла ошибка при добавлении книги");
    }
  };

  return (
    <section id="add-overlay">
      <div className="add-book-container">
        <div id="add-title">
          Найдите книгу в каталоге
          <span
            className="material-symbols-outlined add-close"
            onClick={onClose}
          >
            close
          </span>
        </div>

        {/* Поле поиска */}
        <div className="search-wrapper">
          <input
            type="text"
            className="book-search-input"
            placeholder="Поиск по названию, автору или ISBN..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            autoFocus
          />
          {searchQuery && (
            <button className="clear-search" onClick={() => setSearchQuery("")}>
              X
            </button>
          )}
        </div>

        {/* Статус загрузки/ошибки */}
        {isLoading && <div className="loading-state">Загрузка каталога...</div>}

        {error && (
          <div className="error-state">
            {error}
            <button onClick={() => window.location.reload()}>Повторить</button>
          </div>
        )}

        {/* Список найденных книг */}
        {!isLoading && !error && (
          <>
            <div className="books-list-header">
              <span>Найдено: {filteredBooks.length}</span>
              {searchQuery && filteredBooks.length === 0 && (
                <span className="no-results">Книга не найдена в каталоге</span>
              )}
            </div>

            <div className="books-list">
              {filteredBooks.map((book) => (
                <div
                  key={book.id}
                  className="add-book-card"
                  onClick={() => handleBookSelect(book.id)}
                >
                  <div className="add-book-cover">
                    {book.image_uuid ? (
                      <img
                        src={`/cdn/${book.image_uuid}`}
                        alt={book.title || "Book cover"}
                        loading="lazy"
                      />
                    ) : (
                      <div className="cover-placeholder">📚</div>
                    )}
                  </div>
                  <div className="book-info">
                    <div className="book-title">
                      {book.title || "Без названия"}
                    </div>
                    <div className="book-author">
                      {book.author || "Неизвестный автор"}
                    </div>
                    {book.isbn && (
                      <div className="book-isbn">ISBN: {book.isbn}</div>
                    )}
                    {book.genre && (
                      <span className="book-genre">{book.genre}</span>
                    )}
                  </div>
                  <div className="book-action">
                    <span className="material-symbols-outlined">
                      add_circle
                    </span>
                  </div>
                </div>
              ))}
            </div>

            {(!searchQuery || filteredBooks.length === 0) && (
              <div className="add-new-section">
                <div id="add-buttons">
                  <button
                    id="add-scrape"
                    onClick={() => navigate("/book/scrape")}
                  >
                    Добавить через скрапинг
                  </button>
                  <button
                    id="add-by-hand"
                    onClick={() => navigate("/book/manual")}
                  >
                    Ввести вручную
                  </button>
                </div>
              </div>
            )}
          </>
        )}
      </div>
    </section>
  );
};

export default AddBookTab;
