FROM golang:1.22.5

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@latest
# copy files
COPY . .

# Build
RUN templ generate
RUN go build -o ./bin/main ./cmd/seo-helper-ai/main.go

EXPOSE 9000

# Run
CMD ["./bin/main"]