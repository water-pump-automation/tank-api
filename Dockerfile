FROM golang:1.19.2 as builder

WORKDIR /go/delivery

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./
RUN go build -v -o ./water-tank-api .

FROM alpine:latest

WORKDIR .

RUN apk add libc6-compat

COPY --from=builder /go/delivery/water-tank-api .

COPY ./entry-point.sh /
RUN chmod 755 ./entry-point.sh

EXPOSE 8080

ENTRYPOINT [ "/entry-point.sh" ]
