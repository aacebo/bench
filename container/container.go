package container

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/google/uuid"
)

var imageId = "ami-0bd4d695347c0ef88"
var min int32 = 1
var max int32 = 1

type Container struct {
	ID         *string                 `json:"id"`
	InstanceID *string                 `json:"instance_id"`
	Type       types.InstanceType      `json:"type"`
	State      types.InstanceStateName `json:"state"`
	StartedAt  *time.Time              `json:"started_at"`
}

func New(lang string, problem string) (*Container, error) {
	id := uuid.NewString()
	client := EC2()
	out, err := client.RunInstances(context.Background(), &ec2.RunInstancesInput{
		ImageId:      &imageId,
		InstanceType: types.InstanceTypeT2Medium,
		MinCount:     &min,
		MaxCount:     &max,
		UserData:     userData(id, lang, problem),
	})

	if err != nil {
		return nil, err
	}

	instance := out.Instances[0]
	state := instance.State.Name

	for state == types.InstanceStateNamePending {
		status, err := client.DescribeInstanceStatus(context.Background(), &ec2.DescribeInstanceStatusInput{
			InstanceIds: []string{*instance.InstanceId},
		})

		if err != nil {
			return nil, err
		}

		if len(status.InstanceStatuses) == 1 {
			state = status.InstanceStatuses[0].InstanceState.Name
		}
	}

	log.Info(instance.MetadataOptions)

	return &Container{
		ID:         &id,
		InstanceID: instance.InstanceId,
		Type:       instance.InstanceType,
		State:      state,
		StartedAt:  instance.LaunchTime,
	}, nil
}

func (self *Container) Destroy() error {
	client := EC2()
	_, err := client.TerminateInstances(context.Background(), &ec2.TerminateInstancesInput{
		InstanceIds: []string{*self.InstanceID},
	})

	return err
}
