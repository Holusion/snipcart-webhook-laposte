package shippingRates


type Rate struct{
  Cost float32 `json:"cost"`
  Description string `json:"description"`
}

type Rates struct{
  Rates []Rate `json:"rates"`
}
