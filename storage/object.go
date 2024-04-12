package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/anoideaopen/token/keyvalue"
	"github.com/anoideaopen/token/model"
)

// Errors related to object storage.
var (
	// ErrObjectDatabase represents a generic error related to the database operations.
	ErrObjectDatabase = errors.New("object database error")

	// ErrObjectNotFound is the error returned when an object is not found in the DB.
	ErrObjectNotFound = errors.New("object not found")

	// ErrObjectValidation is the error returned when an object fails to validate.
	ErrObjectValidation = errors.New("object validation failed")

	// ErrObjectMarshaling is the error returned when an error occurs during the marshaling
	// of an object (i.e., converting the object into a format suitable for storage).
	ErrObjectMarshaling = errors.New("object marshaling error")

	// ErrObjectUnmarshaling is the error returned when an error occurs during the unmarshaling
	// of an object (i.e., converting the stored data back into the object form).
	ErrObjectUnmarshaling = errors.New("object unmarshaling error")
)

// Object represents a generic object storage that interacts with a databse.
// It is designed to operate on objects implementing the Object interface from the model package.
//
//go:generate ifacemaker -f object.go -o repository/object.go -i Object -s Object -p repository -y "Repository describes methods, implemented by the storage package."
//go:generate mockgen -package mock -source repository/object.go -destination repository/mock/mock_object.go
type Object struct {
	keyvalue.DB
}

// Load attempts to load an Object from the database.
// The object is identified by the provided query.
func (o *Object) Load(ctx context.Context, q model.ObjectQuery, obj model.Object) error {
	raw, err := o.DB.Get(ctx, keyvalue.Key(q))
	if err != nil {
		if errors.Is(err, keyvalue.ErrNotFound) {
			return ErrObjectNotFound
		}

		return o.wrap(ErrObjectDatabase, err)
	}

	if err := obj.UnmarshalBinary(raw); err != nil {
		return o.wrap(ErrObjectUnmarshaling, err)
	}

	if err := obj.Validate(); err != nil {
		return o.wrap(ErrObjectValidation, err)
	}

	return nil
}

// Save attempts to save an Object object to the database.
// The object is identified by the provided query.
func (o *Object) Save(ctx context.Context, q model.ObjectQuery, obj model.Object) error {
	if err := obj.Validate(); err != nil {
		return o.wrap(ErrObjectValidation, err)
	}

	raw, err := obj.MarshalBinary()
	if err != nil {
		return o.wrap(ErrObjectMarshaling, err)
	}

	if err := o.DB.Set(ctx, keyvalue.Key(q), raw); err != nil {
		return o.wrap(ErrObjectDatabase, err)
	}

	return nil
}

// Delete attempts to delete an Object object from the database.
// The object is identified by the provided query.
func (o *Object) Delete(ctx context.Context, q model.ObjectQuery) error {
	if err := o.DB.Del(ctx, keyvalue.Key(q)); err != nil {
		return o.wrap(ErrObjectDatabase, err)
	}

	return nil
}

// Iter iterates over a set of objects from the database that match a
// given query. For each object that matches, it clones the template object, unmarshals
// the object data into the clone, validates the clone, and then passes the clone to
// the provided iterator function. Iteration stops if the iterator function returns false.
func (o *Object) Iter(
	ctx context.Context,
	q model.ObjectQuery,
	tmpl model.Object,
	cb func(obj model.Object) (stop bool),
) error {
	i, err := o.DB.Iter(ctx, keyvalue.Prefix(q))
	if err != nil {
		return o.wrap(ErrObjectDatabase, err)
	}
	defer i.Close()

	for i.HasNext() {
		_, v, err := i.Next()
		if err != nil {
			return o.wrap(ErrObjectDatabase, err)
		}

		obj := tmpl.Clone()
		if err := obj.UnmarshalBinary(v); err != nil {
			return o.wrap(ErrObjectUnmarshaling, err)
		}

		if err := obj.Validate(); err != nil {
			return o.wrap(ErrObjectValidation, err)
		}

		if cb(obj) { // cb returns stop == true
			break
		}
	}

	return nil
}

func (o *Object) wrap(err, cause error) error {
	return fmt.Errorf("%w: %s", err, cause.Error())
}
