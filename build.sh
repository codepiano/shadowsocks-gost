env GOOS=darwin GOARCH=amd64 go build -o shadowsocks-gost.darwin.amd64
env GOOS=darwin GOARCH=arm64 go build -o shadowsocks-gost.darwin.arm64
env GOOS=linux GOARCH=amd64 go build -o shadowsocks-gost.linux.amd64
env GOOS=windows GOARCH=amd64 go build -o shadowsocks-gost.windows.amd64
