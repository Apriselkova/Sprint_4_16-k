package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	// "github.com/golangci/golangci-lint/pkg/golinters/nilerr"
	// "golang.org/x/mod/sumdb/storage"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

// parsePackage принимает строку с данными, которая содержит количество шагов и продолжительность прогулки
// и возвращает в форматах int, time.Duration и ошибку, если что-то пошло не так
func parsePackage(data string) (int, time.Duration, error) {
	sliceData := strings.Split(data, ",")

	if len(sliceData) != 2 {
		return 0, 0, fmt.Errorf("Неверный формат данных")
	}

	//strconv.Atoi преобразует строку с данными в тип (int)
	//strings.TrimSpace удаляет пробелы и другие разделители в слайсе
	steps, err := strconv.Atoi(strings.TrimSpace(sliceData[0]))
	if err != nil {
		return 0, 0, err
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	//time.ParseDuration разбирает продолжительности времени в виде строки и преобразует его в тип time.Duration.
	//strings.TrimSpace удаляет пробелы и другие разделители в слайсе
	duration, err := time.ParseDuration(strings.TrimSpace(sliceData[1]))
	if err != nil {
		return 0, 0, err
	}
	return steps, duration, nil
}

// DayActionInfo парсит строку с данными с помощью фукции parsePackage(),
// вычисляет дистанцию в километрах и количество потраченных калорий и выводит их
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return ""
	}

	if steps <= 0 {
		var m string
		return m
	}

	distance := (float64(steps) * StepLength) / 1000
	calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	return fmt.Sprintf("Количество шагов: %d. \nДистанция составила %.2f км. \nВы сожгли %.2f ккал.", steps, distance, calories)

}
