# test.py
import os

# Пример уязвимости: использование небезопасной функции
os.system('echo Hello World')  # Эта команда может быть опасной

# Пример уязвимости: небезопасное использование данных
user_input = input("Enter your name: ")
eval(user_input)  # Уязвимость через eval()

# Пример ошибки: использование глобальных переменных без предварительного объявления
x = 10
def test_func():
    print(x)  # Будет ошибка, если переменная не объявлена как глобальная

test_func()
