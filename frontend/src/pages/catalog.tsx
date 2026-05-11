import { useState, useEffect } from "react";
import { useSearchParams } from "react-router-dom";
import SideBar from "../modules/sideBar";
import Header from "../modules/header";
import "../styles/catalog.css";

const API_URL = "http://localhost/api/books";
const ITEMS_PER_PAGE = 10;

const genres = [
  "Все",
  "Фантастика",
  "Детективы",
  "Триллеры",
  "Ужасы",
  "Приключения",
  "Научная фантастика",
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

  const initialGenre = urlGenre && genres.includes(urlGenre) ? urlGenre : "Все";

  const [activeGenre, setActiveGenre] = useState(initialGenre);
  const [currentPageBooks, setCurrentPageBooks] = useState<DisplayBook[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const [currentPage, setCurrentPage] = useState(1);
  const [hasNextPage, setHasNextPage] = useState(false);

  useEffect(() => {
    if (urlGenre && genres.includes(urlGenre)) {
      setActiveGenre(urlGenre);
    } else {
      setActiveGenre("Все");
    }
    setCurrentPage(1);
  }, [urlGenre]);

  useEffect(() => {
    const fetchBooks = async () => {
      try {
        setLoading(true);

        const params = new URLSearchParams({
          limit: (ITEMS_PER_PAGE + 1).toString(),
          offset: ((currentPage - 1) * ITEMS_PER_PAGE).toString(),
        });

        if (activeGenre !== "Все") {
          params.append("genre", activeGenre);
        }

        const res = await fetch(`${API_URL}?${params}`);
        if (!res.ok) throw new Error(`Ошибка сервера: ${res.status}`);

        const data: ApiBook[] = await res.json();

        const hasMore = data.length > ITEMS_PER_PAGE;
        setHasNextPage(hasMore);

        let booksToShow = data.slice(0, ITEMS_PER_PAGE);

        if (activeGenre !== "Все") {
          booksToShow = booksToShow.filter(
            (book: ApiBook) => book.genre !== "Неизвестен",
          );
        }

        const mapped: DisplayBook[] = booksToShow.map((book: ApiBook) => ({
          id: book.id,
          title: book.title?.trim() || "Без названия",
          author: book.author?.trim() || "Неизвестный автор",
          genre: book.genre?.trim() || "Неизвестен",
          cover_url: book.image_uuid
            ? `http://localhost/cdn/${book.image_uuid}`
            : `https://via.placeholder.com/300x450?text=${encodeURIComponent(book.title || "Book")}`,
        }));

        setCurrentPageBooks(mapped);
      } catch (err) {
        setError(
          err instanceof Error ? err.message : "Не удалось загрузить каталог",
        );
      } finally {
        setLoading(false);
      }
    };

    fetchBooks();
  }, [activeGenre, currentPage]);

  const handlePrevPage = () => {
    if (currentPage > 1) {
      setCurrentPage((p) => p - 1);
      window.scrollTo({ top: 0, behavior: "smooth" });
    }
  };

  const handleNextPage = () => {
    if (hasNextPage) {
      setCurrentPage((p) => p + 1);
      window.scrollTo({ top: 0, behavior: "smooth" });
    }
  };

  if (loading && currentPage === 1) {
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
                    onClick={() => {
                      setActiveGenre(genre);
                      setCurrentPage(1);
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
            {currentPageBooks.length > 0 ? (
              currentPageBooks.map((book) => (
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

          {/* Пагинация: показываем если есть куда вернуться ИЛИ если на первой странице 10+ книг и есть следующие */}
          {(currentPage > 1 ||
            (hasNextPage && currentPageBooks.length >= ITEMS_PER_PAGE)) && (
            <div className="pagination">
              <button
                className="pagination-btn"
                onClick={handlePrevPage}
                disabled={currentPage === 1}
              >
                ← Назад
              </button>

              <span className="pagination-info">Страница {currentPage}</span>

              <button
                className="pagination-btn"
                onClick={handleNextPage}
                disabled={!hasNextPage}
              >
                Вперёд →
              </button>
            </div>
          )}
        </div>
      </main>
    </section>
  );
};

export default Catalog;
