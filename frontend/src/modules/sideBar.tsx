import { Link, useLocation } from "react-router-dom";
import HomeBlack from "../assets/Home_Black.svg";
import HomeWhite from "../assets/Home_White.svg";
import LogoIcon from "../assets/LogoWOText.svg";
import FavoritesBlack from "../assets/Favourite_Black.svg";
import FavoritesWhite from "../assets/Favourite_White.svg";
import LibraryBlack from "../assets/Library_Black.svg";
import LibraryWhite from "../assets/Library_White.svg";
import ExitIcon from "../assets/Exit_Black.svg";
import "../styles/sideBar.css";

const SideBar = () => {
  const location = useLocation();
  const currentPath = location.pathname;
  const isCentralActive =
    currentPath === "/central" || currentPath.startsWith("/main/");
  const isLibraryActive =
    currentPath === "/library" || currentPath.startsWith("/library/");
  const isExchangeActive =
    currentPath === "/profile/exchange" ||
    currentPath.startsWith("/profile/exchange/");

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
          <div className={`menu-item ${isExchangeActive ? "active" : ""}`}>
            <Link to="/profile/exchange">
              <img
                src={isExchangeActive ? FavoritesWhite : FavoritesBlack}
                alt="exchange"
              />
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
        <div id="exit_button">
          <img src={ExitIcon} alt="exit" />
        </div>
      </nav>
    </aside>
  );
};

export default SideBar;
