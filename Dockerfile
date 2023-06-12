FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /5-second-rule-go

EXPOSE 8080

CMD [ "/5-second-rule-go" ]