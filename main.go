package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type PeptideSeq struct {
	seqId   string  // sequence Id
	peptide string  // Peptide an sich
	mass    float64 // masse des peptids (Summe)
}

func (ps *PeptideSeq) calucalteMass(amm map[string]uint) error {
	var sum float64

	// ABC
	for _, v := range ps.peptide {
		if v == '-' || v == '*' || v == 'X' {
			continue
		}

		pepMass, ok := amm[string(v)]
		if !ok {
			return fmt.Errorf("invalid key in peptide: %s", string(v))
		}

		sum = sum + float64(pepMass)
	}
	val := (len(ps.peptide) - 2) * 18
	ps.mass = sum - float64(val)
	return nil
}

func main() {
	// öffne fasta datei
	fs, err := os.Open("test.fasta")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer fs.Close()

	aminoacidMassMap := map[string]uint{
		"A": 89,
		"C": 121,
		"D": 133,
		"E": 147,
		"F": 165,
		"G": 75,
		"H": 155,
		"I": 131,
		"K": 146,
		"L": 131,
		"M": 149,
		"N": 132,
		"P": 115,
		"Q": 146,
		"R": 174,
		"S": 105,
		"T": 119,
		"U": 168,
		"V": 117,
		"W": 204,
		"Y": 181,
	}

	fscan := bufio.NewScanner(fs)

	line1Read := false
	var line1, line2 string
	result := make([]PeptideSeq, 5)

	// loop over every line in file
	for fscan.Scan() {
		if !line1Read {
			line1 = fscan.Text()
			line1Read = true
			continue
		}
		line2 = fscan.Text()

		peps := PeptideSeq{
			seqId:   line1,
			peptide: line2,
			mass:    0,
		}
		err := peps.calucalteMass(aminoacidMassMap)
		if err != nil {
			log.Fatal(err.Error())
		}
		line1Read = false
		// done, next line
		result = append(result, peps)
	}

	for _, p := range result {
		fmt.Println(p.peptide, p.mass)
	}
}
