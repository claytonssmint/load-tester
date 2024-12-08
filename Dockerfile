FROM golang:1.23 as build
WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o load-tester ./cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/load-tester .
ENTRYPOINT ["/app/load-tester"]
