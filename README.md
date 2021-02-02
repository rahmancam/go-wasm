# GO WASM
Go web assembly

# Compilation
- To compile run,
```sh
GOOS=js GOARCH=wasm go build -o src/lib.wasm src/main.go
```  

# Server
- To run server, in another terminal
```sh
go run src/server.go
```  