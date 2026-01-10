package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/St-Ivanov/step-by-step/internal/errors"
	"github.com/St-Ivanov/step-by-step/internal/personaldata"
	"github.com/St-Ivanov/step-by-step/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 2 {
		return errors.ErrIncDataEnt
	}
	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return errors.ErrConvToInt
	}
	if steps <= 0 {
		return fmt.Errorf("%w Exactly the steps.", errors.ErrNegativeValue)
	}
	duration, err := time.ParseDuration(data[1])
	if err != nil {
		return errors.ErrConvToTime
	}
	if duration <= 0 {
		return fmt.Errorf("%w Exactly the time.", errors.ErrNegativeValue)
	}
	ds.Steps = steps
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(`Количество шагов: %d.
Дистанция составила %0.2f км.
Вы сожгли %0.2f ккал.
`, ds.Steps, distance, calories), nil
}
