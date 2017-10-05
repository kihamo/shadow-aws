package aws // import "github.com/kihamo/shadow-aws"

//go:generate goimports -w ./
//go:generate /bin/bash -c "cd components/aws/internal && go-bindata-assetfs -pkg=internal templates/..."
