all: leyra

deps:
	{{deps}}

leyra: deps main.go
	go fmt leyra/...
	go build

run: leyra
	./leyra

clean: leyra
	rm leyra