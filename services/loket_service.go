package services

import (
	"encoding/csv"
	"fmt"
	"jm-CICO/models"
	"net/http"
	"strconv"
	"strings"
)

func FetchLoketData(sheetURL string) ([]models.LoketData, error) {
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
		if len(row) < 14 {
			continue
		}
		no, _ := strconv.Atoi(row[1])
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
			No:                          no,
			Loket:                       row[2],
			Samsat:                      row[3],
			JenisLoket:                  row[4],
			KodeLoket:                   row[5],
			IwkbuTlTglTransaksi:         row[6],
			IwkbuTlNoResi:               row[7],
			IwkbuTlNopol:                row[8],
			IwkbuTlMasaLaluAwal:         row[9],
			IwkbuTlMasaLaluAkhir:        row[10],
			IwkbuTlBulanBayar:           bulanBayarTl,
			IwkbuTlBulanMaju:            bulanMajuTl,
			IwkbuTlRupiahPenerimaan:     rupiahTl,
			IwkbuTiTglTransaksi:         row[14],
			IwkbuTiNoResi:               row[15],
			IwkbuTiNopol:                row[16],
			IwkbuTiMasaLaluAwal:         row[17],
			IwkbuTiMasaLaluAkhir:        row[18],
			IwkbuTiBulanBayar:           bulanBayarTi,
			IwkbuTiBulanMaju:            bulanMajuTi,
			IwkbuTiRupiahPenerimaan:     rupiahTi,
			SelisihIwkbuJumlahNopol:     selisihJumlahNopol,
			SelisihIwkbuRupiahPenerimaan: selisihRupiah,
			Po:                          row[24],
			TLKeteranganKonversiIWKBU:   row[25],
			RekapHasil:                  row[26],
			Keterangan:                  row[27],
			Pengisian:                   row[28],
			OutStanding:                 outStanding,
			KodeNopolCi:                 kodeNopolCi,
			KodeNopolCo:                 kodeNopolCo,
			MemastikanNopol:             memastikanNopol,
			MemastikanRupiah:            memastikanRp,
		})
	}
	return data, nil
}