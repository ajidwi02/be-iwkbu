package config

type DetailLocationConfig struct {
	SheetName string
	Ranges    DetailRangeConfig
}

type DetailRangeConfig struct {
	No                   string
	Loket                string
	Samsat               string
	JenisLoket           string
	KodeLoket            string
	IwkbuTahunLalu       IwkbuTahunLaluRange
	IwkbuTahunIni        IwkbuTahunIniRange
	SelisihIwkbu         SelisihIwkbuRange
	Po                   string
	TLKeteranganKonversi string
	RekapHasil           string
	Keterangan           string
	Pengisian            string
	Outstanding          string
	KodeNopolCI          string
	KodeNopolCO          string
	Memastikan           MemastikanRange
}

type IwkbuTahunLaluRange struct {
	TglTransaksi     string
	NoResi           string
	Nopol            string
	MasaLalu         MasaLaluRange
	BulanBayar       string
	BulanMaju        string
	RupiahPenerimaan string
}

type IwkbuTahunIniRange struct {
	TglTransaksi     string
	NoResi           string
	Nopol            string
	MasaLalu         MasaLaluRange
	BulanBayar       string
	BulanMaju        string
	RupiahPenerimaan string
}

type MasaLaluRange struct {
	Awal  string
	Akhir string
}

type SelisihIwkbuRange struct {
	JumlahNopol      string
	RupiahPenerimaan string
}

type MemastikanRange struct {
	Nopol string
	Rp    string
}

var CabangDetails = map[string]DetailLocationConfig{
	"loketcabangjawatengah": {
		SheetName: "1 Lok CAB",
		Ranges: DetailRangeConfig{
			No:         "B11:B",
			Loket:      "C11:C",
			Samsat:     "D11:D",
			JenisLoket: "E11:E",
			KodeLoket:  "F11:F",
			IwkbuTahunLalu: IwkbuTahunLaluRange{
				TglTransaksi: "G11:G",
				NoResi:       "H11:H",
				Nopol:        "I11:I",
				MasaLalu: MasaLaluRange{
					Awal:  "J11:J",
					Akhir: "K11:K",
				},
				BulanBayar:       "L11:L",
				BulanMaju:        "M11:M",
				RupiahPenerimaan: "N11:N",
			},
			IwkbuTahunIni: IwkbuTahunIniRange{
				TglTransaksi: "O11:O",
				NoResi:       "P11:P",
				Nopol:        "Q11:Q",
				MasaLalu: MasaLaluRange{
					Awal:  "R11:R",
					Akhir: "S11:S",
				},
				BulanBayar:       "T11:T",
				BulanMaju:        "U11:U",
				RupiahPenerimaan: "V11:V",
			},
			SelisihIwkbu: SelisihIwkbuRange{
				JumlahNopol:      "W11:W",
				RupiahPenerimaan: "X11:X",
			},
			Po:                   "Y11:Y",
			TLKeteranganKonversi: "Z11:Z",
			RekapHasil:           "AA11:AA",
			Keterangan:           "AB11:AB",
			Pengisian:            "AC11:AC",
			Outstanding:          "AD11:AD",
			KodeNopolCI:          "AE11:AE",
			KodeNopolCO:          "AF11:AF",
			Memastikan: MemastikanRange{
				Nopol: "AG11:AG",
				Rp:    "AH11:AH",
			},
		},
	},
}