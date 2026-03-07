import "../styles/profile.css";

const Profile = () => {
  return (
    <section className="user">
      <div id="aboutUser">
        <h4>Профиль</h4>
        <div>Шиловский Максим</div>
        <div>Таганрог</div>
        <div>
          <button id="editProfile">Редактировать профиль</button>
        </div>
      </div>

      <section className="stats">
        <div id="aboutStats">
          <h4>Статистика</h4>
          <div>Всего книг: 7</div>
          <div>Прочитано: 5</div>
          <div>В процессе: 1</div>
          <div>Доступны для обмена: 5</div>
          <div>Рейтинг: 4,9</div>
        </div>

        <section className="activeTrades">
          <div id="aboutTrades">
            <h4>Активные сделки</h4>
            <div>Лао-цзы</div>
            <div>Облако</div>
          </div>

          <section className="reviews">
            <div>
              <h4>Отзывы</h4>
              <div>Хороший и быстрый обмен</div>
            </div>
          </section>
        </section>
      </section>
    </section>
  );
};

export default Profile;
