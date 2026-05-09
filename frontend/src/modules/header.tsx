import NotificationsIcon from "../assets/Notification.svg";
import SearchIcon from "../assets/Lupa.svg";
import "../styles/header.css";

const Header = () => {
  return (
    <header id="top_menu">
      <div id="top_actions">
        <div id="searchBar">
          <label htmlFor="searchBarInput">
            <img src={SearchIcon} alt="SearchIcon" />
          </label>
          <input
            id="searchBarInput"
            type="text"
            placeholder="Поиск по названия книги, автору, изданию"
          />
        </div>
        <div style={{ display: "flex", alignItems: "center", gap: "20px" }}>
          <a href="/profile" className="profile-link">
            <div id="icon"></div>
            <div id="name">Шиловский Максим</div>
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