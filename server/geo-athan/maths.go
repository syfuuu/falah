/*
    Following maths calculation is inspired by http://praytimes.org/wiki/Code
*/


package main

import (
	"math"
)

type Maths struct {}

func (m Maths) getRadian(degree float64) float64 {
  return (degree * 180) / math.Pi 
}

func (m Maths) getDegree(radian float64) float64 {
  return (radian / 180) * math.Pi 
}

func (m Maths) fix(a float64, mode float64) float64 {
  a = a - mode * (math.Floor(a / mode))
  if a < 0 {
    return a + mode
  }
  return a
}

func (m Maths) Sin(degree float64) float64 {
  return math.Sin(m.getRadian(degree))
}

func (m Maths) Cos(degree float64) float64 {
  return math.Cos(m.getRadian(degree))
}

func (m Maths) Tan(degree float64) float64 {
  return math.Tan(m.getRadian(degree))
}

func (m Maths) Arcsin(x float64) float64 {
  return m.getDegree(math.Asin(x))
}

func (m Maths) Arccos(x float64) float64 {
  return m.getDegree(math.Acos(x))
}

func (m Maths) Arctan(x float64) float64 {
  return m.getDegree(math.Atan(x))
}

func (m Maths) Arctan2(x float64, y float64) float64 {
  return m.getDegree(math.Atan2(y, x))
}

func (m Maths) Arccot(x float64) float64 {
  inv := float64(1) / x
  return m.getDegree(math.Atan(inv))
}

func (m Maths) Fixangle(a float64) float64 {
  return m.fix(a, float64(360))
}

func (m Maths) Fixhour(a float64) float64 {
  return m.fix(a, float64(24))
}
