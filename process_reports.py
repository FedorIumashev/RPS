import os
import json

def process_bandit_json(file_path):
    with open(file_path, "r") as file:
        data = json.load(file)

    issues = []
    # Проверяем, если данные это список (в случае с Bandit)
    if isinstance(data, list):
        for issue in data:
            severity = issue.get("severity", "N/A")
            if severity in ["HIGH", "MEDIUM", "LOW"]:  # Фильтруем только ошибки/предупреждения
                issues.append({
                    "file": issue.get("filename", "N/A"),
                    "line": issue.get("line_number", "N/A"),
                    "severity": severity,
                    "message": issue.get("issue", "N/A")
                })
    return issues

def process_reports(input_directory):
    all_issues = []
    status = "OK"  # Начальный статус — все в порядке

    # Проходим по всем файлам в указанной директории
    for root, _, files in os.walk(input_directory):
        for file in files:
            if file.endswith("-report.json"):  # Обрабатываем только json файлы отчетов
                file_path = os.path.join(root, file)
                
                if "bandit" in file:
                    issues = process_bandit_json(file_path)
                    all_issues.extend(issues)
                # Добавить обработку других отчетов (например, для Gitleaks, Trivy) по аналогии

    # Если найдены ошибки, изменяем статус
    if all_issues:
        status = "ERROR"

    # Выводим информацию в текстовом формате
    if status == "OK":
        print("Статус: Все в порядке. Ошибок не обнаружено.")
    else:
        print("Статус: Обнаружены ошибки или предупреждения.")
        print("\nОшибки и предупреждения:")
        for issue in all_issues:
            print(f"Файл: {issue['file']}, Строка: {issue['line']}, Уровень: {issue['severity']}")
            print(f"Описание: {issue['message']}\n")

# Указываем директорию с отчетами
input_directory = "./reports"

process_reports(input_directory)
