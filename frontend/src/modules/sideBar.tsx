import HomeIcon from "../assets/Home_Black.svg";
import LogoIcon from "../assets/LogoWOText.svg";
import FavoritesIcon from "../assets/Favourite_Black.svg";
import LibraryIcon from "../assets/Library_Black.svg";
import ExitIcon from "../assets/Exit_Black.svg";
import "../styles/sideBar.css";

const SideBar = () => {
  return (
    <aside id="side_menu">
      <nav id="side_nav">
        <a href="/landing" className="logo">
          <img src={LogoIcon} alt="logo" />
        </a>

        <section id="menu_items">
          <div>
            <img src={HomeIcon} alt="home" />
          </div>

          <div>
            <img src={FavoritesIcon} alt="favorites" />
          </div>

          <div>
            <img src={LibraryIcon} alt="favorites" />
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