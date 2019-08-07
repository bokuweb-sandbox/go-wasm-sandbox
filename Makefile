.PHONY: all clean serve

build:
	GOOS=js GOARCH=wasm go build -o "main.wasm" "wasm_main.go"

serve:
	xdg-open 'http://localhost:5002'
	serve -a :5002 || (go get -v github.com/mattn/serve && serve -a :5002)

clean:
	rm -f *.wasm
