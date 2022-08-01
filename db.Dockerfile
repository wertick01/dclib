FROM mysql:latest

COPY ./internals/app/sql/*.sql /docker-entrypoint-initdb.d/