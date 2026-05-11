import FavoritesIcon from "../assets/Favourite_Black.svg";
import "../styles/book.css";
import SideBar from "../modules/sideBar";
import Header from "../modules/header";
import Comment from "../modules/comment";
import CommentArea from "../modules/commentArea";
import type Error from "../types/error";
import type BookT from "../types/book";
import type CommentT from "../types/book";

import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router";

const Book = () => {
  const navigate = useNavigate();
  const { uuid } = useParams();
  const [book, setBook] = useState<BookT>();
  const [bookComments, setBookComments] = useState<CommentT[]>();

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
    const book = result as BookT;
    setBook(book);
  };

  const getComments = async (id: string) => {
    const response = await fetch(`/api/books/${id}/comments`);
    const result = await response.json();
    if (!response.ok) {
      throw new Error((result as Error).error);
    }
    setBookComments(result);
  };

  useEffect(() => {
    getBook(uuid!);
    getComments(uuid!);
  }, [uuid, navigate]);

  return (
    <section id="layout">
      <Header />
      <SideBar />
      <section id="bookBody">
        <div id="cover">
          <img src={`/cdn/${book?.image_uuid}`} alt="book" />
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

          {bookComments?.map((comment) => (
            <Comment
              key={comment.id}
              userID={comment.user_id}
              commentBody={comment.body}
            />
          ))}

          <CommentArea
            BookID={uuid!}
            onCommentSent={() => getComments(uuid!)}
          />
        </section>
      </section>
    </section>
  );
};

export default Book;
