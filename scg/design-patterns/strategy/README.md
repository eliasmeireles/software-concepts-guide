# Design Pattern: Strategy

## Introdução

O padrão de projeto **Strategy** é um dos padrões comportamentais do catálogo GoF (Gang of Four). Ele permite que um
algoritmo seja selecionado em tempo de execução. Esse padrão define um conjunto de algoritmos encapsulados, tornando-os
intercambiáveis, promovendo assim a separação de responsabilidades e facilitando extensibilidade do código.

Em resumo, o padrão **Strategy** é usado quando temos diversas formas de realizar uma tarefa (algoritmos diferentes) e
queremos tornar o código flexível para mudar a implementação dinamicamente.

---

## Estrutura

O padrão Strategy é composto pelos seguintes elementos:

1. **Interface (Strategy):** Define um contrato comum para as estratégias concretas (algoritmos).
2. **Implementações Concretas:** Cada algoritmo ou estratégia é implementado em classes distintas que implementam a
   interface definida.
3. **Contexto:** Contém uma referência a um objeto da interface `Strategy` e delega a execução do comportamento às
   implementações concretas dessa interface.

---

## Exemplo em Go

Vamos criar um exemplo prático onde diferentes métodos de cálculo de frete (como `Correios`, `Transportadora` e
`MotoBoy`) podem ser escolhidos dinamicamente usando o padrão Strategy.

```go
package main

import "fmt"

// Strategy define a interface comum para os algoritmos de cálculo de frete
type FreightCalculator interface {
	CalculateFreight(weight float64) float64
}

// Strategy concreta: Cálculo de frete pelos Correios
type Correios struct{}

func (c *Correios) CalculateFreight(weight float64) float64 {
	return weight * 1.5 // Exemplo de cálculo baseado no peso
}

// Strategy concreta: Cálculo de frete por transportadora
type Transportadora struct{}

func (t *Transportadora) CalculateFreight(weight float64) float64 {
	return weight*1.2 + 5.0 // Taxa fixa + proporção com peso
}

// Strategy concreta: Cálculo de frete por Motoboy
type MotoBoy struct{}

func (m *MotoBoy) CalculateFreight(weight float64) float64 {
	return 10.0 // Taxa fixa, independente do peso
}

// Contexto: Usa uma estratégia para calcular o frete
type FreightContext struct {
	strategy FreightCalculator
}

// Método para configurar dinamicamente a estratégia no contexto
func (f *FreightContext) SetStrategy(strategy FreightCalculator) {
	f.strategy = strategy
}

// Método que delega o cálculo para a estratégia configurada
func (f *FreightContext) Calculate(weight float64) float64 {
	if f.strategy == nil {
		panic("Nenhuma estratégia configurada!")
	}
	return f.strategy.CalculateFreight(weight)
}

func main() {
	// Peso do pacote a ser enviado
	weight := 10.0

	// Contexto
	context := &FreightContext{}

	// Usando Correios
	context.SetStrategy(&Correios{})
	fmt.Printf("Frete pelos Correios: R$%.2f\n", context.Calculate(weight))

	// Usando Transportadora
	context.SetStrategy(&Transportadora{})
	fmt.Printf("Frete pela Transportadora: R$%.2f\n", context.Calculate(weight))

	// Usando MotoBoy
	context.SetStrategy(&MotoBoy{})
	fmt.Printf("Frete pelo Motoboy: R$%.2f\n", context.Calculate(weight))
}

```

---

## Explicação do Código

1. **Interface `FreightCalculator`:** Esta é a interface que define o contrato para diferentes estratégias de cálculo de
   frete.
2. **Implementações Concretas:** Classes como `Correios`, `Transportadora` e `MotoBoy` implementam o cálculo de
   diferentes formas, respeitando o contrato da interface.
3. **Contexto `FreightContext`:** Contém a lógica para configurar dinamicamente a estratégia de cálculo e delega a
   execução para a estratégia configurada.

O código é flexível porque podemos adicionar novas formas de cálculo (estratégias) sem modificar o `Context`.

---

## Prós e Contras

### Prós

- **Reduz Complexidade:** Evita grandes blocos de código condicional (`switch` ou `if-else`) para selecionar o
  algoritmo.
- **Extensibilidade:** Adicionar novas estratégias não exige mudanças no código existente, apenas cria-se uma nova
  implementação.
- **Separação de Responsabilidades:** Cada algoritmo é encapsulado em sua própria classe.

### Contras

- **Maior Número de Classes:** Pode aumentar o número de classes no projeto, tornando o código mais extenso.
- **Configuração Explícita:** O cliente precisa saber qual estratégia usar e configurá-la no contexto.
- **Overhead:** Em casos simples, usar o padrão pode ser considerado um "overdesign".

---

## Conclusão

O padrão **Strategy** é ideal para sistemas que precisam lidar com várias implementações de um mesmo comportamento. Ele
promove boa engenharia de software ao manter responsabilidades claras e possibilitar a extensão do sistema sem alterar
códigos existentes. Contudo, deve ser aplicado com cautela para evitar complexidade desnecessária em situações simples.

---

## Referências

- [Design Patterns Elements of Reusable Object-Oriented Software (GoF)](https://en.wikipedia.org/wiki/Design_Patterns)
- [Documentação Oficial do Go](https://go.dev/doc/)
- [Estratégia - Refactoring Guru](https://refactoring.guru/design-patterns/strategy)
