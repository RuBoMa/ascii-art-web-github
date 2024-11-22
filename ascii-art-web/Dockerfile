# syntax=docker/dockerfile:1

# Specifies the base image for the build
FROM golang:1.23

# set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
RUN go mod download

# Some metadata for ascii art web project
LABEL maintainer="roope.hongisto@gritlab.ax"
LABEL version="1.0"
LABEL description="Web project creating Ascii Art from text"

# Copy the source code.
COPY . ./

# Disable C bindings, Compiles the binary for linux, names the output binary
RUN CGO_ENABLED=0 GOOS=linux go build -o ascii-art-web .

# Tell what port to use
EXPOSE 8080

# Starts the container automatically
CMD [ "./ascii-art-web" ]
