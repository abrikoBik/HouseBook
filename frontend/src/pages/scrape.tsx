import Header from "../modules/header";
import Scraper from "../modules/scraper";
import SideBar from "../modules/sideBar";

const Scrape = () => {
  return (
    <section id="layout">
      <Header />
      <SideBar />
      <section id="scraper">
        <h1>Очень временная страница для скраппинга</h1>
        <Scraper url={""} />
      </section>
    </section>
  );
};

export default Scrape;
