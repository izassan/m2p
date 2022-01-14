build:
	@GOOS=linux GOARCH=amd64 go build -o bin/m2p m2p.go
	@GOOS=windows GOARCH=amd64 go build -o bin/m2p-windows.exe m2p.go
	@GOOS=darwin GOARCH=amd64 go build -o bin/m2p-darwin m2p.go
test:
	@go build m2p.go
	@./m2p testdata/95479584.zip testdata/output

clean:
	@rm -r bin
