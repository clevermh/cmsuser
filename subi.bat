git add .
git commit -m "ultimo commit"
git push
set GOOS= LINUX
set GOARCH = amd64
go build main.go
del main.zip
tar.exe -a -tf main.zip main