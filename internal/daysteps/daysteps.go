// Package daysteps разбирает записи о дневных прогулках и формирует отчёты об активности.
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

// DaySteps хранит данные о дневных прогулках и персональные данные пользователя.
type DaySteps struct {
	// Steps количество пройденных шагов.
	Steps int
	// Duration продолжительность прогулки.
	Duration time.Duration
	// Personal содержит имя, вес и рост пользователя.
	personaldata.Personal
}

// Parse разбирает запись формата "шаги,длительность", например "678,0h50m".
// Количество шагов должно быть положительным целым числом, а длительность —
// положительным значением в формате [time.ParseDuration].
// При ошибке Parse не изменяет ds. При успешном разборе персональные данные сохраняются.
func (ds *DaySteps) Parse(datastring string) (err error) {
	fields := strings.Split(datastring, ",")
	if len(fields) != 2 {
		return errors.New("ожидаются количество шагов и продолжительность")
	}
	steps, err := strconv.Atoi(fields[0])
	if err != nil {
		return fmt.Errorf("неверное количество шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}
	duration, err := time.ParseDuration(fields[1])
	if err != nil {
		return fmt.Errorf("неверная продолжительность: %w", err)
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть положительной")
	}
	ds.Steps, ds.Duration = steps, duration
	return nil
}

// ActionInfo возвращает отчёт о шагах, дистанции в километрах
// и оценке расхода энергии в килокалориях с переводом строки в конце.
// Возвращает пустую строку и ошибку, если расчёт расхода энергии завершился ошибкой.
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories), nil
}
