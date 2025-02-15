# Padrão de Projeto Facade

O **Design Pattern Facade** (Fachada, em português) é um padrão estrutural na programação orientada a objetos. Ele
fornece uma interface simplificada para um conjunto maior de interfaces em um subsistema, facilitando a interação do
cliente com a complexidade do sistema, encapsulando e abstraindo os detalhes internos.

## Objetivo

- Esconder a complexidade de um sistema, expondo apenas uma interface de nível mais alto para que os clientes possam
  usá-lo de maneira simples e descomplicada.

---

## Motivação

Imagine que você está desenvolvendo um sistema bancário e precisa lidar com vários processos internos, como validação de
contas, verificação de saldo e registro de transações. Cada uma dessas funcionalidades pode ser fornecida por
subsistemas diferentes. O padrão Facade, nesse caso, pode ser usado para criar uma interface unificada que permita
realizar todas essas operações sem que o cliente precise entender ou interagir diretamente com o funcionamento interno
desses subsistemas.

---

## Como Funciona?

1. **Facade** atua como um mediador.
2. Ele recebe as solicitações do cliente e delega as tarefas necessárias aos subsistemas apropriados.
3. O cliente não precisa se preocupar com os detalhes internos dos subsistemas.

---

## Exemplo em Go

Abaixo está um exemplo detalhado em Go que utiliza o padrão Facade:

```go
package main

import "fmt"

// Sistema de autenticação
type AuthSystem struct{}

func (a *AuthSystem) Authenticate(user, password string) bool {
	fmt.Println("Autenticando usuário...")
	return user == "admin" && password == "1234"
}

// Sistema de conta bancária
type AccountSystem struct {
	balance float64
}

func (a *AccountSystem) CheckBalance() float64 {
	fmt.Println("Consultando saldo...")
	return a.balance
}

func (a *AccountSystem) Withdraw(amount float64) bool {
	if a.balance >= amount {
		a.balance -= amount
		fmt.Println("Saque realizado com sucesso.")
		return true
	}
	fmt.Println("Saldo insuficiente.")
	return false
}

// Sistema de registro de transações
type TransactionSystem struct{}

func (t *TransactionSystem) RecordTransaction(amount float64) {
	fmt.Printf("Registrando transação de R$%.2f...\n", amount)
}

// Facade
type BankFacade struct {
	auth         *AuthSystem
	account      *AccountSystem
	transactions *TransactionSystem
}

func NewBankFacade(initialBalance float64) *BankFacade {
	return &BankFacade{
		auth:         &AuthSystem{},
		account:      &AccountSystem{balance: initialBalance},
		transactions: &TransactionSystem{},
	}
}

func (b *BankFacade) WithdrawMoney(user, password string, amount float64) {
	if !b.auth.Authenticate(user, password) {
		fmt.Println("Falha na autenticação! Verifique suas credenciais.")
		return
	}

	if b.account.Withdraw(amount) {
		b.transactions.RecordTransaction(amount)
		fmt.Printf("Saque de R$%.2f concluído com sucesso!\n", amount)
	} else {
		fmt.Println("Não foi possível realizar o saque.")
	}
}

func main() {
	facade := NewBankFacade(1000.0)

	// Tentativa de saque
	facade.WithdrawMoney("admin", "1234", 200.0)
}
```

### Explicação do Código

1. **Subsistemas Isolados**:
    - `AuthSystem`: Gerencia autenticação.
    - `AccountSystem`: Controla saldo e saques.
    - `TransactionSystem`: Registra transações.

2. **Fachada**:
    - `BankFacade`: Encapsula os subsistemas, fornecendo uma interface simplificada para a operação de saque, escondendo
      os detalhes do cliente.

3. **Cliente**:
    - O cliente apenas interage com o `BankFacade`, sem precisar conhecer os detalhes internos dos subsistemas.

---

## Vantagens

### **Prós**

- **Simplicidade**: Reduz a complexidade de lidar diretamente com múltiplos subsistemas.
- **Desacoplamento**: Permite isolar os subsistemas internos, protegendo o cliente de mudanças internas.
- **Legibilidade**: Facilita o entendimento e a manutenção do código.

---

## Desvantagens

### **Contras**

- **Sobrecarga de Facade**: A fachada pode se tornar um "Deus Objeto" (objeto com muitas responsabilidades) se não for
  bem projetada.
- **Flexibilidade Reduzida**: Pode sacrificar flexibilidade, pois esconde detalhes que poderiam ser úteis em certos
  casos.

---

## Conclusão

O padrão Facade é amplamente utilizado para simplificar a interação entre o cliente e sistemas complexos. Ele é útil,
especialmente quando o objetivo é aumentar a legibilidade, ocultar a complexidade interna e prover uma interface coesa
para um cliente. No entanto, é essencial evitar que a fachada assuma muitas responsabilidades, tornando-se difícil de
manter.

---

## Referências

1. [Refactoring Guru - Facade Design Pattern](https://refactoring.guru/design-patterns/facade)
2. [GoF - Design Patterns: Elements of Reusable Object-Oriented Software](https://en.wikipedia.org/wiki/Design_Patterns)
3. [Towards Dev - A Deep Dive Into the Facade Pattern](https://towardsdev.com)
