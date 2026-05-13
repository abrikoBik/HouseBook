import Header from "../modules/header";
import ScraperPage from "../modules/scraper";
import SideBar from "../modules/sideBar";

const Scrape = () => {
  return (
    <section id="layout">
      <Header />
      <SideBar />
      <section id="scraper">
        <ScraperPage />
      </section>
    </section>
  );
};

export default Scrape;
