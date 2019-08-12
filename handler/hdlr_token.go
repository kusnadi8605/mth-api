package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	conf "mth-api/config"
	dts "mth-api/datastruct"
	mdl "mth-api/models"
)

//TokenHandler return single data
func TokenHandler(conn *conf.Connection) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var reqToken dts.TokenReq
		var TokenResponse dts.TokenResponse

		body, err := ioutil.ReadAll(req.Body)
		err = json.Unmarshal(body, &reqToken)

		conf.Logf("Request Token : %s", reqToken)

		if err != nil {
			TokenResponse.ResponseCode = "500"
			TokenResponse.ResponseDesc = err.Error()
			json.NewEncoder(w).Encode(TokenResponse)
			conf.Logf("Response GetToken : %s", TokenResponse)
			return
		}

		userName := reqToken.UserName
		password := reqToken.Password

		Token, err := mdl.GetToken(conn, userName, password)

		if err != nil {
			TokenResponse.ResponseCode = "501"
			TokenResponse.ResponseDesc = err.Error()
			json.NewEncoder(w).Encode(TokenResponse)
			conf.Logf("Response GetToken : %s", TokenResponse)
			return
		}

		if len(Token) < 1 {
			TokenResponse.ResponseCode = "501"
			TokenResponse.ResponseDesc = "data not found"
			json.NewEncoder(w).Encode(TokenResponse)
			conf.Logf("Response GetToken : %s", TokenResponse)
			return
		}

		TokenResponse.ResponseCode = "000"
		TokenResponse.ResponseDesc = "Success"
		TokenResponse.Payload = Token
		conf.Logf("Response GetToken : %s", TokenResponse)
		json.NewEncoder(w).Encode(TokenResponse)
	}
}
