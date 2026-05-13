import { useState, useEffect, useRef } from "react";
import NotificationsIcon from "../assets/Notification.svg";
import SearchIcon from "../assets/Lupa.svg";
import "../styles/header.css";
import { jwtDecode } from "jwt-decode";
import { useNavigate } from "react-router";

interface Book {
  id: string;
  title: string;
  author: string;
  cover_url?: string;
}

interface ApiBook {
  id: string;
  title?: string | null;
  author?: string | null;
  cover_url?: string | null;
  image_uuid?: string | null;
}

interface DecodedToken {
  id: string;
  exp?: number;
  iat?: number;
}

// interface Username {
//   name: string;
// }

const Header = () => {
  const navigate = useNavigate();
  const [query, setQuery] = useState("");
  const [allBooks, setAllBooks] = useState<Book[]>([]);
  const [results, setResults] = useState<Book[]>([]);
  const [showDropdown, setShowDropdown] = useState(false);
  const searchRef = useRef<HTMLDivElement>(null);
  const [username, setUsername] = useState<string>();

  useEffect(() => {
    const getUsername = async (id: string) => {
      const response = await fetch(`/api/users/${id}`);
      const respJSON = await response.json();

      if (!response.ok) {
        navigate("/auth");
      } else {
        setUsername(respJSON["name"]);
      }
    };

    const token = localStorage.getItem("token");
    if (token) {
      const decoded = jwtDecode<DecodedToken>(token!);
      localStorage.setItem("UserID", decoded.id);
      getUsername(decoded.id);
    } else {
      navigate("/auth");
    }
  }, [navigate]);

  useEffect(() => {
    const fetchCatalog = async () => {
      try {
        const res = await fetch("/api/books");
        if (!res.ok) throw new Error("Failed to fetch catalog");
        const data: ApiBook[] = await res.json();

        // 🔁 Единая логика формирования URL обложки (как в Catalog.tsx)
        const mapped: Book[] = data.map((b) => ({
          id: b.id,
          title: b.title?.trim() || "Без названия",
          author: b.author?.trim() || "Неизвестный автор",
          cover_url: b.image_uuid
            ? `http://localhost/cdn/${b.image_uuid}` // ← Изменено на /cdn/
            : `https://via.placeholder.com/300x450?text=${encodeURIComponent(b.title || "Book")}`, // ← Заглушка как в каталоге
        }));

        setAllBooks(mapped);
      } catch (err) {
        console.error("Catalog fetch error:", err);
      }
    };
    fetchCatalog();
  }, []);

  useEffect(() => {
    if (query.length < 2) {
      setResults([]);
      setShowDropdown(false);
      return;
    }

    const timer = setTimeout(() => {
      const q = query.toLowerCase();
      const filtered = allBooks
        .filter(
          (b) =>
            b.title.toLowerCase().includes(q) ||
            b.author.toLowerCase().includes(q),
        )
        .slice(0, 3);

      setResults(filtered);
      setShowDropdown(true);
    }, 300);

    return () => clearTimeout(timer);
  }, [query, allBooks]);

  useEffect(() => {
    const handleClickOutside = (e: MouseEvent) => {
      if (searchRef.current && !searchRef.current.contains(e.target as Node)) {
        setShowDropdown(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  // Обработчик ошибок загрузки изображений
  const handleImageError = (e: React.SyntheticEvent<HTMLImageElement>) => {
    e.currentTarget.src = `https://via.placeholder.com/300x450?text=${encodeURIComponent("No Cover")}`;
  };

  return (
    <header id="top_menu">
      <div id="top_actions">
        <div id="searchBar" ref={searchRef}>
          <label htmlFor="searchBarInput">
            <img src={SearchIcon} alt="SearchIcon" />
          </label>
          <input
            id="searchBarInput"
            type="text"
            placeholder="Поиск по названию книги, автору, изданию"
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            onFocus={() => query.length >= 2 && setShowDropdown(true)}
          />

          {showDropdown && (
            <div className="search-dropdown">
              {results.length > 0 ? (
                results.map((book) => (
                  <a
                    href={`/book/${book.id}`}
                    key={book.id}
                    className="search-item"
                    onClick={() => {
                      setShowDropdown(false);
                      setQuery("");
                    }}
                  >
                    <div className="search-item-text">
                      <div className="search-book-title">{book.title}</div>
                      <div className="search-book-author">{book.author}</div>
                    </div>
                    <div className="search-book-cover">
                      <img
                        src={book.cover_url}
                        alt={book.title}
                        onError={handleImageError}
                      />
                    </div>
                  </a>
                ))
              ) : (
                <div className="search-empty">Ничего не найдено</div>
              )}

              <a
                href="/book/scrape"
                className="search-suggestion"
                onClick={() => setShowDropdown(false)}
              >
                Не нашли книгу? <span>Добавить через скрапер →</span>
              </a>
            </div>
          )}
        </div>

        <div style={{ display: "flex", alignItems: "center", gap: "20px" }}>
          <a href="/profile" className="profile-link">
            <div id="icon"></div>
            <div id="name">{username}</div>
          </a>
          <div id="notifications">
            <img src={NotificationsIcon} alt="notifications" />
          </div>
        </div>
      </div>
    </header>
  );
};

export default Header;
