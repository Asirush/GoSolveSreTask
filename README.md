## Installation

To run the application, use the following command:

```bash
docker-compose up
```

If you need to add or edit environment variables for the application, do so in the `docker-compose.yaml` file.

## Description

In this application, I demonstrate how to use the AWS SDK for Golang to deploy and manage infrastructure. I've implemented a few simple methods and considered the application's launch in advance, including a Dockerfile for building and running the project, as well as a docker-compose file to manage environment variables more easily.

## Task
Please find below a task. I would like to ask you send it back within 3 days.

# Intro
This assignment is simple and to the point and should only take and hour or two. We are not looking for a fully functional program, but instead a small sample of code and discussion of approach to get a sense of how you write software and think about infrastructure problems.

# Assignment
Please design a program to do a zero downtime deploy using standard EC2 and related resources such as AMI, ELB, etc. using an AWS SDK or similar package. Please do not use Terraform, Cloudformation, Kubernetes or other such tool.

# Requirements
* Pseudeo code for the full program.
* Full directory structure as if this was indeed a full implementation as a part of a larger application to perform many infrastructure related tasks.
* Discussion of approach.
* A full implementation of one or two interesting parts of the program including tests such as those which require mocking out multiple external calls.

# Notes
* Any programming language can be used.