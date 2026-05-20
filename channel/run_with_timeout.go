package channel

import "time"

// RunWithTimeout запускает функцию fn в отдельной горутине и ждёт её
// результат не дольше, чем timeout.
//
//   - Если fn успела завершиться за отведённое время — возвращается
//     значение, которое она вернула.
//   - Если fn не уложилась в timeout — возвращается defaultVal, при этом
//     RunWithTimeout не должна блокироваться дольше timeout, даже если
//     fn в итоге так и не завершится.
//
// Реализуется через канал для результата и select с time.After(timeout)
// для таймаута. Сама fn остаётся работать в своей горутине — мы её не
// прерываем (в Go нельзя «убить» горутину снаружи); важно лишь, чтобы
// мы не зависли в ожидании её результата.
//
// Пример:
//
//	res := RunWithTimeout(func() int {
//	    time.Sleep(2 * time.Second)
//	    return 42
//	}, time.Second, -1)
//	// fn не уложилась в секунду -> res == -1
func RunWithTimeout(fn func() int, timeout time.Duration, defaultVal int) int {
	return 0
}
