package HW2

// PairInput соответствует объекту nums/target из условия контеста.
type PairInput struct {
	Nums   []int
	Target int
}

// longestPrefix возвращает самый длинный общий префикс строк.
func longestPrefix(words []string) string {
	if len(words) == 0 {
		return ""
	}

	prefix := []rune(words[0])
	if len(prefix) == 0 {
		return ""
	}

	for i, r := range prefix {
		for _, word := range words[1:] {
			wordRunes := []rune(word)
			if i >= len(wordRunes) || wordRunes[i] != r {
				return string(prefix[:i])
			}
		}
	}
	return string(prefix)
}

// pairIndices возвращает индексы двух разных элементов по возрастанию.
func pairIndices(input PairInput) [2]int {
	m := make(map[int]int, len(input.Nums))
	for i, n := range input.Nums {
		diff := input.Target - n
		if prev, ok := m[diff]; ok {
			return [2]int{prev, i}
		}
		m[n] = i
	}
	return [2]int{}
}

// hasDuplicate сообщает, встречается ли какое-либо значение повторно.
func hasDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	m := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		if _, exists := m[n]; exists {
			return true
		}
		m[n] = struct{}{}
	}
	return false
}

// compactSorted создаёт независимый слайс уникальных значений.
func compactSorted(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	result := make([]int, 0, len(nums))
	result = append(result, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			result = append(result, nums[i])
		}
	}
	return result
}

// balancedBrackets проверяет пары и порядок закрытия скобок.
func balancedBrackets(text string) bool {
	if len(text)%2 != 0 {
		return false
	}

	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := make([]rune, 0, len(text))

	for _, v := range text {
		switch v {
		case '(', '{', '[':
			stack = append(stack, v)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[v] {
				return false
			}
			stack = stack[:len(stack)-1]
		}

	}
	return len(stack) == 0
}
