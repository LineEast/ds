package binarytreeanybeta

func Init[T any](compare func(one, two T) int) *Tree[T] {
	return &Tree[T]{
		Compare: compare,
	}
}

func CompareStrings(one, two string) (result int) {
	if one > two {
		return 1
	}

	if one < two {
		return -1
	}

	return
}

