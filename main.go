package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	peptideseq "peptide-analyse/peptide-seq"
	"strconv"
)

func printUsage(s string) {
	fmt.Println(s)
	usage := fmt.Sprintf("Usage: %s min max <file list>", os.Args[0])
	fmt.Println(usage)
}

func printErrorAndExit(s string) {
	fmt.Printf("Error: %s\n", s)
	os.Exit(1)
}

func main() {
	if len(os.Args) <= 2 {
		printUsage("Wrong number of argments")
		os.Exit(1)
	}

	// TODO: Use min and max range
	if _, err := strconv.ParseUint(os.Args[1], 10, 64); err != nil {
		printErrorAndExit(err.Error())
	}
	if _, err := strconv.ParseUint(os.Args[2], 10, 64); err != nil {
		printErrorAndExit(err.Error())
	}

	fs, err := os.Open("test.fasta")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer fs.Close()

	fscan := bufio.NewScanner(fs)

	seqIdRead := false
	var seqId, peptide string
	result := peptideseq.NewPeptideSeqResults()

	// loop over every line in file
	for fscan.Scan() {
		if !seqIdRead {
			seqId = fscan.Text()
			seqIdRead = true
			continue
		}
		peptide = fscan.Text()

		// make new peptide object and calculate mass
		peps := peptideseq.NewPeptideSeq(seqId, peptide)
		err := peps.CalucalteMass()
		if err != nil {
			log.Fatal(err.Error())
		}
		seqIdRead = false

		// done, append and read next line
		result.Append(peps)
	}
	// print all results
	// TODO: Apply min - max range for mass
	for i := 0; i < int(result.Length()); i++ {
		fmt.Print(result.PrintCurrent())
	}
}
