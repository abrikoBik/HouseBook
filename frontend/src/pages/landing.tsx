import React from 'react';
import Header from '../modules/header';
import LogoWText from '../assets/LogoWText.svg';
import Lupa from '../assets/Lupa.svg';
import HomeBlack from '../assets/Home_Black.svg';
import LibraryBlack from '../assets/Library_Black.svg';
import Natification from '../assets/Notification.svg'; 

import '../styles/landing.css';

const Landing: React.FC = () => {
  return (
    <section id="header">
      <Header />
      <section id="landing-hero">
        <img src={LogoWText} alt="HouseBook Main Logo" />
        <h1 id="hero-title">Добро пожаловать в HouseBook</h1>
        <p id="hero-subtitle">Ваш персональный помощник в мире книг</p>
      </section>

      <section id="features-grid">
        <div className="feature-card">
          <div className="card-header">
            <img src={Lupa} alt="Иконка поиска" className="card-icon-inline" />
            <h3>Каталог Книг</h3>
          </div>
          <p>Огромная база данных книг всех жанров и авторов.</p>
        </div>

        <div className="feature-card">
          <div className="card-header">
            <img src={HomeBlack} alt="Иконка дома" className="card-icon-inline" />
            <h3>Своя библиотека</h3>
          </div>
          <p>Отслеживайте прочитанное, ставьте оценки и пишите рецензии.</p>
        </div>

        <div className="feature-card">
          <div className="card-header">
            <img src={LibraryBlack} alt="Иконка библиотеки" className="card-icon-inline" />
            <h3>Меняйся!</h3>
          </div>
          <p>Нету книги, которую хочешь прочитать? Поменяйся!</p>
        </div>

        <div className="feature-card">
          <div className="card-header">
            <img src={Natification} alt="Иконка уведомлений" className="card-icon-inline" />
            <h3>Личный профиль</h3>
          </div>
          <p>Следи за статистикой, отслеживай сделки и делись впечатлениями!</p>
        </div>
      </section>
    </section>
  );
};

export default Landing;