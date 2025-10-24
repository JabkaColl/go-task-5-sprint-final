package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {

	if datastring == "" {
		return errors.New("пустой ввод")
	}

	stringSlice := strings.Split(datastring, ",")
	if len(stringSlice) != 3 {
		return errors.New("неправильное количество параметров")
	}

	stepsStr := stringSlice[0]

	if stepsStr == "" || stepsStr == "-" || stepsStr == "+" {
		return errors.New("неверный формат количества шагов")
	}

	steps, err := strconv.Atoi(stepsStr)

	if err != nil {
		return errors.New("неверный формат количества шагов")
	}

	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}

	t.TrainingType = strings.TrimSpace(stringSlice[1])

	durationStr := strings.TrimSpace(stringSlice[2])

	if durationStr == "" {
		return errors.New("пустая продолжительность")
	}

	if strings.Contains(durationStr, "-") {
		return errors.New("отрицательная продолжительность")
	}

	if strings.Contains(durationStr, " ") {
		return errors.New("пробел между числом и единицей")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return errors.New("неверный формат продолжительности")
	}

	if duration <= 0 {
		return errors.New("продолжительность должна быть положительной")
	}

	t.Steps = steps
	t.Duration = duration

	return nil

}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(
			t.Steps,
			t.Personal.Weight,
			t.Personal.Height,
			t.Duration,
		)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(
			t.Steps,
			t.Personal.Weight,
			t.Personal.Height,
			t.Duration,
		)
	default:
		return "", errors.New("неизвестный тип тренировки: " + t.TrainingType)
	}

	if err != nil {
		return "", err
	}

	durationHours := t.Duration.Hours()

	info := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		durationHours,
		distance,
		speed,
		calories,
	)

	return info, nil
}
