package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

var (
	ErrIncSliceEl    = errors.New("Incorrect number of slice elements.")
	ErrConvToInt     = errors.New("Error converting a string to a number.")
	ErrNegativeSteps = errors.New("Negative or zero number of steps.")
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
		return 0, 0, ErrConvToInt
	}
	if steps <= 0 {
		return 0, 0, ErrNegativeSteps
	}
	walkingTime, err := time.ParseDuration(dataSlice[1])
	if err != nil {
		return 0, 0, ErrConvToTime
	}
	return steps, walkingTime, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, walkingTime, err := parsePackage(data)
	if err != nil {
		return ""
	}
	distance := float64(steps) * stepLength / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, walkingTime)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)
}
