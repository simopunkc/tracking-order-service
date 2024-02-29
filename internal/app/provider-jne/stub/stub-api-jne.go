package stub

const (
	StubJneApiNotFound = `{
    "error": "Cnote No. Not Found.",
    "status": false
	}`
	StubJneApi = `{
    "cnote": {
			"cnote_no": "XXX",
			"reference_number": "XXX",
			"cnote_origin": "XXX",
			"cnote_destination": "XXX",
			"cnote_services_code": "CTCYES",
			"servicetype": "CTCYES19",
			"cnote_cust_no": "XXX",
			"cnote_date": "2022-08-10T08:01:05.000+07:00",
			"cnote_pod_receiver": "XXX",
			"cnote_receiver_name": "XXX",
			"city_name": "XXX ,JAKARTA SELATAN",
			"cnote_pod_date": "2022-08-10T20:23:00.000+07:00",
			"pod_status": "DELIVERED",
			"last_status": "DELIVERED TO [XXX | 10-08-2022 20:23 | JAKARTA ]",
			"cust_type": "066",
			"cnote_amount": "18000",
			"cnote_weight": "1",
			"pod_code": "D01",
			"keterangan": "YANG BERSANGKUTAN",
			"cnote_goods_descr": "JOJO BERHADAPAN DENGAN DIO DENGAN MEMPERTARUHKAN NYAWANYA SE",
			"freight_charge": "18000",
			"shippingcost": "18000",
			"insuranceamount": "0",
			"priceperkg": "18000",
			"signature": "XXX",
			"photo": "XXX",
			"long": "XXX",
			"lat": "XXX",
			"estimate_delivery": "1 Days"
    },
    "detail": [
			{
				"cnote_no": "XXX",
				"cnote_date": "2022-08-10T08:01:05.000+07:00",
				"cnote_weight": "1",
				"cnote_origin": "XXX",
				"cnote_shipper_name": "MATRAMAN",
				"cnote_shipper_addr1": "JL. MATRAMAN RAYA NO.46-48, RT",
				"cnote_shipper_addr2": "UR",
				"cnote_shipper_addr3": null,
				"cnote_shipper_city": "JAKARTA TIMUR",
				"cnote_receiver_name": "XXX",
				"cnote_receiver_addr1": "XXX",
				"cnote_receiver_addr2": "XXX",
				"cnote_receiver_addr3": null,
				"cnote_receiver_city": "XXX ,JAKART"
			}
    ],
    "history": [
			{
				"date": "09-08-2022 19:08",
				"desc": "SHIPMENT PICKED UP BY JNE COURIER [JAKARTA TIMUR]",
				"code": "PU0"
			},
			{
				"date": "10-08-2022 08:01",
				"desc": "SHIPMENT RECEIVED BY JNE COUNTER OFFICER AT [JAKARTA]",
				"code": "RC1"
			},
			{
				"date": "10-08-2022 13:15",
				"desc": "RECEIVED AT SORTING CENTER [JAKARTA]",
				"code": "OP1"
			},
			{
				"date": "10-08-2022 14:25",
				"desc": "SHIPMENT FORWARDED FROM TRANSIT CITY TO DESTINATION CITY [JAKARTA , HUB VETERAN BINTARO]",
				"code": "OP3"
			},
			{
				"date": "10-08-2022 16:29",
				"desc": "RECEIVED AT INBOUND STATION  [JAKARTA , HUB VETERAN BINTARO]",
				"code": "IP2"
			},
			{
				"date": "10-08-2022 19:15",
				"desc": "WITH DELIVERY COURIER  [JAKARTA]",
				"code": "IP3"
			},
			{
				"date": "10-08-2022 20:23",
				"desc": "DELIVERED TO [XXX | 10-08-2022 20:23 | JAKARTA ]",
				"code": "D01"
			}
    ]
	}`
)
