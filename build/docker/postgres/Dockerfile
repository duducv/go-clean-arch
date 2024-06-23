FROM postgres:16.3
ENV POSTGRES_USER=admin
ENV POSTGRES_PASSWORD=as1234
ENV POSTGRES_DB=ticket

COPY ./build/database/schema/init.sql /docker-entrypoint-initdb.d/
