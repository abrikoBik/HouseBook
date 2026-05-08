import React from "react";
import LogoWText from "../assets/LogoWText.svg";
import Lupa from "../assets/Lupa.svg";
import HomeBlack from "../assets/Home_Black.svg";
import LibraryBlack from "../assets/Library_Black.svg";
import Natification from "../assets/Notification.svg";

import "../styles/landing.css";

const Landing: React.FC = () => {
  return (
    <section id="header">
      <section id="landing-hero">
        <img src={LogoWText} alt="HouseBook Main Logo" />
        <h1 id="hero-title">Добро пожаловать в HouseBook</h1>
        <p id="hero-subtitle">Ваш персональный помощник в мире книг</p>
      </section>

      <section id="features-grid">
        <a href="/book/catalog" className="feature-card-link">
          <div className="card-header">
            <img src={Lupa} alt="Иконка поиска" className="card-icon-inline" />
            <h3>Каталог Книг</h3>
          </div>
          <p>Огромная база данных книг всех жанров и авторов.</p>
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

        <a href="/profile/trade" className="feature-card-link">
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
    </section>
  );
};

export default Landing;
