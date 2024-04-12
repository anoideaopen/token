package model

import "encoding"

// ObjectQuery represents a request that helps to load object.
type ObjectQuery string

// Validator is an interface designed to represent structures that can check their
// fields.
type Validator interface {
	// Validate should check the object's internal state for consistency or
	// other correctness criteria. It should return an error if the validation fails.
	Validate() error
}

// Object is an interface designed to represent objects that can be
// marshaled, unmarshaled, and cloned. It must be implemented by all objects
// that need to interact and store their state in the repository.
type Object interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
	Validator

	// Clone should create an exact copy of the object, located in a different
	// memory location from the original. This is necessary for cache
	// optimization to avoid marshalling the object in some cases. Clone is also
	// used as a template for finding an object in the repository.
	Clone() Object
}
