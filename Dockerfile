FROM golang:1.20-alpine3.18 AS build

WORKDIR /app

COPY . .

RUN go build -o ./dist/apollo .

FROM alpine:3.18 AS start

COPY --from=build /app/dist /app

EXPOSE 8080

ENTRYPOINT ["/app/apollo"]