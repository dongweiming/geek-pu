package models

import "fmt"

type GeekError struct {
	Msg string
}

func (e GeekError) Error() string {
	return fmt.Sprintf("%v: %v", e.Msg)
}
