// Package spentenergy оценивает дистанцию, среднюю скорость и расход энергии
// при беге и ходьбе по количеству шагов и физическим параметрам пользователя.
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

// WalkingSpentCalories оценивает расход энергии при ходьбе в килокалориях.
// Вес weight задаётся в килограммах, рост height в метрах. Оценка равна половине
// значения, возвращаемого [RunningSpentCalories]. Если steps, weight, height
// или duration равны нулю или отрицательны, возвращает 0 и ошибку.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return calories * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories оценивает расход энергии при беге в килокалориях как
// произведение веса, средней скорости и длительности в часах. Вес weight задаётся
// в килограммах, рост height в метрах. Если steps, weight, height или duration
// равны нулю или отрицательны, возвращает 0 и ошибку.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть положительным")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть положительным")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть положительной")
	}
	speed := MeanSpeed(steps, height, duration)
	return weight * speed * duration.Minutes() / minInH, nil
}

// MeanSpeed возвращает среднюю скорость в километрах в час с помощью [Distance].
// Рост height задаётся в метрах. Возвращает 0, если duration равна нулю или
// отрицательна либо steps отрицательно.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps < 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

// Distance оценивает расстояние в километрах по количеству шагов steps.
// Длина шага принимается равной 0.45 роста height, заданного в метрах.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}
