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
		conf.Logf("Request Product : %v", prodRequest)

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
			conf.Logf("Response Product : %v", prodResponse)
			return
		}

		prodResponse.ResponseCode = "000"
		prodResponse.ResponseDesc = "Success"
		json.NewEncoder(w).Encode(prodResponse)
		conf.Logf("Response Product : %v", prodResponse)

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
			conf.Logf("Response Product : %v", prodResponse)
			return
		}

		err = mdl.Update(conn, prodRequest)

		if err != nil {

			prodResponse.ResponseCode = "501"
			prodResponse.ResponseDesc = err.Error()
			json.NewEncoder(w).Encode(prodResponse)
			conf.Logf("Response Product : %v", prodResponse)
			return
		}

		prodResponse.ResponseCode = "000"
		prodResponse.ResponseDesc = "Success"
		json.NewEncoder(w).Encode(prodResponse)
		conf.Logf("Response Product : %v", prodResponse)

	}
}

//DetailHandler product
func DetailHandler(conn *conf.Connection) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var actRequest dts.ProductAct
		var prodResponse dts.ProductResponse

		body, err := ioutil.ReadAll(req.Body)
		json.Unmarshal(body, &actRequest)

		//validation
		v := validator.New()
		prods := &dts.ProductAct{
			SKU: actRequest.SKU,
		}
		err = v.Struct(prods)

		if err != nil {
			prodResponse.ResponseCode = "500"
			prodResponse.ResponseDesc = err.Error()
			conf.Logf("Response Product : %v", prodResponse)
			return
		}

		sku := actRequest.SKU
		val, err := mdl.Detail(conn, sku)

		if err != nil {

			prodResponse.ResponseCode = "501"
			prodResponse.ResponseDesc = err.Error()
			prodResponse.Payload = val
			json.NewEncoder(w).Encode(prodResponse)
			conf.Logf("Response Product : %v", prodResponse)
			return
		}

		prodResponse.ResponseCode = "000"
		prodResponse.ResponseDesc = "Success"
		prodResponse.Payload = val
		json.NewEncoder(w).Encode(prodResponse)
		conf.Logf("Response Product : %v", prodResponse)

	}
}

//DeleteHandler product
func DeleteHandler(conn *conf.Connection) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var prodAct dts.ProductAct
		var prodResponse dts.ProductResponse

		body, err := ioutil.ReadAll(req.Body)
		json.Unmarshal(body, &prodAct)

		//validation
		v := validator.New()
		prods := &dts.ProductAct{
			SKU: prodAct.SKU,
		}

		err = v.Struct(prods)
		if err != nil {
			prodResponse.ResponseCode = "500"
			prodResponse.ResponseDesc = err.Error()
			json.NewEncoder(w).Encode(prodResponse)
			conf.Logf("Response Product : %v", prodResponse)
			return
		}

		sku := prodAct.SKU
		err = mdl.Delete(conn, sku)

		if err != nil {

			prodResponse.ResponseCode = "501"
			prodResponse.ResponseDesc = err.Error()
			json.NewEncoder(w).Encode(prodResponse)
			conf.Logf("Response Product : %v", prodResponse)
			return
		}

		prodResponse.ResponseCode = "000"
		prodResponse.ResponseDesc = "Success"
		json.NewEncoder(w).Encode(prodResponse)
		conf.Logf("Response Product : %v", prodResponse)

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
			conf.Logf("Response Product : %v", prodResponse)
			return
		}

		prodResponse.ResponseCode = "000"
		prodResponse.ResponseDesc = "Success"
		prodResponse.Payload = val
		json.NewEncoder(w).Encode(prodResponse)
		conf.Logf("Response Product : %v", prodResponse)

	}
}
