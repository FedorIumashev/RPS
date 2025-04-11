import subprocess
from pathlib import Path
import xml.etree.ElementTree as ET

report_paths = {
    "Bandit": "reports/bandit-report.json",
    "Pip-audit": "reports/pip-audit-report.json",
    "Trivy": "reports/trivy-report.sarif",
    "Gitleaks": "reports/gitleaks-report.sarif"
}

local_dashboards = {
    "Grafana": "http://localhost:3000",
    "Prometheus": "http://localhost:9090",
    "SonarQube": "http://localhost:9000"
}

def generate_xml_report(report_paths, local_dashboards, output_path="final_report.xml"):
    root = ET.Element("SecurityAnalysisReport")
    for name, path in report_paths.items():
        result = ET.SubElement(root, name.replace(" ", ""))
        report_path = Path(path)
        if report_path.exists():
            result.text = str(report_path)
        else:
            result.text = "DataNotFound"

    dashboards = ET.SubElement(root, "Dashboards")
    for name, url in local_dashboards.items():
        dash = ET.SubElement(dashboards, name)
        dash.text = url

    tree = ET.ElementTree(root)
    tree.write(output_path, encoding="utf-8", xml_declaration=True)


def run_docker_compose():
    compose_cmds = [["docker-compose", "up", "--build", "--abort-on-container-exit"],
                    ["docker", "compose", "up", "--build", "--abort-on-container-exit"]]
    for cmd in compose_cmds:
        try:
            return subprocess.Popen(cmd)
        except FileNotFoundError:
            print(f"⚠️ Command not found: {' '.join(cmd)}")
    raise RuntimeError("❌ Neither 'docker-compose' nor 'docker compose' is available on this system.")


def main():
    print("🚀 Launching analysis using Docker...")
    compose_process = run_docker_compose()
    compose_process.wait()

    print("📝 Generating final report...")
    generate_xml_report(report_paths, local_dashboards)

    print("\n✅ Final security report generated: final_report.xml")
    with open("final_report.xml", "r", encoding="utf-8") as file:
        preview = file.read()
        print(preview)


if __name__ == "__main__":
    main()
