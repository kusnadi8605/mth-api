package handler

import (
	"encoding/json"
	"io/ioutil"
	conf "mth-api/config"
	dts "mth-api/datastruct"
	mdl "mth-api/models"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

//CreateHandler ..
//save to mysql and redis
func CreateHandler(conn *conf.Connection) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var prodRequest dts.ProductReq
		var prodResponse dts.ProductResponse

		body, err := ioutil.ReadAll(req.Body)
		json.Unmarshal(body, &prodRequest)

		//validation
		v := validator.New()
		prods := &dts.ProductReq{
			SKU:         prodRequest.SKU,
			ProductName: prodRequest.ProductName,
			ProductDesc: prodRequest.ProductDesc,
			Quantity:    prodRequest.Quantity,
			Price:       prodRequest.Price,
			UserID:      prodRequest.UserID,
		}
		err = v.Struct(prods)

		if err != nil {
			prodResponse.ResponseCode = "500"
			prodResponse.ResponseDesc = err.Error()
			json.NewEncoder(w).Encode(prodResponse)
			conf.Logf("Decode Product : %s", err)
			return
		}

		err = mdl.Create(conn, prodRequest)
		if err != nil {

			prodResponse.ResponseCode = "501"
			prodResponse.ResponseDesc = err.Error()
			json.NewEncoder(w).Encode(prodResponse)
			conf.Logf("Response Product : %v", prodResponse.ResponseDesc)
			return
		}

		prodResponse.ResponseCode = "000"
		prodResponse.ResponseDesc = "Success"
		json.NewEncoder(w).Encode(prodResponse)
		conf.Logf("Response Product : %v", prodResponse.ResponseDesc)

	}
}

//UpdateHandler ..
func UpdateHandler(conn *conf.Connection) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var prodRequest dts.ProductReq
		var prodResponse dts.ProductResponse

		body, err := ioutil.ReadAll(req.Body)
		json.Unmarshal(body, &prodRequest)

		//validation
		v := validator.New()
		prods := &dts.ProductReq{
			SKU:         prodRequest.SKU,
			ProductName: prodRequest.ProductName,
			ProductDesc: prodRequest.ProductDesc,
			Quantity:    prodRequest.Quantity,
			Price:       prodRequest.Price,
			UserID:      prodRequest.UserID,
		}
		err = v.Struct(prods)

		if err != nil {
			prodResponse.ResponseCode = "500"
			prodResponse.ResponseDesc = err.Error()
			json.NewEncoder(w).Encode(prodResponse)
			conf.Logf("Decode Product : %s", err)
			return
		}

		err = mdl.Update(conn, prodRequest)

		if err != nil {

			prodResponse.ResponseCode = "501"
			prodResponse.ResponseDesc = err.Error()
			json.NewEncoder(w).Encode(prodResponse)
			conf.Logf("Response Product : %v", prodResponse.ResponseDesc)
			return
		}

		prodResponse.ResponseCode = "000"
		prodResponse.ResponseDesc = "Success"
		json.NewEncoder(w).Encode(prodResponse)
		conf.Logf("Response Product : %v", prodResponse.ResponseDesc)

	}
}

//DetailHandler product
func DetailHandler(conn *conf.Connection) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var prodRequest dts.ProductReq
		var prodResponse dts.ProductResponse

		body, err := ioutil.ReadAll(req.Body)
		json.Unmarshal(body, &prodRequest)

		//validation
		v := validator.New()
		prods := &dts.ProductReq{
			SKU: prodRequest.SKU,
		}
		err = v.Struct(prods)

		if err != nil {
			prodResponse.ResponseCode = "500"
			prodResponse.ResponseDesc = err.Error()
			json.NewEncoder(w).Encode(prodResponse)
			conf.Logf("Decode Product : %s", err)
			return
		}

		sku := prodRequest.SKU
		val, err := mdl.Detail(conn, sku)

		if err != nil {

			prodResponse.ResponseCode = "501"
			prodResponse.ResponseDesc = err.Error()
			prodResponse.Payload = val
			json.NewEncoder(w).Encode(prodResponse)
			conf.Logf("Response Product : %v", prodResponse.ResponseDesc)
			return
		}

		prodResponse.ResponseCode = "000"
		prodResponse.ResponseDesc = "Success"
		prodResponse.Payload = val
		json.NewEncoder(w).Encode(prodResponse)
		conf.Logf("Response Product : %v", prodResponse.ResponseDesc)

	}
}

//DeleteHandler product
func DeleteHandler(conn *conf.Connection) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var prodRequest dts.ProductReq
		var prodResponse dts.ProductResponse

		body, err := ioutil.ReadAll(req.Body)
		json.Unmarshal(body, &prodRequest)

		//validation
		v := validator.New()
		prods := &dts.ProductReq{
			SKU:         prodRequest.SKU,
			ProductName: prodRequest.ProductName,
			ProductDesc: prodRequest.ProductDesc,
			Quantity:    prodRequest.Quantity,
			Price:       prodRequest.Price,
			UserID:      prodRequest.UserID,
		}

		err = v.Struct(prods)
		if err != nil {
			prodResponse.ResponseCode = "500"
			prodResponse.ResponseDesc = err.Error()
			json.NewEncoder(w).Encode(prodResponse)
			conf.Logf("Decode Product : %s", err)
			return
		}

		sku := prodRequest.SKU
		err = mdl.Delete(conn, sku)

		if err != nil {

			prodResponse.ResponseCode = "501"
			prodResponse.ResponseDesc = err.Error()
			json.NewEncoder(w).Encode(prodResponse)
			conf.Logf("Response Product : %v", prodResponse.ResponseDesc)
			return
		}

		prodResponse.ResponseCode = "000"
		prodResponse.ResponseDesc = "Success"
		json.NewEncoder(w).Encode(prodResponse)
		conf.Logf("Response Product : %v", prodResponse.ResponseDesc)

	}
}

//ListHandler product
func ListHandler(conn *conf.Connection) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var prodResponse dts.ProductResponse

		val, err := mdl.List(conn)

		if err != nil {

			prodResponse.ResponseCode = "501"
			prodResponse.ResponseDesc = err.Error()
			prodResponse.Payload = val
			json.NewEncoder(w).Encode(prodResponse)
			conf.Logf("Response Product : %v", prodResponse.ResponseDesc)
			return
		}

		prodResponse.ResponseCode = "000"
		prodResponse.ResponseDesc = "Success"
		prodResponse.Payload = val
		json.NewEncoder(w).Encode(prodResponse)
		conf.Logf("Response Product : %v", prodResponse.ResponseDesc)

	}
}
