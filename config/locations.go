package config

type LocationConfig struct {
	DataRange             string
	PetugasRange          string
	AnggaranRange         string
	NopolRange            string
	CheckinJmlBulanRange  string
	CheckinRupiahRange    string
	CheckoutNopolRange    string
	CheckoutJmlBulanRange string
	CheckoutRupiahRange   string
}

var Locations = map[string]LocationConfig{
	"loketkedungsapur": {
		DataRange:             "C17:AX22",
		PetugasRange:          "O17:V22",
		AnggaranRange:         "W17:X22",
		NopolRange:            "AA17:AB22",
		CheckinJmlBulanRange:  "AC17:AD22",
		CheckinRupiahRange:    "AE17:AH22",
		CheckoutNopolRange:    "AI17:AJ22",
		CheckoutJmlBulanRange: "AK17:AL22",
		CheckoutRupiahRange:   "AM17:AP22",
	},
	"perwakilansurakarta": {
		DataRange:             "C25:AX31",
		PetugasRange:          "O25:V31",
		AnggaranRange:         "W25:X31",
		NopolRange:            "AA25:AB31",
		CheckinJmlBulanRange:  "AC25:AD31",
		CheckinRupiahRange:    "AE25:AH31",
		CheckoutNopolRange:    "AI25:AJ31",
		CheckoutJmlBulanRange: "AK25:AL31",
		CheckoutRupiahRange:   "AM25:AP31",
	},
	"perwakilanmagelang": {
		DataRange:             "C34:AX41",
		PetugasRange:          "O34:V41",
		AnggaranRange:         "W34:X41",
		NopolRange:            "AA34:AB41",
		CheckinJmlBulanRange:  "AC34:AD41",
		CheckinRupiahRange:    "AE34:AH41",
		CheckoutNopolRange:    "AI34:AJ41",
		CheckoutJmlBulanRange: "AK34:AL41",
		CheckoutRupiahRange:   "AM34:AP41",
	},
	"perwakilanpurwokerto": {
		DataRange:             "C44:AX50",
		PetugasRange:          "O44:V50",
		AnggaranRange:         "W44:X50",
		NopolRange:            "AA44:AB50",
		CheckinJmlBulanRange:  "AC44:AD50",
		CheckinRupiahRange:    "AE44:AH50",
		CheckoutNopolRange:    "AI44:AJ50",
		CheckoutJmlBulanRange: "AK44:AL50",
		CheckoutRupiahRange:   "AM44:AP50",
	},
	"perwakilanpekalongan": {
		DataRange:             "C53:AX62",
		PetugasRange:          "O53:V62",
		AnggaranRange:         "W53:X62",
		NopolRange:            "AA53:AB62",
		CheckinJmlBulanRange:  "AC53:AD62",
		CheckinRupiahRange:    "AE53:AH62",
		CheckoutNopolRange:    "AI53:AJ62",
		CheckoutJmlBulanRange: "AK53:AL62",
		CheckoutRupiahRange:   "AM53:AP62",
	},
	"perwakilanpati": {
		DataRange:             "C65:AX71",
		PetugasRange:          "O65:V71",
		AnggaranRange:         "W65:X71",
		NopolRange:            "AA65:AB71",
		CheckinJmlBulanRange:  "AC65:AD71",
		CheckinRupiahRange:    "AE65:AH71",
		CheckoutNopolRange:    "AI65:AJ71",
		CheckoutJmlBulanRange: "AK65:AL71",
		CheckoutRupiahRange:   "AM65:AP71",
	},
	"perwakilansemarang": {
		DataRange:             "C74:AX77",
		PetugasRange:          "O74:V77",
		AnggaranRange:         "W74:X77",
		NopolRange:            "AA74:AB77",
		CheckinJmlBulanRange:  "AC74:AD77",
		CheckinRupiahRange:    "AE74:AH77",
		CheckoutNopolRange:    "AI74:AJ77",
		CheckoutJmlBulanRange: "AK74:AL77",
		CheckoutRupiahRange:   "AM74:AP77",
	},
	"perwakilansukoharjo": {
		DataRange:             "C80:AX85",
		PetugasRange:          "O80:V85",
		AnggaranRange:         "W80:X85",
		NopolRange:            "AA80:AB85",
		CheckinJmlBulanRange:  "AC80:AD85",
		CheckinRupiahRange:    "AE80:AH85",
		CheckoutNopolRange:    "AI80:AJ85",
		CheckoutJmlBulanRange: "AK80:AL85",
		CheckoutRupiahRange:   "AM80:AP85",
	},
}