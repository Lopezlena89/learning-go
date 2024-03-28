Pasos para iniciar en el mundo de Go

    1.- 
    --Hello world

    package main
    import "fmt"
    func main() {
        fmt.Printf("hello, world\n")
    }

    -- Ir a la carpeta y rescribir run <Nombre del archivo> para compilar

    2.-
    --Tipos de datos

    Al declarar un tipo de  variable solo debe contener valores de ese mismo tipo de variable
    
    package main
    import "fmt"
    var z int = 21 //No es recomendable declarar valores fuera de una funcion y utilizamos la  variable
                    var cuando queremos inicializar por default
    func main() {
	fmt.Println(z)
    }
    3.-

    package main
    import (
        "fmt"
    )
    // Declarando la variable con el tipo de identificador int
    var z int //si no lo declaras el valor es cero por default
    func main() {
        z = 21  //Es mejor declararla dentro de la funcion
        fmt.Println(z)
    }







