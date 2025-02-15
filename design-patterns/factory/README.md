# Padrão de Projeto: Factory (Fábrica)

O padrão de projeto Factory (ou Fábrica) é um padrão criacional que encapsula a criação de objetos, delegando a
instância de objetos para subclasses ou métodos específicos de fábrica. Ele é usado para evitar a complexidade da
criação direta de objetos e para promover a reutilização de código.

Este padrão é amplamente adotado quando você não sabe de antemão qual classe exata precisa ser instanciada ou quando a
lógica de criação de objetos for detalhada e repetitiva.

---

## Motivação

O objetivo do padrão Factory é remover a lógica de criação de objetos do código principal, simplificando e centralizando
esse processo em uma fábrica. É útil em cenários onde diferentes objetos possam compartilhar uma interface comum ou
precisam ser configurados de forma complexa antes de serem usados.

---

## Implementação em Go

Aqui está a documentação com um exemplo prático de uso do padrão Factory em Go:

1. **Definição de uma interface comum**:
   Criamos uma interface que todos os tipos de objetos a serem criados seguirão.

2. **Implementação de classes concretas**:
   Criamos diferentes structs que implementam a interface comum.

3. **Fábrica para criar objetos**:
   A fábrica encapsula a lógica de criação de objetos e decide qual tipo de objeto criar com base em alguma entrada.

### Exemplo em Go

```go
package main

import (
	"fmt"
)

// 1. Interface comum para produtos
type Vehicle interface {
	Drive() string
}

// 2. Implementação de produtos concretos
type Car struct{}

func (c Car) Drive() string {
	return "Dirigindo um carro."
}

type Bike struct{}

func (b Bike) Drive() string {
	return "Pilotando uma bicicleta."
}

// 3. Fábrica para criar objetos Vehicle
func VehicleFactory(vehicleType string) (Vehicle, error) {
	switch vehicleType {
	case "car":
		return Car{}, nil
	case "bike":
		return Bike{}, nil
	default:
		return nil, fmt.Errorf("tipo de veículo desconhecido")
	}
}

func main() {
	// Exemplo de uso da fábrica
	vehicle1, err := VehicleFactory("car")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(vehicle1.Drive())

	vehicle2, err := VehicleFactory("bike")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(vehicle2.Drive())
}
```

---

## Prós e Contras

### Prós

1. **Maior desacoplamento**:
    - O cliente não conhece os detalhes da criação do objeto.
2. **Reutilização de código**:
    - Reduz duplicação da lógica de inicialização de objetos.
3. **Facilidade para adição de novos tipos**:
    - Novos produtos podem ser adicionados sem alterar o código do cliente.

### Contras

1. **Complexidade adicional**:
    - Pode adicionar complexidade à medida que o número de produtos aumenta.
2. **Pode ocultar dependências**:
    - A lógica de criação pode se tornar obscura, dificultando a leitura do código.
3. **Menor transparência**:
    - O cliente não sabe qual implementação exata está sendo usado sem investigar.

---

## Conclusão

O padrão Factory é útil quando o processo de criação de objetos é complexo ou repetitivo e precisamos de flexibilidade
para criar diferentes tipos de objetos. Ele funciona especialmente bem em linguagens fortemente tipadas como Go, onde
interfaces podem ser facilmente exploradas para desacoplar comportamentos. No entanto, para projetos menores ou simples,
pode ser excessivo e adicionar um peso desnecessário.

---

## Links de Referência

- [Padrões de Projeto (GoF)](https://pt.wikipedia.org/wiki/Padr%C3%B5es_de_projeto_de_software)
- [Documentação oficial sobre Interfaces em Go](https://go.dev/tour/methods/9)
- [Exemplos de Design Patterns em Go](https://github.com/tmrts/go-patterns)
