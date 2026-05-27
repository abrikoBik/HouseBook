import { useEffect, useState } from "react";
import type BookT from "../types/book";
import type Error from "../types/error";
import "../styles/avaibleBookList.css";

interface ExchangeableBookT extends BookT {
  user_id: string;
  library_id: string;
}

interface AvailableBookListProps {
  isOpen: boolean;
  onClose: () => void;
  bookId: string;
  currentBook: BookT | undefined; // Добавили проп для передачи данных текущей книги
}

/* ==========================================================================
   Суб-компонент карточки (только владелец и библиотека)
   ========================================================================== */
const ExchangeBookCard = ({ book }: { book: ExchangeableBookT }) => {
  const [userName, setUserName] = useState<string>("Загрузка...");
  const [libraryName, setLibraryName] = useState<string>("Загрузка...");

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const response = await fetch(`/api/users/${book.user_id}`);
        if (response.ok) {
          const data = await response.json();
          setUserName(data.name || "Неизвестный пользователь");
        } else {
          setUserName("Пользователь не найден");
        }
      } catch (err) {
        setUserName("Ошибка загрузки");
      }
    };

    const fetchLibrary = async () => {
      try {
        const response = await fetch(`/api/libraries/${book.library_id}`);
        if (response.ok) {
          const data = await response.json();
          setLibraryName(data.name || "Неизвестная библиотека");
        } else {
          setLibraryName("Библиотека не найдена");
        }
      } catch (err) {
        setLibraryName("Ошибка загрузки");
      }
    };

    if (book.user_id) fetchUser();
    if (book.library_id) fetchLibrary();
  }, [book.user_id, book.library_id]);

  return (
    <div className="exchange-book-item">
      <div className="exchange-book-info">
        <div className="exchange-book-meta">
          <p><strong>Владелец:</strong> {userName}</p>
          <p><strong>Библиотека:</strong> {libraryName}</p>
        </div>
      </div>
    </div>
  );
};

/* ==========================================================================
   Основной компонент Модального Окна
   ========================================================================== */
const AvailableBookList = ({ isOpen, onClose, bookId, currentBook }: AvailableBookListProps) => {
  const [books, setBooks] = useState<ExchangeableBookT[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!isOpen) return;

    const fetchExchangeableBooks = async () => {
      try {
        setIsLoading(true);
        setError(null);
        const response = await fetch(`/api/books/${bookId}/exchangeable`);
        const result = await response.json();

        if (!response.ok) {
          throw new Error((result as Error).error || "Не удалось загрузить список книг");
        }

        setBooks((result || []) as ExchangeableBookT[]);
      } catch (err: any) {
        console.error("Failed to load exchangeable books:", err);
        setError(err.message || "Произошла ошибка при загрузке");
      } finally {
        setIsLoading(false);
      }
    };

    fetchExchangeableBooks();
  }, [isOpen, bookId]);

  if (!isOpen) return null;

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={(e) => e.stopPropagation()}>
        
        <div className="modal-header">
          <div className="modal-title-block">
            <h3>Книги, доступные для обмена</h3>
            {currentBook && (
              <p className="modal-subtitle">
                «{currentBook.title}» — {currentBook.author}
              </p>
            )}
          </div>
          <button className="modal-close-btn" onClick={onClose} aria-label="Закрыть">
            &times;
          </button>
        </div>

        <div className="modal-body">
          {isLoading && <div className="modal-loading">Загрузка списка...</div>}

          {error && <div className="modal-error">{error}</div>}

          {!isLoading && !error && (
            <div className="exchange-books-list">
              {books.length > 0 ? (
                books.map((item) => (
                  <ExchangeBookCard key={item.id} book={item} />
                ))
              ) : (
                <p className="no-books-message">К сожалению, сейчас нет доступных предложений для обмена этой книги.</p>
              )}
            </div>
          )}
        </div>

      </div>
    </div>
  );
};

export default AvailableBookList;