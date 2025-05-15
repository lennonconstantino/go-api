FROM golang:1.24.3

# set working directory
WORKDIR /go/src/app

# copy the source code
COPY . .

# expose the port
EXPOSE 8000

# Build the Go App
RUN go build -o main cmd/main.go

# Run the executable
CMD ["./main"]
