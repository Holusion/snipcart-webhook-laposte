package main

/*
import (
  "os"
  "net/http"
  "net/http/httptest"
  "testing"
  "bytes"
  "fmt"
  "encoding/json"
  "snipcart-webhook-laposte/shippingRates"
)

type Fixture struct{
  CountryCode string
  Quantity int
  Price float32
}
var fixtures []Fixture = []Fixture{
  Fixture{"US",1,1.4},
  //Fixture{"FR",1,0.70}, //Disable usage of authenticated API
}
func TestHandler(t *testing.T) {
  var results shippingRates.Rates
  conn := CreateApiConn(os.Getenv("API_KEY"))
  handler := GetHandler(conn)
  for _, f := range fixtures{
    buf := bytes.NewReader([]byte(format(f)))
    req, err := http.NewRequest("POST", "/", buf)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler.ServeHTTP(rr, req)
    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
    err = json.Unmarshal(rr.Body.Bytes(),&results)
    if err != nil {
      t.Errorf("Error decoding JSON : %s",err)
    }
    if results.Rates[0].Cost != f.Price + BasePrice{
        t.Errorf("handler returned unexpected body: got %f want %f+%f",
            results.Rates[0].Cost, f.Price, BasePrice)
    }
  }
}

func format(f Fixture) string{
  return fmt.Sprintf("{\"content\": {\"shippingAddressCountry\": \"%s\",\"items\": [{ \"quantity\": %d}],\"totalWeight\": %f}}",f.CountryCode,f.Quantity,float32(f.Quantity*10))
}
*/
