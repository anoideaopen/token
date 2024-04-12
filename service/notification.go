package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/anoideaopen/token/model"
	"github.com/anoideaopen/token/storage/repository"
)

var (
	// ErrNotificationDatabase определяет ошибку записи данных в репозитарий сервиса.
	ErrNotificationDatabase = errors.New("notification database error")

	// ErrNotificationValidation сигнализирует о попытке записи уведомления в репозитарий
	// сервиса, которое не прошло валидацию полей.
	ErrNotificationValidation = errors.New("invalid notification validation")
)

// Notification отвечает за работу с различными бухгалтерскими структурами. Он сохраняет или
// уведомляет потребителей относительно важных операций по изменению балансов и другой информации,
// необходимой регуляторам для построения отчетов. Сервисы, которым требуется строгая отчетность,
// должны отправлять эти данные в "бухгалтерию", а Notification, в свою очередь, надежно сохранять
// данные в хранилище.
//
//go:generate ifacemaker -f notification.go -o controller/notification.go -i Notification -s Notification -p controller -y "Controller describes methods, implemented by the service package."
//go:generate mockgen -package mock -source controller/notification.go -destination controller/mock/mock_notification.go
type Notification struct {
	repository.Notification
}

// NotifyBalancesUpdate добавляет новую запись в бухгалтерскую книгу, о движении средств
// пользователя или пользователей. Записи валидируются перед сохранением.
func (n *Notification) NotifyBalancesUpdate(
	ctx context.Context,
	bu model.Notification[model.BalancesUpdate],
) error {
	if err := bu.Validate(); err != nil {
		return fmt.Errorf("%w: %s", ErrNotificationValidation, err.Error())
	}

	if err := n.Notification.SaveBalancesUpdate(ctx, bu); err != nil {
		return fmt.Errorf("%w: %s", ErrNotificationDatabase, err.Error())
	}

	return nil
}
