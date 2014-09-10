git fetch
git checkout %GIT_COMMIT%
git pull

go get github.com/jteeuwen/go-bindata/...
SET PATH=%PATH%;C:\Users\pivotal\Go-Agent\pipelines\Windows-Testing\bin
go-bindata -pkg resources -ignore ".go" -o cf/resources/i18n_resources.go cf/i18n/resources/... cf/i18n/test_fixtures/...

powershell -command set-executionpolicy remotesigned
powershell .\bin\replace-sha.ps1

SET GOPATH=C:\Users\pivotal\Go-Agent\pipelines\Windows-Testing;%CD%\Godeps\_workspace;c:\Users\Administrator\go
c:\Go\bin\go build -v -o cf-windows-amd64.exe ./main
c:\Go\bin\go test -i ./cf/... ./generic/... ./testhelpers/...
c:\Go\bin\go test -cover -v ./cf/... ./generic/... ./testhelpers/...
