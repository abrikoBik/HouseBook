import HomeIcon from "../assets/Home.png";
import LogoIcon from "../assets/Logo.png";
import FavoritesIcon from "../assets/Favorites.png";
import ExitIcon from "../assets/Exit.png";

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
