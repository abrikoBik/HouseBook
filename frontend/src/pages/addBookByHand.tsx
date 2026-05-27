import Header from "../modules/header";
import SideBar from "../modules/sideBar";
import "../styles/addBookByHand.css";
import Placeholder from "../assets/placeholderBook.png";
import BgBook from "../assets/bgBookAdd.png";
import { useState } from "react";

interface BookT {
  author: string;
  cover_type?: string;
  description?: string;
  genre: string;
  id: string;
  isbn: string;
  language?: string;
  page_count?: number;
  publisher: string;
  title: string;
  year?: number;
  image_uuid?: string;
  link?: string;
}

const AddBookByHand = () => {
  const [bookForm, setBookForm] = useState<BookT>({
    author: "",
    cover_type: "",
    description: "",
    genre: "",
    id: "",
    isbn: "",
    language: "",
    page_count: 0,
    publisher: "",
    title: "",
    year: 0,
    image_uuid: "",
    link: "",
  });

  const handleFormChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) => {
    const value = e.target.value;
    const target = e.target.id;
    setBookForm({
      ...bookForm,
      [target]:
        target === "page_count" || target === "year" ? parseInt(value) : value,
    });
  };

  const [imageID, setImageID] = useState<string>();

  const handleCover = async () => {
    const input = document.createElement("input");
    input.type = "file";
    input.accept = "image/*";

    const cover_file = await new Promise<File | null>((resolve) => {
      input.onchange = (e) => {
        const target = e.target as HTMLInputElement;
        resolve(target.files?.[0] || null);
      };

      input.oncancel = () => resolve(null);
      input.click();
    });

    if (cover_file) {
      const cover_data = new FormData();
      cover_data.append("image", cover_file);

      const response = await fetch("http://localhost/cdn/", {
        method: "POST",
        body: cover_data,
      });

      const cover_json = await response.json();

      setImageID(cover_json["id"]);

      console.log(`Обложка ${cover_json["id"]} была успешно загружена!`);
    } else {
      alert("Обложка не был загружен!");
    }
  };

  const sendForm = async () => {
    const response = await fetch("/api/books", {
      method: "POST",
      body: JSON.stringify(bookForm),
    });

    alert(`${bookForm.title} была добавлена!`);

    const book_json = await response.json();

    const put_cover = await fetch(`/api/books/${book_json["id"]}/cover`, {
      method: "PUT",
      body: `{"imageuuid":"${imageID}"}`,
    });

    if (put_cover.ok) {
      alert("Обложка была успешно добавлена!");
    } else {
      await fetch(`/api/books/${book_json["id"]}`, {
        method: "DELETE",
      });
    }
  };

  return (
    <section>
      <Header />
      <SideBar />
      <section id="manual-body">
        <img src={BgBook} alt="Decorative books" id="manual-visual-img" />
        <h2>Добавление новой книги</h2>
        <div id="manual-inputs">
          <div>
            <img
              src={`${imageID ? "/cdn/${imageID}" : Placeholder}`}
              id="manual-cover"
              alt="Book"
            />
            <button onClick={handleCover}>Изменить обложку</button>
          </div>
          <ul>
            <li key="manual-title-box">
              <input
                type="text"
                id="title"
                placeholder="Название"
                onChange={handleFormChange}
              />
            </li>
            <li key="manual-author-box">
              <input
                type="text"
                id="author"
                placeholder="Автор"
                onChange={handleFormChange}
              />
            </li>
            <li key="manual-year-box">
              <input
                type="text"
                id="year"
                placeholder="Год выпуска"
                onChange={handleFormChange}
              />
            </li>
            <li key="manual-genre-box">
              <input
                type="text"
                id="genre"
                placeholder="Жанр"
                onChange={handleFormChange}
              />
            </li>
            <li key="manual-isbn-box">
              <input
                type="text"
                id="isbn"
                placeholder="ISBN"
                onChange={handleFormChange}
              />
            </li>
            <li key="manual-publisher-box">
              <input
                type="text"
                id="publisher"
                placeholder="Издатель"
                onChange={handleFormChange}
              />
            </li>

            <li key="manual-description-box">
              <textarea
                id="description"
                placeholder="Описание"
                onChange={handleFormChange}
              />
            </li>
            <li key="manual-language-box">
              <input
                type="text"
                id="language"
                placeholder="Язык"
                onChange={handleFormChange}
              />
            </li>
            <li key="manual-cover-type-box">
              <input
                type="text"
                id="cover_type"
                placeholder="Тип обложки"
                onChange={handleFormChange}
              />
            </li>
            <li key="manual-page-count-box">
              <input
                type="text"
                id="page_count"
                placeholder="Колличество страниц"
                onChange={handleFormChange}
              />
            </li>
            <li key="manual-link-box">
              <input
                type="text"
                id="link"
                placeholder="Ссылка на книгу в другом источнике"
                onChange={handleFormChange}
              />
            </li>
          </ul>
        </div>
        <button onClick={sendForm}>Создать книгу</button>
      </section>
    </section>
  );
};

export default AddBookByHand;
