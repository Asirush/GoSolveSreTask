package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// LaunchInstances launches EC2 instances based on provided configuration
func LaunchInstances(cfg aws.Config) (string, error) {
	// Initialize the EC2 service client
	svc := ec2.NewFromConfig(cfg)

	// Example: Define the input for RunInstances call
	input := &ec2.RunInstancesInput{
		ImageId:      aws.String("ami-xxxxxxxxxxxxxxxxx"), // Specify the AMI ID
		InstanceType: types.InstanceTypeT2Micro,           // Specify the instance type
		MinCount:     aws.Int32(1),                        // Minimum number of instances to launch
		MaxCount:     aws.Int32(1),                        // Maximum number of instances to launch
	}

	// Call RunInstances to launch an instance
	resp, err := svc.RunInstances(context.TODO(), input)
	if err != nil {
		// Handle the error
		return "", err
	}

	// Check if we got exactly one instance
	if len(resp.Instances) != 1 {
		return "", fmt.Errorf("expected exactly one instance, got %d", len(resp.Instances))
	}

	// Extract the instance ID
	instanceID := *resp.Instances[0].InstanceId

	// Instance launched successfully, return the instance ID
	return instanceID, nil
}
