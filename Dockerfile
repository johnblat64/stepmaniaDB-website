FROM golang:1.20

WORKDIR /app
COPY *.go go.mod go.sum ./
COPY resource ./resources
RUN go mod download
RUN go build -o ./stepmaniadb-website
CMD ["./stepmaniadb-website"]