package dfsbfs

import "testing"

func TestSumNested(t *testing.T) {
	tests := []struct {
		name string
		data []interface{}
		want int
	}{
		{
			// пример из задачи: [1, [2, 3], [[4], 5]]
			//
			//              [корень]
			//             /    |    \
			//            1   [ ]    [ ]
			//               /  \    /  \
			//              2    3 [ ]   5
			//                      |
			//                      4
			//
			// 1 + 2 + 3 + 4 + 5 = 15
			name: "пример из задачи",
			data: []interface{}{
				1,
				[]interface{}{2, 3},
				[]interface{}{
					[]interface{}{4},
					5,
				},
			},
			want: 15,
		},
		{
			// плоский список — корень с 5 листьями.
			//
			//        [корень]
			//        / | | | \
			//       1  2 3 4  5
			//
			// 1 + 2 + 3 + 4 + 5 = 15
			name: "плоский список",
			data: []interface{}{1, 2, 3, 4, 5},
			want: 15,
		},
		{
			// глубокая вложенность: [[[[42]]]]
			//
			//   [корень]
			//      |
			//     [ ]
			//      |
			//     [ ]
			//      |
			//     [ ]
			//      |
			//      42
			name: "глубокая вложенность",
			data: []interface{}{
				[]interface{}{
					[]interface{}{
						[]interface{}{42},
					},
				},
			},
			want: 42,
		},
		{
			// корень без потомков -> сумма 0.
			name: "пустой список",
			data: []interface{}{},
			want: 0,
		},
		{
			// nil трактуется так же, как пустой список.
			name: "nil список",
			data: nil,
			want: 0,
		},
		{
			// список из пустых списков: листьев нет, сумма 0.
			//
			//   [корень]
			//    /    \
			//   [ ]   [ ]
			//          |
			//         [ ]
			name: "список из пустых списков",
			data: []interface{}{
				[]interface{}{},
				[]interface{}{[]interface{}{}},
			},
			want: 0,
		},
		{
			// отрицательные числа: [-1, [-2, [-3, 10]]]
			//
			//      [корень]
			//      /    \
			//    -1    [ ]
			//          /  \
			//        -2   [ ]
			//             /  \
			//           -3   10
			//
			// -1 + -2 + -3 + 10 = 4
			name: "отрицательные числа",
			data: []interface{}{
				-1,
				[]interface{}{-2, []interface{}{-3, 10}},
			},
			want: 4,
		},
		{
			// единственный лист.
			//
			//   [корень]
			//      |
			//      7
			name: "одно число",
			data: []interface{}{7},
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumNested(tt.data)
			if got != tt.want {
				t.Errorf("sumNested(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}
