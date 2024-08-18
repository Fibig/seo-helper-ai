FROM golang:1.22.5

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@latest

# download tailwindcli
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
RUN chmod +x tailwindcss-linux-x64
RUN mv tailwindcss-linux-x64 tailwindcss

# copy files
COPY . .

# build everything
RUN templ generate
RUN tailwindcss -i ./css/tailwind.css -o ./assets/tailwind.css
RUN go build -o ./bin/main ./cmd/seo-helper-ai/main.go

EXPOSE 9000

# Run
CMD ["./bin/main"]