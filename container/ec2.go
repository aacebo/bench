package container

import (
	"bench/logger"
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var client *ec2.Client
var log = logger.New("bench:container")

func EC2() *ec2.Client {
	if client == nil {
		cfg, err := config.LoadDefaultConfig(context.Background())

		if err != nil {
			log.Error(err.Error())
		}

		client = ec2.NewFromConfig(cfg)
	}

	return client
}
