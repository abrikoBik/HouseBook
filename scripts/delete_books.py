from requests import get, delete

def get_books(order: int, limit: int, offset: int) -> list:
    print(f"{order}: getting {limit} books with offset {offset}")
    response = get(
        f"http://localhost/api/books?limit={limit}&offset={offset}"
    )
    response_json = response.json()
    if response.status_code == 200:
        print(f"{order}: got {len(response_json)} books")
        return response_json
    else:
        print(f"{order}: error: {response_json.get("error")}")
        return list()

def delete_book(id: str):
    print(f"{id}: deleting")
    response = delete(
        f"http://localhost/api/books/{id}"
    )
    if response.status_code == 200:
        print(f"{id}: deleted")
    else:
        response_json = response.json()
        print(f"{id}: failed to delete ({response_json.get("error")})")

if __name__ == "__main__":
    id = 0
    while True:
        books = get_books(id, 10, 0)
        if len(books) == 0:
            break
        for book in books:
            delete_book(book.get("id"))
        id += 1
