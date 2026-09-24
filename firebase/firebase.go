package firebase

import (
	"context"
	"fmt"

	firebaseSDK "firebase.google.com/go/v4"
	"cloud.google.com/go/firestore"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type Firebase struct {
	App       *firebaseSDK.App
	Firestore *firestore.Client
	Auth      *auth.Client
}

func Initialize() (*Firebase, error) {
	ctx := context.Background()

	app, err := firebaseSDK.NewApp(
		ctx,
		nil,
		option.WithCredentialsFile(
			"firebase-service-account.json",
		),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"gagal initialize Firebase: %w",
			err,
		)
	}

	// ========================================
	// FIRESTORE
	// ========================================

	firestoreClient, err :=
		app.Firestore(ctx)

	if err != nil {
		return nil, fmt.Errorf(
			"gagal membuat Firestore client: %w",
			err,
		)
	}

	// ========================================
	// FIREBASE AUTH
	// ========================================

	authClient, err :=
		app.Auth(ctx)

	if err != nil {
		firestoreClient.Close()

		return nil, fmt.Errorf(
			"gagal membuat Firebase Auth client: %w",
			err,
		)
	}

	return &Firebase{
		App:       app,
		Firestore: firestoreClient,
		Auth:      authClient,
	}, nil
}