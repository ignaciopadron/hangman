package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	ahorcado "github.com/ignaciopadron/ahorcado/internal/game" // Importar el paquete del juego
)

func main() {
	reader := bufio.NewReader(os.Stdin) // lector para entrada estándar

	fmt.Println("¡Bienvenido al juego del Ahorcado!")
	fmt.Println("Elige el modo de juego:")
	fmt.Println("1. Un jugador (la computadora elige la palabra secreta)")
	fmt.Println("2. Dos jugadores (un jugador ingresa la palabra secreta manualmente)")
	fmt.Print("Selecciona 1 o 2: ")

	// Leer la opción de modo
	opcionModo, _ := reader.ReadString('\n')
	opcionModo = strings.TrimSpace(opcionModo)

	var palabraSecreta string
	if opcionModo == "1" {
		// Modo un jugador: la computadora escoge palabra
		palabraSecreta = ahorcado.EscogerPalabraAleatoria()
		if palabraSecreta == "" {
			fmt.Println("No hay palabras disponibles. Saliendo del juego.")
			return
		}
		fmt.Println("\nModo un jugador iniciado. ¡Empieza a adivinar la palabra!")
	} else if opcionModo == "2" {
		// Modo dos jugadores: pedir palabra al jugador 1
		fmt.Print("Jugador 1, ingresa la palabra secreta: ")
		secretInput, _ := reader.ReadString('\n')
		palabraSecreta = strings.TrimSpace(secretInput)
		if palabraSecreta == "" {
			fmt.Println("No ingresaste ninguna palabra. Saliendo del juego.")
			return
		}
		// Opcional: intentar "ocultar" la palabra secreta desplazando la pantalla:
		fmt.Print("\033[H\033[2J") // código ANSI para limpiar pantalla (puede no funcionar en todos los entornos)
		fmt.Println("¡Palabra secreta guardada! Jugador 2, es tu turno de adivinar.")
	} else {
		fmt.Println("Opción no válida. Ejecuta de nuevo el programa e ingresa '1' o '2'.")
		return
	}

	// Inicializar la partida con la palabra secreta obtenida
	const maxIntentos = 6
	partida := ahorcado.NuevaPartida(palabraSecreta, maxIntentos)

	// Bucle de juego: seguir hasta ganar o perder
	for !partida.Ganado() && !partida.Perdido() {
		// Mostrar estado actual
		fmt.Printf("\nPalabra: %s\n", string(partida.Progreso)) // convierte []rune a string para mostrar
		fmt.Printf("Intentos restantes: %d\n", partida.IntentosRestantes)
		if len(partida.LetrasProbadas) > 0 {
			fmt.Printf("Letras probadas: %s\n", string(partida.LetrasProbadas))
		}

		// Pedir al jugador una letra
		fmt.Print("Ingresa una letra: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Procesar el intento
		mensajeError := partida.IntentarLetra(input)
		if mensajeError != "" {
			// Si hubo un problema (letra inválida o repetida), informamos y continuamos sin avanzar turno.
			fmt.Println(mensajeError)
			continue
		}
		// Si no hubo error, el estado se actualizó (intentosRestantes decrementado o letra revelada).
		// El bucle volverá a chequear las condiciones y continuar.
	}

	// Fuera del bucle: o ganó o perdió
	if partida.Ganado() {
		fmt.Printf("\n🎉 ¡Felicidades! Has adivinado la palabra: %s\n", partida.PalabraSecreta)
	} else if partida.Perdido() {
		fmt.Printf("\n☠️ Te has quedado sin intentos. La palabra secreta era: %s\n", partida.PalabraSecreta)
	}
	fmt.Println("Gracias por jugar. ¡Hasta la próxima!")
}
