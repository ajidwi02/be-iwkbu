// package main

// import (
// 	"encoding/csv"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"strings"
// )

// type LoketData struct {
// 	No                int    `json:"no"`
// 	Loket             string `json:"loket"`
// 	Samsat            string `json:"samsat"`
// 	JenisLoket        string `json:"jenis_loket"`
// 	KodeLoket         string `json:"kode_loket"`
// 	IwkbuTlTglTransaksi    string `json:"iwkbu_tl_tgl_transaksi"`
// 	IwkbuTlNoResi          string `json:"iwkbu_tl_no_resi"`
// 	IwkbuTlNopol           string `json:"iwkbu_tl_nopol"`
// 	IwkbuTlMasaLaluAwal    string `json:"iwkbu_tl_masa_lalu_awal"`
// 	IwkbuTlMasaLaluAkhir   string `json:"iwkbu_tl_masa_lalu_akhir"`
// 	IwkbuTlBulanBayar      int    `json:"iwkbu_tl_bulan_bayar"`
// 	IwkbuTlBulanMaju       int    `json:"iwkbu_tl_bulan_maju"`
// 	IwkbuTlRupiahPenerimaan int   `json:"iwkbu_tl_rupiah_penerimaan"`

// 	IwkbuTiTglTransaksi    string `json:"iwkbu_ti_tgl_transaksi"`
// 	IwkbuTiNoResi          string `json:"iwkbu_ti_no_resi"`
// 	IwkbuTiNopol           string `json:"iwkbu_ti_nopol"`
// 	IwkbuTiMasaLaluAwal    string `json:"iwkbu_ti_masa_lalu_awal"`
// 	IwkbuTiMasaLaluAkhir   string `json:"iwkbu_ti_masa_lalu_akhir"`
// 	IwkbuTiBulanBayar      int    `json:"iwkbu_ti_bulan_bayar"`
// 	IwkbuTiBulanMaju       int    `json:"iwkbu_ti_bulan_maju"`
// 	IwkbuTiRupiahPenerimaan int   `json:"iwkbu_ti_rupiah_penerimaan"`

// 	SelisihIwkbuJumlahNopol      int `json:"selisih_iwkbu_jumlah_nopol"`
// 	SelisihIwkbuRupiahPenerimaan int `json:"selisih_iwkbu_rupiah_penerimaan"`

// 	Po                          string `json:"po"`
//   TLKeteranganKonversiIWKBU  string `json:"tl_keterangan_konversi_iwkbu"`
//   RekapHasil                  string `json:"rekap_hasil"`
//   Keterangan                  string `json:"keterangan"`
//   Pengisian                   string `json:"pengisian"`

// 	OutStanding int `json:"outstanding"`

// 	KodeNopolCi int `json:"kode_nopol_ci"`
// 	KodeNopolCo int `json:"kode_nopol_co"`

// 	MemastikanNopol int `json:"memastikan_nopol"`
// 	MemastikanRupiah int `json:"memastikan_rp"`
// }

// func enableCORS(w http.ResponseWriter){
//   w.Header().Set("Access-Control-Allow-Origin", "*")
//   w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
//   w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// }

// func fetchLoketData() ([]LoketData, error) {
// 	url := "https://docs.google.com/spreadsheets/d/14qJDOG1KR8_BxOMgDbgxosnXtl-45enkNT1e3yXgmTU/gviz/tq?tqx=out:csv&sheet=1%20LOK%20CAB"
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to fetch CSV: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	reader := csv.NewReader(resp.Body)
// 	rows, err := reader.ReadAll()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read CSV: %w", err)
// 	}

// 	var data []LoketData

// 	// semua nilai iwkbutiNopol
// 	var semuaTiNopol []string
// 	for i := 10; i < len(rows); i++ {
// 		if len(rows[i]) > 9 {
// 			semuaTiNopol = append(semuaTiNopol, strings.TrimSpace(rows[i][8]))
// 		}
// 	}

// 	for i := 10; i < len(rows); i++ { // Mulai dari baris ke-11 (index 10)
// 		row := rows[i]
// 		if len(row) < 14 {
// 			continue
// 		}

// 		no, err := strconv.Atoi(row[1])
// 		if err != nil {
// 			continue
// 		}

// 		bulanBayarTl, _ := strconv.Atoi(row[11])
// 		bulanMajuTl, _ := strconv.Atoi(row[12])
// 		rawRupiahTl := strings.ReplaceAll(row[13], ".", "")
// 		rupiahTl, _ := strconv.Atoi(rawRupiahTl)

// 		bulanBayarTi, _ := strconv.Atoi(row[19])
// 		bulanMajuTi, _ := strconv.Atoi(row[20])
// 		rawRupiahTi := strings.ReplaceAll(row[21], ".", "")
// 		rupiahTi, _ := strconv.Atoi(rawRupiahTi)

// 		kodeNopolCi := 0
// 		if row[16] != ""{
// 			kodeNopolCi = 1
// 		}

// 		kodeNopolCo := 0
// 		if row[8] != ""{
// 			kodeNopolCi = 1
// 		}

// 		selisihJumlahNopol := kodeNopolCo - kodeNopolCi
// 		selisihRupiah := rupiahTi - rupiahTl

// 		outStanding := kodeNopolCi - kodeNopolCo

// 		// cek nopol available
// 		iwkbuTiNopol := strings.TrimSpace(row[16])
// 		memastikanNopol := 0
// 		for _, nopol := range semuaTiNopol {
// 			if iwkbuTiNopol != "" && iwkbuTiNopol == nopol {
// 				memastikanNopol = 1
// 				break
// 			}
// 		}

// 		// hitung memastikan_rp
// 		memastikanRp := 0
// 		if memastikanNopol == 1 {
// 			memastikanRp = rupiahTi
// 		}

// 		entry := LoketData{
// 			No:                no,
// 			Loket:             row[2],
// 			Samsat:            row[3],
// 			JenisLoket:        row[4],
// 			KodeLoket:         row[5],
// 			IwkbuTlTglTransaksi: row[6],
// 			IwkbuTlNoResi:       row[7],
// 			IwkbuTlNopol:        row[8],
// 			IwkbuTlMasaLaluAwal:         row[9],
// 			IwkbuTlMasaLaluAkhir:        row[10],
// 			IwkbuTlBulanBayar:   bulanBayarTl,
// 			IwkbuTlBulanMaju:    bulanMajuTl,
// 			IwkbuTlRupiahPenerimaan:       rupiahTl,
// 			IwkbuTiTglTransaksi:    row[14], // Kolom O
// 			IwkbuTiNoResi:          row[15], // Kolom P
// 			IwkbuTiNopol:           row[16], // Kolom Q
// 			IwkbuTiMasaLaluAwal:    row[17], // Kolom R
// 			IwkbuTiMasaLaluAkhir:   row[18], // Kolom S
// 			IwkbuTiBulanBayar:      bulanBayarTi,
// 			IwkbuTiBulanMaju:       bulanMajuTi,
// 			IwkbuTiRupiahPenerimaan: rupiahTi,
// 			KodeNopolCi: kodeNopolCi,
// 			KodeNopolCo: kodeNopolCo,
// 			SelisihIwkbuJumlahNopol : selisihJumlahNopol,
// 			SelisihIwkbuRupiahPenerimaan : selisihRupiah,

// 			Po:                         row[24],
// 			TLKeteranganKonversiIWKBU:  row[25],
// 			RekapHasil:                 row[26],
// 			Keterangan:                 row[27],
// 			Pengisian:                  row[28],
// 			OutStanding : outStanding,		

// 			MemastikanNopol: memastikanNopol,
// 			MemastikanRupiah: memastikanRp,
// 		}
// 		data = append(data, entry)
// 	}
// 	return data, nil
// }

// func loketHandler(w http.ResponseWriter, r *http.Request) {
// 	enableCORS(w)
// 	if r.Method == "OPTIONS"{
// 		return
// 	}
// 	data, err := fetchLoketData()
// 	if err != nil {
// 		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
// 		return
// 	}
// 	result := map[string]interface{}{
// 		"data": data,
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(result)
// }

// func main() {
// 	http.HandleFunc("/loketcabangjawatengah", loketHandler)
// 	fmt.Println("Server started on http://localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }