FROM golang:1.22 as builder

WORKDIR /go/src/water-tank-api

COPY ./go.mod ./go.sum ./
RUN go mod download && go mod verify

COPY . ./
RUN go build -v -o ./water-tank-api ./

FROM alpine:latest

WORKDIR .

RUN apk add libc6-compat

COPY --from=builder /go/src/water-tank-api/water-tank-api .

EXPOSE 8080

ENTRYPOINT [ "./water-tank-api" ]
