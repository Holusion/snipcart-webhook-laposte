package main

import(
  "os"
  "net/http"
  "io/ioutil"
  "log"
  "encoding/json"
  "snipcart-webhook-laposte/monkeyapi"
  "snipcart-webhook-laposte/shippingRates"
)

const BasePrice = 0.1

func GetHandler(c *ApiConn)  http.Handler{
  return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
    var prices []shippingRates.Rate
    if r.ContentLength == 0 {
      //empty body
      log.Printf("Error: No content")
      http.Error(w, "No Content", http.StatusInternalServerError)
      return
    }
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
      log.Printf("Error Reading Body : %s",err)
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    params, err := Parse(body)
    if err != nil {
      log.Printf("Error Parsing Body : %s",err)
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    if params.Content.Country == "FR" {
      if 10 < params.Content.Quantity {
        prices, err = c.ApiRequest("colis",int(params.Content.Weight))
      }else{
        prices, err = c.ApiRequest("lettre",int(params.Content.Weight))
      }
    }else{
      log.Printf("Foreign Request : %. Using monkeyapi",params.Content.Country)
      //As laposte's API doesn't support foreign shipping, we use their human oriented calculator
      countryCode,ok := monkeyapi.Countries[params.Content.Country]
      if !ok {
        log.Printf("Unknown country code : %s. Falling back", params.Content.Country)
        countryCode = monkeyapi.Countries["US"]
      }
      prices,err = monkeyapi.Request(
        countryCode,
        monkeyapi.ShipTypes["Letter"],
        int(params.Content.Weight),
      )
    }
    if err != nil {
      log.Printf("Error calculating rates : %s",err)
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    addEnvelope(prices)
    wrappedPrices := shippingRates.Rates{Rates:prices}

    js, err := json.Marshal(wrappedPrices)
    w.Write(js)
  })
}

func addEnvelope(rates []shippingRates.Rate){
  for i := range rates {
    rates[i].Cost += BasePrice
  }
}

func main() {
  conn := CreateApiConn(os.Getenv("API_KEY"))
  http.Handle("/", GetHandler(conn))
    http.ListenAndServe(":7331", nil)
}
