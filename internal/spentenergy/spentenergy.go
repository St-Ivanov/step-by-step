package spentenergy

import (
	"fmt"
	"time"

	"github.com/St-Ivanov/step-by-step/internal/errors"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func isValidData(steps int, weight, height float64, duration time.Duration) error {
	if steps <= 0 {
		return fmt.Errorf("%w Exactly the steps.", errors.ErrNegativeValue)
	}
	if weight <= 0 {
		return fmt.Errorf("%w Exactly the weight.", errors.ErrNegativeValue)
	}
	if height <= 0 {
		return fmt.Errorf("%w Exactly the height.", errors.ErrNegativeValue)
	}
	if duration <= 0 {
		return fmt.Errorf("%w Exactly the time.", errors.ErrNegativeValue)
	}
	return nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	err := isValidData(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	averageSpeed := MeanSpeed(steps, height, duration)
	return (weight * averageSpeed * duration.Minutes()) / minInH * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	err := isValidData(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	averageSpeed := MeanSpeed(steps, height, duration)
	return (weight * averageSpeed * duration.Minutes()) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	return height * stepLengthCoefficient * float64(steps) / mInKm
}
