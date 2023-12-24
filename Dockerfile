FROM golang:1.21

WORKDIR /app

COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/app /app/cmd/app
EXPOSE 8080
CMD ["/app/bin/app"]