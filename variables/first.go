package main

// func main e a funcao principal a ser executada pelo compilador

//Variables
func main() {
	/*
		//	=========================================================================
			//	Declaracao de variaveis:

				var a int = 5
				var b float32 = 4.32
				const pi float64 = 3.1416123113

				//forma inteligente de declarar (ja sabendo o tipo de dato da variavel)

				x, y := 14, 15

				fmt.Println(a)
				fmt.Println(b)
				fmt.Println(pi)
				fmt.Println(x, ",", y)

		//	=========================================================================

				//Operadores:

				x, y := 5, 6

				fmt.Println("x + y = ", x+y)
				fmt.Println("x - y = ", x-y)
				fmt.Println("x * y = ", x*y)
				fmt.Println("x / y = ", x/y)
				fmt.Println("x mod y = ", x%y)

		//	=========================================================================

				//Booleanos

				isbool := true
				hate := false

				fmt.Println(isbool && hate)
				fmt.Println(isbool || hate)
				fmt.Println(!isbool)
	*/

	//	=========================================================================

	//Pointers

	/*Ao ter um valor declarado em uma variavel, esse dado ocupa um espacio na memoria,
	pra poder modificar uma variavel fora de onde ela foi declarada e poder guardar o valor no espaco da memoria precisamos usar os Poiners

	"&" Para marcar o dado da variavel na funcao main
	"*int" Para marcar o tipo de dado que o espaco de memoria que o espaco de memoria vai receber
	"*x" Para puxar o novo dado da funcao "changeValue" no espaco da memoria da funcao "main"

	BASICAMENTE SE VOCE PRECISA MODIFICAR UM DADO NA VARIAVEL MAIN FORA DA MESMA FUNCAO QUE ELA É, EU PRECISO USAR O "POINTER"
	*/
	/*
			x := 10

			changeValue(&x)
			fmt.Println(x)
		}

		func changeValue(x *int) {
			*x = 7

	*/
	//	=========================================================================

	//Printf:
	/*
		var name string = "Ayrton Jofre"

		//Mostra a quantidade de caracteres incluindo espacos tem na variavel
		fmt.Println(len(name))
		//Utiliza o dado da variavel "name" e concatena com um string declarado na funcao
		fmt.Println(name + " é um cara bacana")

		const pi float64 = 3.14161231
		win := true
		x := 5

		//Print out um float como o numero PI incluindo um ponto especifico dele, podendo arredondar o valor
		fmt.Printf("%.3f \n", pi)

		//Print out o tipo de dado da variavel
		fmt.Printf("%T \n", name)

		//Print out booleans
		fmt.Printf("%t \n", win)

		//Print out integers
		fmt.Printf("%d \n", x)

		//Print out binarios
		fmt.Printf("%b \n", 25)
	*/
}
