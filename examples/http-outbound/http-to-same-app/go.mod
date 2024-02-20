module examples/http-outbound/http-to-same-app

go 1.20

require github.com/fermyon/spin-go-sdk v0.0.0

require github.com/julienschmidt/httprouter v1.3.0 // indirect

replace github.com/fermyon/spin-go-sdk v0.0.0 => ../../../
