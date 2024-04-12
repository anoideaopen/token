package model

import (
	"encoding/json"

	"github.com/jinzhu/copier"
)

// Notification это базовая структура, которая определяет уведомление, которое необходимо
// доставить пользователям и/или сохранить в базе данных.
// В структуре содержится уникальный идентификатор записи, а также тип уведомления.
// структура является шаблонной для отправки уведомлений произвольного типа, но
// compile-time проверкой совпадения типов.
type Notification[T any] struct {
	ID   string `validate:"required"` // Идентификатор уведомления.
	Type string `validate:"required"` // Тип уведомления.
	Body T      `validate:"required"` // Тело уведомления.
}

// Реализация интерфейса model.Object.

func (n *Notification[T]) MarshalBinary() (data []byte, err error) {
	return json.Marshal(n)
}

func (n *Notification[T]) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, n)
}

func (n *Notification[T]) Clone() Object {
	nt := new(Notification[T])
	_ = copier.Copy(nt, n)
	return nt
}

func (n *Notification[T]) Validate() error {
	if validator, ok := any(n.Body).(Validator); ok {
		if err := validator.Validate(); err != nil {
			return err
		}
	}

	return NewValidator().Struct(n)
}

// -----------------------------------
