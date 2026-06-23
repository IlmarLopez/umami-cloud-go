# Umami Cloud Go

Infrastructure as Code (IaC) project built with AWS CDK and Go to automate the deployment of a self-hosted Umami Analytics instance on AWS.

## Overview

Umami Cloud Go provisions cloud infrastructure and deploys Umami Analytics using AWS CDK, Amazon EC2, Docker, and GitHub Actions.

The project follows Infrastructure as Code principles, enabling reproducible deployments across multiple environments while integrating automated provisioning, secret management, and CI/CD workflows.

## Features

* AWS CDK application written in Go
* Automated VPC provisioning
* EC2 instance deployment
* Automated Docker installation
* Automated Umami deployment via Docker Compose
* AWS Secrets Manager integration
* Environment-specific configurations
* CI/CD pipelines with GitHub Actions
* Separate Development and Production environments

## Architecture

```text
Internet
    │
    ▼
Amazon EC2
    │
    ├── Docker
    ├── Docker Compose
    └── Umami Analytics

AWS Secrets Manager
    │
    └── Database and application credentials

Amazon VPC
    └── Public Subnets
```

## Infrastructure Components

### Networking

The project provisions:

* Amazon VPC
* Public Subnets across multiple Availability Zones
* Security Group rules for application and SSH access

### Compute

The infrastructure deploys:

* Amazon EC2 instances running Amazon Linux 2023
* User Data bootstrap scripts for automated configuration
* Environment-based instance sizing

### Security

* AWS Secrets Manager for credential storage
* IAM permissions for secure secret retrieval
* No hardcoded application secrets

### Application Deployment

During instance initialization:

1. Docker is installed automatically.
2. Docker Compose is installed.
3. The Umami repository is cloned.
4. Secrets are retrieved from AWS Secrets Manager.
5. Environment variables are generated dynamically.
6. Umami is deployed through Docker Compose.

## Project Structure

```text
.
├── .github/
│   └── workflows/
│       ├── deploy-dev.yml
│       └── deploy-prod.yml
│
├── cdk/
│   ├── stack.go
│   ├── network.go
│   ├── compute.go
│   └── config/
│       ├── dev.json
│       └── prod.json
│
├── scripts/
│   └── deploy-umami.sh
│
├── cdk.json
├── go.mod
└── README.md
```

## Prerequisites

* Go 1.26+
* Node.js LTS
* AWS CLI
* AWS CDK
* Docker

## Installation

Install Go dependencies:

```bash
go mod download
```

Install AWS CDK:

```bash
npm install -g aws-cdk
```

## Deployment

Generate CloudFormation templates:

```bash
cdk synth
```

Review infrastructure changes:

```bash
cdk diff
```

Deploy the infrastructure:

```bash
cdk deploy
```

## CI/CD

The project includes automated deployment workflows through GitHub Actions:

* Development environment deployment
* Production environment deployment

## Technologies Used

* Go
* AWS CDK
* Amazon EC2
* Amazon VPC
* AWS Secrets Manager
* Docker
* Docker Compose
* GitHub Actions
* CloudFormation

## Learning Objectives

This project was created to explore and apply:

* Infrastructure as Code (IaC)
* Cloud infrastructure automation
* AWS resource provisioning
* Secret management best practices
* CI/CD workflows
* Production-ready deployment processes

## Contributions

This project was developed collaboratively with Ilmar López.

* Eric Frenek López Rosales contributed to the implementation, infrastructure automation, deployment workflows, and feature development.
* Ilmar López provided technical guidance, architecture reviews, and project direction throughout the development process.

The collaboration focused on applying cloud engineering, Infrastructure as Code, and DevOps practices using AWS CDK and Go.

## License

This project is available for educational and portfolio purposes.
