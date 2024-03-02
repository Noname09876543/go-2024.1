package slice

// функция append добавления в конец
func appends(slice []int, elements ...int) []int {
	newSlice := make([]int, len(slice)+len(elements))
	copy(newSlice, slice)
	copy(newSlice[len(slice):], elements)
	return newSlice
}

// AppendElement добавляет элемент в конец слайса.
func AppendElement(ints []int, elem int) []int {
	ints = appends(ints, elem)
	return ints
	//или
	//ints = appends(ints, elem)
	//return ints
}

// функция delets удаление из конеца
func delets(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}
	return slice[:len(slice)-1]
}

// RemoveElement удаляет последний элемент из слайса. Если массив пуст, функция возвращает оригинальный пустой массив.
func RemoveElement(ints []int) []int {
	// Код тут
	ints = delets(ints)
	return ints
	//или
	// if len(ints) == 0 {
	// 	return ints
	// }
	// return ints[:len(ints)-1]

}

// функция incs прибавление 1 к всем эелментам
func incs(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		slice[i]++
	}
	return slice
}

// AddOneToAll увеличивает каждый элемент массива на единицу.
func AddOneToAll(ints []int) []int {
	// Код тут
	ints = incs(ints)
	return ints
	//или
	// for i := 0; i < len(ints); i++ {
	// 	ints[i]++
	// }
	// return ints
}
