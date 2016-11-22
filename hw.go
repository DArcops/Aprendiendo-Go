/*
	pruebas aprendiendo GO
	IMPORTAANTE SIEMPRE PARA COMENZAR DEBO DE EXPORTAR LAS RUTAS DE GOPATH:
	export GOPATH=$HOME/work
		export PATH=$PATH:$GOPATH/bin
*/

package main

import "fmt"

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}
/*
func param() (num int,cad string) {
	num = 2
	cad = " hola\n"
	return
}


var x,y,z int
var a,b,c = 1 , "hola" , false
var prueba,lol = param()*/

func main() {
/*	fmt.Print(returnPi(),"\n")
	recorrido := 1 << 3
	//fmt.Print(recorrido,"\n")
	//fmt.Print(prueba," ",lol,"\n")

	//for sintaxis
	forBasico()

	//fibonacci en Go
//	fmt.Print(fib(5),"\n")*/
	for i := 0 ; i < 10 ; i++ {
		array[i] = i
		fmt.Println(array[i])
	}
	p = array
	fmt.Println("el tamaño de p es: ",len(p),p[1:4])
	mapa := make(map[string]int)
	mapa["hola Diego"] = 13
	for i := range array {
		fmt.Println(i)
	}
	fmt.Println(mapa["hola Diego"])

	type Vertex struct {
    Lat, Long float64
  }

	var m = map[string]Vertex{
    "Bell Labs": Vertex{
        40.68433, -74.39967,
    },
    "Google": Vertex{
        37.42202, -122.08408,
    },
  }
	fmt.Println(m)


pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
        )
    }


}
