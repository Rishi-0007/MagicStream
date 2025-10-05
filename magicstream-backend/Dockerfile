FROM golang:1.22 as builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/magicstream ./cmd/api

FROM gcr.io/distroless/base-debian12
WORKDIR /
COPY --from=builder /out/magicstream /magicstream
ENV PORT=8080
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/magicstream"]
