package stub

const (
	StubJntApiNotFound = `{
		"error_id":"404",
		"error_message":"Invalid AWB number"
	}`
	StubJntApi = `{
		"awb":"XXX",
		"orderid":"XXX",
		"detail":{
			"shipped_date":"2022-08-09 21:29:15",
			"services_code":"EZ",
			"services_type":"",
			"actual_amount":8000,
			"weight":1050,
			"qty":1,
			"itemname":"Laut Bercerita",
			"detail_cost":{
				"shipping_cost":8000,
				"add_cost":0,
				"insurance_cost":0,
				"cod":0,
				"return_cost":0
			},
			"sender":{
				"name":"Matraman",
				"addr":"DKI JAKARTA, JAKARTA, Jl. Matraman Raya No.46-48, RT.12/RW.2, Kab. Manggis, Kec. Matraman, Kota Jakarta Timur",
				"zipcode":"",
				"city":"JAKARTA",
				"geoloc":""
			},
			"receiver":{
				"name":"Aletheia Tan",
				"addr":"DKI JAKARTA, JAKARTA, Apartemen Mediterania Palace Kemayoran, Tower C Unit 06E/AC",
				"zipcode":"10610",
				"city":"JAKARTA",
				"geoloc":""
			},
			"driver":{
				"id":"",
				"name":"XXX",
				"phone":"",
				"photo":""
			},
			"delivDriver":{
				"id":"XXX",
				"name":"XXX",
				"phone":"XXX",
				"photo":""
			}
		},
		"history":[
			{
			"date_time":"2022-08-09 13:25:28",
			"city_name":"JAKARTA",
			"status":"Manifes",
			"status_code":101,
			"storeName":"",
			"nextSiteName":"DC_KEMAYORAN",
			"note":"",
			"receiver":"",
			"driverName":"",
			"driverPhone":"",
			"presenter":"",
			"agentName":"",
			"presentername":""
			},
			{
			"date_time":"2022-08-09 21:29:15",
			"city_name":"JAKARTA",
			"status":"Paket telah diterima oleh DC_PULO_GADUNG",
			"status_code":100,
			"storeName":"DC_PULO_GADUNG",
			"nextSiteName":"DC_PULO_GADUNG",
			"note":"",
			"receiver":"",
			"distributionFlag":"0",
			"driverName":"XXX",
			"driverPhone":"XXX",
			"presenter":"",
			"agentName":"",
			"presentername":""
			},
			{
			"date_time":"2022-08-09 22:40:46",
			"city_name":"JAKARTA",
			"status":"Paket akan dikirimkan ke JAT_GATEWAY",
			"status_code":100,
			"storeName":"DC_PULO_GADUNG",
			"nextSiteName":"JAT_GATEWAY",
			"note":"",
			"receiver":"",
			"distributionFlag":"0",
			"driverName":"",
			"driverPhone":"",
			"presenter":"",
			"agentName":"",
			"presentername":""
			},
			{
			"date_time":"2022-08-09 22:43:43",
			"city_name":"JAKARTA",
			"status":"Paket akan dikirimkan ke JAT_GATEWAY",
			"status_code":100,
			"storeName":"DC_PULO_GADUNG",
			"nextSiteName":"JAT_GATEWAY",
			"note":"Mobil",
			"receiver":"",
			"distributionFlag":"0",
			"driverName":"",
			"driverPhone":"",
			"presenter":"",
			"agentName":"",
			"presentername":""
			},
			{
			"date_time":"2022-08-09 23:08:28",
			"city_name":"JAKARTA",
			"status":"Paket akan dikirimkan ke JAT_GATEWAY",
			"status_code":100,
			"storeName":"DC_PULO_GADUNG",
			"nextSiteName":"JAT_GATEWAY",
			"note":"",
			"receiver":"",
			"distributionFlag":"0",
			"driverName":"",
			"driverPhone":"",
			"presenter":"",
			"agentName":"",
			"presentername":""
			},
			{
			"date_time":"2022-08-10 01:11:20",
			"city_name":"JAKARTA",
			"status":"Paket telah sampai di JAT_GATEWAY",
			"status_code":100,
			"storeName":"JAT_GATEWAY",
			"nextSiteName":"JAT_GATEWAY",
			"note":"Mobil",
			"receiver":"",
			"distributionFlag":"1",
			"driverName":"",
			"driverPhone":"",
			"presenter":"",
			"agentName":"",
			"presentername":""
			},
			{
			"date_time":"2022-08-10 02:53:20",
			"city_name":"JAKARTA",
			"status":"Paket telah sampai di JAT_GATEWAY",
			"status_code":100,
			"storeName":"JAT_GATEWAY",
			"nextSiteName":"JAT_GATEWAY",
			"note":"",
			"receiver":"",
			"distributionFlag":"1",
			"driverName":"",
			"driverPhone":"",
			"presenter":"",
			"agentName":"",
			"presentername":""
			},
			{
			"date_time":"2022-08-10 03:13:49",
			"city_name":"JAKARTA",
			"status":"Paket akan dikirimkan ke DC_KEMAYORAN",
			"status_code":100,
			"storeName":"JAT_GATEWAY",
			"nextSiteName":"DC_KEMAYORAN",
			"note":"Mobil",
			"receiver":"",
			"distributionFlag":"1",
			"driverName":"",
			"driverPhone":"",
			"presenter":"",
			"agentName":"",
			"presentername":""
			},
			{
			"date_time":"2022-08-10 05:24:33",
			"city_name":"JAKARTA",
			"status":"Paket telah sampai di DC_KEMAYORAN",
			"status_code":100,
			"storeName":"DC_KEMAYORAN",
			"nextSiteName":"DC_KEMAYORAN",
			"note":"",
			"receiver":"",
			"distributionFlag":"0",
			"driverName":"",
			"driverPhone":"",
			"presenter":"",
			"agentName":"",
			"presentername":""
			},
			{
			"date_time":"2022-08-10 05:49:46",
			"city_name":"JAKARTA",
			"status":"Paket akan dikirim ke alamat penerima",
			"status_code":100,
			"storeName":"DC_KEMAYORAN",
			"nextSiteName":"DC_KEMAYORAN",
			"note":"",
			"receiver":"",
			"distributionFlag":"0",
			"driverName":"XXX",
			"driverPhone":"XXX",
			"presenter":"",
			"agentName":"",
			"presentername":""
			},
			{
			"date_time":"2022-08-10 10:03:40",
			"city_name":"JAKARTA",
			"status":"Paket telah diterima",
			"status_code":200,
			"storeName":"DC_KEMAYORAN",
			"nextSiteName":"DC_KEMAYORAN",
			"note":"",
			"receiver":"XXX",
			"driverName":"XXX",
			"driverPhone":"",
			"presenter":"TRUE",
			"agentName":"",
			"presentername":"Resepsionis"
			}
		]
 	}`
)
