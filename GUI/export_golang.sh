export GOMODCACHE=/goinfre/$USER/go/pkg/mod
export GOCACHE=/goinfre/$USER/.cache/go-build
export GOBIN=/goinfre/$USER/go/bin
export GOPATH=/goinfre/$USER/go

mkdir -p $GOMODCACHE $GOCACHE $GOBIN

go clean -modcache
go get fyne.io/fyne/v2@latest
go install fyne.io/tools/cmd/fyne@latest
