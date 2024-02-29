package stub

const (
	StubPosApiNotFound = `{
		"response": null
	}`
	StubPosApi = `{
		"response":{
			 "data":[
					{
						 "status":"Diterima Sialamat",
						 "barcode":"XXX",
						 "eventDate":"2022-07-22 14:47:06",
						 "eventName":"SELESAI ANTAR",
						 "officeCode":"XXX",
						 "officeName":"DC SEKEJATI",
						 "description":"Selesai antar di XXX~~ Kantor Pos XXX~~ Diterima oleh XXX (ORANG SERUMAH)"
					},
					{
						 "status":"Antaran",
						 "barcode":"XXX",
						 "eventDate":"2022-07-22 09:19:01",
						 "eventName":"DALAM PROSES",
						 "officeCode":"40000",
						 "officeName":"BANDUNG",
						 "description":"Proses antar di BANDUNG~~ Kantor Pos BANDUNG 40000~~ "
					},
					{
						 "status":"-",
						 "barcode":"XXX",
						 "eventDate":"2022-07-21 18:45:38",
						 "eventName":"DALAM PROSES",
						 "officeCode":"40000",
						 "officeName":"BANDUNG",
						 "description":"Diteruskan ke Hub BANDUNG~~ BANDUNG 40000~~ "
					},
					{
						 "status":"Penerimaan di Loket ",
						 "barcode":"XXX",
						 "eventDate":"2022-07-21 11:59:36",
						 "eventName":"PENERIMAAN DI LOKET",
						 "officeCode":"XXX",
						 "officeName":"Jakartatimurjatinegara",
						 "description":"Penerimaan di loket : Jakartatimurjatinegara~~Kantor Pos Jakartatimurjatinegara XXX~~Pengirim : Matraman~~ Matraman~~ Matraman~~ Matraman~~ Penerima : XXX~~  Produk : PAKET KILAT KHUSUS~~ Berat : 1000gr~~Jenis : Kiriman Pos |gramedia~~Isi Kiriman : -"
					},
					{
						 "status":"Pra Collecting",
						 "barcode":"XXX",
						 "eventDate":"2022-07-21 10:49:13",
						 "eventName":"PRA COLLECTING",
						 "officeCode":"XXX",
						 "officeName":"Jakarta Timur",
						 "description":"Pra Collecting~~Jakarta Timur XXX~~Pengirim : Matraman~~ XXX~~ Jakarta Timur~~ Penerima : XXX~~Bandung~~Produk : KIRIMAN POS~~ Berat : 1050gr~~Jenis : PERCGRAMEDIA02100A~~Isi Kiriman : Para pemimpin adalah mereka yang maju le"
					}
			 ]
		}
 	}`
	StubPosApiStagingNotFound = `{
		"rs_tnt": null
	}`
	StubPosApiStaging = `{
    "rs_tnt": {
			"r_tnt": [
				{
					"barcode": "XXX",
					"officeCode": "XXX",
					"office": "TULUNGAGUNG",
					"eventName": "POSTING LOKET - DUMMY DATA",
					"eventDate": "2018-06-30 16:15:07",
					"description": "XXX"
				},
				{
					"barcode": "XXX",
					"officeCode": "XXX",
					"office": "TULUNGAGUNG",
					"eventName": "MANIFEST SERAH - DUMMY DATA",
					"eventDate": "2018-07-01 12:27:32",
					"description": "XXX"
				},
				{
					"barcode": "XXX",
					"officeCode": "XXX",
					"office": "Ngunut",
					"eventName": "MANIFEST TERIMA - DUMMY DATA",
					"eventDate": "2018-07-02 12:00:01",
					"description": "XXX"
				},
				{
					"barcode": "XXX",
					"officeCode": "XXX",
					"office": "Ngunut",
					"eventName": "PROSES ANTAR - DUMMY DATA",
					"eventDate": "2018-07-02 12:07:04",
					"description": "XXX"
				},
				{
					"barcode": "XXX",
					"officeCode": "XXX",
					"office": "Ngunut",
					"eventName": "SELESAI ANTAR - DUMMY DATA",
					"eventDate": "2018-07-02 18:04:54",
					"description": "XXX"
				}
			]
    }
	}`
)
