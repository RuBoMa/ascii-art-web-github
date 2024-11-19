# syntax=docker/dockerfile:1

FROM golang:1.21

# set destination for COPY
WORKDIR /app

# DOwnload Go modules
COPY go.mod ./
RUN go mod download

# Copy the source code.
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o ascii-art-web

EXPOSE 8080

# Run
CMD [ "./ascii-art-web" ]
