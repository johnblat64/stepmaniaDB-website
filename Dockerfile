FROM golang:1.20

WORKDIR /app
COPY *.go resources go.mod go.sum ./
RUN go mod download
RUN go build -o ./stepmaniadb-website
CMD ["./stepmaniadb-website"]