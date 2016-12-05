package main

import(
  "net/http"
  "fmt"
  "encoding/json"
  "io/ioutil"
  "sort"
  "snipcart-webhook-laposte/shippingRates"
)

type ApiErr struct{
  Code string `json:"code"`
  Message string `json:"message"`
}

type ApiRes struct{
  Channel string `json:"channel"`
  Product string `json:"product"`
  Price float32 `json:"price"`
  Currency string `json:"currency"`
}
type ApiConn struct{
  key string
}
func CreateApiConn(key string) *ApiConn{
  return &ApiConn{key}
}

//Type : colis, chronopost, lettre
//Weight : in grams

func (c *ApiConn) ApiRequest(t string, weight int) ([]shippingRates.Rate,error){
  client := &http.Client{}
  req, err := http.NewRequest(
    "GET",
    fmt.Sprintf("https://api.laposte.fr/tarifenvoi/v1?type=%s&weight=%d", t, weight),
    nil,
  )
  req.Header.Add("X-Okapi-Key", c.key)
  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }
  prices,err := parseResponse(body)
  if err != nil {
    return nil, err
  }
  sort.Sort(ByPrice(prices))
  return resToShippingRates(prices), nil
}

func parseResponse( body []byte) ([]ApiRes, error){
  var data []ApiRes
  err := json.Unmarshal(body, &data)
  if len(data) < 1{
    return nil, fmt.Errorf("Empty response. Original Body : %s", body)
  }
  //fmt.Printf("Body : %s",body)
  return data, err
}

func resToShippingRates(res []ApiRes) []shippingRates.Rate {
  var rates []shippingRates.Rate
  for _,r := range res{
    //Rule out products we don't need
    if r.Channel != "en ligne" ||r.Product == "ECOPLI"{
      continue
    }
    if r.Product == "Lettre prioritaire"{
      rates = append(rates, shippingRates.Rate{
        Cost:r.Price,
        Description:r.Product,
        Delivery: 3,
      })
    }else{
      rates = append(rates, shippingRates.Rate{
        Cost:r.Price,
        Description:r.Product,
        Delivery: 5,
      })
    }
  }
  return rates
}

type ByPrice []ApiRes

func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice)  Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPrice)  Less(i, j int) bool { return a[i].Price < a[i].Price }
