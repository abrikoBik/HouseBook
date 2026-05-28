import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import type BookT from "../types/book";
import type Error from "../types/error";
import "../styles/avaibleBookList.css";

interface ExchangeableBookT extends BookT {
  user_id: string;
  library_id: string;
}

interface AvailableBookListProps {
  isOpen: boolean;
  onClose: () => void;
  bookId: string;
  currentBook: BookT | undefined;
}

/* ==========================================================================
   Суб-компонент карточки книги
   ========================================================================== */
interface ExchangeBookCardProps {
  book: ExchangeableBookT;
  onSelect: (ownerId: string, libraryId: string) => void;
  disabled: boolean;
}

const ExchangeBookCard = ({
  book,
  onSelect,
  disabled,
}: ExchangeBookCardProps) => {
  const [userName, setUserName] = useState<string>("Загрузка...");
  const [libraryName, setLibraryName] = useState<string>("Загрузка...");

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const response = await fetch(`/api/users/${book.user_id}`);
        if (response.ok) {
          const data = await response.json();
          setUserName(data.name || "Неизвестный пользователь");
        } else {
          setUserName("Пользователь не найден");
        }
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
      } catch (err: any) {
        setUserName(err);
      }
    };

    const fetchLibrary = async () => {
      try {
        const response = await fetch(`/api/libraries/${book.library_id}`);
        if (response.ok) {
          const data = await response.json();
          setLibraryName(data.name || "Неизвестная библиотека");
        } else {
          setLibraryName("Библиотека не найдена");
        }
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
      } catch (err: any) {
        setLibraryName(err);
      }
    };

    if (book.user_id) fetchUser();
    if (book.library_id) fetchLibrary();
  }, [book.user_id, book.library_id]);

  return (
    <div
      className={`exchange-book-item ${disabled ? "disabled" : ""}`}
      onClick={() => !disabled && onSelect(book.user_id, book.library_id)}
      style={{ cursor: disabled ? "not-allowed" : "pointer" }}
    >
      <div className="exchange-book-info">
        <div className="exchange-book-meta">
          <p>
            <strong>Владелец:</strong> {userName}
          </p>
          <p>
            <strong>Библиотека:</strong> {libraryName}
          </p>
        </div>
      </div>
    </div>
  );
};

/* ==========================================================================
   Основной компонент Модального Окна
   ========================================================================== */
const AvailableBookList = ({
  isOpen,
  onClose,
  bookId,
  currentBook,
}: AvailableBookListProps) => {
  const navigate = useNavigate();
  const [books, setBooks] = useState<ExchangeableBookT[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [isTrading, setIsTrading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  // Храним признак того, была ли скрыта книга текущего пользователя
  const [hasUserBookHidden, setHasUserBookHidden] = useState<boolean>(false);

  useEffect(() => {
    if (!isOpen) return;

    const fetchExchangeableBooks = async () => {
      try {
        setIsLoading(true);
        setError(null);
        setHasUserBookHidden(false);

        const response = await fetch(`/api/books/${bookId}/exchangeable`);
        const result = await response.json();

        if (!response.ok) {
          throw new Error(
            (result as Error).error || "Не удалось загрузить список книг",
          );
        }

        const fetchedBooks = (result || []) as ExchangeableBookT[];
        const currentUserId = localStorage
          .getItem("UserID")
          ?.toString()
          .trim()
          .toLowerCase();

        // Проверяем совпадение ID владельца с текущим пользователем
        const hasOwnBook = fetchedBooks.some((book) => {
          if (!book.user_id || !currentUserId) return false;
          return book.user_id.toString().trim().toLowerCase() === currentUserId;
        });

        if (hasOwnBook) {
          setHasUserBookHidden(true);
          setBooks([]); // Очищаем массив предложений
        } else {
          setBooks(fetchedBooks);
        }
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
      } catch (err: any) {
        console.error("Failed to load exchangeable books:", err);
        setError(err.message || "Произошла ошибка при загрузке");
      } finally {
        setIsLoading(false);
      }
    };

    fetchExchangeableBooks();
  }, [isOpen, bookId]);

  // Функция создания сделки
  const handleCreateTrade = async (ownerId: string, libraryId: string) => {
    const currentUserId = localStorage.getItem("UserID");

    if (!currentUserId) {
      alert("Сессия истекла. Пожалуйста, авторизуйтесь заново.");
      navigate("/auth");
      return;
    }

    try {
      setIsTrading(true);
      setError(null);

      const tradeResponse = await fetch("/api/trades", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          ownerid: ownerId,
          buyerid: currentUserId,
        }),
      });

      const tradeResult = await tradeResponse.json();

      if (!tradeResponse.ok) {
        throw new Error(
          (tradeResult as Error).error || "Не удалось открыть сделку",
        );
      }

      const tradeId = tradeResult.id;

      const bookResponse = await fetch(`/api/trades/${tradeId}/books`, {
        method: "PATCH",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          bookid: bookId,
          ownerid: ownerId,
          libraryid: libraryId,
        }),
      });

      if (!bookResponse.ok) {
        const bookResult = await bookResponse.json();
        throw new Error(
          (bookResult as Error).error ||
            "Сделка создана, но привязать книгу не удалось",
        );
      }

      alert("Запрос на обмен успешно отправлен!");
      onClose();
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } catch (err: any) {
      console.error("Trade creation error:", err);
      setError(err.message || "Ошибка при оформлении обмена");
    } finally {
      setIsTrading(false);
    }
  };

  if (!isOpen) return null;

  return (
    <div className="modal-overlay" onClick={isTrading ? undefined : onClose}>
      <div className="modal-content" onClick={(e) => e.stopPropagation()}>
        <div className="modal-header">
          <div className="modal-title-block">
            <h3>Книги, доступные для обмена</h3>
            {currentBook && (
              <p className="modal-subtitle">
                «{currentBook.title}» — {currentBook.author}
              </p>
            )}
          </div>
          <button
            className="modal-close-btn"
            onClick={onClose}
            disabled={isTrading}
            aria-label="Закрыть"
          >
            &times;
          </button>
        </div>

        <div className="modal-body">
          {isLoading && (
            <div className="modal-loading">Загрузка списка предложений...</div>
          )}
          {isTrading && (
            <div className="modal-loading" style={{ color: "#007bff" }}>
              Оформление сделки, пожалуйста, подождите...
            </div>
          )}

          {error && <div className="modal-error">{error}</div>}

          {!isLoading && !isTrading && !error && (
            <div className="exchange-books-list">
              {books.length > 0 ? (
                books.map((item) => (
                  <ExchangeBookCard
                    key={item.id}
                    book={item}
                    disabled={isTrading}
                    onSelect={handleCreateTrade}
                  />
                ))
              ) : (
                /* Меняем текст внутри существующего класса no-books-message на основе флага */
                <p className="no-books-message">
                  {hasUserBookHidden
                    ? "У вас уже есть такая книга в библиотеке"
                    : "К сожалению, сейчас нет доступных предложений для обмена этой книги."}
                </p>
              )}
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default AvailableBookList;
