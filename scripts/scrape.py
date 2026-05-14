from urllib.parse import urlparse
from requests import post, put
from multiprocessing import Process
from sys import argv

def determine_service(url: str) -> str:
    parsed = urlparse(url)
    match parsed.netloc.replace("www.", "", 1):
        case "litres.ru":
            return "litres"
        case "labirint.ru":
            return "labirint"
        case "mybook.ru":
            return "mybook"
        case _:
            raise ValueError("unknown service")

def scrape_and_upload(url: str):
    print(f"{url}: scraping")
    book_response = post(
        f"http://localhost/scraper/{determine_service(url)}",
        json = {"url": url},
    )
    book_json = book_response.json()
    if book_response.status_code == 200:
        print(f"{url}: scraped ({book_json.get("title")} by {book_json.get("author")})")
    else:
        print(f"{url}: failed to scrape ({book_json.get("error")})")
    api_response = post(
        "http://localhost/api/books",
        json = book_json,
    )
    api_json = api_response.json()
    if api_response.status_code == 201:
        print(f"{url}: uploaded ({api_json.get("id")})")
    else:
        print(f"{url}: failed to upload ({api_json.get("error")})")
    if book_json.get("image_uuid") is None or api_json.get("id") is None:
        return
    print(f"{url}: updating cover")
    cdn_response = put(
        f"http://localhost/api/books/{api_json.get("id")}/cover",
        json = {"imageuuid": book_json.get("image_uuid")},
    )
    if cdn_response.status_code == 200:
        print(f"{url}: updated cover")
    else:
        print(f"{url}: failed to update cover")

if __name__ == "__main__":
    with open(argv[1]) as urls:
        for url in urls:
            url = url.strip()
            if url == "":
                continue
            process = Process(target=scrape_and_upload, args=(url,))
            # FIX: this might make too many requests; introduce a semaphore?
            process.start()
        process.join()
