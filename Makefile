lint:
	golangci-lint run ./...

test: lint
	go test -coverprofile=c.out -v -race -count=1 ./... \
	&& go tool cover -html=c.out \
	&& rm c.out
