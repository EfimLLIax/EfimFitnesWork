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
	// Длина одного шага в метрах .
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {

	str := strings.Split(data, ",")

	if len(str) != 2 {
		return 0, 0, errors.New("длина слайса не равна 2")
	}
	// Количество шагов
	numberOfSteps, err := strconv.Atoi(str[0])

	if err != nil {
		return 0, 0, err
	}

	if numberOfSteps <= 0 {
		return 0, 0, errors.New("количество шагов меньше или равно 0")
	}
	// Продолжительность прогулки
	durationOfTheWalk, err := time.ParseDuration(str[1])

	if err != nil {
		return 0, 0, err
	}

	return numberOfSteps, durationOfTheWalk, nil

}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию

	numberOfSteps, durationOfTheWalk, err := parsePackage(data)

	if err != nil {
		err = fmt.Errorf("ошибка в ходе выполнения программы: %v", err)
		fmt.Println(err)
		return ""
	}

	if numberOfSteps <= 0 {
		return ""
	}
	// Проиденая дистанция
	distanceInMeters := float64(numberOfSteps) * stepLength

	distanceInKilometers := distanceInMeters / mInKm

	//////////////////////////////////////////////////////////////////********************8
	numberOfCalories, err := spentcalories.WalkingSpentCalories(numberOfSteps, weight, height, durationOfTheWalk)

	return fmt.Sprintf("Количество шагов: %v.\nДистанция составила %.2v км.\nВы сожгли %.2v ккал.", numberOfSteps, distanceInKilometers, numberOfCalories)
}
