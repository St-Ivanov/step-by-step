package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

var (
	ErrIncSliceEl    = errors.New("Incorrect number of slice elements.")
	ErrConvToInt     = errors.New("Error converting a string to a number.")
	ErrNegativeValue = errors.New("Negative or zero number of value.")
	ErrConvToTime    = errors.New("Time conversion error.")
	ErrSportInput    = errors.New("неизвестный тип тренировки")
)

func parseTraining(data string) (int, string, time.Duration, error) {
	dataSlices := strings.Split(data, ",")
	if len(dataSlices) != 3 {
		return 0, "", 0, ErrIncSliceEl
	}
	steps, err := strconv.Atoi(dataSlices[0])
	if err != nil {
		return 0, "", 0, ErrConvToInt
	}
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("%w Exactly the steps.", ErrNegativeValue)
	}
	walkingTime, err := time.ParseDuration(dataSlices[2])
	if err != nil {
		return 0, "", 0, ErrConvToTime
	}
	if walkingTime <= 0 {
		return 0, "", 0, fmt.Errorf("%w Exactly the time.", ErrNegativeValue)
	}
	return steps, dataSlices[1], walkingTime, nil
}

func distance(steps int, height float64) float64 {
	return height * stepLengthCoefficient * float64(steps) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distanceWalk := distance(steps, height)
	return distanceWalk / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, typeSport, walkingTime, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	averageSpeed := meanSpeed(steps, height, walkingTime)
	walkingDistance := distance(steps, height)
	var calories float64
	switch typeSport {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, walkingTime)
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, walkingTime)
	default:
		return "", ErrSportInput
	}
	if err != nil {
		return "", err
	}
	str := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeSport, walkingTime.Hours(), walkingDistance, averageSpeed, calories)
	return str, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("%w Exactly the steps.", ErrNegativeValue)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("%w Exactly the weight.", ErrNegativeValue)
	}
	if height <= 0 {
		return 0, fmt.Errorf("%w Exactly the height.", ErrNegativeValue)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("%w Exactly the time.", ErrNegativeValue)
	}
	averageSpeed := meanSpeed(steps, height, duration)
	calories := (weight * averageSpeed * duration.Minutes()) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("%w Exactly the steps.", ErrNegativeValue)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("%w Exactly the weight.", ErrNegativeValue)
	}
	if height <= 0 {
		return 0, fmt.Errorf("%w Exactly the height.", ErrNegativeValue)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("%w Exactly the time.", ErrNegativeValue)
	}
	averageSpeed := meanSpeed(steps, height, duration)
	calories := (weight * averageSpeed * duration.Minutes()) / minInH * walkingCaloriesCoefficient
	return calories, nil
}
