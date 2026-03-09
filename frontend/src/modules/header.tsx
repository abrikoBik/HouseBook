import NotificationsIcon from "../assets/Notification.svg";
import "../styles/book.css";

const Header = () => {
  return (
    <header id="top_menu">
      <div id="top_actions">
        <div id="profile">
          <div></div>
          <div id="name">Шиловский Максим</div>
        </div>
        <div id="notifications">
          <img src={NotificationsIcon} alt="notifications" />
        </div>
      </div>
    </header>
  );
};

export default Header;
