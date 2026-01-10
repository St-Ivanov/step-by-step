package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/St-Ivanov/step-by-step/internal/errors"
	"github.com/St-Ivanov/step-by-step/internal/personaldata"
	"github.com/St-Ivanov/step-by-step/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return errors.ErrIncDataEnt
	}
	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return errors.ErrConvToInt
	}
	if steps <= 0 {
		return fmt.Errorf("%w Exactly the steps.", errors.ErrNegativeValue)
	}
	duration, err := time.ParseDuration(data[2])
	if err != nil {
		return errors.ErrConvToTime
	}
	if duration <= 0 {
		return fmt.Errorf("%w Exactly the time.", errors.ErrNegativeValue)
	}
	t.Steps = steps
	t.TrainingType = data[1]
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	averageSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var (
		calories float64
		err      error
	)
	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.ErrIncUndTraining
	}
	s := fmt.Sprintf(`Тип тренировки: %s
Длительность: %0.2f ч.
Дистанция: %0.2f км.
Скорость: %0.2f км/ч
Сожгли калорий: %0.2f
`, t.TrainingType, t.Duration.Hours(), distance, averageSpeed, calories)
	return s, nil
}
