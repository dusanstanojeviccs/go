package main

// Gender represents the gender derived from JMBG.
type Gender int

const (
	Male   Gender = iota
	Female Gender = iota
)

// String returns a human-readable gender string.
func (g Gender) String() string {
	if g == Male {
		return "Male"
	}
	return "Female"
}

// Region holds information about a political/geographic region.
type Region struct {
	Code    int
	Name    string
	Country string
}

// regions maps region codes to Region structs.
var regions = map[int]Region{
	// Foreign citizens
	0:  {0, "naturalized citizens which had no republican citizenship", "foreign citizens"},
	1:  {1, "foreigners in Bosnia and Herzegovina", "foreign citizens"},
	2:  {2, "foreigners in Montenegro", "foreign citizens"},
	3:  {3, "foreigners in Croatia", "foreign citizens"},
	4:  {4, "foreigners in Macedonia", "foreign citizens"},
	5:  {5, "foreigners in Slovenia", "foreign citizens"},
	6:  {6, "foreigners in Serbia", "foreign citizens"},
	7:  {7, "foreigners in Serbia/Vojvodina", "foreign citizens"},
	8:  {8, "foreigners in Serbia/Kosovo", "foreign citizens"},
	9:  {9, "naturalized citizens which had no republican citizenship", "foreign citizens"},
	// Bosnia and Herzegovina
	10: {10, "Banja Luka", "Bosnia and Herzegovina"},
	11: {11, "Bihać", "Bosnia and Herzegovina"},
	12: {12, "Doboj", "Bosnia and Herzegovina"},
	13: {13, "Goražde", "Bosnia and Herzegovina"},
	14: {14, "Livno", "Bosnia and Herzegovina"},
	15: {15, "Mostar", "Bosnia and Herzegovina"},
	16: {16, "Prijedor", "Bosnia and Herzegovina"},
	17: {17, "Sarajevo", "Bosnia and Herzegovina"},
	18: {18, "Tuzla", "Bosnia and Herzegovina"},
	19: {19, "Zenica", "Bosnia and Herzegovina"},
	// Montenegro
	21: {21, "Podgorica", "Montenegro"},
	22: {22, "Bar, Ulcinj", "Montenegro"},
	23: {23, "Budva, Kotor, Tivat", "Montenegro"},
	24: {24, "Herceg Novi", "Montenegro"},
	25: {25, "Cetinje", "Montenegro"},
	26: {26, "Nikšić", "Montenegro"},
	27: {27, "Berane, Rožaje, Plav, Andrijevica", "Montenegro"},
	28: {28, "Bijelo Polje, Mojkovac", "Montenegro"},
	29: {29, "Pljevlja, Žabljak", "Montenegro"},
	// Croatia
	30: {30, "Osijek, Slavonija", "Croatia"},
	31: {31, "Bjelovar, Virovitica, Koprivnica, Pakrac, Podravina", "Croatia"},
	32: {32, "Varaždin, Međimurje", "Croatia"},
	33: {33, "Zagreb", "Croatia"},
	34: {34, "Karlovac, Kordun", "Croatia"},
	35: {35, "Gospić, Lika", "Croatia"},
	36: {36, "Rijeka, Pula, Gorski kotar, Istra", "Croatia"},
	37: {37, "Sisak, Banovina", "Croatia"},
	38: {38, "Split, Zadar, Šibenik, Dubrovnik, Dalmacija", "Croatia"},
	39: {39, "Hrvatsko Zagorje", "Croatia"},
	// North Macedonia
	41: {41, "Bitola", "Macedonia"},
	42: {42, "Kumanovo", "Macedonia"},
	43: {43, "Ohrid", "Macedonia"},
	44: {44, "Prilep", "Macedonia"},
	45: {45, "Skopje", "Macedonia"},
	46: {46, "Strumica", "Macedonia"},
	47: {47, "Tetovo", "Macedonia"},
	48: {48, "Veles", "Macedonia"},
	49: {49, "Štip", "Macedonia"},
	// Slovenia
	50: {50, "Slovenia", "Slovenia"},
	// Serbia
	71: {71, "Belgrade", "Serbia"},
	72: {72, "Kragujevac", "Serbia"},
	73: {73, "Niš", "Serbia"},
	74: {74, "Leskovac, Vranje", "Serbia"},
	75: {75, "Zaječar, Bor", "Serbia"},
	76: {76, "Smederevo, Požarevac", "Serbia"},
	77: {77, "Mačva, Kolubara", "Serbia"},
	78: {78, "Čačak, Kraljevo, Kruševac", "Serbia"},
	79: {79, "Užice", "Serbia"},
	// Serbia / Vojvodina
	80: {80, "Novi Sad", "Serbia/Vojvodina"},
	81: {81, "Sombor", "Serbia/Vojvodina"},
	82: {82, "Subotica", "Serbia/Vojvodina"},
	83: {83, "Vrbas", "Serbia/Vojvodina"},
	84: {84, "Kikinda", "Serbia/Vojvodina"},
	85: {85, "Zrenjanin", "Serbia/Vojvodina"},
	86: {86, "Pančevo", "Serbia/Vojvodina"},
	87: {87, "Vršac", "Serbia/Vojvodina"},
	88: {88, "Ruma", "Serbia/Vojvodina"},
	89: {89, "Sremska Mitrovica", "Serbia/Vojvodina"},
	// Serbia / Kosovo
	91: {91, "Priština", "Serbia/Kosovo"},
	92: {92, "Kosovska Mitrovica", "Serbia/Kosovo"},
	93: {93, "Peć", "Serbia/Kosovo"},
	94: {94, "Đakovica", "Serbia/Kosovo"},
	95: {95, "Prizren", "Serbia/Kosovo"},
	96: {96, "Gnjilane, Kosovska Kamenica, Vitna, Novo Brdo", "Serbia/Kosovo"},
}
