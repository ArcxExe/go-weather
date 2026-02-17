package main

import (
	"demo/weather/geo"
	"demo/weather/weather"
	// "flag"
	"fmt"
)

func main() {
	fmt.Print("Погода")
	// city := flag.String("cite", "", "Город пользователя")
	// format := flag.Int("format", 1, "Формат вывода погоды")
	// flag.Parse()
  geoData, err := geo.GetMyLocation("")
  if err != nil {
    fmt.Println(err)
  }
	fmt.Println(geoData)
  fmt.Println(weather.GetWeather(*geoData,1))
}
