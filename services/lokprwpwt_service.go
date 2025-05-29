package services

import (
	"encoding/csv"
	"fmt"
	"jm-CICO/models"
	"net/http"
	"strconv"
	"strings"
)

func FetchLokPrwPWTData(sheetURL string) ([]models.LoketData, error) {
	resp, err := http.Get(sheetURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch CSV: %w", err)
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	var data []models.LoketData
	var semuaTiNopol []string
	for i := 10; i < len(rows); i++ {
		if len(rows[i]) > 9 {
			semuaTiNopol = append(semuaTiNopol, strings.TrimSpace(rows[i][8]))
		}
	}

	for i := 10; i < len(rows); i++ {
		row := rows[i]
		if len(row) < 6 { // minimal sampai kolom F (index 5)
			continue
		}

		no, _ := strconv.Atoi(strings.TrimSpace(row[1])) // kolom B
		bulanBayarTl, _ := strconv.Atoi(row[11])
		bulanMajuTl, _ := strconv.Atoi(row[12])
		rawRupiahTl := strings.ReplaceAll(row[13], ".", "")
		rupiahTl, _ := strconv.Atoi(rawRupiahTl)

		bulanBayarTi, _ := strconv.Atoi(row[19])
		bulanMajuTi, _ := strconv.Atoi(row[20])
		rawRupiahTi := strings.ReplaceAll(row[21], ".", "")
		rupiahTi, _ := strconv.Atoi(rawRupiahTi)

		kodeNopolCi := 0
		if row[16] != "" {
			kodeNopolCi = 1
		}
		kodeNopolCo := 0
		if row[8] != "" {
			kodeNopolCo = 1
		}

		selisihJumlahNopol := kodeNopolCo - kodeNopolCi
		selisihRupiah := rupiahTi - rupiahTl
		outStanding := kodeNopolCi - kodeNopolCo

		iwkbuTiNopol := strings.TrimSpace(row[16])
		memastikanNopol := 0
		for _, nopol := range semuaTiNopol {
			if iwkbuTiNopol != "" && iwkbuTiNopol == nopol {
				memastikanNopol = 1
				break
			}
		}
		memastikanRp := 0
		if memastikanNopol == 1 {
			memastikanRp = rupiahTi
		}

		data = append(data, models.LoketData{
			No:         no,
			Loket:      strings.TrimSpace(row[2]), // kolom C
			Samsat:     strings.TrimSpace(row[3]), // kolom D
			JenisLoket: strings.TrimSpace(row[4]), // kolom E
			KodeLoket:  strings.TrimSpace(row[5]), // kolom F
			IwkbuTlTglTransaksi: strings.TrimSpace(row[6]),
			IwkbuTlNoResi: strings.TrimSpace(row[7]),
			IwkbuTlNopol: strings.TrimSpace(row[8]),
			IwkbuTlMasaLaluAwal: strings.TrimSpace(row[9]),
			IwkbuTlMasaLaluAkhir: strings.TrimSpace(row[10]),
			IwkbuTlBulanBayar:bulanBayarTl,
			IwkbuTlBulanMaju:bulanMajuTl,
			IwkbuTlRupiahPenerimaan:rupiahTl,
			IwkbuTiTglTransaksi:         strings.TrimSpace(row[14]),
			IwkbuTiNoResi:               strings.TrimSpace(row[15]),
			IwkbuTiNopol:                strings.TrimSpace(row[16]),
			IwkbuTiMasaLaluAwal:         strings.TrimSpace(row[17]),
			IwkbuTiMasaLaluAkhir:        strings.TrimSpace(row[18]),
			IwkbuTiBulanBayar:           bulanBayarTi,
			IwkbuTiBulanMaju:            bulanMajuTi,
			IwkbuTiRupiahPenerimaan:     rupiahTi,
			SelisihIwkbuJumlahNopol:     selisihJumlahNopol,
			SelisihIwkbuRupiahPenerimaan: selisihRupiah,
			Po:                          strings.TrimSpace(row[24]),
			TLKeteranganKonversiIWKBU:   strings.TrimSpace(row[25]),
			RekapHasil:                  strings.TrimSpace(row[26]),
			Keterangan:                  strings.TrimSpace(row[27]),
			Pengisian:                   strings.TrimSpace(row[28]),
			OutStanding:                 outStanding,
			KodeNopolCi:                 kodeNopolCi,
			KodeNopolCo:                 kodeNopolCo,
			MemastikanNopol:             memastikanNopol,
			MemastikanRupiah:            memastikanRp,
		})
	}
	return data, nil
}