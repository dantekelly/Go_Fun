package iteration

func Repeat(s string, n int) (c string) {
	for i := 0; i < n; i++ {
		c += s
	}

	return
}
