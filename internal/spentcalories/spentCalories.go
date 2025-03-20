package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

// parseTraining разбивает строку на составляющие и возвращает количество шагов, вид активности и продолжительность.
func parseTraining(data string) (int, string, time.Duration, error) {
	dataSlice := strings.Split(data, ",")

	if len(dataSlice) != 3 {
		return 0, "", 0, fmt.Errorf("Ошибка: ожидается 3 элемент, получено %d", len(dataSlice))
	}
	//strconv.Atoi преобразует строку с данными в тип (int)
	//strings.TrimSpace удаляет пробелы и другие разделители в слайсе
	steps, err := strconv.Atoi(strings.TrimSpace(dataSlice[0]))
	if err != nil {
		return 0, "", 0, err
	}

	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("количество шагов должно быть больше 0")
	}
	//time.ParseDuration разбирает продолжительности времени в виде строки и преобразует его в тип time.Duration.
	//strings.TrimSpace удаляет пробелы и другие разделители в слайсе
	duration, err := time.ParseDuration(strings.TrimSpace(dataSlice[2]))
	if err != nil {
		return 0, "", 0, err
	}
	return steps, strings.TrimSpace(dataSlice[1]), duration, nil
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
func distance(steps int) float64 {
	return float64(steps) * lenStep / float64(mInKm)
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	dist := distance(steps)
	return dist / duration.Hours()
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	meanSpd := meanSpeed(steps, duration)
	return ((runningCaloriesMeanSpeedMultiplier * meanSpd) - runningCaloriesMeanSpeedShift) * weight
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	meanSpeed := meanSpeed(steps, duration)
	return ((float64(walkingCaloriesWeightMultiplier) * weight) + (meanSpeed*meanSpeed/height)*float64(walkingSpeedHeightMultiplier)) * float64(duration) * minInH
}

// TrainingInfo возвращает строку с информацией о тренировке.
func TrainingInfo(data string, weight, height float64) string {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		return fmt.Sprintf("Ошибка: %v", err)
	}

	switch activity {
	case "Бег":
		dist := distance(steps)
		speed := meanSpeed(steps, duration)
		calories := RunningSpentCalories(steps, weight, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
			activity, duration.Hours(), dist, speed, calories)
	case "Ходьба":
		dist := distance(steps)
		speed := meanSpeed(steps, duration)
		calories := WalkingSpentCalories(steps, weight, height, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
			activity, duration.Hours(), dist, speed, calories)
	default:
		return "Неизвестный тип тренировки"
	}
}
