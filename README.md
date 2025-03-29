# 👨‍💻 **Fedor Iumashev** | **RapidPurpleSec** 🚀

Welcome to my **RapidPurpleSec** repository! 🛡️

🔐 **Проект:** Разработка процессов непрерывного тестирования безопасности и оценка состояния безопасности информационных систем.

## 📜 Описание

`RapidPurpleSec` — это многофункциональное консольное приложение, которое позволяет автоматизировать процессы тестирования безопасности и оценки уязвимостей в информационных системах. Этот проект служит основой для создания непрерывных процессов тестирования и мониторинга безопасности в современных DevSecOps-процессах.

### 🎯 Цели проекта:
- ✅ **Непрерывное тестирование безопасности** — автоматизация процесса проверки безопасности кода и инфраструктуры.
- ✅ **Оценка уязвимостей и слабых мест системы** — выявление потенциальных угроз и уязвимостей в системе.
- ✅ **Интеграция с CI/CD** — интеграция с системой непрерывной интеграции и развертывания для автоматизации процессов безопасности.
- ✅ **Автоматическое генерирование отчетов** — создание отчетов о состоянии безопасности системы с возможностью вывода в различные форматы.

## 🛠️ Технологии

- **Язык:** Python, Go
- **Инструменты:**
  - [GitHub Actions](https://github.com/features/actions) для автоматизации процессов CI/CD.
  - [Docker](https://www.docker.com/) для контейнеризации приложений и сервисов.
  - [Jenkins](https://www.jenkins.io/) для автоматизации сборок и развертывания.
  - Инструменты безопасности: Bandit, Nmap, Trivy, Gitleaks, Pip-audit для проверки кода и инфраструктуры на уязвимости.

- **Лицензия:** MIT

## 📦 Установка

Чтобы запустить `RapidPurpleSec` на своей машине, выполните следующие шаги:

1. Клонируйте репозиторий:

    ```bash
    git clone https://github.com/FedorIumashev/RapidPurpleSec.git
    ```

2. Перейдите в папку с проектом:

    ```bash
    cd RapidPurpleSec
    ```

3. Установите все зависимости:

    ```bash
    pip install -r requirements.txt
    ```

4. Убедитесь, что у вас установлены Go и Docker (если используете соответствующие инструменты).

## 🚀 Как использовать

1. Для запуска проекта и сканирования кода используйте команду:
    Если хотите запустить вызов команд rps
    ```bash
    ./rps help 
    ```
    Если хотите быстрый отчёт проверки
    ```bash
    python process_reports.py --directory <path_to_your_directory>
    ```

    Где:
    - `--directory <path_to_your_directory>` — путь к каталогу, который вы хотите проанализировать.

2. Пример запуска проекта для анализа кода:

    ```bash
    ./src main.py
    ```

3. Вы также можете настроить автоматическое сканирование с помощью CI/CD-пайплайнов (например, GitHub Actions, Jenkins).

## 📊 Отчеты

После выполнения тестирования безопасности, вы можете получить отчет в формате JSON или CSV, который будет сохранен в папке `reports`.

Пример команды для обработки отчетов:

```bash
python process_reports.py
