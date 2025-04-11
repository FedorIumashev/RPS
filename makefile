.PHONY: up down build report clean

up:
	docker compose up --build --abort-on-container-exit

down:
	docker compose down

build:
	docker compose build

report:
	python run_analysis.py

clean:
	rm -f final_report.xml
	rm -rf reports/*
