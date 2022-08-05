include .env

.PHONY: up restart migrate g-logs db-shell db-dump migrate

NETWORK = dclib_default
DIR = /home/mrred/'Рабочий стол'/Работа/dclib

COMMAND = version

up:
	docker compose build && docker compose up -d
restart:
	docker compose restart golang
g-logs:
	docker compose logs -f --tail="100" golang
db-shell:
	docker compose exec db bash
db-dump:
	docker exec -t mysql mysqldump -v -h $(DBHOST) -u$(DBUSER) -p$(DBPASS) $(DBNAME)> ./mysql/backup/"`date +"%Y%m%d"`"-dump.sql
migrate:
	docker run --rm -v $(DIR)/internals/app/sql:/migrations --network $(NETWORK) migrate/migrate -path=/migrations/ -database "mysql://$(DBUSER):$(DBPASS)@tcp($(DBHOST):$(DBPORT))/$(DBNAME)" $(COMMAND)