.PHONY: default run build test docs clean

APP_NAME := goapi

default: run-build

run:
	@docker-compose down
	@docker-compose up -d

run-build:
	@docker-compose down
	@docker-compose up -d --build 

