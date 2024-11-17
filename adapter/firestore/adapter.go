package firestore

import (
	"cloud.google.com/go/firestore"
	"github.com/ktm-m/playground-go-firebase/internal/port/outbound"
)

type adapter struct {
	client *firestore.Client
}

func NewAdapter(client *firestore.Client) outbound.FireStorePort {
	return &adapter{
		client: client,
	}
}
