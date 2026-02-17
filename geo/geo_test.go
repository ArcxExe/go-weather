package geo

import "testing"

func TestGetMyLocation(t *testing.T) {
  expeced := Geodata{
    City: "London",
  }
  got,err := GetMyLocation("London")
  if err != nil {
    t.Error(err)
  }
  if got.City != expeced.City {
    t.Error("Вернулся не ожидаемый ответ ")
  }
}
