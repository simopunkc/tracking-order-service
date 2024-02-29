package stub

const (
	StubAnterajaApiNotFound = `{
    "status": 200,
    "info": "OK",
    "content": {
			"waybill_no": "000",
			"history": null,
			"order": null
    }
	}`
	StubAnterajaApi = `{
    "status": 200,
    "info": "OK",
    "content": {
			"waybill_no": "XXX",
			"history": [
				{
					"image_url": null,
					"hub_name": "Hub Sunter",
					"message": {
						"id": "Delivery sukses oleh SATRIA dan paket telah diterima oleh XXX. Terima kasih sudah menggunakan jasa AnterAja #PastiBawaHepi."
					},
					"params": null,
					"tracking_code": 250,
					"timestamp": "2022-08-10T09:28:33.758+0000"
				},
				{
					"image_url": null,
					"hub_name": "Hub Halim",
					"message": {
						"id": "SATRIA sudah ditugaskan dan parcel akan segera diantar ke penerima."
					},
					"params": null,
					"tracking_code": 240,
					"timestamp": "2022-08-10T04:48:59.954+0000"
				},
				{
					"image_url": null,
					"hub_name": "Hub Sunter",
					"message": {
						"id": "Parcel sudah tiba di SS Kota Jkt Sel - Karet Semanggi untuk proses delivery."
					},
					"params": null,
					"tracking_code": 230,
					"timestamp": "2022-08-09T23:17:44.000+0000"
				},
				{
					"image_url": null,
					"hub_name": "Hub Sunter",
					"message": {
						"id": "Parcel menuju ke Staging SS Kota Jkt Sel - Karet Semanggi."
					},
					"params": null,
					"tracking_code": 330,
					"timestamp": "2022-08-09T21:12:53.000+0000"
				},
				{
					"image_url": null,
					"hub_name": "Hub Sunter",
					"message": {
						"id": "Parcel sedang diproses di Hub Jakarta Utara"
					},
					"params": null,
					"tracking_code": 300,
					"timestamp": "2022-08-09T04:31:27.418+0000"
				},
				{
					"image_url": null,
					"hub_name": "Hub Bandung",
					"message": {
						"id": "Parcel menuju ke Hub Cimahi (proses transit)."
					},
					"params": null,
					"tracking_code": 332,
					"timestamp": "2022-08-08T19:05:01.206+0000"
				},
				{
					"image_url": null,
					"hub_name": "Hub Bandung",
					"message": {
						"id": "Parcel sudah tiba di Hub Cimahi."
					},
					"params": null,
					"tracking_code": 220,
					"timestamp": "2022-08-08T14:43:38.608+0000"
				},
				{
					"image_url": null,
					"hub_name": "Hub Bandung",
					"message": {
						"id": "Parcel sudah tiba di SS Kota Bandung - Babakan Tarogong untuk menuju ke hub."
					},
					"params": null,
					"tracking_code": 210,
					"timestamp": "2022-08-08T12:45:25.205+0000"
				},
				{
					"image_url": null,
					"hub_name": "Hub Bandung",
					"message": {
						"id": "Parcel sudah diterima oleh SATRIA."
					},
					"params": null,
					"tracking_code": 200,
					"timestamp": "2022-08-08T11:46:52.802+0000"
				},
				{
					"image_url": null,
					"hub_name": "Hub Bandung",
					"message": {
						"id": "Parsel belum siap di pickup dan akan dijadwalkan ulang."
					},
					"params": null,
					"tracking_code": 400,
					"timestamp": "2022-08-08T10:38:41.290+0000"
				},
				{
					"image_url": null,
					"hub_name": "Hub Bandung",
					"message": {
						"id": "SATRIA sudah ditugaskan dan parcel akan segera di-pickup."
					},
					"params": null,
					"tracking_code": 150,
					"timestamp": "2022-08-08T06:50:13.503+0000"
				},
				{
					"image_url": null,
					"hub_name": null,
					"message": {
						"id": "Pickup sudah di-request oleh shipper, dan SATRIA akan pickup parcel Senin 8 Agustus 2022 sekitar jam 14:00 - 16:00."
					},
					"params": null,
					"tracking_code": 100,
					"timestamp": "2022-08-08T06:47:40.925+0000"
				}
			],
			"order": {
				"booking_id": "XXX",
				"shipper": {
					"addressNotes": null,
					"address": "XXX",
					"phone": "XXX",
					"name": "Gramedia Bandung Festival Citylink",
					"postcode": "40232"
				},
				"waybill": "XXX",
				"receiver": {
					"addressNotes": null,
					"address": "XXX",
					"phone": "XXX",
					"name": "XXX",
					"postcode": "12190"
				},
				"service_fee": 11000,
				"weight": 1050,
				"service_code": "REG",
				"invoice": "XXX",
				"actual_shipper": {
					"proof_images": [],
					"name": null,
					"proof_images_url": [],
					"relationship": null
				},
				"actual_receiver": {
					"proof_images": [],
					"name": "XXX",
					"proof_images_url": [],
					"relationship": "XXX"
				}
			}
    }
	}`
)
