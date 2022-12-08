FROM golang:1.19-alpine as builder

RUN apk add --update \
  build-base \
  gcc \
  git

RUN mkdir -p /opt/golang_builder

COPY . /opt/golang_builder

WORKDIR /opt/golang_builder

RUN go build -o main main.go

FROM alpine:latest

COPY --from=builder /opt/golang_builder/main .

CMD ["./main", "-r", "runWebApp"]