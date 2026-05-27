import "../styles/DropDown.css";
import React, { useState, useRef, useEffect } from "react";

interface DisplayBook {
  id: string;
  title: string;
  author: string;
  cover_url: string;
  exchangeable: boolean;
  status?: "read" | "reading" | "exchangeable";
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
  const [copied, setCopied] = useState(false);
  const [isDropdownOpen, setIsDropdownOpen] = useState(false);

  const dropdownRef = useRef<HTMLDivElement>(null);

  // Конфигурация доступных вариантов для меню
  const statusOptions = [
    { id: "read", label: "Прочитано", class: "status-read" },
    { id: "reading", label: "В процессе", class: "status-reading" },
    {
      id: "exchangeable",
      label: "Доступна для обмена",
      class: "status-exchange",
    },
  ] as const;

  // Определение текущего выбранного значения
  const currentStatusId =
    book.status || (book.exchangeable ? "exchangeable" : "reading");
  const currentStatus =
    statusOptions.find((opt) => opt.id === currentStatusId) || statusOptions[1];

  // Закрытие дропдауна при клике извне
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        dropdownRef.current &&
        !dropdownRef.current.contains(event.target as Node)
      ) {
        setIsDropdownOpen(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  const updateBookState = async (
    statusId: "read" | "reading" | "exchangeable",
  ) => {
    setIsUpdating(true);
    setIsDropdownOpen(false);

    const isExchangeable = statusId === "exchangeable";

    try {
      const response = await fetch(
        `/api/libraries/${library_id}/books/${book.id}`,
        {
          method: "PATCH",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            exchangeable: isExchangeable,
            status: statusId,
          }),
        },
      );

      if (response.ok) {
        onStatusUpdate({
          ...book,
          exchangeable: isExchangeable,
          status: statusId,
        });
      }
    } catch (err) {
      console.error("Ошибка обновления:", err);
    } finally {
      setIsUpdating(false);
    }
  };

  const handleActionButtonClick = () => {
    if (book.exchangeable) {
      updateBookState("reading");
    } else {
      updateBookState("exchangeable");
    }
  };

  const copyToClipboard = () => {
    navigator.clipboard.writeText(book.id);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
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
              {/* Блок ID */}
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

              {/* Блок Статус (исходный вид) */}
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

              {/* Блок Библиотека (новое поле с выпадающим меню) */}
              <div className="detail-item">
                <span className="label">Библиотека:</span>

                <div id="bm-dropdown-container" ref={dropdownRef}>
                  <div
                    id="bm-dropdown-trigger"
                    className={currentStatus.class}
                    onClick={() =>
                      !isUpdating && setIsDropdownOpen(!isDropdownOpen)
                    }
                  >
                    <span>{currentStatus.label}</span>
                    <span className="material-symbols-outlined arrow-icon">
                      {isDropdownOpen ? "expand_less" : "expand_more"}
                    </span>
                  </div>

                  {isDropdownOpen && (
                    <ul id="bm-dropdown-menu">
                      {statusOptions.map((option) => (
                        <li
                          key={option.id}
                          className={`bm-dropdown-option ${option.id === currentStatusId ? "active-option" : ""}`}
                          onClick={() => updateBookState(option.id)}
                        >
                          <span className={`status-dot ${option.class}`}></span>
                          {option.label}
                        </li>
                      ))}
                    </ul>
                  )}
                </div>
              </div>
            </div>

            <div className="modal-actions">
              <button
                className={`action-btn ${book.exchangeable ? "btn-danger" : "btn-success"}`}
                onClick={handleActionButtonClick}
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
