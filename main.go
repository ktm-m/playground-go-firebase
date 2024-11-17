package main

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	firebaseOpt := option.WithCredentialsFile("./config/firebase_service_account_key.json")
	app, err := firebase.NewApp(context.Background(), nil, firebaseOpt)
	if err != nil {
		panic(err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		panic(err)
	}
	defer client.Close()
}
