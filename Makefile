.PHONY: all clean serve

all: main.wasm serve

%.wasm: %.go
	GOOS=js GOARCH=wasm go build -o "$@" "$<"

serve:
	xdg-open 'http://localhost:5002'
	serve -a :5002 || (go get -v github.com/mattn/serve && serve -a :5002)

clean:
	rm -f *.wasm
