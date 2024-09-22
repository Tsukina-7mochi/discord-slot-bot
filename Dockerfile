FROM golang:latest AS build

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /app/main ./cmd/main.go


FROM gcr.io/distroless/static-debian12 AS prod

WORKDIR /app
COPY --from=build /app/main /main
CMD ["/main"]
