package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	if datastring == "" {
		err = errors.New("пустой ввод")
		return
	}

	sliceData := strings.Split(datastring, ",")
	if len(sliceData) != 2 {
		err = errors.New("неправильное количество параметров")
		return
	}

	stepsStr := sliceData[0]
	trimmedSteps := strings.TrimSpace(stepsStr)
	if trimmedSteps != stepsStr {
		err = errors.New("неверный формат количества шагов")
		return
	}

	stepsStr = strings.TrimSpace(sliceData[0])
	if stepsStr == "" || stepsStr == "-" || stepsStr == "+" {
		err = errors.New("неверный формат количества шагов")
		return
	}

	steps, parseErr := strconv.Atoi(stepsStr)
	if parseErr != nil || steps <= 0 {
		err = errors.New("неверный формат количества шагов")
		return
	}
	ds.Steps = steps

	durationStr := strings.TrimSpace(sliceData[1])
	if durationStr == "" {
		err = errors.New("пустая продолжительность")
		return
	}

	if strings.Contains(durationStr, "-") {
		err = errors.New("отрицательная продолжительность")
		return
	}

	if strings.Contains(durationStr, " ") {
		err = errors.New("неверный формат продолжительности")
		return
	}

	duration, parseErr := time.ParseDuration(durationStr)
	if parseErr != nil || duration <= 0 {
		err = errors.New("неверный формат продолжительности")
		return
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", errors.New("продолжительность должна быть положительной")
	}

	if ds.Steps <= 0 {
		return "", errors.New("количество шагов должно быть больше нуля")
	}
	if ds.Weight <= 0 {
		return "", errors.New("вес должен быть больше нуля")
	}
	if ds.Height <= 0 {
		return "", errors.New("рост должен быть больше нуля")
	}

	dist := spentenergy.Distance(ds.Steps, ds.Height)
	cal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	info := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, cal)
	return info, nil
}
