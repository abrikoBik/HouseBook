import { Link, useLocation, useNavigate } from "react-router-dom"; // Добавили useNavigate
import HomeBlack from "../assets/Home_Black.svg";
import HomeWhite from "../assets/Home_White.svg";
import LogoIcon from "../assets/LogoWOText.svg";
import LibraryBlack from "../assets/Library_Black.svg";
import LibraryWhite from "../assets/Library_White.svg";
import ExitIcon from "../assets/Exit_Black.svg";
import "../styles/sideBar.css";

const SideBar = () => {
  const location = useLocation();
  const navigate = useNavigate(); // Инициализируем навигатор
  const currentPath = location.pathname;

  const isCentralActive =
    currentPath === "/central" || currentPath.startsWith("/main/");
  const isLibraryActive =
    currentPath === "/library" || currentPath.startsWith("/library/");

  // Функция для очистки сессии и выхода
  const handleLogout = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("UserID");
    navigate("/"); // Перенаправление на лендинг
  };

  return (
    <aside id="side_menu">
      <nav id="side_nav">
        <a href="/landing" className="logo">
          <img src={LogoIcon} alt="logo" />
        </a>

        <section id="menu_items">
          <div className={`menu-item ${isCentralActive ? "active" : ""}`}>
            <Link to="/central">
              <img src={isCentralActive ? HomeWhite : HomeBlack} alt="home" />
            </Link>
          </div>
          <div className={`menu-item ${isLibraryActive ? "active" : ""}`}>
            <Link to="/library">
              <img
                src={isLibraryActive ? LibraryWhite : LibraryBlack}
                alt="library"
              />
            </Link>
          </div>
        </section>

        {/* НАЧАЛО ИЗМЕНЕНИЙ: Добавлен onClick обработчик */}
        <div id="exit_button" onClick={handleLogout}>
          <img src={ExitIcon} alt="exit" />
        </div>
        {/* КОНЕЦ ИЗМЕНЕНИЙ */}
      </nav>
    </aside>
  );
};

export default SideBar;
