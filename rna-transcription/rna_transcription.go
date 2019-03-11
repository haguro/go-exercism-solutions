//Package strand provides functionality to calculate the transcribed RNA for a given DNA strand.
package strand

var complements = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

//ToRNA returns the RNA complement for a given DNA strand string.
func ToRNA(dna string) string {
	rna := ""
	for i := 0; i < len(dna); i++ {
		rna += complements[string(dna[i])]
	}
	return rna
}
