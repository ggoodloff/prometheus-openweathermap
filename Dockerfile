ARG VERSION_GOLANG=1.21
ARG VERSION_ALPINE=3.19

FROM golang:${VERSION_GOLANG}-alpine${VERSION_ALPINE} AS build

COPY . /project
WORKDIR /project

RUN apk add --no-cache build-base

RUN go test ./...
RUN go build ./cmd/openweathermap

FROM alpine:${VERSION_ALPINE} as oerproxy

COPY --from=build /project/openweathermap /usr/local/bin

EXPOSE 80

ENTRYPOINT ["openweathermap"]
