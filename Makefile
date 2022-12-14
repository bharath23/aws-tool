BIN="./bin"
SRC=$(shell find . -name "*.go")

build:
	$(info ***************************** building awstool ****************************)
	go build -o $(BIN)/

fmt:
	$(info *************************** checking formatting ***************************)
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

test:
	$(info ****************************** running tests ******************************)
	go test -v ./...

install_deps:
	$(info ************************* downloading dependencies ************************)
	go get -v ./...

clean:
	rm -rf $(BIN)
