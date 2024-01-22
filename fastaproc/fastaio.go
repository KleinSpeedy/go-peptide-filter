package fastaproc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	peptideseq "peptide-analyse/peptide-seq"
	"strings"
	"sync"
)

// TODO: Let prefix be specified on program startup
// fasta format usually starts with one >, we have 2
const seqIdPrefix string = ">>"

// comments are indicated by a semicolon and should be put
// between seq ID and peptide seq
const fastaComment string = ";"

type Filter struct {
	out chan<- peptideseq.PeptideSeq
	wg  *sync.WaitGroup
}

func NewFilter(wg *sync.WaitGroup, results chan<- peptideseq.PeptideSeq) *Filter {
	return &Filter{
		out: results,
		wg:  wg,
	}
}

func readSequenceIdentifier(s string) (string, error) {
	if s == "" {
		return "", fmt.Errorf("read an empty string")
	}

	str := strings.TrimPrefix(s, seqIdPrefix)
	if str == s {
		return "", fmt.Errorf("sequence identifier did not have start prefix")
	}

	return str, nil
}

func (fl *Filter) FilterWithRange(filename string, start, end float64, errChan chan<- error) {
	defer fl.wg.Done()

	fd, err := os.Open(filename)
	if err != nil {
		errChan <- err
		return
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	if scanner == nil {
		errChan <- fmt.Errorf("Error allocating new scanner")
		return
	}

	seqIdRead := false
	var seqId, pep string

	for scanner.Scan() {
		// first read sequence id on its own line
		if !seqIdRead {
			seqId, err = readSequenceIdentifier(scanner.Text())
			if err != nil {
				errChan <- err
				return
			}

			seqIdRead = true
			continue
		}
		// after reading sequence id read peptide sequence
		pep = scanner.Text()

		ps := peptideseq.NewPeptideSeq(seqId, pep)
		if err = ps.CalucalteMass(); err != nil {
			log.Fatal("error computing mass")
		}

		// send new peptide if mass is in range
		if ps.MassIsInRange(start, end) {
			fl.out <- ps
		}

		// reset for next iteration
		seqIdRead = false
		seqId = ""
	}
}

func WritePeptideSequencesToFile(name string, results <-chan peptideseq.PeptideSeq, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		ps, ok := <-results
		if !ok {
			break
		}
		fmt.Println(ps.String())
	}
}
