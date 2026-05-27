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
import { useParams, useNavigate } from "react-router-dom";

const Book = () => {
  const navigate = useNavigate();
  const { uuid } = useParams();
  const [book, setBook] = useState<BookT>();
  const [bookComments, setBookComments] = useState<CommentT[]>();

  const getComments = async (id: string) => {
    try {
      const response = await fetch(`/api/books/${id}/comments`);
      const result = await response.json();
      if (!response.ok) {
        throw new Error((result as Error).error);
      }
      setBookComments(result);
    } catch (err) {
      console.error("Failed to load comments:", err);
    }
  };

  useEffect(() => {
    const getBook = async (id: string) => {
      try {
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
      } catch (err) {
        console.error("Failed to load book:", err);
      }
    };

    if (uuid) {
      getBook(uuid);
      getComments(uuid);
    }
  }, [uuid, navigate]);

  const formatBookInfo = () => {
    if (!book) return "";
    const parts = [];
    if (book.cover_type) parts.push(book.cover_type);
    if (book.page_count) parts.push(`${book.page_count} страниц`);

    let infoString = parts.join(", ");
    if (book.isbn) {
      infoString += infoString ? `\nISBN: ${book.isbn}` : `ISBN: ${book.isbn}`;
    }
    return infoString;
  };

  if (!book) return <div className="loading-state">Загрузка...</div>;

  return (
    <section id="layout">
      <Header />
      <SideBar />

      {/* Основной контейнер с отступом под сайдбар */}
      <main id="book-page-container">
        {/* Обертка для ограничения максимальной ширины контента */}
        <div className="content-wrapper">
          {/* Верхняя секция: Обложка + Инфо */}
          <section id="book-hero">
            <div className="hero-cover-wrapper">
              <img
                src={
                  book.image_uuid
                    ? `/cdn/${book.image_uuid}`
                    : "https://via.placeholder.com/300x450?text=No+Cover"
                }
                alt={book.title}
                className="hero-cover"
              />
            </div>

            <div className="hero-info">
              <h1 className="book-title">{book.title}</h1>
              <h2 className="book-author">{book.author}</h2>

              <div className="hero-actions">
                <button id="bronya-btn">Забронировать</button>
                <button id="favorite-btn" aria-label="Добавить в избранное">
                  <img src={FavoritesIcon} alt="favorites" />
                </button>
              </div>
            </div>
          </section>

          {/* Нижняя секция: Карточка с описанием и деталями */}
          <section className="book-content-card">
            <div className="content-grid">
              {/* Левая часть: Описание */}
              <div className="description-col">
                <h3>О книге</h3>
                <p>{book.description || "Описание отсутствует."}</p>
              </div>

              {/* Правая часть: Характеристики (в отдельном блоке) */}
              <div className="details-col">
                <div className="detail-item">
                  <h4>Издательство</h4>
                  <p>{book.publisher || "Не указано"}</p>
                </div>

                {book.language && (
                  <div className="detail-item">
                    <h4>Язык оригинала</h4>
                    <p>{book.language}</p>
                  </div>
                )}

                <div className="detail-item">
                  <h4>Характеристики</h4>
                  <p className="specs-text">{formatBookInfo()}</p>
                </div>
              </div>
            </div>

            {/* Секция комментариев */}
            <div className="comments-section">
              <h3>Отзывы читателей</h3>

              <div className="comments-list">
                {bookComments && bookComments.length > 0 ? (
                  bookComments.map((comment) => (
                    <Comment
                      key={comment.id}
                      userID={comment.user_id}
                      commentBody={comment.body}
                    />
                  ))
                ) : (
                  <p className="no-comments">
                    Пока нет отзывов. Будьте первым!
                  </p>
                )}
              </div>

              <CommentArea
                BookID={uuid!}
                onCommentSent={() => getComments(uuid!)}
              />
            </div>
          </section>
        </div>
      </main>
    </section>
  );
};

export default Book;
