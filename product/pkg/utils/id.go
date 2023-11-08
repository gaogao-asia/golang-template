package utils

import "github.com/lithammer/shortuuid"

// NewID use shortuuid to generate new id
func NewID() string {
	return shortuuid.New()
}
