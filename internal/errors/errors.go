package errors

import "errors"

var (
	ErrIncDataEnt     = errors.New("Incorrect data entry.")
	ErrNegativeValue  = errors.New("Negative or zero number of value.")
	ErrConvToInt      = errors.New("Error converting a string to a number.")
	ErrConvToTime     = errors.New("Time conversion error.")
	ErrIncUndTraining = errors.New("неизвестный тип тренировки")
)
