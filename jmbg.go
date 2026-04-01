// Package jmbg provides validation and parsing of Serbian unique master citizen numbers
// (JMBG - Jedinstveni Matični Broj Građana).
//
// JMBG format: DDMMYYYRRBBBC
//
//   - DD  - Day of birth (01-31)
//   - MM  - Month of birth (01-12)
//   - YYY - Year of birth (last 3 digits)
//   - RR  - Region code
//   - BBB - Unique number (000-499 male, 500-999 female)
//   - C   - Checksum digit
package jmbg

import (
	"fmt"
	"strings"
	"time"
)

// Jmbg holds a parsed and validated JMBG number with all its components.
type Jmbg struct {
	// Raw original input string
	Original string

	// Parsed integer fields
	Day      int
	Month    int
	Year     int
	Region   int
	Unique   int
	Checksum int

	// Original string parts (zero-padded as in the JMBG)
	DayOriginal    string
	MonthOriginal  string
	YearOriginal   string
	RegionOriginal string
	UniqueOriginal string

	// Derived fields
	RegionText string
	Country    string
	gender     Gender
	birthDate  time.Time
}

// Parse parses and validates a JMBG string. Returns a ValidationError if invalid.
func Parse(input string) (*Jmbg, error) {
	input = strings.TrimSpace(input)

	if len(input) != 13 {
		return nil, &ValidationError{
			Err:    ErrInvalidLength,
			Detail: fmt.Sprintf("expected 13 digits, got %d", len(input)),
		}
	}

	for _, c := range input {
		if c < '0' || c > '9' {
			return nil, &ValidationError{Err: ErrInvalidFormat}
		}
	}

	dd := twoDigitAt(input, 0)
	mm := twoDigitAt(input, 2)
	yyy := digitAt(input, 4)*100 + digitAt(input, 5)*10 + digitAt(input, 6)
	rr := twoDigitAt(input, 7)
	bbb := digitAt(input, 9)*100 + digitAt(input, 10)*10 + digitAt(input, 11)
	c := digitAt(input, 12)

	// Determine full year.
	year := 1000 + yyy
	if yyy < 800 {
		year = 2000 + yyy
	}

	// Validate date.
	date := time.Date(year, time.Month(mm), dd, 0, 0, 0, 0, time.UTC)
	if date.Day() != dd || int(date.Month()) != mm || date.Year() != year {
		return nil, &ValidationError{
			Err:    ErrInvalidDate,
			Detail: fmt.Sprintf("%02d/%02d/%d", dd, mm, year),
		}
	}

	// Validate region.
	reg, ok := regions[rr]
	if !ok {
		return nil, &ValidationError{
			Err:    ErrInvalidRegion,
			Detail: fmt.Sprintf("region code %d", rr),
		}
	}

	// Validate checksum (modulo 11).
	sum := 0
	for i := 0; i < 12; i++ {
		sum += digitAt(input, i) * checksumWeights[i]
	}
	remainder := sum % 11

	var expected int
	switch {
	case remainder == 0:
		expected = 0
	case remainder == 1:
		return nil, &ValidationError{Err: ErrInvalidChecksum}
	default:
		expected = 11 - remainder
	}
	if c != expected {
		return nil, &ValidationError{Err: ErrInvalidChecksum}
	}

	gender := Male
	if bbb >= 500 {
		gender = Female
	}

	return &Jmbg{
		Original:       input,
		Day:            dd,
		Month:          mm,
		Year:           year,
		Region:         rr,
		Unique:         bbb,
		Checksum:       c,
		DayOriginal:    input[0:2],
		MonthOriginal:  input[2:4],
		YearOriginal:   input[4:7],
		RegionOriginal: input[7:9],
		UniqueOriginal: input[9:12],
		RegionText:     reg.Name,
		Country:        reg.Country,
		gender:         gender,
		birthDate:      date,
	}, nil
}

// Valid reports whether the given string is a valid JMBG.
func Valid(input string) bool {
	_, err := Parse(input)
	return err == nil
}

// IsMale reports whether the JMBG belongs to a male person.
func (j *Jmbg) IsMale() bool {
	return j.gender == Male
}

// IsFemale reports whether the JMBG belongs to a female person.
func (j *Jmbg) IsFemale() bool {
	return j.gender == Female
}

// Age returns the current age in years.
func (j *Jmbg) Age() int {
	now := time.Now()
	years := now.Year() - j.birthDate.Year()
	if now.YearDay() < j.birthDate.YearDay() {
		years--
	}
	return years
}

// Date returns the birth date.
func (j *Jmbg) Date() time.Time {
	return j.birthDate
}

// String implements the fmt.Stringer interface.
func (j *Jmbg) String() string {
	return j.Original
}

// Gender returns the gender of the person.
func (j *Jmbg) Gender() Gender {
	return j.gender
}

// IsAdult reports whether the person is 18 or older.
func (j *Jmbg) IsAdult() bool {
	return j.Age() >= 18
}
