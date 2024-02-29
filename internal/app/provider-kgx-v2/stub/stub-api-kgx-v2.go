package stub

const (
	StubKgxV2ApiNotFound = `{
    "code": 400,
    "name": "Validation Error",
    "message": "Required attribute not set correctly in Request Body, see details for more info",
    "details": [
			{
				"attribute": "connote_code",
				"error": "0000 is not a valid connote_code"
			}
    ]
	}`
	StubKgxV2Api = `{
    "data": [
			{
				"id": 1,
				"action": "Create",
				"connote_state": "Create",
				"content": "Connote telah dibuat oleh Administrator di lokasi GRAMEDIA.COM MTM",
				"date": "2022-08-09T15:58:25+07:00",
				"created_at": "2022-08-09T15:58:24+07:00",
				"updated_at": "2022-08-09T15:58:24+07:00",
				"location_name": "GRAMEDIA.COM MTM",
				"username": "Administrator",
				"connote_code": "XXX",
				"is_hide": 0,
				"coordinate": ""
			},
			{
				"id": 1,
				"action": "Pickup Request",
				"connote_state": "Pickup Request",
				"content": "Pickup request telah dibuat dengan Nomor FS/22/08/11660",
				"date": "2022-08-09T16:45:00+07:00",
				"created_at": "2022-08-09T17:02:58+07:00",
				"updated_at": "2022-08-09T17:02:58+07:00",
				"location_name": "GRAMEDIA.COM MTM",
				"username": "XXX",
				"connote_code": "XXX",
				"is_hide": 0,
				"coordinate": ""
			},
			{
				"id": 1,
				"action": "Assign",
				"connote_state": "Assign",
				"content": "Kurir Teguh Kunanto ditugaskan untuk pickup",
				"date": "2022-08-09T17:02:58+07:00",
				"created_at": "2022-08-09T17:02:58+07:00",
				"updated_at": "2022-08-09T17:02:58+07:00",
				"location_name": "GRAMEDIA.COM MTM",
				"username": "Teguh Kunanto",
				"connote_code": "XXX",
				"is_hide": 0,
				"coordinate": ""
			},
			{
				"id": 1,
				"action": "OnPick",
				"connote_state": "OnPick",
				"content": "Paket Anda telah melewati proses pickup package oleh Teguh Kunanto di GRAMEDIA.COM MTM",
				"date": "2022-08-09T19:22:20+07:00",
				"created_at": "2022-08-09T19:22:18+07:00",
				"updated_at": "2022-08-09T19:22:18+07:00",
				"location_name": "GRAMEDIA.COM MTM",
				"username": "Teguh Kunanto",
				"connote_code": "XXX",
				"is_hide": 0,
				"coordinate": ""
			},
			{
				"id": 1,
				"action": "Inbound",
				"connote_state": "Inbound",
				"content": "Paket Anda telah melewati proses inbound oleh Yahya di Hub Palmerah (Jakarta Pusat)",
				"date": "2022-08-10T00:02:21+07:00",
				"created_at": "2022-08-10T00:02:20+07:00",
				"updated_at": "2022-08-10T00:02:20+07:00",
				"location_name": "Hub Palmerah (Jakarta Pusat)",
				"username": "Yahya",
				"connote_code": "XXX",
				"is_hide": 0,
				"coordinate": ""
			},
			{
				"id": 1,
				"action": "Outbound",
				"connote_state": "Outbound",
				"content": "Paket Anda telah melewati proses outbound oleh Khalid di Hub Palmerah (Jakarta Pusat)",
				"date": "2022-08-10T00:38:03+07:00",
				"created_at": "2022-08-10T00:38:00+07:00",
				"updated_at": "2022-08-10T00:38:00+07:00",
				"location_name": "Hub Palmerah (Jakarta Pusat)",
				"username": "Khalid",
				"connote_code": "XXX",
				"is_hide": 0,
				"coordinate": ""
			},
			{
				"id": 1,
				"action": "Inbound",
				"connote_state": "Inbound",
				"content": "Paket Anda telah melewati proses inbound oleh Dicky Bagus Tanggara di Hub Palmerah (Jakarta Pusat)",
				"date": "2022-08-10T01:25:02+07:00",
				"created_at": "2022-08-10T01:25:00+07:00",
				"updated_at": "2022-08-10T01:25:00+07:00",
				"location_name": "Hub Palmerah (Jakarta Pusat)",
				"username": "Dicky Bagus Tanggara",
				"connote_code": "XXX",
				"is_hide": 0,
				"coordinate": ""
			},
			{
				"id": 1,
				"action": "Outbound",
				"connote_state": "Outbound",
				"content": "Paket Anda telah melewati proses outbound oleh Miko Kintana di Sortation Palmerah",
				"date": "2022-08-10T03:15:43+07:00",
				"created_at": "2022-08-10T03:15:38+07:00",
				"updated_at": "2022-08-10T03:15:38+07:00",
				"location_name": "Sortation Palmerah",
				"username": "Miko Kintana",
				"connote_code": "XXX",
				"is_hide": 0,
				"coordinate": ""
			},
			{
				"id": 1,
				"action": "Inbound",
				"connote_state": "Inbound",
				"content": "Paket Anda telah melewati proses inbound oleh Danur Sodo Danang Palagan di Hub Fatmawati (Jakarta Selatan)",
				"date": "2022-08-10T07:51:58+07:00",
				"created_at": "2022-08-10T07:51:53+07:00",
				"updated_at": "2022-08-10T07:51:53+07:00",
				"location_name": "Hub Fatmawati (Jakarta Selatan)",
				"username": "Danur Sodo Danang Palagan",
				"connote_code": "XXX",
				"is_hide": 0,
				"coordinate": ""
			},
			{
				"id": 1,
				"action": "WithDeliveryCourier",
				"connote_state": "WithDeliveryCourier",
				"content": "Paket Anda telah melewati proses handover delivery oleh Budi Prayoga di Hub Fatmawati (Jakarta Selatan)",
				"date": "2022-08-10T08:44:41+07:00",
				"created_at": "2022-08-10T08:44:40+07:00",
				"updated_at": "2022-08-10T08:44:40+07:00",
				"location_name": "Hub Fatmawati (Jakarta Selatan)",
				"username": "Budi Prayoga",
				"connote_code": "XXX",
				"is_hide": 0,
				"coordinate": ""
			},
			{
				"id": 1,
				"action": "Delivered",
				"connote_state": "Delivered",
				"content": "Paket Anda Telah diterima oleh XXX sebagai Keluarga Serumah",
				"date": "2022-08-10T10:39:26+07:00",
				"created_at": "2022-08-10T10:39:26+07:00",
				"updated_at": "2022-08-10T10:39:26+07:00",
				"location_name": "XXX",
				"username": "XXX",
				"connote_code": "XXX",
				"is_hide": 0,
				"coordinate": ""
			}
    ],
    "from": 1,
    "to": 11,
    "total": 11,
    "per_page": 20,
    "current_page": 1,
    "last_page": 1,
    "first_page_url": "XXX",
    "prev_page_url": null,
    "next_page_url": null,
    "last_page_url": "XXX",
    "path": "XXX"
	}`
)
