clean:
	go clean --cache

test: clean
	go test -race -cover -coverprofile cov.out ./...
	go tool cover -func cov.out | grep total:

total-coverage: clean test
	go tool cover -func cov.out | grep total | awk '{ print substr($3, 1, length($3)-1) }'

coverage: test
	go tool cover -html=cov.out -o cover.html
	open cover.html
