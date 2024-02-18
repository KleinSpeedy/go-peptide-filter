package fastaproc_test

import (
	"peptide-analyse/fastaproc"
	peptideseq "peptide-analyse/peptide-seq"
	"sync"
	"testing"
	_ "unsafe"
)

func TestFilterWithRangeOk(t *testing.T) {
	// prepare filter channels and sync
	errChan := make(chan error, 1)
	resChan := make(chan peptideseq.PeptideSeq, 10)
	wg := new(sync.WaitGroup)

	filter := fastaproc.NewFilter(wg, resChan)

	wg.Add(1)
	go filter.FilterWithRange("../test_data/test.fasta", 0, 2000, errChan)
	wg.Wait()

	close(resChan)
	close(errChan)

	for err := range errChan {
		if err == nil {
			t.Error("No error expected")
		}
	}
}

func TestFilterWithRangeEmptyLine(t *testing.T) {
	// prepare filter channels and sync
	errChan := make(chan error, 1)
	resChan := make(chan peptideseq.PeptideSeq, 10)
	wg := new(sync.WaitGroup)

	filter := fastaproc.NewFilter(wg, resChan)

	wg.Add(1)
	go filter.FilterWithRange("../test_data/test_fail.fasta", 0, 2000, errChan)
	wg.Wait()

	close(resChan)
	close(errChan)

	for err := range errChan {
		if err == nil {
			t.Error("Error on empty string expected")
		}
	}
}

// readSequenceIdentifier is not exported, make visible for tests

//go:linkname p_readSeqId peptide-analyse/fastaproc.readSequenceIdentifier
func p_readSeqId(string) (string, error)

func TestReadSequenceIdentifier(t *testing.T) {
	const (
		seqId_withSpace = ">>sp_Q9H9K5_MER34_0 .50_HLA-A2501"
		seqId_empty     = ""
		seqId_noId      = "sp_Q9H9K5_MER34_0.50_HLA-A2501"
		seqId_normal    = ">>sp_Q9H9K5_MER34_0.50_HLA-A2501"
	)
	var (
		err error
		s   string
	)

	_, err = p_readSeqId(seqId_withSpace)
	if err == nil {
		t.Error("Read string with space, expected error")
	}

	_, err = p_readSeqId(seqId_empty)
	if err == nil {
		t.Error("Read empty string, expected error")
	}

	_, err = p_readSeqId(seqId_noId)
	if err == nil {
		t.Error("Read string without id sequence, expected error")
	}

	s, err = p_readSeqId(seqId_normal)
	if err != nil {
		t.Errorf("Recived error when nil was expected: %s", err.Error())
	} else if s == seqId_normal {
		t.Error("Returned identical seq id string, expected trimmed string")
	}
}
