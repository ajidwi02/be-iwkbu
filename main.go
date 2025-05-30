package main

import (
	"fmt"
	"jm-CICO/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/loketcabangjawatengah", handlers.LoketCabangHandler) //Oke
	http.HandleFunc("/samsatkendal", handlers.SamsatKendalHandler) //Oke
	http.HandleFunc("/samsatdemak", handlers.SamsatDemakHandler) //Oke
	http.HandleFunc("/samsatpurwodadi", handlers.SamsatPurwodadiHandler) //Oke
	http.HandleFunc("/samsatungaran", handlers.SamsatUngaranHandler) //Oke
	http.HandleFunc("/samsatsalatiga", handlers.SamsatSalatigaHandler) //Oke
	http.HandleFunc("/samsatlokperwsra", handlers.SamsatLokPerwSRAHandler) //Oke
	http.HandleFunc("/samsatsurakarta", handlers.SamsatSurakartaHandler) //Oke
	http.HandleFunc("/samsatklaten", handlers.SamsatKlatenHandler) //Oke
	http.HandleFunc("/samsatsragen", handlers.SamsatSragenHandler) //Oke
	http.HandleFunc("/samsatboyolali", handlers.SamsatBoyolaliHandler) //Oke
	http.HandleFunc("/samsatprambanan", handlers.SamsatPrambananHandler) //Oke
	http.HandleFunc("/samsatdelanggu", handlers.SamsatDelangguHandler) //Oke
	http.HandleFunc("/samsatlokpwkmgl", handlers.SamsatLokPwkMGLHandler) //Oke
	http.HandleFunc("/samsatmagelang", handlers.SamsatMagelangHandler) //Oke
	http.HandleFunc("/samsatpurworejo", handlers.SamsatPurworejoHandler) //Oke
	http.HandleFunc("/samsatkebumen", handlers.SamsatKebumenHandler) //oke
	http.HandleFunc("/samsattemanggung", handlers.SamsatTemanggungHandler) //Oke
	http.HandleFunc("/samsatwonosobo", handlers.SamsatWonosoboHandler) //oke
	http.HandleFunc("/samsatmungkid", handlers.SamsatMungkidHandler) //Oke
	http.HandleFunc("/samsatbagelen", handlers.SamsatBagelenHandler) //Oke
	http.HandleFunc("/samsatlokprwpwt", handlers.SamsatLokPrwPWTHandler) //Oke
	http.HandleFunc("/samsat/purwokerto", handlers.SamsatPurwokertoHandler) //Oke
	http.HandleFunc("/samsat/purbalingga", handlers.SamsatPurbalinggaHandler) //oke
	http.HandleFunc("/samsat/banjarnegara", handlers.SamsatBanjarnegaraHandler) //Oke
	http.HandleFunc("/samsat/majenang", handlers.SamsatMajenangHandler) //Oke
	http.HandleFunc("/samsat/cilacap", handlers.SamsatCilacapHandler) //Oke
	http.HandleFunc("/samsat/wangon", handlers.SamsatWangonHandler)  //Oke
	http.HandleFunc("/samsat/lokprwpkl", handlers.SamsatLokPrwPKLHandler) //oke
	http.HandleFunc("/samsat/pekalongan", handlers.SamsatPekalonganHandler) //Oke
	http.HandleFunc("/samsat/pemalang", handlers.SamsatPemalangHandler) //Oke
	http.HandleFunc("/samsat/tegal", handlers.SamsatTegalHandler) //Oke
	http.HandleFunc("/samsat/brebes", handlers.SamsatBrebesHandler) //Oke
	http.HandleFunc("/samsat/batang", handlers.SamsatBatangHandler) //Oke
	http.HandleFunc("/samsat/kajen", handlers.SamsatKajenHandler) //Oke
	http.HandleFunc("/samsat/slawi", handlers.SamsatSlawiHandler) //Oke
	http.HandleFunc("/samsat/bumiayu", handlers.SamsatBumiayuHandler) //Oke
	http.HandleFunc("/samsat/tanjung", handlers.SamsatTanjungHandler) //Oke
	http.HandleFunc("/samsat/lokprwpti", handlers.SamsatLokPrwPTIHandler) //Oke
	http.HandleFunc("/samsat/pati", handlers.SamsatPatiHandler) //Oke
	http.HandleFunc("/samsat/kudus", handlers.SamsatKudusHandler) //Oke
	http.HandleFunc("/samsat/jepara", handlers.SamsatJeparaHandler) //Oke
	http.HandleFunc("/samsat/rembang", handlers.SamsatRembangHandler) //Oke
	http.HandleFunc("/samsat/blora", handlers.SamsatBloraHandler) //Oke
	http.HandleFunc("/samsat/cepu", handlers.SamsatCepuHandler) //Oke
	http.HandleFunc("/samsat/lokprwsmg", handlers.SamsatLokPrwSMGHandler) //Oke
	http.HandleFunc("/samsat/semarang1", handlers.SamsatSemarang1Handler) //Oke
	http.HandleFunc("/samsat/semarang2", handlers.SamsatSemarang2Handler)  //Oke
	http.HandleFunc("/samsat/semarang3", handlers.SamsatSemarang3Handler) //Oke
	http.HandleFunc("/samsat/lokprwskh", handlers.SamsatLokPrwSKHHandler) //Oke
	http.HandleFunc("/samsat/sukoharjo", handlers.SamsatSukoharjoHandler) //Oke
	http.HandleFunc("/samsat/karanganyar", handlers.SamsatKaranganyarHandler) //Oke
	http.HandleFunc("/samsat/wonogiri", handlers.SamsatWonogiriHandler) //Oke
	http.HandleFunc("/samsat/purwantoro", handlers.SamsatPurwantoroHandler) //Oke
	http.HandleFunc("/samsat/baturetno", handlers.SamsatBaturetnoHandler) //Oke

	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}