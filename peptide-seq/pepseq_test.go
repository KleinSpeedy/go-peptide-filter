package peptideseq_test

import (
	peptideseq "peptide-analyse/peptide-seq"
	"testing"
)

func TestPepseqCalculateMassOk(t *testing.T) {
	psOk := peptideseq.NewPeptideSeq(">>test_valid", "ELVFVPASA")

	err := psOk.CalucalteMass()
	if err != nil {
		t.Errorf("error: %s, expected nil", err.Error())
	}

	if !(psOk.GetMass() > 0) {
		t.Errorf("Calulated mass wrong %f, expected ...", psOk.GetMass())
	}
}

func TestPepseqCalculateMassInvalid(t *testing.T) {
	psInvalid := peptideseq.NewPeptideSeq(">>test_invalid", "ELVXXPABA")

	err := psInvalid.CalucalteMass()
	if err == nil {
		t.Errorf("Got no error but expected calculation to fail")
	}
}

func TestPeptideSeqResultsLength(t *testing.T) {
	res := peptideseq.NewPeptideSeqResults()
	if res.Length() != 0 {
		t.Errorf("Length of pepeseq result is %d, expected 0", res.Length())
	}

	temp := peptideseq.NewPeptideSeq("abc", "ABC")
	res.Append(temp)

	if res.Length() != 1 {
		t.Errorf("Length of pepeseq result is %d, expected 1", res.Length())
	}
}
