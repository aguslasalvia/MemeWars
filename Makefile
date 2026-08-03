.PHONY: build build-frontend build-server clean

build: build-frontend build-server
	mkdir -p bin/static
	cp -r static/dist bin/static/dist
	mkdir -p bin/static/memes

build-frontend:
	cd frontend && bun install && bun run build

build-server:
	go build -o bin/memewars ./cmd/server

clean:
	rm -rf bin
