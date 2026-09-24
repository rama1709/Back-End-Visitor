package firebase

import (
	"context"
	"log"
	"time"
)

type Notification struct {
	Type      string
	Title     string
	Message   string
	Read      bool
	CreatedAt time.Time
}

func (f *Firebase) CreateNotification(
	notification Notification,
) error {
	// ========================================
	// CEK FIREBASE CLIENT
	// ========================================

	if f == nil {
		return nil
	}

	if f.Firestore == nil {
		return nil
	}

	// ========================================
	// TIMEOUT FIREBASE
	// ========================================

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancel()

	// ========================================
	// SIMPAN NOTIFICATION
	// ========================================

	_, _, err := f.Firestore.
		Collection("notifications").
		Add(
			ctx,
			map[string]interface{}{
				"type": notification.Type,

				"title": notification.Title,

				"message": notification.Message,

				"read": notification.Read,

				"createdAt": notification.CreatedAt,
			},
		)

	if err != nil {
		log.Printf(
			"CreateNotification error: %v",
			err,
		)

		return err
	}

	log.Printf(
		"Firebase notification created: %s",
		notification.Title,
	)

	return nil
}
