package utils

import (
	"jm-CICO/models"
	"strconv"
	"strings"
)

// ParseNumber mengkonversi string numerik dengan format ke integer
func ParseNumber(s string) int {
	// Membersihkan karakter non-numerik
	cleaned := strings.NewReplacer(
		".", "",
		",", "",
		"Rp", "",
		" ", "",
	).Replace(s)

	// Konversi ke integer
	num, _ := strconv.Atoi(cleaned)
	return num
}

// ParseInt mengkonversi string ke integer sederhana
func ParseInt(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	num, _ := strconv.Atoi(s)
	return num
}

// IsEmptyRow mengecek apakah sebuah row kosong
func IsEmptyRow(row []string) bool {
	for _, cell := range row {
		if strings.TrimSpace(cell) != "" {
			return false
		}
	}
	return true
}

// IsHeaderRow mengecek apakah sebuah row adalah header
func IsHeaderRow(row []string) bool {
	if len(row) == 0 {
		return false
	}
	
	// Normalisasi dan cek kolom pertama
	firstCell := strings.ToLower(strings.TrimSpace(row[0]))
	return strings.Contains(firstCell, "no") || 
		strings.Contains(firstCell, "nomer") || 
		strings.Contains(firstCell, "nomor")
}

// CleanString membersihkan string dari karakter tidak perlu
func CleanString(s string) string {
	return strings.Trim(strings.TrimSpace(s), `"`)
}

// ParseFloat mengkonversi string ke float64 (opsional)
func ParseFloat(s string) float64 {
	cleaned := strings.NewReplacer(
		",", ".",
		"Rp", "",
		" ", "",
	).Replace(s)

	num, _ := strconv.ParseFloat(cleaned, 64)
	return num
}


func ParseEntry(row []string) (models.Entry, error) {
	for len(row) < 21 {
		row = append(row, "")
	}

	clean := func(s string) string {
		return strings.Trim(strings.TrimSpace(s), `"`)
	}

	no, _ := strconv.Atoi(clean(row[0]))
	
	return models.Entry{
		No:          no,
		LoketKantor: clean(row[2]),   
		Petugas:     clean(row[14]),    
		Anggaran:    ParseNumber(clean(row[20])), 
		Checkin: models.Check{
			Nopol:  ParseInt(clean(row[4])), 
			JumlahBulan: ParseInt(clean(row[5])),
			Rupiah: ParseNumber(clean(row[6])),
		},
		Checkout: models.Check{
			Nopol:  ParseInt(clean(row[7])),
			JumlahBulan: ParseInt(clean(row[8])),
			Rupiah: ParseNumber(clean(row[9])),
		},
		SelisihCICO: models.Check{
			Nopol:  ParseInt(clean(row[10])),
			JumlahBulan: ParseInt(clean(row[11])),
			Rupiah: ParseNumber(clean(row[12])),
		},
	}, nil
}