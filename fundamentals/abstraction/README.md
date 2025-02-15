## **Abstração em Programação**

A **abstração** é um dos pilares da Programação Orientada a Objetos (POO) e um conceito fundamental em programação. Ela
consiste em esconder os detalhes de implementação de um objeto, expondo apenas as funcionalidades essenciais ao usuário.
Isso permite que o programador se concentre no "o quê" (o que o objeto faz) em vez do "como" (como ele faz).

### **O que é Abstração?**

A abstração é o processo de identificar características essenciais de um objeto e ignorar os detalhes irrelevantes. Em
outras palavras, ela permite que você modele objetos do mundo real em termos de suas funcionalidades, sem se preocupar
com a complexidade interna.

### **Quando usar?**

- Quando você quer simplificar a interação com objetos complexos.
- Quando você quer esconder detalhes de implementação para reduzir a complexidade.
- Quando você quer criar uma interface clara e fácil de usar para um sistema ou componente.

### **Exemplo de Abstração em Go**

Go não é uma linguagem puramente orientada a objetos, mas suporta abstração por meio de interfaces e structs. Vamos
criar um exemplo simples de abstração usando uma interface para representar um veículo.

```go
package main

import (
	"fmt"
)

// Definição da interface (abstração)
type Veiculo interface {
	Ligar()
	Acelerar()
	Frear()
}

// Implementação concreta de um carro
type Carro struct {
	modelo string
}

func (c *Carro) Ligar() {
	fmt.Printf("Carro %s ligado.\n", c.modelo)
}

func (c *Carro) Acelerar() {
	fmt.Printf("Carro %s acelerando.\n", c.modelo)
}

func (c *Carro) Frear() {
	fmt.Printf("Carro %s freando.\n", c.modelo)
}

// Implementação concreta de uma moto
type Moto struct {
	modelo string
}

func (m *Moto) Ligar() {
	fmt.Printf("Moto %s ligada.\n", m.modelo)
}

func (m *Moto) Acelerar() {
	fmt.Printf("Moto %s acelerando.\n", m.modelo)
}

func (m *Moto) Frear() {
	fmt.Printf("Moto %s freando.\n", m.modelo)
}

// Função que usa a abstração (interface)
func testarVeiculo(v Veiculo) {
	v.Ligar()
	v.Acelerar()
	v.Frear()
	fmt.Println("-------------------")
}

func main() {
	carro := &Carro{modelo: "Fusca"}
	moto := &Moto{modelo: "Honda CG 150"}

	// Usando a abstração
	testarVeiculo(carro)
	testarVeiculo(moto)
}
```

#### **Saída do Código**

```
Carro Fusca ligado.
Carro Fusca acelerando.
Carro Fusca freando.
-------------------
Moto Honda CG 150 ligada.
Moto Honda CG 150 acelerando.
Moto Honda CG 150 freando.
-------------------
```

### **Prós da Abstração**

1. **Redução da Complexidade**: Esconde detalhes de implementação, permitindo que o programador se concentre na
   funcionalidade.
2. **Reutilização de Código**: Interfaces e classes abstratas podem ser reutilizadas em diferentes contextos.
3. **Facilita a Manutenção**: Mudanças na implementação interna não afetam o código que usa a abstração.
4. **Promove o Desacoplamento**: A abstração permite que diferentes partes do sistema interajam sem conhecer os detalhes
   internos umas das outras.

### **Contras da Abstração**

1. **Curva de Aprendizado**: Pode ser difícil para iniciantes entender o conceito de abstração e como aplicá-lo
   corretamente.
2. **Overhead**: Em alguns casos, a abstração pode adicionar uma camada extra de complexidade desnecessária.
3. **Dificuldade de Debug**: Como os detalhes de implementação estão ocultos, pode ser mais difícil depurar problemas.

### **Conclusão**

A abstração é uma ferramenta poderosa para simplificar sistemas complexos, promovendo a clareza e a organização do
código. Ela permite que os desenvolvedores se concentrem no que um objeto faz, em vez de como ele faz, facilitando a
criação de sistemas modulares e de fácil manutenção. No entanto, é importante usá-la com cuidado para evitar a adição de
complexidade desnecessária.

### **Links de Referência**

- [Abstração em Programação Orientada a Objetos](https://pt.wikipedia.org/wiki/Abstra%C3%A7%C3%A3o_(programa%C3%A7%C3%A3o))
- [Go by Example: Interfaces](https://gobyexample.com/interfaces)
- [Refactoring Guru - Abstração](https://refactoring.guru/pt-br/design-patterns/abstraction)

---

