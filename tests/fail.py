# test.py
import os
import subprocess

# Пример уязвимости
os.system("echo Hello World")  # Уязвимость: использование os.system()

# Пример другой уязвимости
eval("print('Hello, World!')")  # Уязвимость: использование eval
