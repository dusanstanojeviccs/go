# JMBG Go Library

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/Go-%5E1.21-blue)](https://go.dev/)

A Go library for validating and parsing Serbian unique master citizen numbers (JMBG - Jedinstveni Matični Broj Građana).

## Features

- ✅ Validate JMBG numbers with comprehensive checks
- ✅ Extract birth date, region, and gender information
- ✅ Support for all Serbian and ex-Yugoslav regions
- ✅ Calculate age from JMBG
- ✅ Check adult status (18+ years)
- ✅ Zero external dependencies
- ✅ Fully tested with Go testing framework

## Installation

```bash
go get github.com/jmbg-labs/go
```

## Usage

### Basic Validation

```go
import jmbg "github.com/jmbg-labs/go"

// Quick validation
if jmbg.Valid("0101000710009") {
    fmt.Println("Valid JMBG")
}

// Parse and validate
j, err := jmbg.Parse("0101000710009")
if err != nil {
    fmt.Println("Invalid:", err)
    return
}
fmt.Println("Valid:", j.Format())
```

### Extract Information

```go
j, err := jmbg.Parse("0101000710009")
if err != nil {
    log.Fatal(err)
}

// Birth date
date := j.GetDate() // time.Time
fmt.Println(date.Format("2006-01-02")) // 2000-01-01

// Age
fmt.Println(j.GetAge()) // e.g., 26

// Gender
if j.IsMale() {
    fmt.Println("Male")
}
if j.IsFemale() {
    fmt.Println("Female")
}
fmt.Println(j.Gender()) // "Male" or "Female"

// Adult check
fmt.Println(j.IsAdult()) // true/false

// Individual fields
fmt.Println(j.Day)           // 1
fmt.Println(j.Month)         // 1
fmt.Println(j.Year)          // 2000
fmt.Println(j.Region)        // 71
fmt.Println(j.RegionText)    // "Belgrade"
fmt.Println(j.Country)       // "Serbia"
fmt.Println(j.Unique)        // 0
fmt.Println(j.Checksum)      // 9
```

### Original String Parts

```go
j, _ := jmbg.Parse("1505995800002")

fmt.Println(j.Original)        // "1505995800002"
fmt.Println(j.DayOriginal)     // "15"
fmt.Println(j.MonthOriginal)   // "05"
fmt.Println(j.YearOriginal)    // "995"
fmt.Println(j.RegionOriginal)  // "80"
fmt.Println(j.UniqueOriginal)  // "000"

// String conversion
fmt.Println(j.Format())        // "1505995800002"
fmt.Println(j.String())        // "1505995800002"
fmt.Printf("%s\n", j)          // "1505995800002"
```

## JMBG Format

JMBG consists of 13 digits: `DDMMYYYRRBBBC`

| Part | Length | Description |
|------|--------|-------------|
| DD   | 2      | Day of birth (01–31) |
| MM   | 2      | Month of birth (01–12) |
| YYY  | 3      | Year of birth (last 3 digits) |
| RR   | 2      | Region code |
| BBB  | 3      | Unique number (000–499 male, 500–999 female) |
| C    | 1      | Checksum digit |

### Year Decoding

- YYY < 800 → `2000 + YYY` (e.g., 000 → 2000)
- YYY ≥ 800 → `1000 + YYY` (e.g., 978 → 1978)

### Supported Regions

The library supports all Serbian and ex-Yugoslav regions including (beware: ex-Yugoslav regions codes may have changed since the breakup):

- **Serbia** (71-79): Belgrade, Kragujevac, Niš, etc.
- **Serbia/Vojvodina** (80-89): Novi Sad, Subotica, Pančevo, etc.
- **Serbia/Kosovo** (91-96): Priština, Peć, Prizren, etc.
- **Bosnia and Herzegovina** (10-19)
- **Montenegro** (21-29)
- **Croatia** (30-39)
- **Macedonia** (41-49)
- **Slovenia** (50)

## Validation Rules

The library performs comprehensive validation:

1. **Length check** — Must be exactly 13 digits
2. **Format check** — Must contain only numeric characters
3. **Date validation** — Birth date must be valid (including leap year support)
4. **Region validation** — Region code must exist in the registry
5. **Checksum validation** — Modulo 11 algorithm verification

## Error Handling

```go
j, err := jmbg.Parse("1234567890123")
if err != nil {
    fmt.Println(err)
    // Possible messages:
    // - "JMBG string must have exactly 13 digits, got N"
    // - "JMBG must contain only numeric characters"
    // - "Date 'DD/MM/YYYY' is not valid"
    // - "Region 'N' is not valid for JMBG"
    // - "Checksum is not valid"
}
```

## Development

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...
```

## Requirements

- Go ^1.21
- No external dependencies (for production use)

## Contributing

Contributions are welcome! Please ensure:

1. All tests pass (`go test ./...`)
2. Code follows Go conventions (`go fmt`, `go vet`)
3. Add tests for new features
4. Update documentation as needed

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Credits

Developed by [JMBG Labs](https://github.com/jmbg-labs)

## Support

- 🐛 [Report Issues](https://github.com/jmbg-labs/go/issues)
- 📖 [Source Code](https://github.com/jmbg-labs/go)

## Examples

### Validate Multiple JMBGs

```go
package main

import (
	"fmt"
	jmbg "github.com/jmbg-labs/go"
)

func main() {
	jmbgs := []string{"0710003730015", "1705978730032", "invalid"}

	for _, jmbgString := range jmbgs {
		if jmbg.Valid(jmbgString) {
			j, _ := jmbg.Parse(jmbgString)
			fmt.Printf(
				"%s - Born: %s, Region: %s, Gender: %s\n",
				j.Format(),
				j.GetDate().Format("2006-01-02"),
				j.RegionText,
				j.Gender(),
			)
		} else {
			fmt.Printf("%s - Invalid\n", jmbgString)
		}
	}
}
```

### Age Calculation and Adult Check

```go
j, err := jmbg.Parse("0710003730015")
if err != nil {
	log.Fatal(err)
}

age := j.GetAge()
fmt.Printf("Age: %d\n", age)

if j.IsAdult() {
	fmt.Println("Adult (18+)")
} else {
	fmt.Println("Minor")
}
```