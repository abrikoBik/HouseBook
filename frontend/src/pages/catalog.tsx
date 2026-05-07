import { useState } from "react";
import SideBar from "../modules/sideBar";
import Header from "../modules/header";
import "../styles/catalog.css";

const genres = [
  "Все",
  "Одолжить",
  "Фантастика",
  "Детектив",
  "Полки",
  "Приключения",
  "Психология",
  "Научная литература",
  "Художественная литература",
  "Классика",
  "Бизнес",
  "Саморазвитие",
];

const Catalog = () => {
  const [activeGenre, setActiveGenre] = useState("Все");

  return (
    <section id="catalog-body">
      <SideBar />

      <main id="catalog-main">
        <Header />

        <div id="catalog-layout">
          <section id="catalog-header-section">
            <div className="header-content">
              <h1>Каталог</h1>

              <div className="nav-arrows">
                <button className="arrow-btn left">&larr;</button>
                <button className="arrow-btn right">&rarr;</button>
              </div>
            </div>

            <nav className="genre-nav">
              <ul>
                {genres.map((genre) => (
                  <li
                    key={genre}
                    className={activeGenre === genre ? "active" : ""}
                    onClick={() => setActiveGenre(genre)}
                  >
                    {genre}
                  </li>
                ))}
              </ul>
            </nav>
          </section>

          <section id="book-grid-placeholder">
            <p>Здесь будет сетка книг...</p>
          </section>
        </div>
      </main>
    </section>
  );
};

export default Catalog;
