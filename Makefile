all: leyra

deps:
	go get gopkg.in/leyra/echo.v1
	go get github.com/joho/godotenv

leyra: deps main.go
	go fmt leyra/...
	go build

run: leyra
	./leyra

clean: leyra
	rm leyra