# Stage 1: Build
FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/server/

# Stage 2: Runtime
FROM alpine:3.20

RUN apk add --no-cache ca-certificates

COPY --from=build /server /server
COPY --from=build /app/database/migrations/ /app/migrations/

WORKDIR /app
EXPOSE 8080

CMD ["/server"]
