package main

import "fmt"


func fibonacci(x int) int{ 
	if (x==0){
		return 0
	}
	if (x==1){
		return 1
	}
	if (x==2){
		return 1
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

func variadicFunctions(x ...int) int{
	var total int = 0
	var i int = 0

	for (i < len(x)){
		if (x[i] > total){
			total = x[i]
		}

		i = i + 1
	}
	return total
}

func generadorImpares() func() uint {
	i := uint(1) 
	return func() uint {
		var par = i
		i += 2
		return par
	}
}

func intercambia(x *int, y *int){
	aux := *x
	*x = *y
	*y = aux

}

func main() {
	var hastaQueNumeroFibonacci int = 6
	arregloParaNumeroMayor := []int{1, 4, 2, 6, 8 , 15, 2 , 6, 1}
	siguienteImpar := generadorImpares()
	var a,b int

	fmt.Println("Serie Fibonacci en la posicion ", hastaQueNumeroFibonacci)
	fmt.Println(fibonacci(hastaQueNumeroFibonacci))
	fmt.Println("")

	fmt.Println("Numero mayor")
	fmt.Println(variadicFunctions(arregloParaNumeroMayor...))
	fmt.Println("")

	fmt.Println("Numeros impares")
	fmt.Println(siguienteImpar())
	fmt.Println(siguienteImpar())
	fmt.Println(siguienteImpar())
	fmt.Println(siguienteImpar())
	fmt.Println("")

	fmt.Print("Dame el valor para a: ")
	fmt.Scanln(&a)
	fmt.Print("Dame el valor para b: ")
	fmt.Scanln(&b)
	fmt.Println("")
	
	fmt.Println("Intercambiar numeros")
	intercambia(&a,&b)
	fmt.Println("El valor de a es ", a)
	fmt.Println("El valor de b es ", b)

}