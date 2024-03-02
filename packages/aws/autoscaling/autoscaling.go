package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
)

func CreateAutoScalingGroup(cfg aws.Config, autoScalingGroupName, launchConfigurationName string, minSize, maxSize, desiredCapacity int32, targetGroupARNs []string, vpcZoneIdentifier string) (string, error) {
	svc := autoscaling.NewFromConfig(cfg)

	input := &autoscaling.CreateAutoScalingGroupInput{
		AutoScalingGroupName:    aws.String(autoScalingGroupName),
		LaunchConfigurationName: aws.String(launchConfigurationName),
		MinSize:                 aws.Int32(minSize),
		MaxSize:                 aws.Int32(maxSize),
		DesiredCapacity:         aws.Int32(desiredCapacity),
		VPCZoneIdentifier:       aws.String(vpcZoneIdentifier),
		TargetGroupARNs:         targetGroupARNs, // Include target group ARNs
	}

	_, err := svc.CreateAutoScalingGroup(context.TODO(), input)
	if err != nil {
		return "", err
	}

	// Auto Scaling group created successfully, return its name
	return autoScalingGroupName, nil
}
