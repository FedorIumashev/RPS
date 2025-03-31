# insecure_code.py

import subprocess

def insecure_command():
    # Эта команда может быть уязвимой для атак, так как параметры передаются в командную строку напрямую
    subprocess.run(f"echo {input('Enter a command: ')}", shell=True)
