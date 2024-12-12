FROM golang:1.23.4-alpine3.21
WORKDIR /app
COPY . .
RUN go build -o pixie-repro
CMD ["/app/pixie-repro"] 
