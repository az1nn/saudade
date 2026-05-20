run: build
	@./bin/app

build:
	@go build -o bin/app .

dev:
	@air

css:
	tailwindcss -i views/css/app.css -o public/styles.css --watch
