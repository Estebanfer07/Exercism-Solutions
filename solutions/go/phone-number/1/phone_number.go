package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type PhoneNumber struct {
	AreaCode   string
	FirstPart  string
	SecondPart string
	FullNumber string
}

var (
	cleanRegex = regexp.MustCompile(`[.\-() +]`)
	phoneRegex = regexp.MustCompile(`^([2-9]\d{2})([2-9]\d{2})(\d{4})$`)
)

func Clean(phoneNumber string) (PhoneNumber, error) {
	phoneNumber = cleanRegex.ReplaceAllString(phoneNumber, "")

	if len(phoneNumber) == 11 && strings.HasPrefix(phoneNumber, "1") {
		phoneNumber = strings.TrimPrefix(phoneNumber, "1")
	}

	match := phoneRegex.FindStringSubmatch(phoneNumber)

	if match == nil {
		return PhoneNumber{}, errors.New("invalid phone number")
	}

	return PhoneNumber{
		AreaCode:   match[1],
		FirstPart:  match[2],
		SecondPart: match[3],
		FullNumber: match[0],
	}, nil
}

func Number(phoneNumber string) (string, error) {

	number, err := Clean(phoneNumber)

	return number.FullNumber, err
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Clean(phoneNumber)

	if err != nil {
		return "", errors.New("Invalid phone number")
	}

	return number.AreaCode, nil
}

func Format(phoneNumber string) (string, error) {
	number, err := Clean(phoneNumber)

	if err != nil {
		return "", errors.New("Invalid phone number")
	}

	return fmt.Sprintf("(%s) %s-%s", number.AreaCode, number.FirstPart, number.SecondPart), nil
}
