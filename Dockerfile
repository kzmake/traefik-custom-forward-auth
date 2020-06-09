FROM golang:1.14-alpine AS builder

RUN set -ex \
  && apk add --no-cache -q --no-progress git curl g++ gcc libgcc linux-headers make

WORKDIR /go/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -o randauth service/randauth/cmd/main.go


FROM scratch

COPY --from=builder /go/app/randauth .

EXPOSE 3000

CMD ["/randauth"]
