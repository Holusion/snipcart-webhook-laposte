package monkeyapi

import (
    "testing"
    "io/ioutil"
)

func TestParse(t *testing.T) {
  fileData, err := ioutil.ReadFile("../../../fixtures/laposte_client_calc.html")
  if err != nil {
    t.Fatalf("Can't open fixture : %s",err)
  }
  p, err := parseResponse(fileData)
  if err != nil {
    t.Fatalf("Can't parse HTML : %s",err)
  }
  if p.Cost != 6.25 {
    t.Fatalf("Invalid price : %f", p.Cost)
  }
}
