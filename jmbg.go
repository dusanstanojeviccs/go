// Package main provides validation and parsing of Serbian unique master citizen numbers
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
package main

import (
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

// Parse parses and validates a JMBG string. Returns an error if invalid.
func Parse(input string) (*Jmbg, error) {
	input = strings.TrimSpace(input)
	
	if len(input) != 13 {
		return nil, newError("JMBG string must have exactly 13 digits, got %d", len(input))
	}

	for _, c := range input {
		if c < '0' || c > '9' {
			return nil, newError("JMBG must contain only numeric characters")
		}
	}

	digit := func(i int) int { return int(input[i] - '0') }
	twoDigit := func(i int) int { return digit(i)*10 + digit(i+1) }

	dd := twoDigit(0)
	mm := twoDigit(2)
	yyy := digit(4)*100 + digit(5)*10 + digit(6)
	rr := twoDigit(7)
	bbb := digit(9)*100 + digit(10)*10 + digit(11)
	c := digit(12)

	// Determine full year
	year := 1000 + yyy
	if yyy < 800 {
		year = 2000 + yyy
	}

	// Validate date
	date := time.Date(year, time.Month(mm), dd, 0, 0, 0, 0, time.UTC)
	if date.Day() != dd || int(date.Month()) != mm || date.Year() != year {
		return nil, newError("Date '%02d/%02d/%d' is not valid", dd, mm, year)
	}

	// Validate region
	reg, ok := regions[rr]
	if !ok {
		return nil, newError("Region '%d' is not valid for JMBG", rr)
	}

	// Validate checksum (modulo 11)
	w := []int{7, 6, 5, 4, 3, 2, 7, 6, 5, 4, 3, 2}
	sum := 0
	for i := 0; i < 12; i++ {
		sum += digit(i) * w[i]
	}
	remainder := sum % 11
	var expectedChecksum int
	if remainder == 0 {
		expectedChecksum = 0
	} else if remainder == 1 {
		return nil, newError("Checksum is not valid")
	} else {
		expectedChecksum = 11 - remainder
	}
	if c != expectedChecksum {
		return nil, newError("Checksum is not valid")
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

// Valid returns true if the given string is a valid JMBG.
func Valid(input string) bool {
	_, err := Parse(input)
	return err == nil
}

// IsMale returns true if the JMBG belongs to a male person.
func (j *Jmbg) IsMale() bool {
	return j.gender == Male
}

// IsFemale returns true if the JMBG belongs to a female person.
func (j *Jmbg) IsFemale() bool {
	return j.gender == Female
}

// GetAge returns the current age in years.
func (j *Jmbg) GetAge() int {
	now := time.Now()
	years := now.Year() - j.birthDate.Year()
	if now.YearDay() < j.birthDate.YearDay() {
		years--
	}
	return years
}

// GetDate returns the birth date.
func (j *Jmbg) GetDate() time.Time {
	return j.birthDate
}

// Format returns the JMBG as a 13-digit string.
func (j *Jmbg) Format() string {
	return j.Original
}

// String implements the fmt.Stringer interface.
func (j *Jmbg) String() string {
	return j.Original
}

// Gender returns the gender of the person.
func (j *Jmbg) Gender() Gender {
	return j.gender
}

// IsAdult returns true if the person is 18 or older.
func (j *Jmbg) IsAdult() bool {
	return j.GetAge() >= 18
}
