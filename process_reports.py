import json
import os

def print_bandit_report(report_path):
    with open(report_path, 'r') as f:
        data = json.load(f)
    
    print(f"\n[Bandit] Total Issues: {len(data['results'])}")
    for issue in data['results']:
        print(f"File: {issue['filename']}, Line: {issue['line_number']}, Issue: {issue['issue']}")
        
def print_pip_audit_report(report_path):
    with open(report_path, 'r') as f:
        data = json.load(f)

    print("\n[Pip-Audit] Total Issues: {0}".format(len(data['dependencies'])))
    
    # Запрашиваем у пользователя, хочет ли он подробный вывод
    show_details = input("\nDo you want to see detailed vulnerability information? (y/n): ").strip().lower()

    for dep in data['dependencies']:
        if dep['vulns']:
            print(f"\nPackage: {dep['name']} (Version: {dep['version']})")
            if show_details == 'y':
                for vuln in dep['vulns']:
                    print(f"  Vulnerability ID: {vuln['id']}")
                    print(f"  Severity: {vuln['severity']}")
                    print(f"  Description: {vuln.get('description', 'No description available')}")
                    print(f"  URL: {vuln.get('url', 'No URL available')}")
            else:
                print("  (Detailed information skipped. Use 'y' to view details.)")
        else:
            print(f"Package: {dep['name']} (Version: {dep['version']}) - No vulnerabilities found")




# Указание на директорию отчётов
reports_dir = './reports'

# Печать отчётов
print_bandit_report(os.path.join(reports_dir, 'bandit-report.json'))
print_pip_audit_report(os.path.join(reports_dir, 'pip-audit-report.json'))
