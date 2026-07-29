package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	splitData := strings.Split(data, ",")

	if len(splitData) != 3 {
		return 0, "", 0, errors.New("invalid data - splitData")
	}

	steps, err := strconv.Atoi(splitData[0])
	if err != nil {
		return 0, "", 0, errors.New("invalid data - steps")
	}

	actions := splitData[1]
	timeString := splitData[2]

	timeTraining, err := time.ParseDuration(timeString)
	if err != nil {
		return 0, "", 0, errors.New("invalid data timeTraining")
	}

	return steps, actions, timeTraining, nil

}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию

	calculatedDistance := height * stepLengthCoefficient * float64(steps) / mInKm

	return calculatedDistance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	averageSpeed := distance(steps, height) / duration.Hours()

	return averageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, actions, timeTraining, err := parseTraining(data)
	if err != nil {
		return "", err
	}

	log.Printf("data: %v", data)

	switch actions {
	case "Ходьба":
		distanceWalking := distance(steps, height)
		caloriesWalking, err := WalkingSpentCalories(steps, weight, height, timeTraining)
		if err != nil {
			log.Println(err)
		}
		speedWalking := meanSpeed(steps, height, timeTraining)

		result := fmt.Sprintf(
			"Тип тренировки: Ходьба\n"+
				"Длительность: %.2f ч.\n"+
				"Дистанция: %.2f км.\n"+
				"Скорость: %.2f км/ч\n"+
				"Сожгли калорий: %f", timeTraining.Hours(), distanceWalking, speedWalking, caloriesWalking)

		return result, nil

	case "Бег":
		distanceRunning := distance(steps, height)
		calories, err := RunningSpentCalories(steps, weight, height, timeTraining)
		if err != nil {
			log.Println(err)
		}
		speedRunning := meanSpeed(steps, height, timeTraining)
		result := fmt.Sprintf(
			"Тип тренировки: Бег\n"+
				"Длительность: %.2f ч.\n"+
				"Дистанция: %.2f км.\n"+
				"Скорость: %.2f км/ч\n"+
				"Сожгли калорий: %f", timeTraining.Hours(), distanceRunning, speedRunning, calories)

		return result, nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 {
		return 0, errors.New("invalid spent for running")
	}
	averageSpeed := meanSpeed(steps, height, duration)

	timeMinute := duration.Minutes()

	quantityCaloriesRunning := (weight * averageSpeed * timeMinute) / minInH

	return quantityCaloriesRunning, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps < 0 || weight < 0 || height < 0 {
		return 0, errors.New("invalid spent for walking")
	}
	averageSpeed := meanSpeed(steps, height, duration)

	timeMinute := duration.Minutes()

	quantityCaloriesWalking := (weight * averageSpeed * timeMinute) / minInH * walkingCaloriesCoefficient

	return quantityCaloriesWalking, nil
}
