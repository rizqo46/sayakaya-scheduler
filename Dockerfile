FROM golang:1.22.3-alpine AS build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /binary ./cmd/app/main.go

FROM scratch

COPY --from=build /binary /binary
COPY --from=build /usr/local/go/lib/time/zoneinfo.zip /
ENV ZONEINFO=/zoneinfo.zip

ENTRYPOINT ["/binary"]