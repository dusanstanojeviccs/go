# Release Notes

## v0.1.0 (2026-03-30)

### 🎉 Initial Release

First public release of the JMBG Go library for validating and parsing Serbian unique master citizen numbers (JMBG - Jedinstveni Matični Broj Građana).

### ✨ Features

- **Validation**: Comprehensive JMBG validation with length, format, date, region, and checksum checks
- **Parsing**: Extract all components from JMBG numbers (day, month, year, region, unique number, checksum)
- **Birth Date**: Get birth date as `time.Time` with proper century calculation
- **Age Calculation**: Calculate current age from JMBG
- **Gender Detection**: Determine gender from the unique number field (000-499 male, 500-999 female)
- **Adult Check**: Verify if person is 18+ years old
- **Region Support**: Full support for all Serbian and ex-Yugoslav regions including:
  - Serbia (71-79)
  - Serbia/Vojvodina (80-89)
  - Serbia/Kosovo (91-96)
  - Bosnia and Herzegovina (10-19)
  - Montenegro (21-29)
  - Croatia (30-39)
  - Macedonia (41-49)
  - Slovenia (50)
- **Zero Dependencies**: No external dependencies required
- **Type Safety**: Strongly typed API with proper error handling
- **Original Preservation**: Maintains original string parts with zero-padding

### 📦 Installation

```bash
go get github.com/jmbg-labs/go@v0.1.0
```

### 🔧 Usage Example

```go
import jmbg "github.com/jmbg-labs/go"

// Validate and parse
j, err := jmbg.Parse("0101000710009")
if err != nil {
    log.Fatal(err)
}

fmt.Println(j.GetDate())     // 2000-01-01
fmt.Println(j.GetAge())      // Current age
fmt.Println(j.Gender())      // "Male" or "Female"
fmt.Println(j.RegionText)    // "Belgrade"
fmt.Println(j.IsAdult())     // true/false
```

### 📋 Requirements

- Go 1.21 or higher
- No external dependencies

### 📄 License

MIT License - see [LICENSE](LICENSE) file for details

### 🙏 Credits

Developed by [JMBG Labs](https://github.com/jmbg-labs)
