package peptideseq

// Responsible for all operations analysing peptide sequences
// and calculating their respective masses

import (
	"fmt"
	"peptide-analyse/aminoacid"
	"strconv"
)

// represents a peptide sequence with its identifier
// and calculated mass
// read from a fasta file
type PeptideSeq struct {
	seqId   string  // sequence identifier
	peptide string  // peptide aminoacid sequence string
	mass    float64 // calculated mass of peptide
}

// represents a slice of read peptides
type PeptideSeqResults struct {
	seqeuences []PeptideSeq
	len        uint

	head uint
}

func NewPeptideSeq(seq string, peptide string) PeptideSeq {
	return PeptideSeq{
		seqId:   seq,
		peptide: peptide,
		mass:    0,
	}
}

func NewPeptideSeqResults() PeptideSeqResults {
	return PeptideSeqResults{
		seqeuences: make([]PeptideSeq, 0),
		len:        0,
		head:       0,
	}
}

// calculate the mass of the peptide sequence
// we calculate the sum of each aminoacids mass
// and substract it by a constant value
func (ps *PeptideSeq) CalucalteMass() error {
	var sum float64

	for _, v := range ps.peptide {
		// TODO: Is it ok to skip these
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

func (ps *PeptideSeq) GetMass() float64 {
	return ps.mass
}

// stringify contents of peptide sequence structure
// Note: does not conform to fasta specific structure
func (ps *PeptideSeq) String() string {
	return fmt.Sprintf("ID: %s\nPeptide: %s\n Mass: %f\n",
		ps.seqId, ps.peptide, ps.mass)
}

// appends a new peptide to results
func (res *PeptideSeqResults) Append(ps PeptideSeq) {
	res.seqeuences = append(res.seqeuences, ps)
	res.len++
}

// get number of peptides stored in results slice
func (res *PeptideSeqResults) Length() uint {
	return res.len
}

// prints the peptide in a fasta conform way
func (res *PeptideSeqResults) PrintCurrent() string {
	seqLine := res.seqeuences[res.head].seqId
	mass := "; " + strconv.FormatFloat(res.seqeuences[res.head].mass, 'f', -1, 64)
	pepLine := res.seqeuences[res.head].peptide
	res.head++

	return fmt.Sprintf("%s\n%s\n%s\n", seqLine, mass, pepLine)
}
