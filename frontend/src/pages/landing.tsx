import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom"; // Используем Link для плавной навигации
import LogoWText from "../assets/LogoWText.svg";
import Lupa from "../assets/Lupa.svg";
import HomeBlack from "../assets/Home_Black.svg";
import LibraryBlack from "../assets/Library_Black.svg";
import Natification from "../assets/Notification.svg";

import "../styles/landing.css";

const Landing: React.FC = () => {
  // Состояние для проверки авторизации
  const [isAuthorized, setIsAuthorized] = useState<boolean>(false);

  useEffect(() => {
    // Проверяем наличие токена при загрузке страницы
    const token = localStorage.getItem("token");
    if (token) {
      setIsAuthorized(true);
    }
  }, []);

  return (
    <section id="header">
      <section id="landing-hero">
        <img src={LogoWText} alt="HouseBook Main Logo" />
        <h1 id="hero-title">Добро пожаловать в HouseBook</h1>
        <p id="hero-subtitle">Ваш персональный помощник в мире книг</p>
      </section>

      <section id="features-grid">
        <a href="/central" className="feature-card-link">
          <div className="card-header">
            <img src={Lupa} alt="Иконка поиска" className="card-icon-inline" />
            <h3>Рекомендации</h3>
          </div>
          <p>Лучшие предложения только для вас!</p>
        </a>

        <a href="/library" className="feature-card-link">
          <div className="card-header">
            <img
              src={HomeBlack}
              alt="Иконка дома"
              className="card-icon-inline"
            />
            <h3>Своя библиотека</h3>
          </div>
          <p>Отслеживайте прочитанное, ставьте оценки и пишите рецензии.</p>
        </a>

        <a href="/profile/exchange" className="feature-card-link">
          <div className="card-header">
            <img
              src={LibraryBlack}
              alt="Иконка библиотеки"
              className="card-icon-inline"
            />
            <h3>Меняйся!</h3>
          </div>
          <p>Нету книги, которую хочешь прочитать? Обменяйся!</p>
        </a>

        <a href="/profile" className="feature-card-link">
          <div className="card-header">
            <img
              src={Natification}
              alt="Иконка уведомлений"
              className="card-icon-inline"
            />
            <h3>Личный профиль</h3>
          </div>
          <p>Следи за статистикой, отслеживай сделки и делись впечатлениями!</p>
        </a>
      </section>

      {/* Кнопки показываются ТОЛЬКО если пользователь НЕ авторизован */}
      {!isAuthorized && (
        <section id="landing-actions">
          <Link to="/auth?mode=login" className="btn-landing btn-login">
            Войти
          </Link>
          <Link to="/auth?mode=register" className="btn-landing btn-register">
            Зарегистрироваться
          </Link>
        </section>
      )}
    </section>
  );
};

export default Landing;
