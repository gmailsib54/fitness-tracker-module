package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	splitString := strings.Split(data, ",")

	if len(splitString) < 2 {
		return 0, 0, errors.New("invalid data")
	}

	steps, err := strconv.Atoi(splitString[0])

	if err != nil {
		return 0, 0, err
	}

	if steps <= 0 {
		return 0, 0, errors.New("there isn't a single step")
	}

	timeString := splitString[1]

	userTime, err := time.ParseDuration(timeString)

	if err != nil {
		return 0, 0, err
	}

	return steps, userTime, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию

	steps, userTime, err := parsePackage(data)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	if steps == 0 {
		return ""
	}

	distance := float64(steps) * stepLength

	kilometer := distance / mInKm

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, userTime)
	if err != nil {
		return ""
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила: %.2f км.\n"+
			"Вы сожгли: %.2f ккал.", steps, kilometer, calories)

	return result

}
