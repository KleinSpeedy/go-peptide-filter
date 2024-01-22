package fastaproc

import (
	"bufio"
	"fmt"
	"os"
	peptideseq "peptide-analyse/peptide-seq"
	"strings"
	"sync"
)

type Filter struct {
	out chan<- peptideseq.PeptideSeq
	wg  *sync.WaitGroup
}

type OutputFunc func(peptideseq.PeptideSeq) error

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

	str := strings.TrimPrefix(s, peptideseq.FastaSeqIdPrefix)
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
			errChan <- err
			return
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

func WritePeptideSequences(out OutputFunc, results <-chan peptideseq.PeptideSeq, errChan chan error, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		ps, ok := <-results
		if !ok {
			break
		}
		err := out(ps)
		if err != nil {
			errChan <- err
			return
		}
	}
}
