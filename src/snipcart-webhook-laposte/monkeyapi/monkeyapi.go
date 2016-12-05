package monkeyapi

import(
  "regexp"
  "snipcart-webhook-laposte/shippingRates"
  "fmt"
  "strconv"
  "strings"
  "net/http"
  "io/ioutil"
)

var ShipTypes map[string]int = map[string]int{
  "Letter" : 3,
}

var Countries map[string]int = map[string]int{
  "US" : 220,
}
const URL = "https://www.laposte.fr/particulier/outils/calculateur-de-tarif/lettres_et_po/%d-%d-1-%d-1"
func Request(country int, shipType int, weight int) ([]shippingRates.Rate,error){
  resp, err := http.Post(fmt.Sprintf(URL,shipType,country,weight), "", nil)
  if err != nil {
    return nil,err
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil,err
  }
  rate,err := parseResponse(body)
  return []shippingRates.Rate{*rate},err
}

var re *regexp.Regexp = regexp.MustCompile(`<span class="h1">([\d,]*) €</span>`)

func parseResponse(htmlCode []byte) (*shippingRates.Rate,error){
  res := re.FindSubmatch(htmlCode)
  //replace "," with "."
  str := strings.Replace(string(res[1]), ",", ".", 1)
  price, err := strconv.ParseFloat(str,32)
  if err != nil {
    return nil, err
  }
  return &shippingRates.Rate{Cost:float32(price), Description:"La Poste Letre Verte", Delivery: 8},nil
}
