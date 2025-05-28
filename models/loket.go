package models

type LoketData struct {
	No                           int    `json:"no"`
	Loket                        string `json:"loket"`
	Samsat                       string `json:"samsat"`
	JenisLoket                   string `json:"jenis_loket"`
	KodeLoket                    string `json:"kode_loket"`
	IwkbuTlTglTransaksi          string `json:"iwkbu_tl_tgl_transaksi"`
	IwkbuTlNoResi                string `json:"iwkbu_tl_no_resi"`
	IwkbuTlNopol                 string `json:"iwkbu_tl_nopol"`
	IwkbuTlMasaLaluAwal          string `json:"iwkbu_tl_masa_lalu_awal"`
	IwkbuTlMasaLaluAkhir         string `json:"iwkbu_tl_masa_lalu_akhir"`
	IwkbuTlBulanBayar            int    `json:"iwkbu_tl_bulan_bayar"`
	IwkbuTlBulanMaju             int    `json:"iwkbu_tl_bulan_maju"`
	IwkbuTlRupiahPenerimaan      int    `json:"iwkbu_tl_rupiah_penerimaan"`
	IwkbuTiTglTransaksi          string `json:"iwkbu_ti_tgl_transaksi"`
	IwkbuTiNoResi                string `json:"iwkbu_ti_no_resi"`
	IwkbuTiNopol                 string `json:"iwkbu_ti_nopol"`
	IwkbuTiMasaLaluAwal          string `json:"iwkbu_ti_masa_lalu_awal"`
	IwkbuTiMasaLaluAkhir         string `json:"iwkbu_ti_masa_lalu_akhir"`
	IwkbuTiBulanBayar            int    `json:"iwkbu_ti_bulan_bayar"`
	IwkbuTiBulanMaju             int    `json:"iwkbu_ti_bulan_maju"`
	IwkbuTiRupiahPenerimaan      int    `json:"iwkbu_ti_rupiah_penerimaan"`
	SelisihIwkbuJumlahNopol      int    `json:"selisih_iwkbu_jumlah_nopol"`
	SelisihIwkbuRupiahPenerimaan int    `json:"selisih_iwkbu_rupiah_penerimaan"`
	Po                           string `json:"po"`
	TLKeteranganKonversiIWKBU    string `json:"tl_keterangan_konversi_iwkbu"`
	RekapHasil                   string `json:"rekap_hasil"`
	Keterangan                   string `json:"keterangan"`
	Pengisian                    string `json:"pengisian"`
	OutStanding                  int    `json:"outstanding"`
	KodeNopolCi                  int    `json:"kode_nopol_ci"`
	KodeNopolCo                  int    `json:"kode_nopol_co"`
	MemastikanNopol              int    `json:"memastikan_nopol"`
	MemastikanRupiah             int    `json:"memastikan_rp"`
}