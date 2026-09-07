package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func validateParameters(steps int, weight, height float64, duration time.Duration) error {
	if steps <= 0 {
		return fmt.Errorf("steps can't be equal 0 or less")
	}

	if weight <= 0 {
		return fmt.Errorf("weight can't be equal 0 or less")
	}

	if height <= 0 {
		return fmt.Errorf("height can't be equal 0 or less")
	}

	if duration <= 0 {
		return fmt.Errorf("duration can't be equal 0 or less")
	}

	return nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validateParameters(steps, weight, height, duration); err != nil {
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	spentCalories := (weight * meanSpeed * duration.Minutes()) / minInH * walkingCaloriesCoefficient

	return spentCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validateParameters(steps, weight, height, duration); err != nil {
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	spentCalories := (weight * meanSpeed * duration.Minutes()) / minInH
	return spentCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 {
		return 0
	}

	if duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)
	meanSpeed := distance / duration.Hours()

	return meanSpeed
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 {
		return 0
	}

	distance := (stepLengthCoefficient * height) * float64(steps)
	distanceInKm := distance / mInKm

	return distanceInKm
}
