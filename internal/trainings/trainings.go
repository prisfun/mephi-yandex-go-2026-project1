// Package trainings разбирает записи о беге и ходьбе и формирует отчёты о тренировках.
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

// Training хранит показатели тренировки и персональные данные пользователя.
type Training struct {
	// Steps количество шагов за тренировку.
	Steps int
	// TrainingType вид активности: "Бег" или "Ходьба".
	TrainingType string
	// Duration продолжительность тренировки.
	Duration time.Duration
	// Personal содержит имя, вес и рост пользователя.
	personaldata.Personal
}

// Parse разбирает запись формата "шаги,тип,длительность", например
// "3456,Ходьба,3h00m". Количество шагов должно быть положительным целым числом,
// а длительность — положительным значением в формате [time.ParseDuration].
// Тип активности сохраняется без проверки; его поддержку проверяет [Training.ActionInfo].
//
// При ошибке Parse не изменяет t. При успешном разборе персональные данные сохраняются.
func (t *Training) Parse(datastring string) (err error) {
	fields := strings.Split(datastring, ",")
	if len(fields) != 3 {
		return errors.New("ожидаются количество шагов, тип тренировки и продолжительность")
	}
	steps, err := strconv.Atoi(fields[0])
	if err != nil {
		return fmt.Errorf("неверное количество шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}
	duration, err := time.ParseDuration(fields[2])
	if err != nil {
		return fmt.Errorf("неверная продолжительность: %w", err)
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть положительной")
	}
	t.Steps, t.TrainingType, t.Duration = steps, fields[1], duration
	return nil
}

// ActionInfo возвращает отчёт о тренировке с переводом строки в конце.
// Длительность указывается в часах, дистанция в километрах, скорость в километрах
// в час, а оценка расхода энергии в килокалориях.
// Возвращает пустую строку и ошибку, если тип тренировки не поддерживается
// или расчёт расхода энергии завершился ошибкой.
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, calories), nil
}
