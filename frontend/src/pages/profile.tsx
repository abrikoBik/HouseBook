import { useState, useEffect } from "react";
import SideBar from "../modules/sideBar";
import Header from "../modules/header";
import "../styles/profile.css";

interface UserProfile {
  id: string;
  name: string;
  username: string;
  city: string;
  avatar_url?: string;
  stats: {
    total_books: number;
    read: number;
    reading: number;
    available_for_exchange: number;
    rating: number;
  };
  active_trades: Array<{
    id: string;
    book_title: string;
    book_cover_url?: string;
    partner_username: string;
    status: string;
  }>;
  reviews: Array<{
    id: string;
    author_username: string;
    text: string;
    rating: number;
  }>;
}

// ============================================================================
// TODO: API — Эту функцию заменить на реальный fetch при подключении бэкенда
// ============================================================================
const mockFetchUserProfile = async (): Promise<UserProfile> => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve({
        id: "1",
        name: "Шиловский Максим",
        username: "@krutoi.chelik",
        city: "Таганрог",
        avatar_url: "",
        stats: {
          total_books: 7,
          read: 5,
          reading: 1,
          available_for_exchange: 5,
          rating: 4.9,
        },
        active_trades: [
          {
            id: "1",
            book_title: "Лао-цзы",
            book_cover_url: "",
            partner_username: "@fawatafa",
            status: "Ожидает подтверждения",
          },
          {
            id: "2",
            book_title: "Облако",
            book_cover_url: "",
            partner_username: "@genius1337",
            status: "Передача книги",
          },
        ],
        reviews: [
          {
            id: "1",
            author_username: "@zayanozlaya",
            text: "Хороший и самое главное быстрый обмен, пока ждала у меня умерли все близкие",
            rating: 5,
          },
          {
            id: "2",
            author_username: "@mezulina",
            text: "Прочитала предыдущий отзыв, надеялась произойдет тоже самое",
            rating: 1,
          },
        ],
      });
    }, 500);
  });
};

const Profile = () => {
  const [user, setUser] = useState<UserProfile | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const loadProfile = async () => {
      try {
        setLoading(true);
        // TODO: API — Заменить на реальный вызов, например:
        // const response = await fetch(`/api/users/${currentUserId}`);
        // const data = await response.json();
        const data = await mockFetchUserProfile();
        setUser(data);
      } catch (error) {
        console.error("Failed to load profile:", error);
      } finally {
        setLoading(false);
      }
    };

    loadProfile();
  }, []);

  if (loading) {
    return (
      <section className="profile-page">
        <SideBar />
        <main className="profile-content">
          <Header />
          <div style={{ textAlign: "center", padding: "100px" }}>
            Загрузка профиля...
          </div>
        </main>
      </section>
    );
  }

  if (!user) {
    return (
      <section className="profile-page">
        <SideBar />
        <main className="profile-content">
          <Header />
          <div style={{ textAlign: "center", padding: "100px", color: "red" }}>
            Не удалось загрузить профиль
          </div>
        </main>
      </section>
    );
  }

  return (
    <section className="profile-page">
      <SideBar />
      <main className="profile-content">
        <Header />
        <div className="profile-container">
          <div className="profile-header">
            <div className="user-info">
              <div className="avatar-placeholder">Фото</div>
              <div className="user-details">
                <h1>{user.name}</h1>
                <p className="username">{user.username}</p>
                <p className="city">📍 {user.city}</p>
                <button className="edit-btn">Редактировать профиль</button>
              </div>
            </div>

            <div className="stats-card">
              <h2>Статистика</h2>
              <div className="stat-item">
                <span className="stat-icon">📚</span>
                <span>Всего книг: {user.stats.total_books}</span>
              </div>
              <div className="stat-item">
                <span className="stat-icon">✅</span>
                <span>Прочитано: {user.stats.read}</span>
              </div>
              <div className="stat-item">
                <span className="stat-icon">⏳</span>
                <span>В процессе: {user.stats.reading}</span>
              </div>
              <div className="stat-item">
                <span className="stat-icon">🔄</span>
                <span>Доступны для обмена: {user.stats.available_for_exchange}</span>
              </div>
              <div className="stat-item">
                <span className="stat-icon">⭐</span>
                <span>Рейтинг: {user.stats.rating}</span>
              </div>
            </div>
          </div>

          <section className="trades-section">
            <h2>Активные сделки</h2>
            {user.active_trades.length > 0 ? (
              <div className="trades-list">
                {user.active_trades.map((trade) => (
                  <div key={trade.id} className="trade-card">
                    <div className="book-cover-placeholder"></div>
                    <div className="trade-info">
                      <h3>{trade.book_title}</h3>
                      <p>⇄ Обмен с – <strong>{trade.partner_username}</strong></p>
                      <p className="status">◷ {trade.status}</p>
                    </div>
                    <button className="cancel-btn">✕ Отменить</button>
                  </div>
                ))}
              </div>
            ) : (
              <p className="empty-state">Нет активных сделок</p>
            )}
          </section>

          <section className="reviews-section">
            <h2>Отзывы</h2>
            {user.reviews.length > 0 ? (
              <div className="reviews-list">
                {user.reviews.map((review) => (
                  <div key={review.id} className="review-card">
                    <div className="review-header">
                      <div className="stars">
                        {[...Array(5)].map((_, i) => (
                          <span key={i} className={i < review.rating ? "star filled" : "star"}>
                            ☆
                          </span>
                        ))}
                      </div>
                    </div>
                    <p className="review-text">{review.text}</p>
                    <p className="review-author">– {review.author_username}</p>
                  </div>
                ))}
              </div>
            ) : (
              <p className="empty-state">Пока нет отзывов</p>
            )}
          </section>
        </div>
      </main>
    </section>
  );
};

export default Profile;