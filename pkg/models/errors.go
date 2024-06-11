package models

import "errors"

var (
	ErrorUnvalidDescription = errors.New("description is not valid")
	ErrorUnvalidTopic       = errors.New("topic is not valid")
	ErrTodoNotFound         = errors.New("todo not found")
)
