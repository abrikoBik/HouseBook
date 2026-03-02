import NotificationsIcon from "../assets/Notifications.png";
import ProfileIcon from "../assets/ProfileIcon.png";
import "../styles/book.css";

const Header = () => {
  return (
    <header id="top_menu">
      <div id="top_actions">
        <div id="profile">
          <img id="img1" src={ProfileIcon} alt="profile" />
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
