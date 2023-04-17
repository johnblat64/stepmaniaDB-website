package main

type Chart struct {
	Chartid      string `json:"chartId" `
	Chartname    string `json:"name" `
	Stepstype    string `json:"stepsType" `
	Description  string `json:"description" `
	Chartstyle   string `json:"chartStyle" `
	Difficulty   string `json:"difficulty" `
	Meter        int    `json:"meter" `
	Credit       string `json:"credit" `
	StopsCount   int    `json:"stopsCount" db:"stops_count"`
	DelaysCount  int    `json:"delaysCount" db:"delays_count"`
	WarpsCount   int    `json:"warpsCount" db:"warps_count"`
	ScrollsCount int    `json:"scrollsCount" db:"scrolls_count"`
	FakesCount   int    `json:"fakesCount" db:"fakes_count"`
	SpeedsCount  int    `json:"speedsCount" db:"speeds_count"`
}

type Bpm struct {
	Value float32 `json:"value" db:"song_bpm"`
}

type TimeSignature struct {
	Numerator   int `json:"numerator" db:"time_signature_numerator"`
	Denominator int `json:"denominator" db:"time_signature_denominator"`
}

type Song struct {
	Songid         string          `json:"songId"         `
	Title          string          `json:"title"      `
	Artist         string          `json:"artist"     `
	Bpms           []Bpm           `json:"bpms"           `
	Timesignatures []TimeSignature `json:"timeSignatures" `
	Charts         []Chart         `json:"charts"`
	PackId         string          `json:"packId"         db:"packid"`
	PackName       string          `json:"packName"       db:"name"`
}

type SongPage struct {
	Page      int `json:"pageNum"`
	PageSize  int `json:"pageSize"`
	PageCount int `json:"pageCount"`
	//ResultsCount int    `json:"resultsCount"`
	Songs []Song `json:"songs"`
}

// struct for counting songs
type Count struct {
	Count int `db:"count"`
}
