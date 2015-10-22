all: leyra

deps:
	{{ range $key, $value := .Goget}}go get {{ $value }}
	{{ end }}

env:
	cp env.example .env

leyra: env deps main.go
	go fmt leyra/...
	go build -v -o server

run: leyra
	@echo "Starting your application..."
	./server

clean: server
	rm server
