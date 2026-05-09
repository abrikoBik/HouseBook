import { useState, useEffect } from "react";
import { useSearchParams } from "react-router-dom";
import SideBar from "../modules/sideBar";
import Header from "../modules/header";
import "../styles/catalog.css";

const API_URL = "http://localhost/api/books";

// Обновленный список из 12 жанров (как на главной)
const genres = [
  "Все",
  "Фантастика",
  "Детективы", // Обрати внимание: во множественном числе, как в Central
  "Триллеры",
  "Ужасы",
  "Приключения",
  "Научная фантастика", // Соответствует Sci-Fi
  "История",
  "Комедия",
  "Документальное",
  "Романы",
  "Экшен",
  "Драма",
];

interface ApiBook {
  id: string;
  title: string | null;
  author: string | null;
  genre: string | null;
  image_uuid: string | null;
}

interface DisplayBook {
  id: string;
  title: string;
  author: string;
  genre: string;
  cover_url: string;
}

const Catalog = () => {
  const [searchParams] = useSearchParams();
  const urlGenre = searchParams.get("genre");

  // Проверяем, есть ли такой жанр в нашем новом списке
  const initialGenre = urlGenre && genres.includes(urlGenre) ? urlGenre : "Все";

  const [activeGenre, setActiveGenre] = useState(initialGenre);
  const [books, setBooks] = useState<DisplayBook[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchBooks = async () => {
      try {
        setLoading(true);
        const res = await fetch(API_URL);
        if (!res.ok) throw new Error(`Ошибка сервера: ${res.status}`);

        const data: ApiBook[] = await res.json();

        const mapped: DisplayBook[] = data.map((book) => ({
          id: book.id,
          title: book.title?.trim() || "Без названия",
          author: book.author?.trim() || "Неизвестный автор",
          genre: book.genre?.trim() || "Другое",
          // Заглушка для обложки, если нет реального URL
          cover_url: book.image_uuid
            ? `http://localhost/api/images/${book.image_uuid}`
            : `https://via.placeholder.com/300x450?text=${encodeURIComponent(book.title || "Book")}`,
        }));

        setBooks(mapped);
      } catch (err) {
        setError(
          err instanceof Error ? err.message : "Не удалось загрузить каталог",
        );
      } finally {
        setLoading(false);
      }
    };

    fetchBooks();
  }, []);

  // Синхронизация состояния при изменении URL (кнопка назад/вперед браузера)
  useEffect(() => {
    if (urlGenre && genres.includes(urlGenre)) {
      setActiveGenre(urlGenre);
    } else {
      setActiveGenre("Все");
    }
  }, [urlGenre]);

  const filteredBooks =
    activeGenre === "Все"
      ? books
      : books.filter((b) => b.genre === activeGenre);

  if (loading) {
    return (
      <section id="catalog-body">
        <SideBar />
        <main id="catalog-main">
          <Header />
          <div
            id="catalog-layout"
            style={{ textAlign: "center", padding: "100px 0" }}
          >
            <p>Загрузка каталога...</p>
          </div>
        </main>
      </section>
    );
  }

  if (error) {
    return (
      <section id="catalog-body">
        <SideBar />
        <main id="catalog-main">
          <Header />
          <div
            id="catalog-layout"
            style={{
              textAlign: "center",
              padding: "100px 0",
              color: "#d32f2f",
            }}
          >
            <p>Ошибка: {error}</p>
          </div>
        </main>
      </section>
    );
  }

  return (
    <section id="catalog-body">
      <SideBar />
      <main id="catalog-main">
        <Header />
        <div id="catalog-layout">
          <section id="catalog-header-section">
            <h1>Каталог</h1>

            {activeGenre !== "Все" && (
              <p style={{ marginBottom: "10px", color: "#666" }}>
                Фильтр: <strong>{activeGenre}</strong>
                <button
                  onClick={() => {
                    setActiveGenre("Все");
                    window.history.pushState({}, "", "/book/catalog");
                  }}
                  style={{
                    marginLeft: "10px",
                    border: "none",
                    background: "none",
                    color: "#E0616A",
                    cursor: "pointer",
                    textDecoration: "underline",
                  }}
                >
                  Сбросить
                </button>
              </p>
            )}

            <nav className="genre-nav">
              <ul>
                {genres.map((genre) => (
                  <li
                    key={genre}
                    className={activeGenre === genre ? "active" : ""}
                    onClick={() => {
                      setActiveGenre(genre);
                      if (genre === "Все") {
                        window.history.pushState({}, "", "/book/catalog");
                      } else {
                        window.history.pushState(
                          {},
                          "",
                          `/book/catalog?genre=${encodeURIComponent(genre)}`,
                        );
                      }
                    }}
                  >
                    {genre}
                  </li>
                ))}
              </ul>
            </nav>
          </section>

          <section id="book-grid">
            {filteredBooks.length > 0 ? (
              filteredBooks.map((book) => (
                <a href={`/book/${book.id}`} key={book.id}>
                  <div className="book-card">
                    <div className="book-cover">
                      <img src={book.cover_url} alt={book.title} />
                    </div>
                    <div className="book-info">
                      <h3 className="book-title">{book.title}</h3>
                      <p className="book-author">{book.author}</p>
                    </div>
                  </div>
                </a>
              ))
            ) : (
              <div className="no-books">
                <p>Книги в категории "{activeGenre}" не найдены</p>
              </div>
            )}
          </section>
        </div>
      </main>
    </section>
  );
};

export default Catalog;
