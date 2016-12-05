package main

import (
    "testing"
    "io/ioutil"
)

func TestParse(t *testing.T) {
  fileData, err := ioutil.ReadFile("../../fixtures/test_1.json")
  if err != nil {
    t.Fatalf("Can't open fixture : %s",err)
  }
  p, err := Parse(fileData)
  if err != nil {
    t.Fatalf("Can't parse JSON data : %s",err)
  }
  if p.EventName != "shippingrates.fetch" {
    t.Fatalf("Invalid result. Expected p.EventName to equal \"shippingrates.fetch\". Got : \"%s\" : ",p.EventName)
  }
  if p.Content.Weight != 20.00 {
    t.Fatalf("Invalid result. Expected p.Content.Weight to equal \"20.00\". Got : \"%f\" : ",p.Content.Weight)
  }
}
