package args

type Characters struct {
	Number    bool
	Lowercase bool
	Uppercase bool
	Special   bool
}

type Args struct {
	Length     int
	Characters *Characters
}
