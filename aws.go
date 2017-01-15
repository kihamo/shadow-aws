package aws // import "github.com/kihamo/shadow-aws"

//go:generate goimports -w ./
//go:generate sh -c "cd components/aws && go-bindata-assetfs -pkg=aws templates/..."
