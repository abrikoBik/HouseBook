import { useEffect, useState, useCallback } from "react";
import Header from "../modules/header";
import SideBar from "../modules/sideBar";
import "../styles/library.css";

interface LibraryType {
  id: string;
  name: string;
  user_id: string;
}

interface LibBook {
  id: string;
  exchange: boolean;
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
}

const Library = () => {
  const [activeTab, setActiveTab] = useState(0);
  
  const [userLibs, setUserLibs] = useState<LibraryType[]>([]);
  const [displayBooks, setDisplayBooks] = useState<DisplayBook[]>([]);
  const [isLoading, setIsLoading] = useState(false);

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

  const fetchLibBookIds = async (libId: string): Promise<LibBook[]> => {
    const response = await fetch(`/api/libraries/${libId}/books`);
    if (!response.ok) return [];
    return await response.json();
  };

  const fetchBooksForLibrary = useCallback(async (libId: string) => {
    setIsLoading(true);
    setDisplayBooks([]);

    try {
      const libBooksList = await fetchLibBookIds(libId);

      if (libBooksList.length === 0) {
        setIsLoading(false);
        return;
      }

      const booksDetailsPromises = libBooksList.map(async (item: LibBook) => {
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
          };
        } catch (err) {
          console.warn(`Failed to load book ${item.id}`, err);
          return null;
        }
      });

      const results = await Promise.all(booksDetailsPromises);
      
      const validBooks = results.filter((b): b is DisplayBook => b !== null);
      setDisplayBooks(validBooks);

    } catch (error) {
      console.error("Error loading books for library:", error);
    } finally {
      setIsLoading(false);
    }
  }, []);

  useEffect(() => {
    fetchLibraries();
  }, [fetchLibraries]);

  useEffect(() => {
    if (activeTab === 0) {
      // TODO: fetch all books
      setDisplayBooks([]); 
      setIsLoading(false);
    } else {
      const selectedLibIndex = activeTab - 1;
      if (userLibs[selectedLibIndex]) {
        fetchBooksForLibrary(userLibs[selectedLibIndex].id);
      }
    }
  }, [activeTab, userLibs, fetchBooksForLibrary]);

  return (
    <div>
      <SideBar />
      <Header />

      <section id="libraryLayout">
        <section id="info">
          <h4>Моя библиотека</h4>
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
            <div style={{ padding: "20px", textAlign: "center", width: "100%" }}>
              Загрузка книг...
            </div>
          ) : displayBooks.length > 0 ? (
            displayBooks.map((book: DisplayBook) => (
              <div className="libBookCard" key={book.id}>
                <div className="libBookCover">
                  {book.cover_url ? (
                    <img src={book.cover_url} alt={book.title} loading="lazy" />
                  ) : (
                    <div className="cover-placeholder"></div>
                  )}
                </div>
                <div className="libBookTitle">{book.title}</div>
                <div className="libBookAuthor">{book.author}</div>
              </div>
            ))
          ) : (
            !isLoading && (
              <div style={{ padding: "20px", textAlign: "center", width: "85vw", color: "#666" }}>
                {activeTab === 0 
                  ? "Выберите библиотеку или добавьте книги." 
                  : "В этой библиотеке пока нет книг."}
              </div>
            )
          )}
        </section>
      </section>
    </div>
  );
};

export default Library;