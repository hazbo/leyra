all: leyra

deps:
	go get gopkg.in/leyra/echo.v1
	go get gopkg.in/leyra/godotenv.v1
	go get gopkg.in/leyra/mysql.v1
	go get gopkg.in/leyra/gorm.v1
	go get gopkg.in/leyra/toml.v1

leyra: deps main.go
	go fmt leyra/...
	go build

run: leyra
	./leyra

clean: leyra
	rm leyra