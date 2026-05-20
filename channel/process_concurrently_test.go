package channel

import (
	"reflect"
	"sync/atomic"
	"testing"
	"time"
)

func TestProcessConcurrently_Correctness(t *testing.T) {
	square := func(x int) int { return x * x }

	tests := []struct {
		name    string
		items   []int
		workers int
		fn      func(int) int
		want    []int
	}{
		{
			name:    "квадраты с 2 воркерами",
			items:   []int{1, 2, 3, 4},
			workers: 2,
			fn:      square,
			want:    []int{1, 4, 9, 16},
		},
		{
			name:    "воркеров больше, чем элементов",
			items:   []int{10, 20, 30},
			workers: 10,
			fn:      func(x int) int { return x + 1 },
			want:    []int{11, 21, 31},
		},
		{
			name:    "один воркер — sequential",
			items:   []int{5, 6, 7},
			workers: 1,
			fn:      func(x int) int { return -x },
			want:    []int{-5, -6, -7},
		},
		{
			name:    "workers <= 0 трактуется как 1",
			items:   []int{1, 2, 3},
			workers: 0,
			fn:      func(x int) int { return x * 10 },
			want:    []int{10, 20, 30},
		},
		{
			name:    "пустой слайс",
			items:   []int{},
			workers: 4,
			fn:      square,
			want:    []int{},
		},
		{
			name:    "nil слайс",
			items:   nil,
			workers: 4,
			fn:      square,
			want:    []int{},
		},
		{
			name:    "один элемент",
			items:   []int{9},
			workers: 4,
			fn:      square,
			want:    []int{81},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProcessConcurrently(tt.items, tt.workers, tt.fn)

			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessConcurrently(%v, %d, fn) = %v, want %v", tt.items, tt.workers, got, tt.want)
			}
		})
	}
}

// Проверяет, что задания реально считаются параллельно: 4 элемента
// по 100мс каждый с 4 воркерами должны уложиться сильно быстрее, чем
// последовательные 400мс.
func TestProcessConcurrently_Parallelism(t *testing.T) {
	const itemDur = 100 * time.Millisecond

	slow := func(x int) int {
		time.Sleep(itemDur)
		return x
	}

	items := []int{1, 2, 3, 4}

	start := time.Now()
	got := ProcessConcurrently(items, 4, slow)
	elapsed := time.Since(start)

	// При полной параллельности — ~100мс, при последовательности — ~400мс.
	// Берём порог 250мс: с запасом отделяет параллельный случай от sequential.
	if elapsed > 250*time.Millisecond {
		t.Errorf("4 элемента по %v с 4 воркерами заняли %v — похоже на последовательное выполнение", itemDur, elapsed)
	}

	if !reflect.DeepEqual(got, items) {
		t.Errorf("got %v, want %v", got, items)
	}
}

// Проверяет, что одновременно работает не более workers горутин:
// замеряем максимальное число одновременно исполняющихся вызовов fn.
func TestProcessConcurrently_RespectsWorkerLimit(t *testing.T) {
	const workers = 3
	const items = 12

	var inFlight int32
	var maxInFlight int32

	fn := func(x int) int {
		cur := atomic.AddInt32(&inFlight, 1)
		for {
			prev := atomic.LoadInt32(&maxInFlight)
			if cur <= prev || atomic.CompareAndSwapInt32(&maxInFlight, prev, cur) {
				break
			}
		}
		time.Sleep(20 * time.Millisecond)
		atomic.AddInt32(&inFlight, -1)
		return x
	}

	input := make([]int, items)
	for i := range input {
		input[i] = i
	}

	ProcessConcurrently(input, workers, fn)

	if got := atomic.LoadInt32(&maxInFlight); got > workers {
		t.Errorf("одновременно работало %d горутин, лимит = %d", got, workers)
	}
}
