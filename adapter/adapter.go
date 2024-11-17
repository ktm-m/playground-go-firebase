package adapter

import (
	"cloud.google.com/go/firestore"
	fireStoreAdapter "github.com/ktm-m/playground-go-firebase/adapter/firestore"
	"github.com/ktm-m/playground-go-firebase/internal/port/outbound"
)

type Adapter struct {
	FirebaseAdapter outbound.FireStorePort
}

func NewAdapter(client *firestore.Client) *Adapter {
	return &Adapter{
		FirebaseAdapter: fireStoreAdapter.NewAdapter(client),
	}
}
