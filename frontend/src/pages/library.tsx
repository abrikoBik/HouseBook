import { useState } from "react";
import Header from "../modules/header";
import SideBar from "../modules/sideBar";
import "../styles/library.css";

const Library = () => {
  const [activeTab, setActiveTab] = useState(0);

  return (
    <div>
      <SideBar />
      <Header />

      <section id="libraryLayout">
        <section id="info">
          <h4>Моя библиотека</h4>
          <nav>
            <ul>
              <li
                className={activeTab === 0 ? "active" : ""}
                onClick={() => setActiveTab(0)}
              >
                Все (кол-во)
              </li>
              <li
                className={activeTab === 1 ? "active" : ""}
                onClick={() => setActiveTab(1)}
              >
                Прочитанные (кол-во)
              </li>
              <li
                className={activeTab === 2 ? "active" : ""}
                onClick={() => setActiveTab(2)}
              >
                В процессе (кол-во)
              </li>
              <li
                className={activeTab === 3 ? "active" : ""}
                onClick={() => setActiveTab(3)}
              >
                Хочу прочитать (кол-во)
              </li>
              <li
                className={activeTab === 4 ? "active" : ""}
                onClick={() => setActiveTab(4)}
              >
                Избранные (кол-во)
              </li>
              <li
                className={activeTab === 5 ? "active" : ""}
                onClick={() => setActiveTab(5)}
              >
                Доступны для обмена (кол-во)
              </li>
            </ul>
          </nav>
        </section>
        <section id="books">
          <div>Книга 1</div>
          <div>Книга 2</div>
        </section>
      </section>
    </div>
  );
};

export default Library;
