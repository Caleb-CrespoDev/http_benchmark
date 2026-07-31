FROM golang:1.26.5 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
RUN CGO_ENABLED=0 go build -o /bench-go-http .

FROM gcr.io/distroless/static-debian12
COPY --from=build /bench-go-http /bench-go-http
EXPOSE 4000
ENTRYPOINT ["/bench-go-http"]
