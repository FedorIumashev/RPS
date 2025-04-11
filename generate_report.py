import os
from xml.etree.ElementTree import Element, SubElement, tostring
from xml.dom.minidom import parseString

report = Element('report')

def check_file(path, tool, message):
    if os.path.isfile(path):
        SubElement(report, 'tool', name=tool, status="OK", path=path)
    else:
        SubElement(report, 'tool', name=tool, status="DataNotFound", message=message)

check_file("./reports/bandit-report.json", "Bandit", "No Bandit report")
check_file("./reports/pip-audit-report.json", "Pip-audit", "No pip-audit report")
check_file("./reports/trivy-report.sarif", "Trivy", "No containers found")
check_file("./reports/gitleaks-report.sarif", "Gitleaks", "Git data not found")

SubElement(report, 'link', name="Grafana", href="http://localhost:3000")
SubElement(report, 'link', name="SonarQube", href="http://localhost:9000")
SubElement(report, 'link', name="Prometheus", href="http://localhost:9090")

with open("final_report.xml", "w") as f:
    xml_str = parseString(tostring(report)).toprettyxml()
    f.write(xml_str)

print("📄 Отчёт создан: final_report.xml")
