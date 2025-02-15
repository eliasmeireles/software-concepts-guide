# **Fundamentos de Programação: Polimorfismo**

## **O que é Polimorfismo?**

**Polimorfismo** é um princípio fundamental da programação orientada a objetos (POO) que permite que objetos de
diferentes classes sejam tratados como objetos de uma classe comum. O termo "polimorfismo" vem do grego, que significa "
muitas formas", e no contexto da POO, isso se refere à capacidade de diferentes tipos de objetos responderem de maneira
diferente à mesma mensagem ou chamada de método.

Em outras palavras, o polimorfismo permite que você tenha um único método ou operação que pode ser utilizado em objetos
de diferentes classes, proporcionando flexibilidade e reutilização de código.

### **Tipos de Polimorfismo**:

1. **Polimorfismo de Sobrecarga (Overloading)**: Quando duas ou mais funções ou métodos têm o mesmo nome, mas diferem no
   número ou tipo de parâmetros.

2. **Polimorfismo de Sobrescrita (Overriding)**: Quando um método da classe base é reescrito por um método na classe
   derivada, mantendo a mesma assinatura.

3. **Polimorfismo Paramétrico**: Usado em linguagens como Go e C++, permite que funções e tipos operem em uma variedade
   de tipos sem precisar saber de antemão quais tipos serão usados.

No Go, como não é uma linguagem puramente orientada a objetos, o polimorfismo ocorre principalmente por meio de *
*interfaces**.

---

## **Polimorfismo em GoLang**

Go é uma linguagem que não utiliza classes no estilo tradicional das linguagens orientadas a objetos, mas permite
polimorfismo por meio de **interfaces**. Em Go, um tipo implementa uma interface simplesmente ao fornecer as
implementações dos métodos definidos pela interface. Não há necessidade de declaração explícita de implementação como em
outras linguagens como Java ou C#.

### **Exemplo de Polimorfismo em Go**

Aqui está um exemplo básico de como o polimorfismo pode ser implementado em Go utilizando interfaces.

#### **Exemplo Prático**:

```go
package main

import "fmt"

// Definindo a interface Animal
type Animal interface {
	Falar() string
}

// Implementação da estrutura Cachorro
type Cachorro struct{}

func (c Cachorro) Falar() string {
	return "Au Au"
}

// Implementação da estrutura Gato
type Gato struct{}

func (g Gato) Falar() string {
	return "Miau"
}

// Função que aceita qualquer tipo que implemente a interface Animal
func OuvirAnimal(a Animal) {
	fmt.Println(a.Falar())
}

func main() {
	// Criando instâncias de Cachorro e Gato
	var a Animal

	a = Cachorro{}
	OuvirAnimal(a)

	a = Gato{}
	OuvirAnimal(a)
}
```

#### **Explicação do Código**:

1. Definimos uma interface `Animal`, que possui um método `Falar()`.
2. Criamos duas estruturas (`Cachorro` e `Gato`), cada uma implementando a interface `Animal` com o método `Falar()`.
3. A função `OuvirAnimal` aceita um parâmetro do tipo `Animal` e chama o método `Falar()`, demonstrando o polimorfismo,
   pois dependendo do tipo de objeto passado (Cachorro ou Gato), o comportamento do método será diferente.
4. No `main`, criamos uma variável do tipo `Animal` e atribuímos a ela tanto o `Cachorro` quanto o `Gato`, permitindo
   que o mesmo código trate diferentes tipos de animais de maneira uniforme.

#### **Saída**:

```bash
Au Au
Miau
```

---

## **Prós e Contras do Polimorfismo**

### **Prós**:

- **Reusabilidade de Código**: O polimorfismo permite que você escreva código que funciona com diferentes tipos de
  objetos, sem precisar duplicar o código para cada tipo específico.
- **Flexibilidade**: O código que usa polimorfismo é mais flexível e pode ser facilmente estendido. Novos tipos de
  objetos podem ser introduzidos sem modificar as funções que operam com a interface.
- **Abstração**: O polimorfismo, combinado com interfaces, permite esconder os detalhes de implementação e trabalhar
  apenas com abstrações, tornando o código mais limpo e fácil de manter.
- **Manutenção**: O uso de polimorfismo pode simplificar a manutenção do código, já que os métodos podem ser atualizados
  sem afetar o restante do sistema, desde que as interfaces permaneçam consistentes.

### **Contras**:

- **Complexidade**: Pode adicionar complexidade ao código, especialmente em sistemas grandes, onde muitos tipos
  implementam as mesmas interfaces. Pode ser difícil rastrear como os diferentes objetos interagem.
- **Desempenho**: Em alguns casos, o uso de polimorfismo pode reduzir o desempenho devido à busca dinâmica de métodos,
  especialmente quando o código precisa fazer muitas chamadas de métodos em objetos diferentes.
- **Dificuldade em Depurar**: O polimorfismo pode tornar o processo de depuração mais difícil, pois o comportamento do
  programa pode ser alterado dependendo do tipo real do objeto passado para um método.
- **Acoplamento Excessivo**: O uso excessivo de interfaces pode resultar em acoplamento excessivo entre as diferentes
  partes do sistema, tornando o código mais difícil de refatorar ou modificar sem impactar outras áreas.

---

## **Conclusão**

O **polimorfismo** é uma característica fundamental da programação orientada a objetos e oferece muitos benefícios, como
a reusabilidade de código, flexibilidade e abstração. Em Go, o polimorfismo é implementado de maneira simples e
eficiente usando interfaces, permitindo que diferentes tipos compartilhem comportamentos comuns sem precisar de herança
complexa.

Apesar das suas vantagens, como qualquer conceito de design, o polimorfismo deve ser usado com cuidado. A introdução de
polimorfismo pode aumentar a complexidade do código e o tempo de depuração, além de potencialmente afetar o desempenho
em sistemas críticos.

No geral, o polimorfismo é uma técnica poderosa que pode tornar seu código mais flexível e extensível, desde que seja
usado de forma balanceada e consciente.

---

## **Links de Referência**

- [Go Interfaces - Documentação oficial](https://golang.org/doc/effective_go.html#interfaces)
- [Polimorfismo - Artigo na Wikipedia](https://pt.wikipedia.org/wiki/Polimorfismo_(computa%C3%A7%C3%A3o))
- [Polimorfismo em Go](https://www.calhoun.io/understanding-polymorphism-in-go/)
- [Como usar Interfaces em Go](https://www.digitalocean.com/community/tutorials/understanding-interfaces-in-go)

---
