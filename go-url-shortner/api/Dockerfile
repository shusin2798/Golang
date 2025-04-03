FROM golang:alpine as buider

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN go build -o main .

#stage2

FROM alpine:latest

RUN adduser -S -D -H /app appuser

USER appuser

COPY . /app

COPY --from=builder /build/main /app

WORKDIR /app

EXPOSE 3000

CMD [".main"]