from crontab import CronTab

# создаем новую задачу для cron
cron = CronTab(user='geges')

# Добавляем команду
job = cron.new(command='python scraper.py')

# ставим расписание на каждый час
job.setall('0 * * * *')

# выводим результат в файл
cron.write()
