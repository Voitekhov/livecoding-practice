package slice

// ChunkSlice разбивает src на под-слайсы длиной не более size.
// Последний под-слайс может быть короче, если длина src не кратна size.
// Если size <= 0, возвращается nil.
// Для nil или пустого слайса возвращается nil.
func ChunkSlice(src []int, size int) [][]int {
	// TODO: реализовать разбиение слайса на части по size элементов
	return nil
}
