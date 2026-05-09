import Header from "../modules/header";
import SideBar from "../modules/sideBar";
import "../styles/main.css";

const Main = () => {
  return (
    <section id="layout">
      <Header />
      <SideBar />
      <section id="status404">
        <h1>40004</h1>
        <p>Страница не была найдена =(</p>
        <a href="/">На главную</a>
      </section>
    </section>
  );
};

export default Main;
