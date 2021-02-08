package greet

//Primera letra de lo que sea es mayuscula se exporta
//Se no es mayuscula no se exporta

//Variables a nivel de paquetes
//Se pueden usar en todos los archivos dentro del paquete
//No se puede usar la forma corta de declaración :=
var emoji = "hola"

//GreetEnglish retorna saludo en ingles
func English() string {
	return "Hi " + emoji
}

//Italian retorna saludo en italiano
func Italian() string {
	return "Ciao " + emoji
}
