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

	str := strings.Split(data, ",")

	if len(str) != 3 {
		return 0, "", 0, errors.New("длина слайса не равна трем")
	}

	numberOfSteps, err := strconv.Atoi(str[0])

	if err != nil {
		return 0, "", 0, err
	}

	if numberOfSteps <= 0 {
		return 0, "", 0, errors.New("количество шагов меньше или равно нулю")
	}

	durationOfTheActivity, err := time.ParseDuration(str[2])

	if err != nil {
		return 0, "", 0, err
	}

	if durationOfTheActivity <= 0 {
		return 0, "", 0, errors.New("продолжительность равно или меньше нуля")
	}

	return numberOfSteps, str[1], durationOfTheActivity, nil
}

func distance(steps int, height float64) float64 {

	stepLength := height * stepLengthCoefficient

	distanceTraveledInMeters := stepLength * float64(steps)

	distanceTraveledInKilometers := distanceTraveledInMeters / mInKm

	return distanceTraveledInKilometers

}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {

	if duration <= 0 {
		err := errors.New("продолжительность меньше или равна нулю")
		log.Println(err)
		return 0
	}

	if height <= 0 {
		err := errors.New("рост меньше или равен нулю")
		log.Println(err)
		return 0
	}

	distanceTravelForSpeed := distance(steps, height)

	averageSpeed := distanceTravelForSpeed / float64(duration.Hours())

	return averageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {

	numberOfSteps, typeOfTraining, durationOfTheActivity, err := parseTraining(data)

	if err != nil {
		log.Println(err)
		return "", err
	}

	if weight <= 0 {
		return "", errors.New("отрицательный или нулевой вес")
	}

	if height <= 0 {
		return "", errors.New("отрицательный или нулевой рост")
	}

	switch typeOfTraining {
	case "Бег":
		distance := distance(numberOfSteps, height)
		averageSpeed := meanSpeed(numberOfSteps, height, durationOfTheActivity)
		RunningSpentCalories, err := RunningSpentCalories(numberOfSteps, weight, height, durationOfTheActivity)

		if err != nil {
			log.Println(err)
			return "", err
		}

		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			typeOfTraining, float64(durationOfTheActivity.Hours()), distance, averageSpeed, RunningSpentCalories), nil

	case "Ходьба":
		distance := distance(numberOfSteps, height)
		averageSpeed := meanSpeed(numberOfSteps, height, durationOfTheActivity)
		WalkingSpentCalories, err := WalkingSpentCalories(numberOfSteps, weight, height, durationOfTheActivity)

		if err != nil {
			log.Println(err)
			return "", err
		}

		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			typeOfTraining, float64(durationOfTheActivity.Hours()), distance, averageSpeed, WalkingSpentCalories), nil

	default:
		return "", errors.New("неизвестный тип тренировки")
	}

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		return 0, errors.New("количество шагов меньше или равно нулю")
	}

	if duration <= 0 {
		return 0, errors.New("продолжительность меньше или равна нулю")
	}

	if weight <= 0 {
		return 0, errors.New("отрицательный или нулевой вес")
	}

	if height <= 0 {
		return 0, errors.New("отрицательный или нулевой рост")
	}

	averageSpeed := meanSpeed(steps, height, duration)

	durationInMinutes := float64(duration.Minutes())

	numberOfCalories := (weight * averageSpeed * durationInMinutes) / minInH

	return numberOfCalories, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		return 0, errors.New("количество шагов меньше или равно нулю")
	}

	if duration <= 0 {
		return 0, errors.New("продолжительность меньше или равна нулю")
	}

	if weight <= 0 {
		return 0, errors.New("отрицательный или нулевой вес")
	}

	if height <= 0 {
		return 0, errors.New("отрицательный или нулевой рост")
	}

	averageSpeed := meanSpeed(steps, height, duration)

	durationInMinutes := float64(duration.Minutes())

	numberOfCalories := (weight * averageSpeed * durationInMinutes) / minInH

	numberOfCaloriesIsSpecified := numberOfCalories * walkingCaloriesCoefficient

	return numberOfCaloriesIsSpecified, nil

}
