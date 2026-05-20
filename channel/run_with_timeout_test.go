package channel

import (
	"testing"
	"time"
)

func TestRunWithTimeout(t *testing.T) {
	t.Run("функция успевает — возвращается её результат", func(t *testing.T) {
		got := RunWithTimeout(func() int {
			return 42
		}, 100*time.Millisecond, -1)

		if got != 42 {
			t.Errorf("RunWithTimeout(быстрая fn) = %d, want 42", got)
		}
	})

	t.Run("функция не успевает — возвращается defaultVal", func(t *testing.T) {
		got := RunWithTimeout(func() int {
			time.Sleep(200 * time.Millisecond)
			return 42
		}, 50*time.Millisecond, -1)

		if got != -1 {
			t.Errorf("RunWithTimeout(медленная fn) = %d, want -1 (defaultVal)", got)
		}
	})

	t.Run("функция возвращает 0 — это не defaultVal", func(t *testing.T) {
		// Проверяет, что мы отличаем «функция вернула 0» от «случился таймаут».
		got := RunWithTimeout(func() int {
			return 0
		}, 100*time.Millisecond, 999)

		if got != 0 {
			t.Errorf("RunWithTimeout(fn возвращает 0) = %d, want 0", got)
		}
	})

	t.Run("успевает впритык", func(t *testing.T) {
		got := RunWithTimeout(func() int {
			time.Sleep(20 * time.Millisecond)
			return 7
		}, 200*time.Millisecond, -1)

		if got != 7 {
			t.Errorf("RunWithTimeout(20мс при таймауте 200мс) = %d, want 7", got)
		}
	})

	t.Run("не блокируется дольше timeout", func(t *testing.T) {
		const timeout = 50 * time.Millisecond
		const slack = 100 * time.Millisecond // запас на планировщик горутин

		start := time.Now()
		_ = RunWithTimeout(func() int {
			time.Sleep(2 * time.Second) // намеренно долго
			return 1
		}, timeout, -1)
		elapsed := time.Since(start)

		if elapsed > timeout+slack {
			t.Errorf("RunWithTimeout заблокировался на %v, что больше timeout %v + запас %v", elapsed, timeout, slack)
		}
	})

	t.Run("отрицательный результат корректно прокидывается", func(t *testing.T) {
		got := RunWithTimeout(func() int {
			return -50
		}, 100*time.Millisecond, 0)

		if got != -50 {
			t.Errorf("RunWithTimeout(fn возвращает -50) = %d, want -50", got)
		}
	})
}
