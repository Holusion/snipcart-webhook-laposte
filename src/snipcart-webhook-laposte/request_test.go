package main

import (
    "testing"
    "io/ioutil"
)

func TestRequest(t *testing.T) {
  fileData, err := ioutil.ReadFile("../../fixtures/laposte_body_test.json")
  if err != nil {
    t.Fatalf("Can't open fixture : %s",err)
  }
  p, err := parseResponse(fileData)
  if err != nil {
    t.Fatalf("Can't parse JSON data : %s",err)
  }
  if len(p) != 5 {
    t.Fatalf("Expected 5 results. Got %d",len(p))
  }
  if p[0].Product != "Lettre verte" {
    t.Fatalf("Expected p[0].Product to equal \"Lettre Verte\". Got : %s ",p[0].Product)
  }
}
