windows: OUTPUT=go-2x2-solver.exe
linux: OUTPUT=go-2x2-solver

build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm ./cmd/main.go
	go build -o build/bin/$(OUTPUT) ./cmd/main.go

run: build
	./build/bin/$(OUTPUT)
