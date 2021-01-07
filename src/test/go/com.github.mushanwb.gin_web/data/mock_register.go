package data

var RegisterData = map[string]interface{}{
	"name":     "John",
	"phone":    "13098858711",
	"password": "1234567",
}

var RegisterBadData = map[string]interface{}{
	"name":     "John",
	"phone":    "130988587",
	"password": "12345王",
}
