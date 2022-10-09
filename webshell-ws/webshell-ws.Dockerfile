ARG ALPINE_VERSION=3.16.2
FROM alpine:${ALPINE_VERSION}
LABEL maintainer="Ray Krishardi Layadi <raykrishardi@gmail.com>"

RUN mkdir /app

COPY webshWSApp /app
COPY websh /app
RUN mv "/app/websh" /usr/local/bin

CMD ["/app/webshWSApp"]