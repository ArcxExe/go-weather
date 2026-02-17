package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var NoCity = errors.New("NoCity")

type Geodata struct {

  City string `json:"city"` 
}
type CitePopualtionResponce struct {
  Err bool `json:"error"`
}
func GetMyLocation(city string) (*Geodata , error) {
  if city != "" {
    return &Geodata{
      City: city,
    } , nil
  }
  resp , err := http.Get("https://ipapi.co/json")
  if err != nil {
    return nil , err
  }
  // if resp.StatusCode != 200 {
  //   return nil , errors.New("Not 200")
  // }
  body , err := io.ReadAll(resp.Body)
  if err != nil {
    return nil , err
  }
  var geo Geodata
  err= json.Unmarshal(body,&geo)
  if err != nil {
  
    fmt.Println(err)
    panic(err)
  }
  return &geo , nil
}
func CheckSity(city string) bool {
  postBoby, _ := json.Marshal(map[string]string{
    "city": city,
  })
  resp,err := http.Post("https://coutriesnow.space/api/v0.1/countries/population/cities","application/json",bytes.NewBuffer(postBoby))
  if err != nil {
    return false
  }
  defer resp.Body.Close()
  boby , err := io.ReadAll(resp.Body)
  if err != nil {
    return false
  }
  var populationResponce CitePopualtionResponce 
  err = json.Unmarshal(boby,&populationResponce)
  if err != nil {
    panic(err)
  }
  return !populationResponce.Err
}
