package lesson2

// FizzBuzz возвращает FizzBuzz, Fizz, Buzz или пустую строку для обычного числа.
// Для n <= 0 возвращает invalid.
func FizzBuzz(n int) string {
	// TODO: Реализовать через switch без выражения.
	switch {
	case n%3 == 0 && n%5 != 0 && n != 0:
		return "Fizz"
	case n%5 == 0 && n%3 != 0 && n != 0:
		return "Buzz"
	case n%15 == 0 && n != 0:
		return "FizzBuzz"
	case n <= 0:
		return "invalid"
	default:
		return ""
	}
}

// Sum возвращает сумму. Для пустого среза и nil возвращает 0.
func Sum(values []int64) int64 {
	// TODO: Пройти по значениям циклом.
	if values == nil || len(values) == 0 {
		return 0
	}
	sum := int64(0)
	for _, value := range values {
		sum += value
	}
	return sum
}

// Min возвращает минимум. Для пустого среза и nil возвращает 0 по условию задачи.
func Min(values []int64) int64 {
	// TODO: Учесть пустой ввод, затем сравнить элементы.
	if values == nil || len(values) == 0 {
		return 0
	}
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}

// Max возвращает максимум. Для пустого среза и nil возвращает 0 по условию задачи.
func Max(values []int64) int64 {
	// TODO: Учесть пустой ввод, затем сравнить элементы.
	if values == nil || len(values) == 0 {
		return 0
	}
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

// Average возвращает среднее как float64. Для пустого среза и nil возвращает 0.
func Average(values []int64) float64 {
	// TODO: Выполнить преобразования до деления.
	if values == nil || len(values) == 0 {
		return 0
	}
	sum := Sum(values)
	return float64(sum) / float64(len(values))
}

// Category возвращает invalid, small, regular или large по сумме в копейках.
func Category(amount int64) string {
	// TODO: Проверить границы через switch без выражения.
	switch {
	case amount >= 0 && amount <= 9999:
		return "small"
	case amount >= 10000 && amount <= 49999:
		return "regular"
	case amount >= 50000:
		return "large"
	default:
		return "invalid"
	}
}

// RuneCount считает кодовые точки Unicode в строке через range.
func RuneCount(text string) int {
	// TODO: Не использовать len(text) как количество кодовых точек.
	cnt := 0
	for range text {
		cnt++
	}
	return cnt

}
