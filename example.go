package  main

import (
	"fmt"
	"encoding/json"
	"github.com/supalik/fthreeclient/goformclient"
	"github.com/google/uuid"
	"github.com/supalik/fthreeclient/models"
	"net/http"
	"strings"
	"unsafe"
	"strconv"
)

var (
	getf3httpClient = getform3client()
	getCustomClient = getform3CustomClient()
)
func getform3CustomClient() goformclient.Client{
	customClient := http.Client{}
	client := goformclient.NewGenerator().
		SetHttpClient(&customClient).
                Generate()
	return client


}
func  getform3client() goformclient.Client{
	client := goformclient.NewGenerator().
		DisableTimeouts(true).
		SetMaxIdleConns(20).
		Generate()
	commonHeader := make(http.Header)
	commonHeader.Set("Authorization", "ABCD")
	return client
}

func main(){
	getDatawithCustomClient()
	getData()
	uuidWithHyphen := uuid.New()
        uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	raw_data := ` {
	  "data": {
	    "type": "accounts",
	    "id": "4538c8264f6640f5bbfe734cf1f5c981",
	    "version": 0,
	    "organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
	    "attributes": {
	      "country": "GB",
	      "base_currency": "GBP",
	      "account_number": "41426819",
	      "bank_id": "400300",
	      "bank_id_code": "GBDSC",
	      "bic": "NWBKGB22",
	      "iban": "GB11NWBK40030041426819",
	          "name": [
	        "Samantha Holder1"
	      ],
	      "status": "confirmed"
	    }
	  }
	}`

	inputData :=models.Data{}
	fmt.Printf("\n Raw Data : %v \n", raw_data)
	json.Unmarshal([]byte(raw_data), &inputData)
	inputData.DataStruct.ID = uuid
	inputData.DataStruct.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	fmt.Printf("Data ID  ex: %v \n", inputData.DataStruct.ID)
	createData(inputData)
	deleteData(uuid, inputData.DataStruct.Version)
}

func createData(reqData models.Data){
	commonHeaders := make(http.Header)
        commonHeaders.Set("Content-Type", "application/json")
	ret := unsafe.Sizeof(reqData)
	str := fmt.Sprint(ret)
        fmt.Println(str)
        commonHeaders.Set("Content-Length", str)

        fmt.Printf("\nData to create an Account record %v \n", reqData)
	resp, err := getf3httpClient.Post("http://localhost:8080/v1/organisation/accounts", commonHeaders, reqData)
        if err !=nil{
                panic(err)
        }
	fmt.Println(resp.Status())
	fmt.Println(resp.StatusCode())
	fmt.Println(resp.String())
}

func getDatawithCustomClient(){

	fmt.Printf("Get using custom client")
	resp, err := getCustomClient.Get("http://localhost:8080/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", nil)
	if err !=nil{
		panic(err)
	}
	fmt.Printf("Get Response Code: %d \n", resp.StatusCode())
	bytes := resp.Bytes()
	fmt.Printf("\nGet Response Data: %s", string(bytes))
}
func getData(){

	resp, err := getf3httpClient.Get("http://localhost:8080/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", nil)
	if err !=nil{
		panic(err)
	}
	fmt.Printf("Get Response Code: %d \n", resp.StatusCode())
	bytes := resp.Bytes()
	fmt.Printf("\nGet Response Data: %s", string(bytes))
}

func deleteData(uuid string, v int){

        fmt.Printf("\nDelete Account data for version %d", v)
	url := "http://localhost:8080/v1/organisation/accounts/" + uuid + "/" + "?" + "version=" + strconv.Itoa(v)
	resp, err := getf3httpClient.Delete(url, nil)
        if err !=nil{
                panic(err)
        }
        fmt.Printf("\nDelete response Code: %d \n", resp.StatusCode())

}
