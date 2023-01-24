all: open run

build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build

open:
	open http://localhost:8000

run: build
	./reservas