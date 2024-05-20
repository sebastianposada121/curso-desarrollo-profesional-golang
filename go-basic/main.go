/*
Agrupar funcionalidades relacionadas
Encapsulamiento, Reusabilidad, Organización, Gestión de dependencias
*/
package main

import (
	"basic/custom"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

// constantes
const greeting = "Hello"

func main() {
	declarateVarConst()
	typeOfData()
	reflectTypeOf()
	pointers()
	conditionals()
	arraysAndSlice()
	cicleAndIterators()
	maps()
	functions(3, 4)
	structures()
	interfaces()
	packages()
}

// ------------------------------------ Declaracion de variables y constantes---------------------------------//
func declarateVarConst() {
	fmt.Println("--------- Declaracion de variables y constantes ---------")
	// inferencia
	var name string = "Vanesa"
	// rapida o corta
	secondName := "Camila"
	fmt.Println(greeting, name, secondName)
	fmt.Printf("My name is %s %s \n", name, secondName)
}

// --------------------------------------- Tipos de datos --------------------------------//
func typeOfData() {
	fmt.Println("--------- Tipos de datos ---------")
	var smallText string = "hello"
	var largeText string = `Aislamiento de dependencias: Los módulos Go proporcionan un aislamiento claro de las dependencias de un proyecto, lo que significa que cada proyecto puede tener sus propias dependencias y versiones específicas sin interferir con otros proyectos en el mismo sistema.`
	var status bool = true
	var float_32 float32 = 32.434
	var float_64 float64 = 64.34
	var year int = 2023
	var int_8 int8 = 123   // -128 to 127
	var int_16 int16 = 123 // -2^15 to 2^15
	var uint_8 uint8 = 43  // 0 to 255
	fmt.Println(smallText, "\n", largeText, "\n", status, float_32, float_64, year, int_8, int_16, uint_8)
}

/*
reflect TypeOf
*/
func reflectTypeOf() {
	text := ""
	fmt.Println("------- Declaracion de variables y constantes -------")
	fmt.Println(reflect.TypeOf(text))

}

/*
punteros: obtener el valor binario del espacio en memoria del valor
*/
func pointers() {
	fmt.Println("------- punteros &var --------")
	color := "red"
	fmt.Println(color, &color)
}

/*
condicionales if else, switch
*/
func conditionals() {

	fmt.Println("------- condicionales if else, switch --------")
	// x == y igual
	// x != y diferente
	// x < y menor
	// x <= y menor o igual
	// x > y mayor
	// x => y mayor o igual
	// se pueden declarar vaiables dentro de un if
	age := 24
	if maxAge := 18; age >= maxAge {
		fmt.Println("Eres un anciano...", maxAge)
	} else {
		fmt.Println("Eres un cachorro...", maxAge)
	}

	rol := "developer"

	switch rol {
	case "developer":
		fmt.Println("Eres de la gerencia de ti")
	case "administrator":
		fmt.Println("Eres de finanzas")
	default:
		fmt.Println("No haces parte de la organizacion")
	}
}

/*
Arreglos y slice
Arreglos: decirle el largo o el capacidad de valores a contener
Slice: No necesitamos decirle la cantidad de valores
*/
func arraysAndSlice() {
	// array
	fmt.Println("------- Arreglos y slice --------")
	var countries [2]string
	countries[0] = "Brazil"
	countries[1] = "Colombia"
	fmt.Println(countries, "total", len(countries))
	// slice
	var countriesSlice = make([]string, 4)
	countriesSlice[0] = "Colombia"
	countriesSlice[1] = "Mexico"
	countriesSlice[3] = "Peru"
	// agregar elemento
	countriesSlice = append(countriesSlice, "Dinamarca")
	// eliminar elemento
	countriesSlice = append(countriesSlice[:4], countriesSlice[4+1:]...)
	fmt.Println(countriesSlice, "total", len(countriesSlice))
}

// -------------------------------- Ciclos e iteraciones ----------------------- //

func cicleAndIterators() {
	fmt.Println("------- Ciclos e iteraciones --------")
	// declaracion normal
	i := 10
	for i < 1 {
		fmt.Println(i)
		i++
	}
	// declaracion rapida
	table := 5
	fmt.Println("tabla #", table)
	for j := 0; j < 10; j++ {
		// salir ciclo
		if table == 0 {
			continue
		}
		fmt.Println(table, "x", j, "=", table*j)
	}

}

/*
mapas estructura de datos que proporciona una asociación entre pares clave-val
*/
func maps() {
	foods := make(map[string]string)
	foods["Frijoles"] = "Una belleza"
	fmt.Println(foods, foods["Frijoles"])
	// validar si existe valor en el map
	pets := map[string]string{
		"paco":   "perro",
		"lulu":   "gato",
		"tayson": "perro",
	}
	// validar si existe en el mapa
	loro, exist := pets["loro"]
	fmt.Println(pets)
	fmt.Println(loro, exist)
	// elimina elemento
	delete(pets, "paco")
	fmt.Println(pets)
	for pname, ptype := range pets {
		fmt.Println(pname, ptype)
	}
}

//------------------------------------ funciones ---------------------------------//
/*
funciones porcion de codigo que realiza una tarea en especifico
argumentos a y b
*/
func functions(a int, b int) {
	fmt.Println("------- funciones ---------")
	fmt.Println("multiplicación", a*b)
	fmt.Println("suma", a+b)
	fmt.Println("resta", a-b)

	fmt.Println(functionReturn("sebas"))
	name, lastName, age := functionMultipleReturn()
	fmt.Println("name", name, "lastName", lastName, "age", age)
	fmt.Println(anonimousFunctions(19, 34))
	fmt.Println()

	// funciones que retornan una funcion
	Clousure := clousure(2)
	for i := 1; i <= 10; i++ {
		fmt.Printf("2 x %v = %v \n", i, Clousure())
	}

	// crear gorutinas
	myChanel := make(chan string)
	go func() {
		myChanel <- gorutinas()
	}()
	fmt.Println(<-myChanel, "Continuar con la ejecucion")

	// recursion
	fmt.Println("------recursion-----")
	recursion(0)

	// panic y defer
	panicAndDefer(false)
}

// retornar un solo valor
func functionReturn(name string) string {
	fmt.Println("-----------Funcion  que retorna un solo valor---------")
	return "Nombre:" + name
}

// retornar muchos valores
func functionMultipleReturn() (string, string, int) {
	fmt.Println("----------Funcion  que retorna varios valores------------")
	return "Juan", "Perez", 43
}

// funcion anonima
var anonimousFunctions = func(a int, b int) int {
	fmt.Println("------funcion de anonima-----")
	return a + b
}

// funcion que devuelve una funcion
func clousure(value int) func() int {
	fmt.Println("------clousure-----")
	number := value
	secuence := 0
	return func() int {
		secuence++
		return number * secuence
	}
}

/*
	gorutinas o canales
	ejecutar varios procesos
*/

func gorutinas() string {
	fmt.Println("------gorutinas-----")
	// pausar procesos durante 10segundas
	time.Sleep(time.Second * 5)
	//  crear canalaes
	return "Hola nueva gorutina"
}

/*
recursividad
funcion que se llama asi misma
*/
func recursion(value int) {
	data := value + 1
	fmt.Println(data)
	if data < 10 {
		recursion(data)
	}
}

/*
	defer y panic
	panic mostrar mensaje de error y terminar el proceso de compilacion
*/

func panicAndDefer(error bool) {
	defer fmt.Println("este es el mensaje final de la ejecucion")
	fmt.Println("----panic y defer----")
	if error {
		panic("esto esun panic")
	}
}

//------------------------------------ estructuras---------------------------------//
/*
Empieza el nombre en mayuscula para que sea publica la estructura
*/

type People struct {
	Id       int
	Name     string
	Email    string
	Age      int
	Category Category
}

type Category struct {
	Id     int
	Name   string
	Salary int
}

func structures() {
	user := People{
		Id:    1,
		Name:  "Kevin",
		Email: "kevin@pirani.co",
		Age:   34,
		Category: Category{
			Id:     1,
			Name:   "barber",
			Salary: 323232,
		},
	}

	category := Category{
		Id:     1,
		Name:   "dev",
		Salary: 323232,
	}
	user2 := new(People)
	user2.Id = 1
	user2.Name = "sebas"
	user2.Category = category
	fmt.Println(user, reflect.TypeOf(user))
	fmt.Println(user2, reflect.TypeOf(user2))
}

// ----------------------- interfaces ------------------------//

func interfaces() {
	myStruct := ExampleStructure{}
	fmt.Println(myStruct.myInterfaceStruct())
}

type ExampleInterface interface {
	sum() int
	calculate(n1 int, n2 int) int
}

type ExampleStructure struct {
}

func (*ExampleStructure) myInterfaceStruct() string {
	return "implementando funcion"
}

//------------------------------------ Paquetes---------------------------------//

/*
time: manejo del tiempo
*/
func packages() {
	timeP()
	stringsP()
	mathRand()
	osP()
	logP()
	custom.Greeting()
}

func timeP() {
	fmt.Println("-------------------- time --------------------------")
	fmt.Println(time.Now())

	date := time.Now()

	fmt.Println("Año", date.Year(), "Mes", date.Month(), "dia", date.Day())

	fmt.Println(formatDate(date))

	fmt.Println("Mas 20 días:")
	date1 := date.Add(time.Hour * 24 * 20)
	fmt.Println(formatDate(date1))
}
func formatDate(date time.Time) string {
	value := fmt.Sprintf("%v %v de %v de %v a las %v:%v:%v", date.Weekday(), date.Day(), date.Month(), date.Year(), date.Hour(), date.Minute(), date.Second())
	return value
}

/*
strings: manejo de textos
*/
func stringsP() {
	fmt.Println("-------------------- strings --------------------------")
	text := "Hola esTa es uNa CaDena "
	fmt.Println("minuscula", strings.ToLower(text))
	fmt.Println("mayuscula", strings.ToUpper(text))
	fmt.Println("convertir a array", strings.Split(text, " "))
	fmt.Println("buscar indice de una palabra", strings.Index(text, "hola"))
	fmt.Println("repetir texto", strings.Repeat(text, 10))
	fmt.Println("remplazar texto", strings.Replace(text, "CaDena", "texto", -1))
	fmt.Println("mostrar palabra del 0 al 4", string(text[0:4]))
}

/*
mathRand: manejo de aleatorios
*/

func mathRand() {
	fmt.Println("--------------- math rand ------------------------")
	random := rand.Intn(101)
	fmt.Println(random)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	random2 := rand.Intn(10-0) + 0
	fmt.Println(random2)
}

/*
OS argumentos por consola
*/

func osP() {
	fmt.Println("--------------- os ------------------------")
	name := flag.String("name", "", "Name:")
	age := flag.Int("age", 18, "Age:")
	flag.Parse()

	fmt.Println(*name, *age)
}

/*
Log mostrar mensajes
*/

func logP() {
	fmt.Println("--------------- log ------------------------")
	err := errors.New("nuevo error..")
	// detiene la ejecucion
	// log.Fatal("Soy tu mayor error")
	// log.Panic("Soy tu mayor error")
	fmt.Println(err)

}
