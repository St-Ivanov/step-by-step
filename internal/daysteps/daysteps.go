package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/St-Ivanov/step-by-step/internal/spentcalories"
)

var (
	ErrIncSliceEl    = errors.New("Incorrect number of slice elements.")
	ErrConvToInt     = errors.New("Error converting a string to a number.")
	ErrNegativeValue = errors.New("Negative or zero number of value.")
	ErrConvToTime    = errors.New("Time conversion error.")
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	dataSlice := strings.Split(data, ",")
	if len(dataSlice) != 2 {
		return 0, 0, ErrIncSliceEl
	}
	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		log.Println(ErrConvToInt)
		return 0, 0, ErrConvToInt
	}
	if steps <= 0 {
		log.Println(fmt.Errorf("%w Exactly the steps.", ErrNegativeValue))
		return 0, 0, fmt.Errorf("%w Exactly the steps.", ErrNegativeValue)
	}
	walkingTime, err := time.ParseDuration(dataSlice[1])
	if err != nil {
		log.Println(ErrConvToTime)
		return 0, 0, ErrConvToTime
	}
	if walkingTime <= 0 {
		log.Println(fmt.Errorf("%w Exactly the time.", ErrNegativeValue))
		return 0, 0, fmt.Errorf("%w Exactly the time.", ErrNegativeValue)
	}
	return steps, walkingTime, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, walkingTime, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	distance := float64(steps) * stepLength / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, walkingTime)
	if err != nil {
		log.Println(err)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)
}
