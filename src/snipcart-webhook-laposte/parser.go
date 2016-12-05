package main
import(
  "encoding/json"
)
type Params struct {
  EventName string `json:"eventName"`
  Mode  string `json:"mode"`
  Content Content  `json:"content"`
}
type Content struct{
  Token string `json:"token"`
  Country string `json:"shippingAddressCountry"`
  Weight float32 `json:"totalWeight"`
  Items []Item `json:"items"`
  Quantity int //Not in JSON
}

type Item struct{
  Id string `json:"id"`
  Quantity int `json:"quantity"`
}

func Parse(data []byte) (*Params,error){
  var p Params
  err := json.Unmarshal(data, &p)
  if err != nil {
    return nil, err
  }
  p.Content.Quantity = 0
  for _, i := range p.Content.Items {
    p.Content.Quantity += i.Quantity
  }
  return &p, nil
}
