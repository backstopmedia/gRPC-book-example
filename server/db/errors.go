package db

import "fmt"

// NotFoundErr is returned when a certain resource cannot be found
type NotFoundErr struct {
	// ResourceType should be assigned to the name of the resource type that was
	// missing, for example: starship or vehicle
	ResourceType string

	// Resource ID that could not be found
	ID string
}

// NotFound returns an initialized NotFoundErr type
func NotFound(rt, id string) NotFoundErr {
	return NotFoundErr{ResourceType: rt, ID: id}
}

// Error implements error
func (err NotFoundErr) Error() string {
	return fmt.Sprintf("%s with id '%s' not found", err.ResourceType, err.ID)
}
