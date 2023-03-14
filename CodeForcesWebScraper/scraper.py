from typing import List, Dict

import psycopg2
import requests
from bs4 import BeautifulSoup


# парсим задачи с сайта codeforces
def parse_problemset() -> List[Dict]:
    url = "https://codeforces.com/problemset?order=BY_SOLVED_DESC"
    response = requests.get(url)
    soup = BeautifulSoup(response.content, "html.parser")

    problems = []
    for row in soup.find_all("tr"):
        cols = row.find_all("td")
        if len(cols) > 0:
            name = cols[0].text.strip()
            num_solutions = int(cols[2].text.strip())
            tags = [tag.text.strip() for tag in cols[1].find_all("a")]
            difficulty = int(cols[3].text.strip())
            problems.append({"name": name, "num_solutions": num_solutions, "tags": tags, "difficulty": difficulty})
    return problems


# подключаемся к БД и добавляем данные о задачах
def update_database(problems: List[Dict]):
    conn = psycopg2.connect(database="codeforces", user="postgres", password="password", host="localhost", port="5432")
    cur = conn.cursor()

    for problem in problems:
        cur.execute("SELECT id FROM problems WHERE name=%s", (problem["name"],))
        result = cur.fetchone()
        if result:
            cur.execute("UPDATE problems SET num_solutions=%s, tags=%s, difficulty=%s WHERE id=%s",
                        (problem["num_solutions"], problem["tags"], problem["difficulty"], result[0]))
        else:
            cur.execute("INSERT INTO problems (name, num_solutions, tags, difficulty) VALUES (%s, %s, %s, %s)",
                        (problem["name"], problem["num_solutions"], problem["tags"], problem["difficulty"]))

    conn.commit()
    cur.close()
    conn.close()
