package peptideseq

// Responsible for all operations analysing peptide sequences
// and calculating their respective masses

import (
	"fmt"
	"peptide-analyse/aminoacid"
	"strconv"
)

// TODO: Let prefix be specified on program startup
// fasta format usually starts with one >, we have 2
const FastaSeqIdPrefix string = ">>"

// comments are indicated by a semicolon and should be put
// between seq ID and peptide seq
const FastaComment string = "; "

// represents a peptide sequence with its identifier
// and calculated mass
// read from a fasta file
type PeptideSeq struct {
	seqId   string  // sequence identifier
	peptide string  // peptide aminoacid sequence string
	mass    float64 // calculated mass of peptide
}

func NewPeptideSeq(seq string, peptide string) PeptideSeq {
	return PeptideSeq{
		seqId:   seq,
		peptide: peptide,
		mass:    0,
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

	ps.mass = sum - float64(len(ps.peptide)-1)*18.02
	return nil
}

func (ps *PeptideSeq) GetMass() float64 {
	return ps.mass
}

func (ps *PeptideSeq) MassIsInRange(start, end float64) bool {
	return (ps.mass >= start) && (ps.mass <= end)
}

// stringify contents of peptide sequence structure
// Note: does not conform to fasta specific structure
func (ps *PeptideSeq) String() string {
	return fmt.Sprintf("ID: %s\nPeptide: %s\nMass: %f",
		ps.seqId, ps.peptide, ps.mass)
}

// Prints the peptide with its sequence ID in a fasta conform way
// does not print the mass as a comment
func (ps *PeptideSeq) Write() string {
	return fmt.Sprintf("%s\n%s\n", (FastaSeqIdPrefix + ps.seqId), ps.peptide)
}

// Prints the peptide with its sequence ID in a fasta conform way
// adds the peptide mass as a comment
func (ps *PeptideSeq) WriteWithComment() string {
	id := (FastaSeqIdPrefix + ps.seqId)
	massStr := FastaComment + strconv.FormatFloat(ps.mass, 'f', -1, 64)
	return fmt.Sprintf("%s\n%s\n%s\n", id, massStr, ps.peptide)
}
