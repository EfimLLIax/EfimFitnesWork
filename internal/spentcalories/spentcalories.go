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
	// Разеление строки на слайсы
	str := strings.Split(data, ",")

	// Количество элементов
	if len(str) != 3 {
		return 0, "", 0, errors.New("длина слайса не равна 3")
	}

	// Количество шагов
	numberOfSteps, err := strconv.Atoi(str[0])

	if err != nil {
		return 0, "", 0, err
	}

	if numberOfSteps <= 0 {
		return 0, "", 0, errors.New("количество шагов должно быть больше 0")
	}

	//Продолжительность активности
	durationOfTheActivity, err := time.ParseDuration(str[2])

	if err != nil {
		return 0, "", 0, err
	}

	return numberOfSteps, str[1], durationOfTheActivity, nil
}

func distance(steps int, height float64) float64 {
	// Длинна шага
	stepLength := height * stepLengthCoefficient
	// Проиденное растояние в метрах
	distanceTraveledInMeters := stepLength * float64(steps)
	// Проиденное растояние в километрах
	distanceTraveledInKilometers := distanceTraveledInMeters / mInKm

	return distanceTraveledInKilometers

}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {

	if duration <= 0 {
		return 0
	}
	// Вычисление дистанций
	distanceTravelForSpeed := distance(steps, height)
	// Вычисление средней скорости
	averageSpeed := distanceTravelForSpeed / float64(duration.Hours())

	return averageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	//
	numberOfSteps, typeOfTraining, durationOfTheActivity, err := parseTraining(data)

	if err != nil {
		log.Println(err)
		return "", err
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

		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2v\nДистанция: %.2v\nСкорость: %.2v\nСожгли калорий: %.2v",
			typeOfTraining, float64(durationOfTheActivity.Hours()), distance, averageSpeed, RunningSpentCalories), nil

	case "Ходьба":
		distance := distance(numberOfSteps, height)
		averageSpeed := meanSpeed(numberOfSteps, height, durationOfTheActivity)
		WalkingSpentCalories, err := WalkingSpentCalories(numberOfSteps, weight, height, durationOfTheActivity)

		if err != nil {
			log.Println(err)
			return "", err
		}

		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2v\nДистанция: %.2v\nСкорость: %.2v\nСожгли калорий: %.2v",
			typeOfTraining, float64(durationOfTheActivity.Hours()), distance, averageSpeed, WalkingSpentCalories), nil

	default:
		err = errors.New("неизвестный тип тренировки")
		return "", err
	}

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверка входных данных
	if steps <= 0 {
		return 0, errors.New("количество шагов меньше или равно 0")
	}

	if duration <= 0 {
		return 0, errors.New("продолжительность меньше или равна 0")
	}
	// Средняя скорость
	averageSpeed := meanSpeed(steps, height, duration)
	// Продолжительность в минуах
	durationInMinutes := float64(duration.Minutes())
	// Количество калорий
	numberOfCalories := (weight * averageSpeed * durationInMinutes) / minInH

	return numberOfCalories, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверка входных данных
	if steps <= 0 {
		return 0, errors.New("количество шагов меньше или равно 0")
	}

	if duration <= 0 {
		return 0, errors.New("продолжительность меньше или равна 0")
	}

	// Средняя скорость
	averageSpeed := meanSpeed(steps, height, duration)
	// Продолжительность в минуах
	durationInMinutes := float64(duration.Minutes())
	// Количество калорий
	numberOfCalories := (weight * averageSpeed * durationInMinutes) / minInH

	numberOfCaloriesIsSpecified := numberOfCalories * walkingCaloriesCoefficient

	return numberOfCaloriesIsSpecified, nil

}
