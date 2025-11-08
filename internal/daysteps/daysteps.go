package daysteps

import (
	"errors"
	"fmt"
	"log"
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

	// Продллжительность меньше или равно 0
	if durationOfTheWalk <= 0 {
		return 0, 0, errors.New("время равно 0")
	}

	return numberOfSteps, durationOfTheWalk, nil

}

func DayActionInfo(data string, weight, height float64) string {

	numberOfSteps, durationOfTheWalk, err := parsePackage(data)

	if err != nil {
		log.Println(err)
		return ""
	}

	if numberOfSteps <= 0 {
		return ""
	}
	// Проиденая дистанция
	distanceInMeters := float64(numberOfSteps) * stepLength

	distanceInKilometers := distanceInMeters / mInKm

	numberOfCalories, err := spentcalories.WalkingSpentCalories(numberOfSteps, weight, height, durationOfTheWalk)

	if err != nil {
		log.Println(err)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		numberOfSteps, distanceInKilometers, numberOfCalories)
}
