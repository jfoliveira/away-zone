# Build apps binary, from source
FROM golang:1.22 AS builder

ARG EXECUTABLE_NAME

ENV BINARY_NAME=${EXECUTABLE_NAME}
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY *.go ./
RUN go build -o ${EXECUTABLE_NAME}


FROM scratch as health-checker
WORKDIR /
COPY --from=builder /app/health-checker /health-checker
ENTRYPOINT ["/health-checker"]

FROM scratch as web
WORKDIR /
COPY --from=builder /app/web-app /web
EXPOSE 8000

ENTRYPOINT ["/web"]
