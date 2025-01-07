package resources

import (
	"context"
	"sync"

	"cloud.google.com/go/firestore"
	_ "github.com/go-sql-driver/mysql"
)

var (
	firestoreInstance *firestore.Client
	once              sync.Once
)

// GetFirestoreClient ensures a single Firestore client
func GetFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	var err error
	once.Do(func() {
		firestoreInstance, err = firestore.NewClient(ctx, "project-id")
	})
	return firestoreInstance, err
}
