package main

import "fmt"

// Valor Zero
/*
    - Declaração vs. inicialização vs. atribuição de valor. Variáveis: caixas postais.
    - O que é valor zero?
    - Os zeros:
        - ints: 0
        - floats: 0.0
        - booleans: false
        - strings: ""
        - pointers, functions, interfaces, slices, channels, maps: nil
    - Use := sempre que possível.
    - Use var para package-level scope.
*/
var x int
var y string
var z bool
func main(){
    fmt.Printf("%v\n%v\n%v",x,y,z)
}
