FROM golang:1.27.4 AS development
WORKDIR /app

COPY go.mod ./
RUN go mod download

CMD ["go", "run", "./cmd/controller"]

FROM golang:1.27.4 AS build
WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/controller

FROM scratch AS production
COPY --from=build /app/server /server
ENTRYPOINT ["/server"]
