FROM postgres:11

ENV entrypoint /docker-entrypoint-initdb.d/

RUN mkdir -p ${entrypoint}

COPY create.sql ${entrypoint}
