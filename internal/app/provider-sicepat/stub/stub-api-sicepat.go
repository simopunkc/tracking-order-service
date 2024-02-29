package stub

const (
	MockSicepatApiLastUpdated = "2022-08-11 03:55"
	MockSicepatApiNotFound    = `{
    "sicepat": {
			"status": {
				"code": 400,
				"description": "Can't get waybill from database"
			}
    }
	}`
	MockSicepatApi = `{
		"sicepat": {
			"status": {
				"code": 200,
				"description": "OK"
			},
			"result": {
				"waybill_number": "XXXXXXXXXXX",
				"kodeasal": "XXXXXXXXXXX",
				"kodetujuan": "XXXXXXXXXXXX",
				"service": "SIUNT",
				"weight": 1,
				"partner": "PT. Gramedia Asri Media",
				"sender": "Matraman",
				"sender_address": "DKI Jakarta",
				"receiver_address": "XXXXXXXXX, Kota Bengkulu",
				"receiver_name": "XXXXXXXXXXX",
				"realprice": 0,
				"totalprice": 38000,
				"POD_receiver": "",
				"POD_receiver_time": "",
				"send_date": "2022-08-11 05:10",
				"track_history": [
					{
						"date_time": "2022-08-10 18:23",
						"status": "PICKREQ",
						"city": "Terima permintaan pick up dari [PT. Gramedia Asri Media]"
					},
					{
						"date_time": "2022-08-10 21:13",
						"status": "PICK",
						"city": "Paket telah di pick up oleh [SIGESIT - Ahmad Revaldi]"
					},
					{
						"date_time": "2022-08-10 22:10",
						"status": "IN",
						"city": "Paket telah di input (manifested) di Jakarta Timur []"
					},
					{
						"date_time": "2022-08-10 22:10",
						"status": "IN",
						"city": "Paket telah di input (manifested) di Jakarta Timur [SiCepat Ekspres Matraman]"
					},
					{
						"date_time": "2022-08-11 03:55",
						"status": "OUT",
						"city": "Paket keluar dari Jakarta Timur [Jaktim Matraman]"
					}
				],
				"last_status": {
					"date_time": "2022-08-11 03:55",
					"status": "OUT",
					"city": "Paket keluar dari Jakarta Timur [Jaktim Matraman]"
				},
				"perwakilan": "BKS",
				"pop_sigesit_img_path": "XXX",
				"pod_sigesit_img_path": null,
				"pod_sign_img_path": null,
				"pod_img_path": null,
				"manifested_img_path": "XXX"
			}
		}
	}`
)
