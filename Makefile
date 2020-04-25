test:
	go test -coverprofile=cover.out ./...

cover:
	go tool cover -html=cover.out
