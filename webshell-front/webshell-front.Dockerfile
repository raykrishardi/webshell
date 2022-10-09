ARG ALPINE_VERSION=3.16.2
FROM alpine:${ALPINE_VERSION}
LABEL maintainer="Ray Krishardi Layadi <raykrishardi@gmail.com>"

RUN mkdir /app

COPY webshFrontApp /app
COPY static /app/static
COPY templates /app/templates

WORKDIR /app

CMD ["/app/webshFrontApp"]