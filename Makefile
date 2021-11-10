update_chain_client:
	GOPRIVATE=github.com/poktbridge/* go get -u github.com/poktbridge/protocols/sign && go mod tidy
production_bin:
	GOPRIVATE=signer/* GOOS=linux CGO_ENABLED=0 go build -a -ldflags="-w -s" ./cmd/signer