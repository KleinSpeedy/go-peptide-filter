package peptideseq

import (
	"fmt"
	"peptide-analyse/aminoacid"
	"strconv"
)

type PeptideSeq struct {
	seqId   string  // sequence Id
	peptide string  // Peptide an sich
	mass    float64 // masse des peptids (Summe)
}

type PeptideSeqResults struct {
	seqeuences []PeptideSeq
	len        uint

	head uint
}

func NewPeptideSeq(seq string, peptide string) PeptideSeq {
	pepseq := PeptideSeq{
		seqId:   seq,
		peptide: peptide,
		mass:    0,
	}

	return pepseq
}

func NewPeptideSeqResults() PeptideSeqResults {
	return PeptideSeqResults{
		seqeuences: make([]PeptideSeq, 0),
		len:        0,
		head:       0,
	}
}

func (ps *PeptideSeq) CalucalteMass() error {
	var sum float64

	// ABC
	for _, v := range ps.peptide {
		if v == '-' || v == '*' || v == 'X' {
			continue
		}

		pepMass, err := aminoacid.GetAminoacidMass(byte(v))
		if err != nil {
			return err
		}

		sum = sum + float64(pepMass)
	}
	val := (len(ps.peptide) - 2) * 18
	ps.mass = sum - float64(val)
	return nil
}

func (ps *PeptideSeq) String() string {
	return fmt.Sprintf("ID: %s\nPeptide: %s\n Mass: %f\n",
		ps.seqId, ps.peptide, ps.mass)
}

func (res *PeptideSeqResults) Append(ps PeptideSeq) {
	res.seqeuences = append(res.seqeuences, ps)
	res.len++
}

func (res *PeptideSeqResults) Length() uint {
	return res.len
}

func (res *PeptideSeqResults) PrintCurrent() string {
	seqLine := res.seqeuences[res.head].seqId
	mass := "; " + strconv.FormatFloat(res.seqeuences[res.head].mass, 'f', -1, 64)
	pepLine := res.seqeuences[res.head].peptide
	res.head++

	return fmt.Sprintf("%s\n%s\n%s\n", seqLine, mass, pepLine)
}
