import psycopg2
from telegram.ext import Updater, CommandHandler

# подключаемся к БД
conn = psycopg2.connect(
    host="localhost",
    database="codeforces",
    user="postgres",
    password="password"
)


# команда /start
def start(update, context):
    context.bot.send_message(chat_id=update.effective_chat.id, text="Привет! Введите уровень сложности(например: 800) "
                                                                    "и тему(например: strings) задачи")


# команда /get
def get_problems(update, context):
    user_input = update.message.text.split()

    difficulty = user_input[0]
    topic = " ".join(user_input[1:])

    cur = conn.cursor()
    cur.execute(
        "SELECT name, number, difficulty, topics, solutions FROM problems WHERE difficulty = %s AND %s = ANY(topics) "
        "LIMIT 10",
        (difficulty, topic)
    )
    result = cur.fetchall()

    if result:
        for item in result:
            message = f"{item[0]} ({item[1]})\nDifficulty: {item[2]}\nTopics: {', '.join(item[3])}\nNumber of Solutions: {item[4]}"
            context.bot.send_message(chat_id=update.effective_chat.id, text=message)
    else:
        context.bot.send_message(chat_id=update.effective_chat.id, text="Не нашлось подходящих задач")


# команда /search
def search(update, context):
    name = update.message.text.split()[1]

    cur = conn.cursor()
    cur.execute(
        "SELECT name, number, difficulty, topics, solutions FROM problems WHERE name ILIKE %s",
        ('%' + name + '%')
    )
    result = cur.fetchall()

    if result:
        for item in result:
            message = f"{item[0]} ({item[1]})\nDifficulty: {item[2]}\nTopics: {', '.join(item[3])}\nNumber of Solutions: {result[4]}"
            context.bot.send_message(chat_id=update.effective_chat.id, text=message)
    else:
        context.bot.send_message(chat_id=update.effective_chat.id, text="Не нашлось подходящих задач")


token = open("/Users/geges/Documents/Dev/CodeForcesWebScraper/bot_token.txt", "r").__str__()
updater = Updater(token, use_context=True)
dispatcher = updater.dispatcher
dispatcher.add_handler(CommandHandler("start", start))
dispatcher.add_handler(CommandHandler("get", get_problems))
dispatcher.add_handler(CommandHandler("search", search))

updater.start_polling()
