VERSION = 0.1.0
BUILD = go build -o
SRC = m2p.go


build_all: release_linux release_windows release_darwin

build_linux:
	@GOOS=linux GOARCH=amd64 $(BUILD) bin/m2p $(SRC)

build_windows:
	@GOOS=windows GOARCH=amd64 $(BUILD) bin/m2p.exe $(SRC)

build_darwin:
	@GOOS=darwin GOARCH=amd64 $(BUILD) bin/m2p $(SRC)

release_linux: build_linux
	@cd bin && tar czvf m2p-$(VERSION)-linux_amd64.tar.gz m2p

release_windows: build_windows
	@zip -j bin/m2p-$(VERSION)-windows_amd64.zip bin/m2p.exe

release_darwin: build_darwin
	@cd bin && tar czvf m2p-$(VERSION)-darwin_amd64.tar.gz m2p

test:
	@go build m2p.go
	@./m2p testdata/95479584.zip testdata/output

test_name:
	@go build m2p.go
	@./m2p testdata/name_test testdata/output

clean:
	@rm -r bin
