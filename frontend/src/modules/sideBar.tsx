import HomeIcon from "../assets/Home_Black.svg";
import LogoIcon from "../assets/LogoWOText.svg";
import FavoritesIcon from "../assets/Favourite_Black.svg";
import LibraryIcon from "../assets/Library_Black.svg";
import ExitIcon from "../assets/Exit_Black.svg";
import "../styles/sideBar.css";

const sideBar = () => {
  return (
    <aside id="side_menu">
      <nav>
        <div className="menu_item">
          <img src={LogoIcon} alt="logo" />
        </div>

        <div className="menu_item">
          <img src={HomeIcon} alt="home" />
        </div>

        <div className="menu_item">
          <img src={FavoritesIcon} alt="favorites" />
        </div>

        <div className="menu_item">
          <img src={ExitIcon} alt="exit" />
        </div>
      </nav>
    </aside>
  );
};

export default sideBar;
