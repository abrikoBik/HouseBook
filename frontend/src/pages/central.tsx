import { useState, useEffect } from "react";
import SideBar from "../modules/sideBar";
import Header from "../modules/header";
import "../styles/central.css";

// Импортируем иконки категорий
import ComedyIcon from "../assets/Comedy.png";
import DocumentaryIcon from "../assets/Documentary.png";
import NovelIcon from "../assets/Novel.png";
import HistoryIcon from "../assets/History.png";
import SciFiIcon from "../assets/Sci-Fi.png";
import ActionIcon from "../assets/Action.png";
import DramaIcon from "../assets/Drama.png";
import AdventureIcon from "../assets/Adventure.png";
import ThrillerIcon from "../assets/Thriller.png";
import FantasyIcon from "../assets/Fantasy.png";
import HorrorIcon from "../assets/Horror.png";
import DetectiveIcon from "../assets/Detective.png";

interface Book {
  id: string;
  title: string;
  author: string;
  cover_url?: string;
}

interface Category {
  id: string;
  name: string;
  icon_url: string;
}

const Central = () => {
  const [recommendedBooks, setRecommendedBooks] = useState<Book[]>([]);
  const [loading, setLoading] = useState(true);

  const categories: Category[] = [
    { id: "Фантастика", name: "Фантастика", icon_url: FantasyIcon },
    { id: "Детектив", name: "Детективы", icon_url: DetectiveIcon },
    { id: "Триллеры", name: "Триллеры", icon_url: ThrillerIcon },
    { id: "Ужасы", name: "Ужасы", icon_url: HorrorIcon },
    { id: "Приключения", name: "Приключения", icon_url: AdventureIcon },
    { id: "Наука", name: "Научная фантастика", icon_url: SciFiIcon },
    { id: "История", name: "История", icon_url: HistoryIcon },
    { id: "Комедия", name: "Комедия", icon_url: ComedyIcon },
    { id: "Документальное", name: "Документальное", icon_url: DocumentaryIcon },
    { id: "Романы", name: "Романы", icon_url: NovelIcon },
    { id: "Экшен", name: "Экшен", icon_url: ActionIcon },
    { id: "Драма", name: "Драма", icon_url: DramaIcon },
  ];

  useEffect(() => {
    const fetchRandomBooks = async () => {
      try {
        setLoading(true);
        const res = await fetch("http://localhost/api/books");
        if (!res.ok) throw new Error("Failed to fetch books");

        const allBooks: Book[] = await res.json();

        const shuffled = [...allBooks].sort(() => 0.5 - Math.random());
        const randomFive = shuffled.slice(0, 5).map((book) => ({
          id: book.id,
          title: book.title || "Без названия",
          author: book.author || "Неизвестный автор",
          cover_url: book.cover_url || "",
        }));

        setRecommendedBooks(randomFive);
      } catch (error) {
        console.error("Error loading recommendations:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchRandomBooks();
  }, []);

  return (
    <section className="central-page">
      <SideBar />
      <main className="central-content">
        <Header />
        <div className="central-container">
          {/* --- ЧАСТЬ 1: РЕКОМЕНДАЦИИ --- */}
          <section className="recommendations-section">
            <h1>Рекомендации</h1>
            <p className="subtitle">
              Вам также может быть интересно. Мы подобрали для вас эти материалы
              на основе ваших текущих интересов и активности на нашем сайте.
            </p>

            <button className="expand-btn">Развернуть →</button>

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
                        {book.cover_url ? (
                          <img src={book.cover_url} alt={book.title} />
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

          {/* --- ЧАСТЬ 2: КАТЕГОРИИ С ИКОНКАМИ --- */}
          <section className="categories-section">
            <h2 className="section-title">Категории книг</h2>

            <div className="categories-grid">
              {categories.map((category) => (
                <a
                  href={`/book/catalog?genre=${encodeURIComponent(category.id)}`}
                  key={category.id}
                  className="category-card"
                >
                  <img
                    src={category.icon_url}
                    alt={category.name}
                    className="category-icon"
                  />
                  <p className="category-name">{category.name}</p>
                </a>
              ))}
            </div>
          </section>
        </div>
      </main>
    </section>
  );
};

export default Central;
