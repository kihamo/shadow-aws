package instance

import (
	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow-aws/components/aws/internal"
)

func NewComponent() shadow.Component {
	return &internal.Component{}
}
