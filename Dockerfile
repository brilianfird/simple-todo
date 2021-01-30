FROM golang:alpine AS build-env

WORKDIR $HOME/app

COPY main ./main
COPY conf.json .
COPY go.mod .

WORKDIR main

RUN go build

FROM alpine
WORKDIR /
COPY --from=build-env /app/main/main /
COPY --from=build-env /app/conf.json /
ENTRYPOINT ./main