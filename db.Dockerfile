FROM mysql:8.0.30

COPY ./internals/app/sql/*.sql /docker-entrypoint-initdb.d/