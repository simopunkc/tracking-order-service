## Build ##
FROM golang:alpine AS build
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go build -o tracking-order cmd/cron/main.go

## Deploy ##
FROM gcr.io/distroless/base-debian12:nonroot
WORKDIR /app
COPY --from=build /app/.env /app/.env
COPY --from=build /app/tracking-order /app/tracking-order
USER nonroot:nonroot