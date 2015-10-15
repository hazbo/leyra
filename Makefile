all: leyra

deps:
	go get github.com/labstack/echo
	go get github.com/joho/godotenv

leyra: deps main.go
	go fmt leyra/...
	go build

run: leyra
	./leyra

clean: leyra
	rm leyra