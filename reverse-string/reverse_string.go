//Package reverse provides functionality to reverse strings
package reverse

//String returnes a reversed version of a given string
func String(s string) string {
	sr := []rune(s)
	rs := make([]rune, len(sr))
	for i, j := len(sr)-1, 0; i >= 0; i-- {
		rs[j] = sr[i]
		j++
	}
	return string(rs)
}
