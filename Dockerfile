FROM golang:1.19-alpine as builder

RUN apt-get update && apt-get -y dist-upgrade
RUN apt install -y netcat

ENV GO111MODULE=on
RUN mkdir -p /go/src/github.com/Avyukth/dicedb
WORKDIR /go/src/github.com/Avyukth/dicedb
ADD . /go/src/github.com/Avyukth/dicedb

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o dice main.go


FROM gcr.io/distroless/static:nonroot
WORKDIR /app/server

COPY --from=builder /go/src/github.com/Avyukth/dicedb ./

EXPOSE 7379

RUN chmod +x ./dice

CMD [ "./dice" ]
