package main

import (
	"strings"
	"testing"
	"time"
)

func TestValidJmbgCanBeParsed(t *testing.T) {
	j, err := Parse("0710003730015")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if j == nil {
		t.Error("expected non-nil Jmbg instance")
	}
}

func TestValidJmbgReturnsTrue(t *testing.T) {
	if !Valid("0710003730015") {
		t.Error("expected Valid to return true for valid JMBG")
	}
}

func TestInvalidJmbgReturnsFalse(t *testing.T) {
	if Valid("1234567890123") {
		t.Error("expected Valid to return false for invalid JMBG")
	}
}

func TestInvalidLengthThrowsException(t *testing.T) {
	_, err := Parse("123456789")
	if err == nil {
		t.Fatal("expected error for invalid length")
	}
	if !strings.Contains(err.Error(), "13 digits") {
		t.Errorf("expected error message about 13 digits, got: %v", err)
	}
}

func TestNonNumericJmbgThrowsException(t *testing.T) {
	_, err := Parse("01019907100ab")
	if err == nil {
		t.Fatal("expected error for non-numeric input")
	}
	if !strings.Contains(err.Error(), "numeric") {
		t.Errorf("expected error message about numeric characters, got: %v", err)
	}
}

func TestInvalidDateThrowsException(t *testing.T) {
	_, err := Parse("3201990710009")
	if err == nil {
		t.Fatal("expected error for invalid date")
	}
	if !strings.Contains(err.Error(), "not valid") {
		t.Errorf("expected error message about invalid date, got: %v", err)
	}
}

func TestInvalidRegionThrowsException(t *testing.T) {
	_, err := Parse("0710003660015")
	if err == nil {
		t.Fatal("expected error for invalid region")
	}
	if !strings.Contains(err.Error(), "Region") && !strings.Contains(err.Error(), "66") {
		t.Errorf("expected error message about region 66, got: %v", err)
	}
}

func TestInvalidChecksumThrowsException(t *testing.T) {
	_, err := Parse("0710003730025")
	if err == nil {
		t.Fatal("expected error for invalid checksum")
	}
	if !strings.Contains(err.Error(), "Checksum") {
		t.Errorf("expected error message about checksum, got: %v", err)
	}
}

func TestIsMaleReturnsTrue(t *testing.T) {
	j, err := Parse("0710003730015")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !j.IsMale() {
		t.Error("expected IsMale to return true")
	}
}

func TestIsMaleReturnsFalse(t *testing.T) {
	j, err := Parse("0710003735017")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if j.IsMale() {
		t.Error("expected IsMale to return false")
	}
}

func TestIsFemaleReturnsTrue(t *testing.T) {
	j, err := Parse("0710003735017")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !j.IsFemale() {
		t.Error("expected IsFemale to return true")
	}
}

func TestIsFemaleReturnsFalse(t *testing.T) {
	j, err := Parse("0710003730015")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if j.IsFemale() {
		t.Error("expected IsFemale to return false")
	}
}

func TestGetAge(t *testing.T) {
	j, err := Parse("0710003730015")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	birthDate := time.Date(2003, 10, 7, 0, 0, 0, 0, time.UTC)
	now := time.Now()
	expectedAge := now.Year() - birthDate.Year()
	if now.YearDay() < birthDate.YearDay() {
		expectedAge--
	}
	if j.GetAge() != expectedAge {
		t.Errorf("GetAge() = %d, want %d", j.GetAge(), expectedAge)
	}
}

func TestGetDate(t *testing.T) {
	j, err := Parse("0710003730015")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	date := j.GetDate()
	if date.Year() != 2003 || date.Month() != time.October || date.Day() != 7 {
		t.Errorf("GetDate() = %v, want 2003-10-07", date)
	}
}

func TestFormat(t *testing.T) {
	j, err := Parse("0710003730015")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if j.Format() != "0710003730015" {
		t.Errorf("Format() = %q, want %q", j.Format(), "0710003730015")
	}
}

func TestToString(t *testing.T) {
	j, err := Parse("0710003730015")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if j.String() != "0710003730015" {
		t.Errorf("String() = %q, want %q", j.String(), "0710003730015")
	}
}

func TestMagicGetters(t *testing.T) {
	j, err := Parse("2902992710005")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if j.Original != "2902992710005" {
		t.Errorf("Original = %q, want %q", j.Original, "2902992710005")
	}
	if j.Day != 29 {
		t.Errorf("Day = %d, want 29", j.Day)
	}
	if j.DayOriginal != "29" {
		t.Errorf("DayOriginal = %q, want %q", j.DayOriginal, "29")
	}
	if j.Month != 2 {
		t.Errorf("Month = %d, want 2", j.Month)
	}
	if j.MonthOriginal != "02" {
		t.Errorf("MonthOriginal = %q, want %q", j.MonthOriginal, "02")
	}
	if j.Year != 1992 {
		t.Errorf("Year = %d, want 1992", j.Year)
	}
	if j.YearOriginal != "992" {
		t.Errorf("YearOriginal = %q, want %q", j.YearOriginal, "992")
	}
	if j.Region != 71 {
		t.Errorf("Region = %d, want 71", j.Region)
	}
	if j.RegionOriginal != "71" {
		t.Errorf("RegionOriginal = %q, want %q", j.RegionOriginal, "71")
	}
	if j.RegionText != "Belgrade" {
		t.Errorf("RegionText = %q, want %q", j.RegionText, "Belgrade")
	}
	if j.Country != "Serbia" {
		t.Errorf("Country = %q, want %q", j.Country, "Serbia")
	}
	if j.Unique != 0 {
		t.Errorf("Unique = %d, want 0", j.Unique)
	}
	if j.UniqueOriginal != "000" {
		t.Errorf("UniqueOriginal = %q, want %q", j.UniqueOriginal, "000")
	}
	if j.Checksum != 5 {
		t.Errorf("Checksum = %d, want 5", j.Checksum)
	}
}

func TestTrimWhitespace(t *testing.T) {
	j, err := Parse("  0710003730015  ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if j.Format() != "0710003730015" {
		t.Errorf("Format() = %q, want %q", j.Format(), "0710003730015")
	}
}

func TestYearCalculationFor2000s(t *testing.T) {
	j, err := Parse("0101000710009")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if j.Year != 2000 {
		t.Errorf("Year = %d, want 2000", j.Year)
	}
}

func TestYearCalculationFor1900s(t *testing.T) {
	j, err := Parse("1705978730032")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if j.Year != 1978 {
		t.Errorf("Year = %d, want 1978", j.Year)
	}
}

func TestDifferentRegions(t *testing.T) {
	j1, err := Parse("2902992710005")
	if err != nil {
		t.Fatalf("unexpected error for Belgrade: %v", err)
	}
	if j1.RegionText != "Belgrade" {
		t.Errorf("RegionText = %q, want %q", j1.RegionText, "Belgrade")
	}
	if j1.Country != "Serbia" {
		t.Errorf("Country = %q, want %q", j1.Country, "Serbia")
	}

	j2, err := Parse("1505995800002")
	if err != nil {
		t.Fatalf("unexpected error for Novi Sad: %v", err)
	}
	if j2.RegionText != "Novi Sad" {
		t.Errorf("RegionText = %q, want %q", j2.RegionText, "Novi Sad")
	}
	if j2.Country != "Serbia/Vojvodina" {
		t.Errorf("Country = %q, want %q", j2.Country, "Serbia/Vojvodina")
	}
}

func TestBoundaryUniqueNumbers(t *testing.T) {
	j1, err := Parse("1505995800002")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !j1.IsMale() {
		t.Error("expected male for unique 0")
	}
	if j1.Unique != 0 {
		t.Errorf("Unique = %d, want 0", j1.Unique)
	}

	j2, err := Parse("1505995804997")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !j2.IsMale() {
		t.Error("expected male for unique 499")
	}
	if j2.Unique != 499 {
		t.Errorf("Unique = %d, want 499", j2.Unique)
	}

	j3, err := Parse("1505995805004")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !j3.IsFemale() {
		t.Error("expected female for unique 500")
	}
	if j3.Unique != 500 {
		t.Errorf("Unique = %d, want 500", j3.Unique)
	}

	j4, err := Parse("1505995809999")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !j4.IsFemale() {
		t.Error("expected female for unique 999")
	}
	if j4.Unique != 999 {
		t.Errorf("Unique = %d, want 999", j4.Unique)
	}
}

func TestLeapYearDate(t *testing.T) {
	j, err := Parse("2902992710005")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if j.Day != 29 {
		t.Errorf("Day = %d, want 29", j.Day)
	}
	if j.Month != 2 {
		t.Errorf("Month = %d, want 2", j.Month)
	}
	if j.Year != 1992 {
		t.Errorf("Year = %d, want 1992", j.Year)
	}
}

func TestInvalidLeapYearDateThrowsException(t *testing.T) {
	_, err := Parse("2902979758318")
	if err == nil {
		t.Fatal("expected error for invalid leap year date")
	}
	if !strings.Contains(err.Error(), "not valid") {
		t.Errorf("expected error message about invalid date, got: %v", err)
	}
}
