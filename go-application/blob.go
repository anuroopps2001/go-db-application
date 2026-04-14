package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

type BlobClient struct {
	client    *azblob.Client
	container string
	account   string
}

func NewBlobClient() (*BlobClient, error) {
	account := os.Getenv("AZURE_STORAGE_ACCOUNT")

	key := os.Getenv("AZURE_STORAGE_KEY")

	container := os.Getenv("AZURE_STORAGE_CONTAINER")

	if account == "" || key == "" || container == "" {
		return nil, fmt.Errorf("missing blob env vars")
	}

	cred, err := azblob.NewSharedKeyCredential(account, key)

	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://%s.blob.core.windows.net/", account)

	client, err := azblob.NewClientWithSharedKeyCredential(url, cred, nil)
	if err != nil {
		return nil, err
	}

	return &BlobClient{
		client:    client,
		container: container,
		account:   account,
	}, nil
}

func (b *BlobClient) Upload(ctx context.Context, fileName string, data []byte) (string, error) {

	_, err := b.client.UploadBuffer(ctx, b.container, fileName, data, nil)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.blob.core.windows.net/%s/%s",
		b.account,
		b.container,
		fileName,
	)

	return url, nil
}
