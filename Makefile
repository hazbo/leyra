all: leyra

deps:
	go get gopkg.in/leyra/echo.v1
	go get gopkg.in/leyra/godotenv.v1
	go get gopkg.in/leyra/mysql.v1
	go get gopkg.in/leyra/go-sqlite3.v1
	go get gopkg.in/leyra/gorm.v1
	go get gopkg.in/leyra/toml.v1
	go get gopkg.in/leyra/grace.v1/gracehttp
	go get gopkg.in/leyra/blackfriday.v1
	go get gopkg.in/leyra/sessions.v1

env:
	cp env.example .env

leyra: env deps main.go
	go fmt leyra/...
	go build -v -o server

run: leyra
	./server

clean: server
	rm server
