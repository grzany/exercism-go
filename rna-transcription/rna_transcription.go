//Package strand provides function for RNA transcription
package strand

// ToRNA Given a DNA strand, returns its RNA complement (per RNA transcription).
func ToRNA(dna string) string {

	var rna string
	r := make(map[string]string)

	r["G"] = "C"
	r["C"] = "G"
	r["T"] = "A"
	r["A"] = "U"

	for i := 0; i < len(dna); i++ {
		rna = rna + r[string(dna[i])]
	}

	return rna
}
