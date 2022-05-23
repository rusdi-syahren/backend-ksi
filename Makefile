.PHONY : test format

ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")

hello:
	echo "welcome to halomobi api"

clean:
	[ -f app-osx ] && rm app-osx || true
	[ -f app-linux ] && rm app-linux || true
	[ -f app32.exe ] && rm app32.exe || true
	[ -f app64.exe ] && rm app64.exe || true
	[ -f coverage.txt ] && rm coverage.txt || true
	rm ./coverages/*.txt



app-osx: main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o $@

app-linux: main.go
	GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $@

app64.exe: main.go
	GOOS=windows GOARCH=amd64 go build -ldflags '-s -w' -o $@

app32.exe: main.go
	GOOS=windows GOARCH=386 go build -ldflags '-s -w' -o $@

app-windows: app64.exe app32.exe

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

# Docker Build
docker: Dockerfile
	docker build -t boilerplate-go:latest .

test:
	$(foreach pkg, $(ALL_PACKAGES),\
	go test -race -short $(pkg);)