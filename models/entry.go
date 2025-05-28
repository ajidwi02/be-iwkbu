package models

type Entry struct {
	No          int    `json:"no"`
	LoketKantor string `json:"loket_kantor"`
	Petugas     string `json:"petugas"`
	Anggaran    int    `json:"anggaran"`
	Checkin     Check  `json:"checkin"`
	Checkout    Check  `json:"checkout"`
	SelisihCICO Check  `json:"selisih_cico"`
}

type Check struct {
	Nopol       int `json:"nopol"`
	JumlahBulan int `json:"jml_bln"`
	Rupiah      int `json:"rupiah"`
}

type Subtotal struct {
	Anggaran          int `json:"anggaran"`
	CheckinNopol      int `json:"checkin_nopol"`
	CheckinJmlBln     int `json:"checkin_jml_bln"`
	CheckinRupiah     int `json:"checkin_rupiah"`
	CheckoutNopol     int `json:"checkout_nopol"`
	CheckoutJmlBln    int `json:"checkout_jml_bln"`
	CheckoutRupiah    int `json:"checkout_rupiah"`
	SelisihCICONopol  int `json:"selisih_cico_nopol"`
	SelisihCICOJmlBln int `json:"selisih_cico_jml_bln"`
	SelisihCICORupiah int `json:"selisih_cico_rupiah"`
}
