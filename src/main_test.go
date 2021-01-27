package main

import "testing"

func TestSoma_4(t *testing.T){
  resultado := soma(2,2)
  if resultado != 4{
    t.Errorf("Soma esperada: %d, obtida: %d", 4, resultado)
  }
}

func TestSoma_5(t *testing.T){
  resultado := soma(2,-3)
  if resultado != -1{
    t.Errorf("Soma esperada: %d, obtida: %d", 4, resultado)
  }
}

func TestSoma_10(t *testing.T){
  resultado := soma(-2,12)
  if resultado != 10{
    t.Errorf("Soma esperada: %d, obtida: %d", 4, resultado)
  }
}


func TestSoma_2(t *testing.T){
  resultado := soma(10,-8)
  if resultado != 10{
    t.Errorf("Soma esperada: %d, obtida: %d", 2, resultado)
  }
}
