FROM golang:alphine

RUN apk update

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o go-cloud-run

ENTRYPOINT [ "./go-cloud-run" ]