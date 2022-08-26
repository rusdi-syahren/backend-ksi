# .PHONY : test format

# ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")

# hello:
# 	echo "welcome to halomobi api"

# clean:
# 	[ -f app-osx ] && rm app-osx || true
# 	[ -f app-linux ] && rm app-linux || true
# 	[ -f app32.exe ] && rm app32.exe || true
# 	[ -f app64.exe ] && rm app64.exe || true
# 	[ -f coverage.txt ] && rm coverage.txt || true
# 	rm ./coverages/*.txt



# app-osx: main.go
# 	GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o $@

# app-linux: main.go
# 	GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $@

# app64.exe: main.go
# 	GOOS=windows GOARCH=amd64 go build -ldflags '-s -w' -o $@

# app32.exe: main.go
# 	GOOS=windows GOARCH=386 go build -ldflags '-s -w' -o $@

# app-windows: app64.exe app32.exe

# format:
# 	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

# # Docker Build
# docker: Dockerfile
# 	docker build -t boilerplate-go:latest .

# test:
# 	$(foreach pkg, $(ALL_PACKAGES),\
# 	go test -race -short $(pkg);)


# Definitions
ROOT                    := $(PWD) # direktori folder
GO_HTML_COV             := ./coverage.html #nama file coverage.html
GO_TEST_OUTFILE         := ./c.out #output gotest
GOLANG_DOCKER_IMAGE     := golang:1.4.3-alpine #menjalankan go dengan versi tertentu
CODECLIMATE_DEV			:= ${CODECLIMATE_DEV} #report id codeclimate
CC_PREFIX				:= github.com/rusdi-syahren/backend-ksi #prefix url repo kita

.PHONY: clean build packing

# custom logic for code climate, gross but necessary
coverage:
	# download CC test reported
	docker run -w /app -v ${ROOT}:/app ${GOLANG_DOCKER_IMAGE} \
		/bin/bash -c \
		"curl -L https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64 > ./cc-test-reporter"
	
	# update perms
	docker run -w /app -v ${ROOT}:/app ${GOLANG_DOCKER_IMAGE} chmod +x ./cc-test-reporter

	# run before build
	docker run -w /app -v ${ROOT}:/app \
		 -e CODECLIMATE_DEV=${CODECLIMATE_DEV} \
		${GOLANG_DOCKER_IMAGE} ./cc-test-reporter before-build

	# run testing
	docker run -w /app -v ${ROOT}:/app ${GOLANG_DOCKER_IMAGE} go test ./... -coverprofile=${GO_TEST_OUTFILE}
	docker run -w /app -v ${ROOT}:/app ${GOLANG_DOCKER_IMAGE} go tool cover -html=${GO_TEST_OUTFILE} -o ${GO_HTML_COV}

	#upload coverage result
	$(eval PREFIX=${CC_PREFIX})
ifdef prefix
	$(eval PREFIX=${prefix})
endif
	# upload data to CC
	docker run -w /app -v ${ROOT}:/app \
		-e CODECLIMATE_DEV=${CODECLIMATE_DEV} \
		${GOLANG_DOCKER_IMAGE} ./cc-test-reporter after-build --prefix ${PREFIX}

test:
	@go test ./... -coverprofile=./coverage.out & go tool cover -html=./coverage.out
	
build:
	@GOOS=linux GOARCH=amd64
	@echo ">> Building GRPC..."
	@go build -o phonebook-service-grpc ./cmd/grpc
	@echo ">> Finished"

run:
	@./phonebook-service-grpc
