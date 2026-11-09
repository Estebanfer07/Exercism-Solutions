package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	result := initial

	for _, value := range s {
		result = fn(result, value)
	}

	return result
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	result := initial

	for i := len(s) - 1; i >= 0; i-- {
		result = fn(s[i], result)
	}

	return result
}

func (s IntList) Filter(fn func(int) bool) IntList {
	result := IntList{}

	for _, value := range s {
		if fn(value) {
			result = append(result, value)
		}
	}

	return result
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	result := IntList{}

	for _, value := range s {
		result = append(result, fn(value))
	}

	return result
}

func (s IntList) Reverse() IntList {
	result := IntList{}

	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}

	return result
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	result := append(IntList{}, s...)

	for _, lst := range lists {
		result = result.Append(lst)
	}

	return result
}
