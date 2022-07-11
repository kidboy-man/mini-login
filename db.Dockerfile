FROM postgres
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgrespw
ENV POSTGRES_DB db_user_mini
COPY init.sql /docker-entrypoint-initdb.d/
EXPOSE 5432