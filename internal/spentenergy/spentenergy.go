package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть положительным числом")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть положительным числом")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть положительным числом")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть положительной")
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	caloriesCorrect := ((weight * meanSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient

	return caloriesCorrect, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть положительным числом")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть положительным числом")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть положительным числом")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть положительной")
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories, nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 {
		return 0
	}

	if duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)

	hours := duration.Hours()
	speed := distance / hours

	return speed

}

func Distance(steps int, height float64) float64 {
	lenStep := height * stepLengthCoefficient
	distanceKm := (float64(steps) * lenStep) / mInKm
	return distanceKm
}
