// src/pages/scrape.tsx
import { useState, useRef } from "react";
import SideBar from "../modules/sideBar";
import Header from "../modules/header";
import "../styles/scraper.css";

import type Error from "../types/error";
import type BookT from "../types/book";

// --- API функции ---

const scrape = async (link: string) => {
  if (!link || link.length === 0) throw new Error("Введите ссылку");
  if (!URL.canParse(link)) throw new Error("Некорректный URL");

  const url = new URL(link);
  let service: string;

  switch (url.hostname.toLowerCase()) {
    case "www.labirint.ru":
    case "labirint.ru":
      service = "labirint";
      break;
    case "www.litres.ru":
    case "litres.ru":
      service = "litres";
      break;
    case "www.mybook.ru":
    case "mybook.ru":
      service = "mybook";
      break;
    default:
      throw new Error("Сервис не поддерживается");
  }

  const response = await fetch(`/scraper/${service}`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ url: url.toString() }),
  });

  const result = await response.json();
  if (!response.ok)
    throw new Error((result as Error).error || "Ошибка скрапинга");
  return result as BookT;
};

const postScrapedBook = async (book: Partial<BookT>) => {
  const response = await fetch(`/api/books`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(book),
  });

  const responseText = await response.text();

  if (!response.ok) {
    try {
      const json = JSON.parse(responseText);
      throw new Error((json as Error).error || "Ошибка сохранения");
    } catch {
      throw new Error(responseText || `Ошибка ${response.status}`);
    }
  }

  try {
    return JSON.parse(responseText) as BookT;
  } catch {
    return { id: "unknown", title: "", author: "", genre: "" } as BookT;
  }
};

const updateCoverImage = async (
  bookID: string,
  coverID: string | undefined,
) => {
  if (!coverID || !bookID) return;

  try {
    const response = await fetch(`/api/books/${bookID}/cover`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ imageuuid: coverID }),
    });

    if (!response.ok) {
      const result = await response.json().catch(() => ({}));
      console.warn("⚠️ Не удалось обновить обложку:", (result as Error).error);
    }
  } catch (err) {
    console.warn("⚠️ Ошибка при обновлении обложки:", err);
  }
};

// --- Компонент страницы ---

type Step = "idle" | "scraping" | "preview" | "saving" | "success" | "error";

// ✅ Компонент НЕ принимает пропы — поэтому ошибки TS не будет
const ScraperPage = () => {
  const [url, setUrl] = useState("");
  const [step, setStep] = useState<Step>("idle");
  const [error, setError] = useState<string | null>(null);
  const [scrapedBook, setScrapedBook] = useState<BookT | null>(null);
  const [savedBook, setSavedBook] = useState<BookT | null>(null);

  const [editableBook, setEditableBook] = useState<{
    title: string;
    author: string;
    genre: string;
    description?: string;
  }>({ title: "", author: "", genre: "", description: "" });

  const textareaRef = useRef<HTMLTextAreaElement>(null);

  const isScraping = step === "scraping";
  const isSaving = step === "saving";
  const isPreview = step === "preview";
  const isSuccess = step === "success";
  const isError = step === "error";
  const isIdle = step === "idle";
  const canReset = isPreview || isSuccess || isError;
  const isFormValid =
    editableBook.title.trim().length >= 2 &&
    editableBook.author.trim().length >= 2;

  const handleScrape = async () => {
    const inputUrl = url.trim();
    if (!inputUrl) {
      setError("Введите ссылку на книгу");
      return;
    }

    setStep("scraping");
    setError(null);
    setScrapedBook(null);

    try {
      console.log("🔍 Scraping:", inputUrl);
      const book = await scrape(inputUrl);

      console.log("✅ Scraped book:", book);

      if (!book || (!book.title && !book.author)) {
        throw new Error("Пустой ответ от скрапера");
      }

      setScrapedBook(book);
      setEditableBook({
        title: book.title?.trim() || "",
        author: book.author?.trim() || "",
        genre: book.genre?.trim() || "",
        description: book.description || "",
      });
      setStep("preview");
    } catch (err) {
      console.error("❌ Scrape error:", err);
      const message = err instanceof Error ? err.message : "Неизвестная ошибка";
      setError(`Не удалось получить данные: ${message}`);
      setStep("error");
    }
  };

  const handleSave = async () => {
    if (!scrapedBook) {
      setError("Нет данных для сохранения");
      setStep("error");
      return;
    }

    if (!isFormValid) {
      setError("Заполните название и автора (мин. 2 символа)");
      return;
    }

    setStep("saving");
    setError(null);

    try {
      const bookToSave: Partial<BookT> = {
        id: scrapedBook.id,
        image_uuid: scrapedBook.image_uuid,
        isbn: scrapedBook.isbn,
        title: editableBook.title.trim(),
        author: editableBook.author.trim(),
        genre: editableBook.genre.trim() || undefined,
        description: editableBook.description?.trim() || undefined,
      };

      console.log("💾 Saving book:", bookToSave);

      const saved = await postScrapedBook(bookToSave);
      console.log("✅ Saved:", saved);

      setSavedBook(saved);

      if (scrapedBook.image_uuid && saved.id) {
        await updateCoverImage(saved.id, scrapedBook.image_uuid);
      }

      setStep("success");
    } catch (err) {
      console.error("❌ Save error:", err);
      const message = err instanceof Error ? err.message : "Ошибка сохранения";
      setError(message);
      setStep("error");
    }
  };

  const handleReset = () => {
    setUrl("");
    setScrapedBook(null);
    setSavedBook(null);
    setError(null);
    setEditableBook({ title: "", author: "", genre: "", description: "" });
    setStep("idle");
    setTimeout(() => textareaRef.current?.focus(), 100);
  };

  return (
    <section id="scraper-body">
      <SideBar />
      <main id="scraper-main">
        <Header />
        <div id="scraper-layout">
          <section id="scraper-header-section">
            <h1>Добавить книгу</h1>
            <p className="subtitle">
              Вставьте ссылку на книгу с поддерживаемых сайтов:
              <strong> Labirint.ru, Litres.ru, MyBook.ru</strong>
            </p>
          </section>

          <div className="scraper-input-section">
            <textarea
              ref={textareaRef}
              className="scraper-url-input"
              placeholder="https://www.labirint.ru/books/..."
              value={url}
              onChange={(e) => setUrl(e.target.value)}
              disabled={isScraping || isSaving}
              onKeyDown={(e) => {
                if (e.ctrlKey && e.key === "Enter") handleScrape();
              }}
            />
            <div className="scraper-actions">
              <button
                className="btn-primary"
                onClick={handleScrape}
                disabled={isScraping || isSaving || !url.trim()}
              >
                {isScraping ? (
                  <>
                    <span className="spinner"></span>
                    Сканируем...
                  </>
                ) : (
                  "Найти книгу"
                )}
              </button>
              {canReset && (
                <button className="btn-secondary" onClick={handleReset}>
                  Сбросить
                </button>
              )}
            </div>
            <p className="hint-short">
              💡 Нажмите Ctrl+Enter для быстрого поиска
            </p>
          </div>

          {error && (
            <div className="alert alert-error">
              <strong>Ошибка:</strong> {error}
              <button className="alert-close" onClick={() => setError(null)}>
                ×
              </button>
            </div>
          )}

          {isSuccess && savedBook && (
            <div className="alert alert-success">
              ✅ Книга успешно добавлена!
              <br />
              <small>ID: {savedBook.id}</small>
              <div style={{ marginTop: "12px" }}>
                <a
                  href={`/book/${savedBook.id}`}
                  className="btn-primary"
                  style={{ padding: "8px 16px", fontSize: "0.9rem" }}
                >
                  Открыть книгу →
                </a>
              </div>
            </div>
          )}

          {isPreview && scrapedBook && (
            <div className="scraper-preview">
              <h3>Предпросмотр</h3>
              <div className="preview-grid">
                <div className="preview-cover">
                  {scrapedBook.image_uuid ? (
                    <img
                      src={`http://localhost/cdn/${scrapedBook.image_uuid}`}
                      alt="Обложка"
                      onError={(e) => {
                        (e.target as HTMLImageElement).src =
                          "https://via.placeholder.com/300x450?text=No+Cover";
                      }}
                    />
                  ) : (
                    <div className="cover-placeholder"></div>
                  )}
                </div>
                <div className="preview-form">
                  <div className="form-group">
                    <label>Название *</label>
                    <input
                      type="text"
                      value={editableBook.title}
                      onChange={(e) =>
                        setEditableBook({
                          ...editableBook,
                          title: e.target.value,
                        })
                      }
                      placeholder="Введите название"
                    />
                    {editableBook.title.trim().length > 0 &&
                      editableBook.title.trim().length < 2 && (
                        <span className="form-hint">Минимум 2 символа</span>
                      )}
                  </div>
                  <div className="form-group">
                    <label>Автор *</label>
                    <input
                      type="text"
                      value={editableBook.author}
                      onChange={(e) =>
                        setEditableBook({
                          ...editableBook,
                          author: e.target.value,
                        })
                      }
                      placeholder="Введите автора"
                    />
                    {editableBook.author.trim().length > 0 &&
                      editableBook.author.trim().length < 2 && (
                        <span className="form-hint">Минимум 2 символа</span>
                      )}
                  </div>
                  <div className="form-group">
                    <label>Жанр</label>
                    <input
                      type="text"
                      value={editableBook.genre}
                      onChange={(e) =>
                        setEditableBook({
                          ...editableBook,
                          genre: e.target.value,
                        })
                      }
                      placeholder="Например: Фантастика"
                      list="genre-suggestions"
                    />
                    <datalist id="genre-suggestions">
                      <option value="Фантастика" />
                      <option value="Детективы" />
                      <option value="Романы" />
                      <option value="История" />
                      <option value="Документальное" />
                    </datalist>
                  </div>
                  {editableBook.description && (
                    <div className="form-group">
                      <label>Описание</label>
                      <textarea
                        className="form-textarea"
                        value={editableBook.description}
                        onChange={(e) =>
                          setEditableBook({
                            ...editableBook,
                            description: e.target.value,
                          })
                        }
                        rows={3}
                        placeholder="Краткое описание"
                      />
                    </div>
                  )}
                  {scrapedBook.isbn && (
                    <div className="form-group">
                      <label>ISBN</label>
                      <input
                        type="text"
                        value={scrapedBook.isbn}
                        disabled
                        className="form-readonly"
                      />
                    </div>
                  )}
                  <button
                    className="btn-primary btn-save"
                    onClick={handleSave}
                    disabled={isSaving || !isFormValid}
                  >
                    {isSaving ? (
                      <>
                        <span className="spinner"></span>
                        Сохраняем...
                      </>
                    ) : (
                      "Сохранить в каталог"
                    )}
                  </button>
                </div>
              </div>
            </div>
          )}

          {isIdle && (
            <div className="scraper-hints">
              <h4>Как это работает?</h4>
              <ol>
                <li>Вставьте ссылку на книгу</li>
                <li>Нажмите "Найти книгу" — мы автоматически считаем данные</li>
                <li>Проверьте и отредактируйте информацию при необходимости</li>
                <li>Нажмите "Сохранить" — книга появится в каталоге</li>
              </ol>
              <p className="hint-note">
                💡 Поддерживаются прямые ссылки на страницы книг. Если скрапинг
                не сработал — проверьте, что ссылка ведёт именно на книгу, а не
                на поиск или категорию.
              </p>
            </div>
          )}
        </div>
      </main>
    </section>
  );
};

export default ScraperPage;
