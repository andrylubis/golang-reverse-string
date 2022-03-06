package golang_reverse_string

// importing fmt

// function, which takes a string as
// argument and return the reverse of string.
func ReverseString(s string, nama string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return "Hasilnya = " + nama + " " + string(rns)
}
