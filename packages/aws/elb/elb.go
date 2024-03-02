package elb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
)

// InitializeAWSConfig initializes and returns an AWS config
func InitializeAWSConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	return cfg
}

// CreateApplicationLoadBalancer creates an Application Load Balancer and returns its ARN
func CreateApplicationLoadBalancer(cfg aws.Config, name, vpcID string, subnetIDs []string) (string, error) {
	svc := elasticloadbalancingv2.NewFromConfig(cfg)

	input := &elasticloadbalancingv2.CreateLoadBalancerInput{
		Name:          aws.String(name),
		Subnets:       subnetIDs,
		Scheme:        types.LoadBalancerSchemeEnumInternetFacing,
		Type:          types.LoadBalancerTypeEnumApplication,
		IpAddressType: types.IpAddressTypeIpv4,
	}

	result, err := svc.CreateLoadBalancer(context.TODO(), input)
	if err != nil {
		return "", err
	}
	if len(result.LoadBalancers) == 0 {
		return "", fmt.Errorf("no load balancer created")
	}
	return *result.LoadBalancers[0].LoadBalancerArn, nil
}

// CreateTargetGroup creates a target group for the specified VPC and returns its ARN
func CreateTargetGroup(cfg aws.Config, name, vpcID string) (string, error) {
	svc := elasticloadbalancingv2.NewFromConfig(cfg)

	input := &elasticloadbalancingv2.CreateTargetGroupInput{
		Name:       aws.String(name),
		Protocol:   types.ProtocolEnumHttp, // Corrected to enum type
		Port:       aws.Int32(80),
		VpcId:      aws.String(vpcID),
		TargetType: types.TargetTypeEnumInstance, // Corrected to enum type, use TargetTypeEnumIp for IP addresses
	}

	result, err := svc.CreateTargetGroup(context.TODO(), input)
	if err != nil {
		return "", err
	}
	if len(result.TargetGroups) == 0 {
		return "", fmt.Errorf("no target group created")
	}
	return *result.TargetGroups[0].TargetGroupArn, nil
}

// CreateListener creates a listener for the specified load balancer and target group
func CreateListener(cfg aws.Config, loadBalancerArn, targetGroupArn string) (string, error) {
	svc := elasticloadbalancingv2.NewFromConfig(cfg)

	input := &elasticloadbalancingv2.CreateListenerInput{
		DefaultActions: []types.Action{
			{
				Type:           types.ActionTypeEnumForward, // Corrected to enum type
				TargetGroupArn: aws.String(targetGroupArn),
			},
		},
		LoadBalancerArn: aws.String(loadBalancerArn),
		Port:            aws.Int32(80),
		Protocol:        types.ProtocolEnumHttp, // Corrected to enum type
	}

	result, err := svc.CreateListener(context.TODO(), input)
	if err != nil {
		return "", err
	}
	if len(result.Listeners) == 0 {
		return "", fmt.Errorf("no listener created")
	}
	return *result.Listeners[0].ListenerArn, nil
}
