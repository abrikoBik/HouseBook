from urllib.parse import urlparse
from requests import post
from multiprocessing import Process

def determine_service(url: str) -> str:
    parsed = urlparse(url)
    match parsed.netloc.replace("www.", "", 1):
        case "litres.ru":
            return "litres"
        case "labirint.ru":
            return "labirint"
        case "mybook.ru":
            return "mybook"

def scrape_and_upload(url: str):
    print(f"{url}: scraping")
    book_response = post(
        f"http://localhost/scraper/{determine_service(url)}",
        json = {"url": url},
    )
    book_json = book_response.json()
    print(f"{url}: scraped ({book_json.get("title")} by {book_json.get("author")})")
    api_response = post(
        f"http://localhost/api/books",
        json = book_json,
    )
    api_json = api_response.json()
    print(f"{url}: uploaded ({api_json.get("id")})")

if __name__ == '__main__':
    with open("urls.txt") as urls:
        for url in urls:
            url = url.strip()
            if url == "":
                continue
            process = Process(target=scrape_and_upload, args=(url,))
            process.start()
        process.join()
