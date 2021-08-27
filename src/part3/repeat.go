// package iteration

package iteration

func Repeat(character string, times int) string {
	var repeat string
	for i := 0; i < times; i++ {
		repeat += character
	}
	return repeat
}
