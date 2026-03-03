import "../styles/404.css";
import Header from "../modules/header";
import SideBar from "../modules/sideBar";

const NotFound = () => {
  return (
    <section id="layout">
      <Header />
      <SideBar />
      <section id="status">
        <h1>404</h1>
        <p>Страница не была найдена =(</p>
        <a href="/">На главную</a>
      </section>
    </section>
  );
};

export default NotFound;
