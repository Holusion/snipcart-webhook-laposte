package shippingRates


type Rate struct{
  Cost float32 `json:"cost"`
  Description string `json:"description"`
  Delivery int `json:"guaranteedDaysToDelivery"`
}

type Rates struct{
  Rates []Rate `json:"rates"`
}
