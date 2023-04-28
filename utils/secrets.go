package utils

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"context"
	"fmt"
	"os"
)

const (
	awsSecretKeyID      = "AWS_SECRET_KEY_ID"
	awsSecretAccessKey  = "AWS_SECRET_ACCESS_KEY"
	spotifyClientId     = "SPOTIFY_CLIENT_ID"
	spotifyClientSecret = "SPOTIFY_CLIENT_SECRET"
)

func getClient() (*secretmanager.Client, error) {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret manager client: %v", err)
	}
	return client, nil
}

// getContext initializes a new context
func getContext() context.Context {
	return context.Background()
}

// getAWSCredentials retrieves the AWS credentials from Google Secrets Manager
func getAWSCredentials() (string, string, error) {
	ctx := getContext()
	client, err := getClient()
	if err != nil {
		return "", "", fmt.Errorf("failed to create secret manager client: %v", err)
	}
	defer client.Close()

	accessKeyIDResp, err := client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{Name: os.Getenv(awsSecretKeyID)})
	if err != nil {
		return "", "", fmt.Errorf("failed to retrieve AWS access key ID: %v", err)
	}

	secretAccessKeyResp, err := client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{Name: os.Getenv(awsSecretAccessKey)})
	if err != nil {
		return "", "", fmt.Errorf("failed to retrieve AWS secret access key: %v", err)
	}
	return string(accessKeyIDResp.Payload.Data), string(secretAccessKeyResp.Payload.Data), nil
}

func GetSpotifyCredentials() (string, string, error) {
	ctx := getContext()
	client, err := getClient()
	if err != nil {
		return "", "", fmt.Errorf("failed to create secret manager client: %v", err)
	}
	defer client.Close()

	spotifyClientIdResp, err := client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{Name: os.Getenv(spotifyClientId)})
	if err != nil {
		return "", "", fmt.Errorf("failed to retrieve AWS access key ID: %v", err)
	}

	spotifyClientSecretResp, err := client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{Name: os.Getenv(spotifyClientSecret)})
	if err != nil {
		return "", "", fmt.Errorf("failed to retrieve AWS secret access key: %v", err)
	}
	return string(spotifyClientIdResp.Payload.Data), string(spotifyClientSecretResp.Payload.Data), nil
}
