lint:
	golangci-lint run ./...

test: lint
	go test -coverprofile=c.out -v ./... \
	&& go tool cover -html=c.out \
	&& rm c.out

race: lint
	go test -coverprofile=rc.out -v ./... \
	&& go tool cover -html=rc.out \
	&& rm rc.out
