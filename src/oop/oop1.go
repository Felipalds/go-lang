/**
 * Exemplo de orientação a objetos em Go
 */
 
package main

import (
   "fmt"
   "math"
)

/***********************************
 * Define um tipo de dados ponto2D *
 ***********************************/ 
type point2D struct {
   x, y int8
}

func (p point2D) Getx() int8 {
	return p.x
}
 
func (p *point2D) Setx(x int8) {
	p.x = x
}

func (p *point2D) Translatexy(dx, dy int8) {
   p.x += dx
   p.y += dy
}

func (p1 point2D) Dist2D(p2 point2D) (dist int8) {
   dist = int8(math.Sqrt(math.Pow(float64(p2.x - p1.x), 2.0) + math.Pow(float64(p2.y - p1.y), 2.0)))
   return dist
}

/*******************************************************
 * Função Clear()                                      *
 * Limpa a tela com comandos de escape do terminal     *
 *******************************************************/ 
func Clear (){
   fmt.Print("\033[H\033[2J") // escape codes para limpar a tela (Unix)
}

/*****************
 * Função main() *
 *****************/
func main() {
   var (
      x, y int8
   )
   
   Clear()
   fmt.Println("### Exemplo de classe em Go ###")
   
   // leitura dos dados para o objeto p1
   fmt.Println("\t1 Lendo dados para a estrutura p1:")
   fmt.Print("Valor de x: ")
   fmt.Scan(&x)
   
   fmt.Print("Valor de y: ")
   fmt.Scan(&y)
   
   // instancia o objeto p1
   p1 := point2D{x, y}
   
   // escrevendo os dados armazenados em P1
   fmt.Println("\t2 Escrevendo dados armazenados na estrutura p1")
   fmt.Println("p1: ", p1)
   fmt.Println()
   
   // leitura dos dados para o objeto p2
   fmt.Println("\t3 Lendo dados para a estrutura p2:")
   fmt.Print("Valor de x: ")
   fmt.Scan(&x)
   
   fmt.Print("Valor de y: ")
   fmt.Scan(&y)
   
   // instanciando objeto p2
   p2 := point2D{x, y}
   
   // escrevendo os dados armazenados em P2
   
   fmt.Println("\t4 Escrevendo dados armazenados na estrutura p2")
   fmt.Println("p2: ", p2)
   fmt.Println()
   
   // transladando p1
   fmt.Println("\t5. Transladando p1 por (1, 2):")
   p1.Translatexy(1, 2)
   fmt.Println("p1: ", p1)
   fmt.Println()
   
   // escrevendo os dados armazenados em P2
   fmt.Println("\t6 Escrevendo dados armazenados na estrutura")
   fmt.Println("p2: ", p2)
   fmt.Println()
   
   // Calcula a distância entre p1 e p2
   fmt.Println("\t7 Calculando a distância entre p1 e p2:")
   dist := p1.Dist2D(p2)
   fmt.Println("Distância entre p1 e p2 = ", dist)
   
}

