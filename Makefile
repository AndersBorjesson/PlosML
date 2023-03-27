all : windows linux

linux : main.go
	GO_ENABLED=1  GOOS=linux go build .

windows : main.go
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows  go build --ldflags '-extldflags "-Wl,--allow-multiple-definition"'

native : main.go
	go build .