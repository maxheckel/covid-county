FROM postgres:11.5-alpine

# copy sql scripts that execute on boot
# NOTE: we name the files ###_*.sql so we can control the order in which they are executed
COPY ./schema.sql /docker-entrypoint-initdb.d/001_schema.sql
COPY ./tables.sql /docker-entrypoint-initdb.d/003_tables.sql

ENV POSTGRES_DB="covid_county"
