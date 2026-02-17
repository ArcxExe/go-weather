package weather

import (
	"demo/weather/geo"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetWeather(geo geo.Geodata , formant int) string {
  baseUrl , err := url.Parse("https://wttr.in/" + geo.City)

  if nil != err {
    return ""
  }
  params := url.Values{}
  params.Add("format",fmt.Sprint(formant))
  baseUrl.RawQuery = params.Encode()
  resp , err := http.Get(baseUrl.String())
  if err != nil {
    return ""
  }
  body , err := io.ReadAll(resp.Body)
  if err != nil {
    fmt.Println(err.Error())
    return ""
  }
  return string(body)

}

