.PHONY: build

run:
	air

watch_templ:
	templ generate -watch

watch_css:
	npx tailwindcss -i ./css/tailwind.css -o ./public/css/tailwind.css --watch

build:
	templ generate
	npx tailwindcss -i ./css/tailwind.css -o ./public/css/tailwind.css
	go build -o ./bin/main.exe ./cmd/seo-helper-ai/main.go