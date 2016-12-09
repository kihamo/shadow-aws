package aws

import (
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/resource/config"
	"github.com/kihamo/shadow/resource/logger"
)

type Resource struct {
	application *shadow.Application
	awsConfig   *aws.Config
	config      *config.Resource
	logger      logger.Logger

	mutex    sync.RWMutex
	services map[string]interface{}
}

type AwsArnParse struct {
	Arn          string
	Partition    string
	Service      string
	Region       string
	Account      string
	Resource     string
	ResourceType string
}

func (r *Resource) GetName() string {
	return "aws"
}

func (r *Resource) Init(a *shadow.Application) error {
	r.application = a
	r.services = map[string]interface{}{}

	resourceConfig, err := a.GetResource("config")
	if err != nil {
		return err
	}

	r.config = resourceConfig.(*config.Resource)

	return nil
}

func (r *Resource) Run() error {
	if resourceLogger, err := r.application.GetResource("logger"); err == nil {
		r.logger = resourceLogger.(*logger.Resource).Get(r.GetName())
	} else {
		r.logger = logger.NopLogger
	}

	r.awsConfig = aws.NewConfig().
		WithCredentials(credentials.NewStaticCredentials(r.config.GetString("aws.key"), r.config.GetString("aws.secret"), "")).
		WithRegion(r.config.GetString("aws.region"))

	if r.config.GetBool("debug") {
		r.awsConfig.WithLogLevel(aws.LogDebug)
	}

	fields := map[string]interface{}{
		"region": *r.awsConfig.Region,
	}

	credentials, err := r.awsConfig.Credentials.Get()
	if err == nil {
		fields["key"] = credentials.AccessKeyID
		fields["secret"] = credentials.SecretAccessKey
	}

	r.logger.Info("Connect AWS", fields)

	return nil
}

func (r *Resource) GetSNS() *sns.SNS {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.services["sns"]; !ok {
		r.services["sns"] = sns.New(session.New(r.awsConfig))
	}

	return r.services["sns"].(*sns.SNS)
}

func (r *Resource) ParseArn(arn string) *AwsArnParse {
	// http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-arns

	parts := strings.Split(arn, ":")
	result := AwsArnParse{
		Arn:       parts[0],
		Partition: parts[1],
		Service:   parts[2],
		Region:    parts[3],
		Account:   parts[4],
	}

	if len(parts) > 6 {
		result.Resource = parts[5]
		result.ResourceType = parts[6]
	} else {
		path := strings.Split(parts[5], "/")

		result.Resource = path[0]
		result.ResourceType = strings.Join(path[1:], "/")
	}

	return &result
}

func (r *Resource) GetServices() map[string]interface{} {
	r.mutex.RLock()
	r.mutex.RUnlock()

	return r.services
}
