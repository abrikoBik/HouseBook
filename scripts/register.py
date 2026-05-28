from requests import post
from multiprocessing import Process
from sys import argv

def split_and_register(user_str: str):
    # [0] = e-mail, [1] = login, [2] = name, [3] = password
    split = [s.strip() for s in user_str.split(",")]
    if len(split) != 4:
        return
    print(f"{split[1]}: registering")
    response = post(
        "http://localhost/api/users",
        json = {
            "email": split[0],
            "login": split[1],
            "name": split[2],
            "password": split[3],
        },
    )
    response_json = response.json()
    if response.status_code == 201:
        print(f"{split[1]}: registered ({response_json.get("id")})")
    else:
        print(f"{split[1]}: failed to register ({response_json.get("error")})")

if __name__ == "__main__":
    with open(argv[1]) as users:
        for user in users:
            user = user.strip()
            if user[0] == "#" or user == "":
                continue
            process = Process(target=split_and_register, args=(user,))
            process.start()
        process.join()
