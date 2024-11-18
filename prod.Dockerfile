# Build environment
# -----------------
FROM golang:1.22-alpine as build-env
WORKDIR /brickstorage

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/api ./cmd/api \
    && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate


# Deployment environment
# ----------------------
FROM alpine

COPY --from=build-env /brickstorage/bin/api /brickstorage/
COPY --from=build-env /brickstorage/bin/migrate /brickstorage/
COPY --from=build-env /brickstorage/migrations /brickstorage/migrations

EXPOSE 3023
CMD ["/brickstorage/api"]