import React, { useState } from "react";

interface DisplayBook {
  id: string;
  title: string;
  author: string;
  cover_url: string;
  exchangeable: boolean;
}

interface BookModalProps {
  book: DisplayBook;
  library_id: string;
  onClose: () => void;
  onStatusUpdate: (updatedBook: DisplayBook) => void;
}

const BookModal: React.FC<BookModalProps> = ({
  book,
  library_id,
  onClose,
  onStatusUpdate,
}) => {
  const [isUpdating, setIsUpdating] = useState(false);
  const [copied, setCopied] = useState(false); // Состояние для анимации кнопки копирования

  const toggleExchangeStatus = async () => {
    setIsUpdating(true);
    try {
      const response = await fetch(
        `/api/libraries/${library_id}/books/${book.id}`,
        {
          method: "PATCH",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ exchangeable: !book.exchangeable }),
        },
      );

      if (response.ok) {
        onStatusUpdate({ ...book, exchangeable: !book.exchangeable });
      }
    } catch (err) {
      console.error("Ошибка обновления:", err);
    } finally {
      setIsUpdating(false);
    }
  };

  // Функция копирования
  const copyToClipboard = () => {
    navigator.clipboard.writeText(book.id);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000); // Сброс статуса через 2 секунды
  };

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-container" onClick={(e) => e.stopPropagation()}>
        <button className="modal-close-icon" onClick={onClose}>
          <span className="material-symbols-outlined">close</span>
        </button>

        <div className="modal-grid">
          <div className="modal-image-section">
            <img src={book.cover_url} alt={book.title} />
          </div>

          <div className="modal-info-section">
            <div className="modal-header-text">
              <h2>{book.title}</h2>
              <h3>{book.author}</h3>
            </div>

            <div className="modal-details">
              <div className="detail-item">
                <span className="label">ID книги:</span>
                <div
                  className="copy-wrapper"
                  onClick={copyToClipboard}
                  title="Копировать ID"
                >
                  <span className="value">#{book.id.slice(0, 8)}...</span>
                  <span
                    className={`material-symbols-outlined copy-icon ${copied ? "copied" : ""}`}
                  >
                    {copied ? "done" : "content_copy"}
                  </span>
                </div>
              </div>

              <div className="detail-item">
                <span className="label">Статус:</span>
                <span
                  className={`status-pill ${book.exchangeable ? "active" : "inactive"}`}
                >
                  {book.exchangeable
                    ? "Доступна для обмена"
                    : "Личная коллекция"}
                </span>
              </div>
            </div>

            <div className="modal-actions">
              <button
                className={`action-btn ${book.exchangeable ? "btn-danger" : "btn-success"}`}
                onClick={toggleExchangeStatus}
                disabled={isUpdating}
              >
                <span className="material-symbols-outlined">
                  {book.exchangeable ? "block" : "swap_horiz"}
                </span>
                {isUpdating
                  ? "Загрузка..."
                  : book.exchangeable
                    ? "Убрать из обмена"
                    : "Сделать доступной"}
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default BookModal;
