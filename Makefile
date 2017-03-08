TEST?=./...
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

default: fmt test testrace vet protoc

# test runs the test suite and vets the code
test: get-deps fmtcheck
	@golint ./...
	@echo "==> Running Tests"
	@go list $(TEST) | xargs -n1 go test -timeout=60s -parallel=10 $(TESTARGS)

# testrace runs the race checker
testrace:
	@go list $(TEST) | xargs -n1 go test -race $(TESTARGS)

# dev creates binaries for testing locally. These are put
# into ./bin/ as well as $GOPATH/bin
dev: get-deps fmt 
	@EFR_DEV=1 sh -c "'$(CURDIR)/scripts/build.sh'"

bin: get-deps
	sh -c "'$(CURDIR)/scripts/buildbin.sh'"

# vet runs the Go source code static analysis tool `vet` to find
# any common errors.
vet:
	@echo "==> Running Go Vet"
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

get-deps:
	@echo "==> Fetching dependencies"
	@go get -v $(TEST)
	@go get -u github.com/golang/lint/golint
	

fmt:
	gofmt -w $(GOFMT_FILES)

protoc:
	protoc -I ./proto/tickets ./proto/tickets/tickets.proto --go_out=plugins=grpc:proto/tickets 
	protoc -I ./proto/users ./proto/users/users.proto --go_out=plugins=grpc:proto/users
	sed -i '' 's/"Email,omitempty/"-/g' proto/tickets/tickets.pb.go 

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"