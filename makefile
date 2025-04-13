.PHONY: up down build report clean


up:
	docker compose up --build

down:
	docker compose down

build:
	docker compose build

report:
	python run_analysis.py

clean:
	rm -f final_report.xml
	rm -rf reports/*

logs:
	docker compose logs --follow

all: 
	build up report clean