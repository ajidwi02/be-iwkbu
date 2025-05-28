package services

import (
	"encoding/csv"
	"fmt"
	"jm-CICO/config"
	"jm-CICO/models"
	"jm-CICO/utils"
	"log"
	"net/http"
	"strings"
)

func FetchData(config config.LocationConfig) ([]models.Entry, error) {
	sheetID := "1mBRaSCNk1TTaeJ6AJDhO1kEb7tAJSFZ1ZPhW13nVjq4"
	sheetName := "REPORT"

	// Fetch data utama
	mainURL := fmt.Sprintf("https://docs.google.com/spreadsheets/d/%s/gviz/tq?tqx=out:csv&sheet=%s&range=%s",
		sheetID, sheetName, config.DataRange,)

	mainResp, err:= http.Get(mainURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch main data: %w", err)
	}
	defer mainResp.Body.Close()

	mainReader := csv.NewReader(mainResp.Body)
	mainReader.FieldsPerRecord = -1
	mainData, err := mainReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed read main data: %w", err)
	}

	// Fetch kolom anggaran sesuai lokasi
	anggaranURL := fmt.Sprintf(
		"https://docs.google.com/spreadsheets/d/%s/gviz/tq?tqx=out:csv&sheet=%s&range=%s",
		sheetID, sheetName, config.AnggaranRange,
	)
	anggaranResp, err := http.Get(anggaranURL)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data anggaran: %v", err)
	}
	defer anggaranResp.Body.Close()
	anggaranReader := csv.NewReader(anggaranResp.Body)
	anggaranReader.FieldsPerRecord = -1
	anggaranData, err := anggaranReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("gagal membaca data anggaran: %v", err)
	}

	// fetch kolom NOPOL sesuai lokasi
	nopoURL := fmt.Sprintf("https://docs.google.com/spreadsheets/d/%s/gviz/tq?tqx=out:csv&sheet=%s&range=%s", sheetID, sheetName, config.NopolRange,)

	nopolResp, err := http.Get(nopoURL)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data nopol: %v", err)
	}
	defer nopolResp.Body.Close()

	nopolReader := csv.NewReader(nopolResp.Body)
	nopolReader.FieldsPerRecord = -1
	nopolData, err := nopolReader.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("gagal membaca data nopol: %v", err)
	}

	// Fetch kolom PETUGAS sesuai lokasi
	petugasURL := fmt.Sprintf(
		"https://docs.google.com/spreadsheets/d/%s/gviz/tq?tqx=out:csv&sheet=%s&range=%s", sheetID, sheetName, config.PetugasRange,
	)

	petugasResp, err := http.Get(petugasURL)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil kolom petugas: %v", err)
	}
	defer petugasResp.Body.Close()

	petugasReader := csv.NewReader(petugasResp.Body)
	petugasReader.FieldsPerRecord = -1
	petugasData, err := petugasReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("gagal membaca data petugas: %v", err)
	}

	// Fetch data Jml Bulan (Checkin jml_bln) dari kolom AC:AD
	jmlBlnURL := fmt.Sprintf(
		"https://docs.google.com/spreadsheets/d/%s/gviz/tq?tqx=out:csv&sheet=%s&range=%s",
		sheetID, sheetName, config.CheckinJmlBulanRange,
	)

	jmlBlnResp, err := http.Get(jmlBlnURL)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data jmlBln: %v", err)
	}
	defer jmlBlnResp.Body.Close()

	jmlBlnReader := csv.NewReader(jmlBlnResp.Body)
	jmlBlnReader.FieldsPerRecord = -1
	jmlBlnData, err := jmlBlnReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("gagal membaca data jmlBln: %v", err)
	}

	// Fetch data check_in rupiah 
	rupiahURL := fmt.Sprintf("https://docs.google.com/spreadsheets/d/%s/gviz/tq?tqx=out:csv&sheet=%s&range=%s", sheetID, sheetName, config.CheckinRupiahRange)

	rupiahResp, err:= http.Get(rupiahURL)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data checkin rupiah: %v", err)
	}
	defer rupiahResp.Body.Close()

	rupiahReader := csv.NewReader(rupiahResp.Body)
	rupiahReader.FieldsPerRecord = -1
	rupiahData, err:= rupiahReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("gagal membaca data checkin rupiah: %v", err)
	}

	// Fetch data checkout nopol
	checkoutNopolURL := fmt.Sprintf("https://docs.google.com/spreadsheets/d/%s/gviz/tq?tqx=out:csv&sheet=%s&range=%s", sheetID, sheetName, config.CheckoutNopolRange)
	checkoutNopolResp, err := http.Get(checkoutNopolURL)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data checkout nopol: %v", err)
	}
	defer checkoutNopolResp.Body.Close()
	checkoutNopolReader := csv.NewReader(checkoutNopolResp.Body)
	checkoutNopolReader.FieldsPerRecord = -1
	checkoutNopolData, err := checkoutNopolReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("gagal membaca data checkout nopol: %v", err)
	}	

	// Fetch data checkout jml_bln
	checkoutJmlBlnURL := fmt.Sprintf("https://docs.google.com/spreadsheets/d/%s/gviz/tq?tqx=out:csv&sheet=%s&range=%s", sheetID, sheetName, config.CheckoutJmlBulanRange)
	checkoutJmlBlnResp, err := http.Get(checkoutJmlBlnURL)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data checkout jml_bln: %v", err)
	}
	defer checkoutJmlBlnResp.Body.Close()
	checkoutJmlBlnReader := csv.NewReader(checkoutJmlBlnResp.Body)
	checkoutJmlBlnReader.FieldsPerRecord = -1
	checkoutJmlBlnData, err := checkoutJmlBlnReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("gagal membaca data checkout jml_bln: %v", err)
	}


	// Fetch data checkout rupiah
	checkoutRupiahURL := fmt.Sprintf("https://docs.google.com/spreadsheets/d/%s/gviz/tq?tqx=out:csv&sheet=%s&range=%s", sheetID, sheetName, config.CheckoutRupiahRange)
	checkoutRupiahResp, err := http.Get(checkoutRupiahURL)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data checkout rupiah: %v", err)
	}
	defer checkoutRupiahResp.Body.Close()
	checkoutRupiahReader := csv.NewReader(checkoutRupiahResp.Body)
	checkoutRupiahReader.FieldsPerRecord = -1
	checkoutRupiahData, err := checkoutRupiahReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("gagal membaca data checkout rupiah: %v", err)
	}

	var entries []models.Entry
	for i, row:= range mainData{
		if utils.IsEmptyRow(row) || utils.IsHeaderRow(row) || i >= len(anggaranData) || i >= len(petugasData) || i >= len(nopolData) || i >= len(jmlBlnData) || i > len(rupiahData) || i > len(checkoutNopolData) || i > len(checkoutJmlBlnData) || i > len(checkoutRupiahData) {
			continue
		}
		
		entry, err := utils.ParseEntry(row)
		if err!= nil {
			log.Printf("failed fetch the row: %v | Error: %v", row, err)
			continue
		}

		// Gabungkan nama petugas dari O→V (kolom 0 sampai 7)
		fullName := ""
		for _, col := range petugasData[i] {
			part := strings.TrimSpace(col)
			if part != "" {
				fullName += part + " "
			}
		}
		entry.Petugas = strings.TrimSpace(fullName)

		// Ambil anggaran dari kolom W dan X
		anggaranStr := strings.Join(anggaranData[i], "")
		entry.Anggaran = utils.ParseNumber(anggaranStr)

		// Ambil nopol checkin dari kolom AA dan AB
		nopolStr := strings.Join(nopolData[i], "")
		entry.Checkin.Nopol = utils.ParseInt(nopolStr)

		// Masukkan Rubian dari kolom AC:AD
		jmlBlnStr := strings.Join(jmlBlnData[i], "")
		entry.Checkin.JumlahBulan = utils.ParseInt(jmlBlnStr)

		// Masukkan ciRupiah dari kolom
		entry.Checkin.Rupiah = utils.ParseNumber(strings.Join(rupiahData[i], ""))

		// Ambil nopol checkout dari kolom AI dan AJ
		checkoutNopolStr := strings.Join(checkoutNopolData[i], "")
		entry.Checkout.Nopol = utils.ParseInt(checkoutNopolStr)

		// Ambil jml_bln checkout dari kolom AK dan AL
		checkoutJmlBlnStr := strings.Join(checkoutJmlBlnData[i], "")
		entry.Checkout.JumlahBulan = utils.ParseInt(checkoutJmlBlnStr)

		// Ambil checkout rupiah dari kolom AM:AP
		checkoutRupiahStr := strings.Join(checkoutRupiahData[i], "")
		entry.Checkout.Rupiah = utils.ParseNumber(checkoutRupiahStr)

		// hitungan selisih CICO
		entry.SelisihCICO.Nopol = entry.Checkout.Nopol - entry.Checkin.Nopol

		entry.SelisihCICO.JumlahBulan = entry.Checkout.JumlahBulan - entry.Checkin.JumlahBulan

		entry.SelisihCICO.Rupiah = entry.Checkout.Rupiah - entry.Checkin.Rupiah
		
		entries = append(entries, entry)
	}
	return entries, nil
}

func CalculateSubtotal(entries []models.Entry) models.Subtotal {
	var st models.Subtotal

	for _, entry := range entries{
		st.Anggaran += entry.Anggaran
		st.CheckinNopol += entry.Checkin.Nopol
		st.CheckinJmlBln += entry.Checkin.JumlahBulan
		st.CheckinRupiah += entry.Checkin.Rupiah
		st.CheckoutNopol += entry.Checkout.Nopol
		st.CheckoutJmlBln += entry.Checkout.JumlahBulan
		st.CheckoutRupiah += entry.Checkout.Rupiah
		st.SelisihCICONopol += entry.SelisihCICO.Nopol
		st.SelisihCICOJmlBln += entry.SelisihCICO.JumlahBulan
		st.SelisihCICORupiah += entry.SelisihCICO.Rupiah
	}
	return st
}