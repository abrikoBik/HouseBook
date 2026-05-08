import { useState, useEffect } from "react";
import SideBar from "../modules/sideBar";
import Header from "../modules/header";
import "../styles/catalog.css";

const API_URL = "http://localhost/api/books";

const genres = [
  "Все",
  "Одолжить",
  "Фантастика",
  "Детектив",
  "Полки",
  "Приключения",
  "Психология",
  "Научная литература",
  "Художественная литература",
  "Классика",
  "Бизнес",
  "Саморазвитие",
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
  const [activeGenre, setActiveGenre] = useState("Все");
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
          cover_url: `Book}`,
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
            <nav className="genre-nav">
              <ul>
                {genres.map((genre) => (
                  <li
                    key={genre}
                    className={activeGenre === genre ? "active" : ""}
                    onClick={() => setActiveGenre(genre)}
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
                <p>Книги в этой категории не найдены</p>
              </div>
            )}
          </section>
        </div>
      </main>
    </section>
  );
};

export default Catalog;
