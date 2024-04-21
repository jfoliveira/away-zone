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

# use a shell so env var is expanded
CMD [ "sh", "-c", "./$BINARY_NAME" ]
