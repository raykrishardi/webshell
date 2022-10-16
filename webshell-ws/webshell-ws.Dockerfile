ARG ALPINE_VERSION=3.16.2
FROM alpine:${ALPINE_VERSION}
LABEL maintainer="Ray Krishardi Layadi <raykrishardi@gmail.com>"

# Arguments to be passed from docker build command
ARG ROOT_CMD="websh"

RUN mkdir /app

COPY webshWSApp /app
COPY "${ROOT_CMD}" /app
RUN mv "/app/${ROOT_CMD}" /usr/local/bin

CMD ["/app/webshWSApp"]