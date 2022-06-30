// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package scripts

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// lint:ignore unused
func getCredentialsUsingMFA() {
	// lint:ignore forbidigo
	fmt.Println("Please setup MFA_TOKEN, SERIAL_NUMBER, AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY first.")
	mfaToken := "123456"
	serialNumber := "arn:aws:iam::654321:mfa/test@test.com"

	// access key id and secret access key of your IAM user.
	accessKeyID := "FAKE-ACCESS-KEY-ID"
	// lint:ignore gosec
	secretAccessKey := "FAKE-SECRET-ACCESS-KEY"

	os.Setenv("AWS_ACCESS_KEY_ID", accessKeyID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", secretAccessKey)
	cfg, err := awsconfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("failed to load config: ", err.Error())
	}

	stsSvc := sts.NewFromConfig(cfg)
	durationLongest := int32(129600)
	getSessionTokenInput := sts.GetSessionTokenInput{
		DurationSeconds: &durationLongest,
		SerialNumber:    aws.String(serialNumber),
		TokenCode:       aws.String(mfaToken),
	}

	tempToken, err := stsSvc.GetSessionToken(context.TODO(), &getSessionTokenInput)
	if err != nil {
		fmt.Println("GetSessionToken failed: ", err)
	}

	fmt.Println("temp aws_access_key_id =", *tempToken.Credentials.AccessKeyId)
	fmt.Println("temp aws_secret_access_key =", *tempToken.Credentials.SecretAccessKey)
	fmt.Println("temp aws_session_token =", *tempToken.Credentials.SessionToken)
}
