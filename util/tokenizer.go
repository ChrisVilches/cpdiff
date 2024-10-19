package util

func StringFieldsKeepWhitespace(s string) func(func(int, int) bool) {
	return func(yield func(int, int) bool) {
		state := 0

		idx := 0

		for i, c := range s {
			switch state {
			case 0:
				if c != ' ' {
					state = 1
				}
			case 1:
				if c == ' ' {
					state = 2
				}
			case 2:
				if c != ' ' {
					if !yield(idx, i) {
						return
					}

					idx++
					state = 1
				}
			}
		}

		if state != 0 {
			yield(idx, len(s))
		}
	}
}
