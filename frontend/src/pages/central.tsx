import { useState, useEffect } from "react";
import SideBar from "../modules/sideBar";
import Header from "../modules/header";
import "../styles/central.css";
import type BookT from "../types/book";

const Central = () => {
  const [recommendedBooks, setRecommendedBooks] = useState<BookT[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchRandomBooks = async () => {
      try {
        setLoading(true);
        const res = await fetch("http://localhost/api/books/recommended");
        if (!res.ok) throw new Error("Failed to fetch books");
        const books = await res.json();
        setRecommendedBooks(books);
      } catch (error) {
        console.error("Error loading recommendations:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchRandomBooks();
  }, []);

  // ... (остальной код без изменений) ...

  return (
    <section className="central-page">
      <SideBar />
      <main className="central-content">
        <Header />
        <div className="central-container">
          <section className="recommendations-section">
            <h1>Рекомендации</h1>
            <p className="subtitle">
              Вам также может быть интересно. Мы подобрали для вас эти материалы
              на основе ваших текущих интересов и активности на нашем сайте.
            </p>

            <a href="/book/catalog" className="expand-btn">
              В полный каталог →
            </a>

            {loading ? (
              <div className="loading-books">Загрузка рекомендаций...</div>
            ) : (
              <div className="books-carousel">
                {recommendedBooks.map((book) => (
                  <a
                    href={`/book/${book.id}`}
                    key={book.id}
                    className="book-card-link"
                  >
                    <div className="book-card">
                      <div className="book-cover">
                        {book.image_uuid ? (
                          <img
                            src={`/cdn/${book.image_uuid}`}
                            alt={book.title}
                          />
                        ) : (
                          <div className="cover-placeholder"></div>
                        )}
                      </div>
                      <div className="book-info">
                        <h3>{book.title}</h3>
                        <p>{book.author}</p>
                      </div>
                    </div>
                  </a>
                ))}
              </div>
            )}
          </section>
        </div>
      </main>
    </section>
  );
};

export default Central;