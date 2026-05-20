package mapslice

// GroupByLength группирует строки из words по их длине (в байтах).
// Ключ результирующего map — длина строки, значение — слайс строк такой длины
// в том порядке, в котором они встречались в words. Для nil или пустого
// слайса возвращает пустой (не nil) map.
func GroupByLength(words []string) map[int][]string {
	// TODO: реализовать через map[int][]string, аккумулируя строки по len(word)
	return nil
}
