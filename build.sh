mkdir bin
mkdir bin/windows
mkdir bin/linux
mkdir bin/mac
GOOS=windows GOARCH=amd64 go build  -o bin/windows/tier1.exe ./tier1/main.go
GOOS=windows GOARCH=amd64 go build  -o bin/windows/tier2.exe ./tier2/main.go
GOOS=windows GOARCH=amd64 go build  -o bin/windows/tier3.exe ./tier3/main.go
GOOS=linux   GOARCH=amd64 go build  -o bin/linux/tier1 ./tier1/main.go
GOOS=linux   GOARCH=amd64 go build  -o bin/linux/tier2 ./tier2/main.go
GOOS=linux   GOARCH=amd64 go build  -o bin/linux/tier3 ./tier3/main.go
GOOS=darwin  GOARCH=amd64 go build  -o bin/mac/tier1 ./tier1/main.go
GOOS=darwin  GOARCH=amd64 go build  -o bin/mac/tier2 ./tier2/main.go
GOOS=darwin  GOARCH=amd64 go build  -o bin/mac/tier3 ./tier3/main.go

cd bin
zip traces_windows.zip -r windows
tar zcvf traces_mac.tar.gz mac/
tar zcvf traces_linux.tar.gz linux/
