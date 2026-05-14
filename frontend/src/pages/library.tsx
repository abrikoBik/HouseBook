import { useEffect, useState } from "react";
import Header from "../modules/header";
import SideBar from "../modules/sideBar";
import "../styles/library.css";

interface LibraryType {
  id: string;
  name: string;
  user_id: string;
}

interface LibBook {
  bookID: string;
  exchange: boolean;
}

const Library = () => {
  const [activeTab, setActiveTab] = useState(0);

  // TODO: Make one big array which will store all libs
  const [inProcessLib, setInProcessLib] = useState<LibBook[]>([]);
  const [readyForExchangeLib, setReadyForExchangeLib] = useState<LibBook[]>([]);
  const [readLib, setReadLib] = useState<LibBook[]>([]);

  const fetchLibraries = async () => {
    const id = localStorage.getItem("UserID");
    const user_libs_resp = await fetch(`/api/users/${id}/libraries`);
    const user_libs: LibraryType[] = await user_libs_resp.json();
    if (user_libs_resp.ok) {
      console.log(user_libs);

      const return_libs: string[] = [];
      user_libs.forEach((el) => {
        return_libs.push(el.id);
      });

      return return_libs;
    }
  };

  const fetchLib = async (id: string) => {
    const lib_books_resp = await fetch(`/api/libraries/${id}/books`);
    const lib_books = await lib_books_resp.json();
    return lib_books;
  };

  useEffect(() => {
    const loadData = async () => {
      const libs = await fetchLibraries();

      if (libs && libs.length >= 3) {
        const [readData, inProcessData, exchangeData] = await Promise.all([
          fetchLib(libs[0]),
          fetchLib(libs[1]),
          fetchLib(libs[2]),
        ]);

        if (readData[0].id == null) {
          setReadLib([]);
        } else {
          setReadLib(readData || []);
        }

        if (inProcessData[0].id == null) {
          setInProcessLib([]);
        } else {
          setInProcessLib(inProcessData || []);
        }

        if (exchangeData[0].id == null) {
          setReadyForExchangeLib([]);
        } else {
          setReadyForExchangeLib(exchangeData || []);
        }

        console.log(readData);
        console.log(inProcessData);
        console.log(exchangeData);
      }
    };

    loadData();
  }, []);

  return (
    <div>
      <SideBar />
      <Header />

      <section id="libraryLayout">
        <section id="info">
          <h4>Моя библиотека</h4>
          <nav>
            <ul>
              <li
                className={activeTab === 0 ? "active" : ""}
                onClick={() => setActiveTab(0)}
              >
                {`Все (${(readLib?.length || 0) + (inProcessLib!.length || 0) + (readyForExchangeLib!.length || 0)})`}
              </li>
              <li
                className={activeTab === 1 ? "active" : ""}
                onClick={() => setActiveTab(1)}
              >
                {`Прочитанные (${readLib?.length || 0})`}
              </li>
              <li
                className={activeTab === 2 ? "active" : ""}
                onClick={() => setActiveTab(2)}
              >
                {`В процессе (${inProcessLib!.length || 0})`}
              </li>
              <li
                className={activeTab === 3 ? "active" : ""}
                onClick={() => setActiveTab(3)}
              >
                {`Доступны для обмена (${readyForExchangeLib!.length || 0})`}
              </li>
            </ul>
          </nav>
        </section>
        <section id="books">
          <div>Книга 1</div>
        </section>
      </section>
    </div>
  );
};

export default Library;
