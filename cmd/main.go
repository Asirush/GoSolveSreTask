package main

import (
	"context"
	"os"
	"strconv"
	"strings"

	autoscaling "gosolve-sre-task/packages/aws/autoscaling"
	ec2 "gosolve-sre-task/packages/aws/ec2"
	elb "gosolve-sre-task/packages/aws/elb"

	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	// Load the AWS SDK configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	autoScalingGroupName := getEnv("AUTO_SCALING_GROUP_NAME", "my-auto-scaling-group")
	launchConfigurationName := getEnv("LAUNCH_CONFIGURATION_NAME", "my-launch-configuration")
	minSize := getEnvAsInt32("MIN_SIZE", 1)
	maxSize := getEnvAsInt32("MAX_SIZE", 3)
	desiredCapacity := getEnvAsInt32("DESIRED_CAPACITY", 2)
	targetGroupARNs := getEnvAsSlice("TARGET_GROUP_ARNS", []string{}, ",")
	vpcZoneIdentifier := getEnv("VPC_ZONE_IDENTIFIER", "subnet-xxxx")
	loadBalancerName := getEnv("LOAD_BALANCER_NAME", "my-application-load-balancer")
	vpcID := getEnv("VPC_ID", "vpc-xxxxxxx")
	subnetIDs := getEnvAsSlice("SUBNET_IDS", []string{"subnet-xxxx", "subnet-yyyy"}, ",")

	// Assuming LaunchInstances is correctly implemented in your ec2 package.
	// Add necessary parameters as required by your implementation.
	instanceID, err := ec2.LaunchInstances(cfg /*, additional parameters */)
	if err != nil {
		panic("error launching EC2 instances, " + err.Error())
	}
	println("Launched EC2 Instance: ", instanceID)

	// Setup ELB with corrected parameters
	lbArn, err := elb.CreateApplicationLoadBalancer(cfg, loadBalancerName, vpcID, subnetIDs)
	if err != nil {
		panic("error setting up ELB, " + err.Error())
	}
	println("Created Load Balancer with ARN: ", lbArn)

	// Create an Auto Scaling group
	asgName, err := autoscaling.CreateAutoScalingGroup(cfg, autoScalingGroupName, launchConfigurationName, minSize, maxSize, desiredCapacity, targetGroupARNs, vpcZoneIdentifier)
	if err != nil {
		panic("error configuring auto scaling, " + err.Error())
	}
	println("Created Auto Scaling Group: ", asgName)
}

// functions for env variables parcing
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsSlice(key string, fallback []string, separator string) []string {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return fallback
	}
	return strings.Split(valueStr, separator)
}

func getEnvAsInt32(key string, fallback int32) int32 {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseInt(valueStr, 10, 32); err == nil {
		return int32(value)
	}
	return fallback
}
