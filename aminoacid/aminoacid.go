package aminoacid

import "fmt"

// aminoacid map assigning mass to
// single-letter aminoacid codes
var aminoacidMassMap = map[byte]uint{
	'A': 89,
	'C': 121,
	'D': 133,
	'E': 147,
	'F': 165,
	'G': 75,
	'H': 155,
	'I': 131,
	'K': 146,
	'L': 131,
	'M': 149,
	'N': 132,
	'P': 115,
	'Q': 146,
	'R': 174,
	'S': 105,
	'T': 119,
	'U': 168,
	'V': 117,
	'W': 204,
	'Y': 181,
}

// check whether the one-letter code is a aminoacid
func IsAminoacid(aa byte) bool {
	// ignore error
	_, ok := aminoacidMassMap[aa]
	return ok
}

// get mass from aminoacid mass map
// returns error if aminoacid does not exist
func GetAminoacidMass(aa byte) (uint, error) {
	mass, ok := aminoacidMassMap[aa]
	if !ok {
		return 0, fmt.Errorf("not an aminoacid: %s\n", string(aa))
	}
	return mass, nil
}
