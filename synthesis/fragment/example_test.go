package fragment_test

import (
	"fmt"
	"github.com/TimothyStiles/poly/synthesis/fragment"
)

// This example shows how to use the fragmenter to fragment a gene in
// preparation for synthesis.
func Example_basic() {
	lacZ := "ATGACCATGATTACGCCAAGCTTGCATGCCTGCAGGTCGACTCTAGAGGATCCCCGGGTACCGAGCTCGAATTCACTGGCCGTCGTTTTACAACGTCGTGACTGGGAAAACCCTGGCGTTACCCAACTTAATCGCCTTGCAGCACATCCCCCTTTCGCCAGCTGGCGTAATAGCGAAGAGGCCCGCACCGATCGCCCTTCCCAACAGTTGCGCAGCCTGAATGGCGAATGGCGCCTGATGCGGTATTTTCTCCTTACGCATCTGTGCGGTATTTCACACCGCATATGGTGCACTCTCAGTACAATCTGCTCTGATGCCGCATAG"
	fragments, _, _ := fragment.Fragment(lacZ, 95, 105)

	fmt.Println(fragments)
	// Output: [ATGACCATGATTACGCCAAGCTTGCATGCCTGCAGGTCGACTCTAGAGGATCCCCGGGTACCGAGCTCGAATTCACTGGCCGTCGTTTTACAACGTCGTGACTGG CTGGGAAAACCCTGGCGTTACCCAACTTAATCGCCTTGCAGCACATCCCCCTTTCGCCAGCTGGCGTAATAGCGAAGAGGCCCGCACCGATCGCCCTTCCCAACA AACAGTTGCGCAGCCTGAATGGCGAATGGCGCCTGATGCGGTATTTTCTCCTTACGCATCTGTGCG TGCGGTATTTCACACCGCATATGGTGCACTCTCAGTACAATCTGCTCTGATGCCGCATAG]
}

// This example shows how to generate a new overhang onto a list of overhangs.
func ExampleNextOverhang() {
	primerOverhangs := []string{"ATAA"}
	primerOverhangs = append(primerOverhangs, fragment.NextOverhang(primerOverhangs))
	primerOverhangs = append(primerOverhangs, fragment.NextOverhang(primerOverhangs))
	primerOverhangs = append(primerOverhangs, fragment.NextOverhang(primerOverhangs))

	fmt.Println(primerOverhangs)
	// Output: [ATAA AAAT AATA AAGA]
}
