package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/anoideaopen/token/keyvalue"
	"github.com/anoideaopen/token/model"
)

// ErrNotificationDatabase represents a generic error related to the database operations.
var ErrNotificationDatabase = errors.New("notofocation database error")

// Notification is a structure which encapsulates the keyvalue.DB to interact with
// notification structures in database.
//
//go:generate ifacemaker -f notification.go -o repository/notification.go -i Notification -s Notification -p repository -y "Repository describes methods, implemented by the storage package."
//go:generate mockgen -package mock -source repository/notification.go -destination repository/mock/mock_notification.go
type Notification struct {
	Object
}

// SaveBalancesUpdate stores notification record to the notification database.
func (n *Notification) SaveBalancesUpdate(
	ctx context.Context,
	bu model.Notification[model.BalancesUpdate],
) error {
	if err := n.Object.Save(ctx, model.ObjectQuery(
		keyvalue.Join(bu.Type, bu.ID),
	), &bu); err != nil {
		return fmt.Errorf("%w: %s", ErrNotificationDatabase, err)
	}

	return nil
}
