package common

func RemoveEle[T comparable](list []T, ele T) []T {
	newSlice := make([]T, 0, len(list))

	for i := range list {
		if list[i] != ele {
			newSlice = append(newSlice, list[i])

			continue
		}

		newSlice = append(newSlice, list[i+1:]...)

		break
	}

	return newSlice
}
