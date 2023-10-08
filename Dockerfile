FROM golang:1.21.2-alpine AS build

WORKDIR /app

COPY ./src .
RUN go build -ldflags "-s -w"

FROM alpine:3

WORKDIR /usr/local/bin
COPY --from=build /app/clarity-api .

CMD ["clarity-api"]