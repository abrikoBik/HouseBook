import type Error from "../types/error";
import type BookT from "../types/book";
import { useEffect, useRef, useState } from "react";

const scrape = async (link: string) => {
  if (link.length == 0) return;
  if (!URL.canParse(link)) {
    throw new Error("invalid url");
  }
  const url = new URL(link);
  let service: string;
  switch (url.hostname.toLowerCase()) {
    case "www.labirint.ru":
    case "labirint.ru":
      service = "labirint";
      break;
    case "www.litres.ru":
    case "litres.ru":
      service = "litres";
      break;
    case "www.mybook.ru":
    case "mybook.ru":
      service = "mybook";
      break;
    default:
      throw new Error("unknown service");
  }
  const response = await fetch(`/scraper/${service}`, {
    method: "POST",
    body: JSON.stringify({ url: url.toString() }),
  });
  const result = await response.json();
  if (!response.ok) {
    throw new Error((result as Error).error);
  }
  return result as BookT;
};

const postScrapedBook = async (book: BookT) => {
  const response = await fetch(`/api/books`, {
    method: "POST",
    body: JSON.stringify(book),
  });
  const result = await response.json();
  if (!response.ok) {
    throw new Error((result as Error).error);
  }
  return result as BookT;
};

const updateCoverImage = async (
  bookID: string,
  coverID: string | undefined,
) => {
  if (coverID === undefined) return;
  const response = await fetch(`/api/books/${bookID}/cover`, {
    method: "PUT",
    body: JSON.stringify({ imageuuid: coverID }),
  });
  const result = await response.json();
  if (!response.ok) {
    throw new Error((result as Error).error);
  }
  return result as BookT;
};

const Scraper = ({ url }: { url: string }) => {
  const textareaRef = useRef<HTMLTextAreaElement>(null);
  const handleScrapeSubmit = async () => {
    const url = textareaRef.current?.value;
    // TODO: try-catch
    try {
      const book = await scrape(url!);
      console.log(book);
      alert("👍 (книжка есть)");
      const apiResponse = await postScrapedBook(book!);
      console.log(apiResponse);
      alert(`👍 (книжка теперь на сервере, ${apiResponse?.id})`);
      const coverResponse = await updateCoverImage(
        apiResponse?.id,
        book?.image_uuid,
      );
      alert(`👍 (обложка теперь на сервере, ${coverResponse?.image_uuid})`);
    } catch (err: unknown) {
      alert(`👎 (${err})`);
      console.error(err as Error);
    }
  };

  const [book, setBook] = useState<BookT>();
  console.log(book, setBook);
  useEffect(() => {
    const scrapeBook = async () => {
      // const name = await scrape(url);
      // setBook(name);
    };
    scrapeBook();
  }, [url]);

  return (
    <div id="ScrapeArea">
      <textarea ref={textareaRef} placeholder="(ссылка на книгу)" />
      <button onClick={handleScrapeSubmit}>Скраппинг</button>
    </div>
  );
};

export default Scraper;
