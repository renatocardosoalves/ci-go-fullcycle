FROM golang:latest AS builder
WORKDIR /app
COPY *.go ./
RUN go mod init ci && go build -o math

FROM scratch
WORKDIR /app
COPY --from=builder /app/math .
CMD [ "./math" ]