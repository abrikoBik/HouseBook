import { useEffect, useState, useCallback } from "react";
import Header from "../modules/header";
import SideBar from "../modules/sideBar";
import BookModal from "../modules/bookModal";
import AddBookTab from "../modules/addBookTab";
import "../styles/library.css";

interface LibraryType {
  id: string;
  name: string;
  user_id: string;
}

interface LibBook {
  id: string;
  exchangeable: boolean;
}

interface ApiBook {
  id: string;
  title: string | null;
  author: string | null;
  image_uuid: string | null;
}

interface DisplayBook {
  id: string;
  title: string;
  author: string;
  cover_url: string;
  exchangeable: boolean;
}

const Library = () => {
  const [activeTab, setActiveTab] = useState(0);
  const [addBookActive, setAddBookActive] = useState(false); // Изначально false — модалка закрыта

  const [userLibs, setUserLibs] = useState<LibraryType[]>([]);
  const [displayBooks, setDisplayBooks] = useState<DisplayBook[]>([]);
  const [selectedBook, setSelectedBook] = useState<DisplayBook | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  const handleBookUpdate = (updatedBook: DisplayBook) => {
    setDisplayBooks((prev) =>
      prev.map((b) => (b.id === updatedBook.id ? updatedBook : b)),
    );
    setSelectedBook(updatedBook);
  };

  const fetchLibraries = useCallback(async () => {
    const id = localStorage.getItem("UserID");
    if (!id) return;

    try {
      const response = await fetch(`/api/users/${id}/libraries`);
      if (response.ok) {
        const data: LibraryType[] = await response.json();
        setUserLibs(data);
      } else {
        console.error("Failed to fetch libraries");
      }
    } catch (error) {
      console.error("Error fetching libraries:", error);
    }
  }, []);

  const fetchBookDetails = async (
    item: LibBook,
  ): Promise<DisplayBook | null> => {
    try {
      const res = await fetch(`/api/books/${item.id}`);
      if (!res.ok) return null;

      const bookData: ApiBook = await res.json();

      return {
        id: bookData.id,
        title: bookData.title?.trim() || "Без названия",
        author: bookData.author?.trim() || "Неизвестный автор",
        cover_url: bookData.image_uuid
          ? `http://localhost/cdn/${bookData.image_uuid}`
          : `https://via.placeholder.com/300x450?text=No+Cover`,
        exchangeable: item.exchangeable,
      };
    } catch (err) {
      console.warn(`Failed to load book ${item.id}`, err);
      return null;
    }
  };

  const fetchLibBookIds = async (libId: string): Promise<LibBook[]> => {
    const response = await fetch(`/api/libraries/${libId}/books`);
    if (!response.ok) return [];
    return await response.json();
  };

  const getBooksForLibrary = async (libId: string): Promise<DisplayBook[]> => {
    const libBooksList = await fetchLibBookIds(libId);
    if (libBooksList.length === 0) return [];

    const booksDetailsPromises = libBooksList.map((item: LibBook) =>
      fetchBookDetails(item),
    );
    const results = await Promise.all(booksDetailsPromises);

    return results.filter((b): b is DisplayBook => b !== null);
  };

  const fetchBooksForLibrary = useCallback(async (libId: string) => {
    setIsLoading(true);
    setDisplayBooks([]);

    try {
      const validBooks = await getBooksForLibrary(libId);
      setDisplayBooks(validBooks);
    } catch (error) {
      console.error("Error loading books for library:", error);
    } finally {
      setIsLoading(false);
    }
  }, []);

  const handleAddBookClose = () => {
    setAddBookActive(false);
    // Обновляем список книг после добавления
    if (activeTab === 0) {
      fetchAllBooks();
    } else {
      const selectedLibIndex = activeTab - 1;
      if (userLibs[selectedLibIndex]) {
        fetchBooksForLibrary(userLibs[selectedLibIndex].id);
      }
    }
  };

  const fetchAllBooks = useCallback(async () => {
    if (userLibs.length === 0) {
      setDisplayBooks([]);
      setIsLoading(false);
      return;
    }

    setIsLoading(true);
    setDisplayBooks([]);

    try {
      const allLibsBooksPromises = userLibs.map((lib) =>
        fetchLibBookIds(lib.id),
      );
      const allLibsBooksArrays = await Promise.all(allLibsBooksPromises);
      const flatLibBooksList: LibBook[] = allLibsBooksArrays.flat();

      if (flatLibBooksList.length === 0) {
        setDisplayBooks([]);
        return;
      }

      const booksDetailsPromises = flatLibBooksList.map((item: LibBook) =>
        fetchBookDetails(item),
      );
      const results = await Promise.all(booksDetailsPromises);

      const validBooks = results.filter((b): b is DisplayBook => b !== null);
      setDisplayBooks(validBooks);
    } catch (error) {
      console.error("Error loading all books:", error);
    } finally {
      setIsLoading(false);
    }
  }, [userLibs]);

  useEffect(() => {
    fetchLibraries();
  }, [fetchLibraries]);

  useEffect(() => {
    setIsLoading(true);

    if (activeTab === 0) {
      fetchAllBooks();
    } else {
      const selectedLibIndex = activeTab - 1;
      if (userLibs[selectedLibIndex]) {
        fetchBooksForLibrary(userLibs[selectedLibIndex].id);
      } else {
        setIsLoading(false);
        setDisplayBooks([]);
      }
    }
  }, [activeTab, userLibs, fetchBooksForLibrary, fetchAllBooks]);

  return (
    <div>
      <SideBar />
      <Header />

      {addBookActive && <AddBookTab onClose={handleAddBookClose} />}

      <section id="libraryLayout">
        <section id="info">
          <h4>Моя библиотека</h4>

          {/* 🔘 Кнопка добавления книги */}
          <button
            className="add-book-btn"
            onClick={() => setAddBookActive(true)}
            title="Добавить книгу"
          >
            <span className="material-symbols-outlined">add</span>
            <span>Добавить книгу</span>
          </button>

          <nav>
            <ul>
              <li
                className={activeTab === 0 ? "active" : ""}
                onClick={() => setActiveTab(0)}
              >
                Все
              </li>

              {userLibs.map((lib, index) => (
                <li
                  key={lib.id}
                  className={activeTab === index + 1 ? "active" : ""}
                  onClick={() => setActiveTab(index + 1)}
                >
                  {lib.name}
                </li>
              ))}
            </ul>
          </nav>
        </section>

        <section id="books">
          {isLoading ? (
            <div
              style={{ padding: "20px", textAlign: "center", width: "100%" }}
            >
              Загрузка книг...
            </div>
          ) : displayBooks.length > 0 ? (
            displayBooks.map((book: DisplayBook) => (
              <div
                className="libBookCard"
                key={book.id}
                onClick={() => setSelectedBook(book)}
              >
                <div className="libBookCover">
                  {book.cover_url ? (
                    <img src={book.cover_url} alt={book.title} loading="lazy" />
                  ) : (
                    <div className="cover-placeholder"></div>
                  )}
                </div>
                <div className="libBookTitle">{book.title}</div>
                <div className="libBookAuthor">{book.author}</div>
                <div
                  id="exchangableSign"
                  className={`${book.exchangeable ? "exchangeAvaible" : "exchangeNotAvaible"}`}
                >
                  <span className="material-symbols-outlined">swap_horiz</span>
                  <div style={{ fontWeight: "bold" }}>
                    {book.exchangeable ? "Доступна" : "Недоступна"}
                  </div>
                </div>
              </div>
            ))
          ) : (
            !isLoading && (
              <div
                style={{
                  padding: "20px",
                  textAlign: "center",
                  width: "85vw",
                  color: "#666",
                }}
              >
                {activeTab === 0
                  ? "У вас пока нет книг ни в одной библиотеке."
                  : "В этой библиотеке пока нет книг."}
              </div>
            )
          )}
        </section>
      </section>

      {selectedBook && (
        <BookModal
          book={selectedBook}
          library_id={
            userLibs.length > 0
              ? activeTab === 0
                ? userLibs[0].id
                : userLibs[activeTab - 1].id
              : ""
          }
          onClose={() => setSelectedBook(null)}
          onStatusUpdate={handleBookUpdate}
        />
      )}
    </div>
  );
};

export default Library;
