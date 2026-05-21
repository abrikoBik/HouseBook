import { useState, useEffect } from "react";
import SideBar from "../modules/sideBar";
import Header from "../modules/header";
import "../styles/profile.css";

interface UserAPI {
  email: string;
  finished_reading_id: string;
  id: string;
  image_uuid: string;
  login: string;
  name: string;
  now_reading_id: string;
  rating: number;
  trade_ready_id: string;
}

const Profile = () => {
  const [user, setUser] = useState<UserAPI | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // Получение userId из localStorage
  const getUserId = (): string | null => {
    if (typeof window !== "undefined") {
      return localStorage.getItem("UserID");
    }
    return null;
  };

  useEffect(() => {
    const loadProfile = async () => {
      try {
        setLoading(true);
        setError(null);

        const userId = getUserId();
        if (!userId) {
          throw new Error("Пользователь не авторизован");
        }

        const response = await fetch(`/api/users/${userId}`);
        if (!response.ok) {
          throw new Error("Ошибка загрузки профиля");
        }

        const data: UserAPI = await response.json();
        setUser(data);
      } catch (error) {
        console.error("Failed to load profile:", error);
        setError("Не удалось загрузить профиль");
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

  if (error || !user) {
    return (
      <section className="profile-page">
        <SideBar />
        <main className="profile-content">
          <Header />
          <div style={{ textAlign: "center", padding: "100px", color: "red" }}>
            {error || "Не удалось загрузить профиль"}
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
              <div className="avatar-placeholder">
                {user.image_uuid ? (
                  <img src={`/api/images/${user.image_uuid}`} alt="Avatar" />
                ) : (
                  "Фото"
                )}
              </div>
              <div className="user-details">
                <h1>{user.name}</h1>
                <p className="username">@{user.login}</p>
                <p className="city">📍 Таганрог</p>
                <button className="edit-btn">Редактировать профиль</button>
              </div>
            </div>

            <div className="stats-card">
              <h2>Статистика</h2>
              <div className="stat-item">
                <span className="stat-icon">⭐</span>
                <span>Рейтинг: {user.rating.toFixed(1)}</span>
              </div>
              <div className="stat-item">
                <span className="stat-icon">📖</span>
                <span>Читает сейчас: {user.now_reading_id || "—"}</span>
              </div>
              <div className="stat-item">
                <span className="stat-icon">✅</span>
                <span>Прочитано: {user.finished_reading_id || "—"}</span>
              </div>
              <div className="stat-item">
                <span className="stat-icon">🔄</span>
                <span>Готов к обмену: {user.trade_ready_id || "—"}</span>
              </div>
            </div>
          </div>
        </div>
      </main>
    </section>
  );
};

export default Profile;
