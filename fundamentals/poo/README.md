# Programação Orientada a Objetos (POO)

## O que é Programação Orientada a Objetos?

A Programação Orientada a Objetos (POO) é um paradigma de programação que organiza o código em "objetos", que encapsulam
dados e comportamento. Por meio de POO, você pode modelar problemas do mundo real, criando estruturas reutilizáveis e
organizadas. POO utiliza conceitos principais como **classes**, **objetos**, **herança**, **polimorfismo**, *
*encapsulamento** e **abstração**.

Embora o Go não suporte POO no sentido clássico (não há suporte para classes), ele fornece recursos que permitem
programar com orientações semelhantes ao paradigma POO, como **structs**, **métodos** e **interfaces**.

---

## Conceitos Fundamentais

1. **Classes e Objetos**:
    - Em linguagens como Java ou C++, classes são os modelos que definem objetos.
    - No Go, usamos `structs` (estruturas) para representar esses modelos.

2. **Encapsulamento**:
    - Refere-se ao ato de esconder os detalhes internos de um objeto e expor apenas os necessários.
    - Em Go, isso é conseguido controlando a visibilidade por meio de maiúsculas (exportado) ou minúsculas (não
      exportado) no início do nome de variáveis e métodos.

3. **Herança**:
    - É a possibilidade de criar novas estruturas a partir de outras. Em Go, não existe herança direta, mas a
      composição (inserir uma `struct` em outra) é incentivada.

4. **Polimorfismo**:
    - Polimorfismo permite que diferentes tipos sejam tratados de forma uniforme.
    - No Go, isso é alcançado por meio de **interfaces**.

5. **Abstração**:
    - Trata-se de focar no essencial e esconder os detalhes.
    - No Go, interfaces ajudam na criação de abstrações.

---

## Exemplo em Go

Vamos ilustrar um exemplo de POO em Go, utilizando os principais conceitos:

```go
package main

import "fmt"

// Definindo uma abstração com interface
type Animal interface {
	Falar() string
}

// Structs representando diferentes "objetos"
type Cachorro struct{}
type Gato struct{}

// Implementação dos métodos da interface Animal
func (c Cachorro) Falar() string {
	return "Au Au"
}

func (g Gato) Falar() string {
	return "Miau"
}

// Struct com composição para modelar um Dono
type Dono struct {
	Nome   string
	Animal Animal // Composição: Dono tem um Animal
}

func main() {
	// Criando instâncias dos objetos
	cachorro := Cachorro{}
	gato := Gato{}

	// Criando donos com seus respectivos animais
	dono1 := Dono{Nome: "João", Animal: cachorro}
	dono2 := Dono{Nome: "Maria", Animal: gato}

	// Utilizando polimorfismo para acessar o método Falar
	fmt.Printf("%s tem um cachorro que faz: %s\n", dono1.Nome, dono1.Animal.Falar())
	fmt.Printf("%s tem um gato que faz: %s\n", dono2.Nome, dono2.Animal.Falar())
}
```

---

## Prós e Contras da POO em Go

### Prós

1. **Modularidade**:
    - O Go permite organizar o código em pacotes e structs, facilitando o reúso e manutenção.
2. **Polimorfismo através de interfaces**:
    - O uso de interfaces oferece grande flexibilidade sem comprometer desempenho.
3. **Código limpo e legível**:
    - O Go favorece composições simples em vez de hierarquias muito complexas.
4. **Altamente performático**:
    - O Go é uma linguagem leve e eficiente, mesmo com usos que simulam POO.

### Contras

1. **Ausência de herança clássica**:
    - Não há suporte para herança direta, o que pode ser um problema para quem vem de outras linguagens.
2. **Maior foco em composição**:
    - Embora eficaz, a composição pode levar a estruturas menos intuitivas para iniciantes em comparação com a herança.
3. **Falta de alguns recursos de POO tradicionais**:
    - Não há suporte para sobrecarga de métodos ou construtores.

---

## Conclusão

A Programação Orientada a Objetos em Go é implementada com um enfoque próprio que prioriza simplicidade e eficiência.
Embora Go não forneça diretamente recursos como herança clássica, suas ferramentas, como interfaces e composição,
tornam-no altamente poderoso e adequado para sistemas modernos. Com a prática e o conhecimento das boas práticas da
linguagem, é possível utilizar conceitos de POO de forma prática e eficiente no ecossistema Go.

---

## Links para Referências

1. [Documentação oficial do Go](https://go.dev/doc/)
2. [Tour do Go](https://go.dev/tour/welcome/1)
3. [Interfaces no Go](https://go.dev/doc/effective_go#interfaces)
