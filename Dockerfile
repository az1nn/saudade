# syntax=docker/dockerfile:1

# ---------- build stage ----------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install templ code generator
RUN go install github.com/a-h/templ/cmd/templ@v0.2.663

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Regenerate templ files (in case sources are newer than generated files)
RUN templ generate ./views/...

# Build the production binary (embeds public/ via //go:embed)
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/app .

# ---------- runtime stage ----------
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/bin/app ./app

EXPOSE 8080

CMD ["./app"]
