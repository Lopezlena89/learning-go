package main

import "fmt"

// const(
// 	i = 100
// 	pi = 3.1415
// 	prefix = "Go_"
// )

// var(
// 	i int
// 	pi float32
// 	prefix string
// )

func main() {

	// _, d := 42, 43
	// fmt.Println(d)

	// s := "hello"
	// c := []byte(s) // convertimos una string a un tipo []byte
	// c[0] = 'c'
	// s2 := string(c) // volvemos a convertirlo a un tipo string
	// fmt.Printf("%s\n", s2)

	// s := "hello,"
	// m := " world"
	// a := s + m
	// fmt.Printf("%s\n", a)

	// s := "hello"
	// s = "c" + s[1:] // no puede cambiar los valores de una cadena por su índices pero puedes obtener sus valores.
	// fmt.Printf("%s\n", s)

	// err := errors.New("emit macho dwarf: elf header corrupted")
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// a := [...]int{1, 2, 3, 4, 5}
	// fmt.Printf("%d\n", a)

	// arregloSimple := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	// fmt.Println(arregloSimple)
	// definimos un segmento con 10 elementos donde el tipo es byte
	// var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	// definimos dos segmento de tipo []byte
	// var a, b []byte

	// a apunta  a los elementos del tercero al quinto elemento en el array ar.
	// a = ar[2:5]
	// ahora tiene los elementos ar[2],ar[3] y ar[4]

	// b es otro slice del array ar
	// b = ar[3:5]
	// ahora b tiene los elementos ar[3] y ar[4]

	// definimos un arreglo
	// var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// definimos dos segmentos
	// var aSlice, bSlice []byte

	// Algunas operaciones útiles
	// aSlice = array[:3] // es igual a aSlice = array[0:3] aSlice tiene los elementos a,b,c
	// aSlice = array[5:] // es igual a aSlice = array[5:10] aSlice tiene los elementos f,g,h,i,j
	// aSlice = array[:]  // es igual a aSlice = array[0:10] aSlice tiene todos los elementos

	// segmento desde segmento
	// aSlice = array[3:7]  // aSlice tiene los elementos d,e,f,g，len=4，cap=7
	// bSlice = aSlice[1:3] // bSlice contiene aSlice[1], aSlice[2], entonces este tendrá los elementos e,f
	// bSlice = aSlice[:3]  // bSlice contiene aSlice[0], aSlice[1], aSlice[2], entonces este tiene d,e,f
	// bSlice = aSlice[0:5] // bSlice se puede expandir, ahora bSlice contiene  d,e,f,g,h
	// bSlice = aSlice[:]   // bSlice tiene los mismos elementos que aSlice, que son d,e,f,g
	// usamos string tipo de llave (key), int como el tipo de valor, y debemos usar `make` para inicializarlo.
	// var numbers map[string] int
	// otra forma de definir un mapa
	// numbers := make(map[string]int)
	// numbers["one"] = 1 // asignamos el valor para la clave
	// numbers["ten"] = 10
	// numbers["three"] = 3

	// fmt.Println("El tercer número es: ", numbers["three"]) // tomamos el valor
	// Esto imprime: El tercer número es: 3
	// Inicialice un mapa
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	// map nos devuelve dos valores. El segundo valor ，ok es false， si la clave no
	//existe true de otra forma.
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# se encuentra en el map y su ranking es ", csharpRating)
	} else {
		fmt.Println("No tenemos un ranking asociado con C# en este map")
	}

	// delete(rating, "C")  // borramos el elemento con la clave "c"
}
