FROM alpine:3.7
RUN apk update
RUN apk add curl
RUN curl -D - -s  -o /dev/null http://example.com

RUN curl -o cloud-sql-proxy https://storage.googleapis.com/cloud-sql-connectors/cloud-sql-proxy/v2.1.2/cloud-sql-proxy.linux.amd64
RUN chmod +x cloud-sql-proxy

RUN ./cloud-sql-proxy portfolio-fab99:us-central1:portfolio-db
