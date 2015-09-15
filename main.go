package aws // import "github.com/kihamo/shadow-aws"

//go:generate goimports -w ./
//go:generate sh -c "cd service && go-bindata-assetfs -pkg=service templates/..."
