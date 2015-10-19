all: leyra

deps:
	{{goget}}

leyra: deps main.go
	go fmt leyra/...
	go build

run: leyra
	./leyra

clean: leyra
	rm leyra