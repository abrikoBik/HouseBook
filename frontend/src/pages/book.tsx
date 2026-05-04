import BookCover from "../assets/Book.png";
import FavoritesIcon from "../assets/Favourite_Black.svg";
import "../styles/book.css";
import SideBar from "../modules/sideBar";
import Header from "../modules/header";
import Comment from "../modules/comment";
import type Error from "../types/error";
import { type default as BookT } from "../types/book";

import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router";

const Book = () => {
  const navigate = useNavigate();
  const { uuid } = useParams();
  const [book, setBook] = useState<BookT>();

  const getBook = async (id: string) => {
    const response = await fetch(`/api/books/${id}`);
    const result = await response.json();
    if (!response.ok) {
      switch (response.status) {
        case 400:
        case 404:
          navigate("/404", { replace: true });
          break;
        default:
          throw new Error((result as Error).error);
      }
    }
    setBook(result as BookT);
  };

  useEffect(() => {
    getBook(uuid!);
  }, [uuid]);

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
<<<<<<< HEAD
              <p>
                Долгожданное продолжение «Шантарама» — одного из самых
                поразительных романов XXI века. <br />
                Эта исповедь человека, выбравшегося из бездны, разошлась
                многомиллионными тиражами и заслужила сравнения с произведениями
                Мелвилла и Хемингуэя. <br />
                Спустя два года после гибели мафиозного босса Кадербхая и ухода
                возлюбленной Карлы Лин по прозвищу Шантарам должен выполнить
                последнее поручение своего покойного покровителя. <br />
                Ему предстоит завоевать доверие горного мудреца, уцелеть в
                разгорающейся войне мафиозных кланов и, главное, вновь обрести
                любовь и веру.
              </p>
=======
              <p>{book?.description}</p>
>>>>>>> a6fd1c4836fd0fc8173664d471d8d480cacc4ed0
            </div>

            <section className="kniga">
              <div id="about">
                <h4>Издатель</h4>
                <div>{book?.publisher}</div>
              </div>
              {book?.language != null ? (
                <div id="language">
                  <h4>Язык</h4>
                  <div>{book?.language}</div>
                </div>
              ) : (
                <></>
              )}

              <div id="type">
                <h4>Тип обложки</h4>
                <div>
                  {/*TODO: склонение страниц, убрать пустые поля*/}
                  {`${book?.cover_type == null ? "" : book?.cover_type} ${book?.page_count != null ? `, ${book?.page_count} страниц ` : ""} \n ISBN: ${book?.isbn}`}
                </div>
              </div>
            </section>
          </div>

          <div id="comments">
            <h4>Шиловский Максим</h4>
            <p>
              Очень сильная книга с интересным сюжетом и непривычными событиями
              для европейского, русского менталитета. <br />
              Читая эту книгу, я пересказывал почти каждую главу родственникам.
              Книга содержит в себе много философский мыслей и индийской
              мудрости, но при этом не перегружает читателя. <br />
              Чтение «Тень горы» показалось более захватывающим, чем «Шантарам»
              из-за меньшего содержания военных действий. Буду рекомендовать
              книгу знакомым.
            </p>
          </div>

          <div>
            <h4>Попов Даниил</h4>
            <p>
              Хуй пизда гавно какашечки Хуй пизда гавно какашечки Хуй пизда
              гавно какашечки Хуй пизда гавно какашечки Хуй пизда гавно
              какашечки Хуй пизда гавно какашечки Хуй пизда гавно хуй =)
              какашечки Хуй пизда гавно какашечки Хуй пизда гавно какашечки Хуй
              пизда гавно какашечки Хуй пизда гавно какашечки Хуй пизда гавно
              какашечки! Гавно, гавно, гавно, гавно, гавно, гавно, гавно, гавно,
              гавно, гавно, гавно, гавно, гавно, гавно.
            </p>
          </div>

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
