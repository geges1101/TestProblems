from binance.client import Client

import random

import keys
import tests

client = Client(keys.api_key, keys.api_secret)


# Проверяем корректность вхожных данных
def validate_input(data):
    if not isinstance(data, dict):
        raise ValueError("Input data should be a dictionary")

    required_keys = ["volume", "number", "amountDif", "side", "priceMin", "priceMax"]
    for key in required_keys:
        if key not in data:
            raise ValueError(f"Missing required key '{key}' in input data.")

    if not isinstance(data["volume"], (int, float)) or data["volume"] <= 0:
        raise ValueError("Volume must be a positive number")

    if not isinstance(data["number"], int) or data["number"] <= 0:
        raise ValueError("Number of orders must be a positive integer")

    if not isinstance(data["amountDif"], (int, float)) or data["amountDif"] <= 0:
        raise ValueError("Amount difference must be a positive number")

    if data["side"] not in ["BUY", "SELL"]:
        raise ValueError("Side must be either 'BUY' or 'SELL'")

    if not isinstance(data["priceMin"], (int, float)) or data["priceMin"] < 0:
        raise ValueError("Minimum price must be a non-negative number")

    if not isinstance(data["priceMax"], (int, float)) or data["priceMax"] <= 0:
        raise ValueError("Maximum price must be a positive number")

    if data["priceMax"] <= data["priceMin"]:
        raise ValueError("Maximum price must be greater than minimum price")

    return True


def create_order(symbol, side, quantity, price):
    order = client.create_order(
        symbol=symbol,
        side=side,
        type='LIMIT',
        timeInForce='GTC',
        quantity=quantity,
        price=price
    )
    return order


def create_orders(data):
    volume = data['volume']
    number = data['number']
    amount_dif = data['amountDif']
    side = data['side']
    price_min = data['priceMin']
    price_max = data['priceMax']

    # Не продолжаем работу с некорректными данными
    validate_input(data)

    # Разбиваем общий объем на равные части для каждого ордера
    volume_per_order = volume / number

    for i in range(number):
        # Генерируем случайные цены и объемы для ордеров
        price = round(random.uniform(price_min, price_max), 2)
        volume_with_dif = round(random.uniform(volume_per_order - amount_dif, volume_per_order + amount_dif), 2)

        # Округляем объем до шага торговли
        step_size = client.get_symbol_info('BTCUSDT')['filters'][1]['stepSize']
        quantity = round(volume_with_dif / price, int(step_size.find('1') - 1))

        order = create_order('BTCUSDT', side, quantity, price)
        print(f'Created order: {order}')


# проверяем тесты
create_orders(tests.test1)
create_orders(tests.test2)
create_orders(tests.test3)
create_orders(tests.test4)
create_orders(tests.test5)
create_orders(tests.test6)
create_orders(tests.test7)
create_orders(tests.test8)
create_orders(tests.test9)
create_orders(tests.test10)
