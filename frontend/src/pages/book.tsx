import BookCover from "../assets/Book.png";
import FavoritesIcon from "../assets/Favorites.png";
import "../styles/book.css";
import SideBar from "../modules/sideBar";
import Header from "../modules/header";
import Comment from "../modules/comment";
import type Error from "../types/error";
import type BookT from "../types/book";
import { useEffect, useState } from "react";

const Book = () => {
  var id = "beb2c6cd-1674-4616-8b8a-b6a281e15089";

  const [book, setBook] = useState<BookT>();

  async function getBook(uuid: string) {
    try {
      const response = await fetch(`/api/books/${uuid}`);
      const result = await response.json();
      console.log(response);
      setBook(result);
    } catch (error: any) {
      console.error(error as Error);
    }
  }

  useEffect(() => {
    getBook(id);
  }, []);

  return (
    <section id="layout">
      <Header />
      <SideBar />
      <section id="bookBody">
        <div id="cover">
          <img src={BookCover} alt="book" />
          <div>
            <div>
              <h1 className={"desc"}>{book?.title}</h1>
              <h4 className={"desc"}>{book?.author}</h4>
            </div>
            <div>
              <button id="bronya">Забронировать</button>
              <div id="favoriteIcon">
                <img src={FavoritesIcon} alt="favorites" />
              </div>
            </div>
          </div>
        </div>

        <section className="bookDescription">
          <div id="description">
            <div>
              <h4>Описание</h4>
              <p>{book?.description}</p>
            </div>

            <section className="kniga">
              <div id="about">
                <h4>Издатель</h4>
                <div>{book?.publisher}</div>
              </div>

              <div id="language">
                <h4>Язык</h4>
                <div>{book?.language}</div>
              </div>

              <div id="type">
                <h4>Тип обложки</h4>
                <div>
                  {/*TODO: склонение страниц*/}
                  {book?.cover_type}, {book?.page_count} страниц <br />
                  ISBN: {book?.isbn}
                </div>
              </div>
            </section>
          </div>

          <Comment />

          <div id="textarea">
            <textarea />
            <div>
              <div>U</div>
              <div>I</div>
              <div>A</div>
              <button>Отправить</button>
            </div>
          </div>
        </section>
      </section>
    </section>
  );
};

export default Book;
