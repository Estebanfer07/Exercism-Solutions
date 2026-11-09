package proteintranslation

import "errors"

var AminoAcidMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

var ErrInvalidBase = errors.New("ErrInvalidBase")
var ErrStop = errors.New("ErrStop")

func FromRNA(rna string) ([]string, error) {

	aminos := []string{}

	for idx := 0; idx < len(rna); idx += 3 {

		if idx+2 > len(rna)-1 {
			return []string{}, ErrInvalidBase
		}

		codon := string(rna[idx : idx+3])
		amino, ok := AminoAcidMap[codon]

		if !ok {
			return []string{}, ErrInvalidBase
		}

		if amino == "STOP" {
			break
		}

		aminos = append(aminos, amino)
	}

	return aminos, nil
}

func FromCodon(codon string) (string, error) {
	amino, ok := AminoAcidMap[codon]

	if !ok {
		return "", ErrInvalidBase
	}

	if amino == "STOP" {
		return "", ErrStop
	}

	return amino, nil
}
