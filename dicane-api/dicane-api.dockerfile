# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY dicaneApp /app

CMD ["/app/dicaneApp"]