# Padrão de Projeto: Decorator

## Introdução

O padrão de projeto **Decorator** é um padrão estrutural que permite adicionar funcionalidades a objetos de forma
dinâmica, sem modificar o código da classe original. Ele oferece uma alternativa flexível à herança para estender o
comportamento de classes. O Decorator envolve o objeto original com um "envelope", chamado de decorador, que adiciona o
comportamento desejado.

### Quando usar?

- Quando precisamos adicionar responsabilidades ou funcionalidades a um objeto de forma dinâmica.
- Quando queremos evitar uma explosão de subclasses para cada combinação possível de funcionalidade (um problema comum
  com abordagem via herança).
- Quando a modificação do código existente não é viável ou desejada.

---

## Estrutura

1. **Componente (Interface ou Classe Abstrata):** Define os métodos que serão implementados pelos objetos concretos e
   pelos Decorators.
2. **Componente Concreto:** Implementa a interface Componente, representando o comportamento original do objeto.
3. **Decorator (Classe Base):** Mantém uma referência ao Componente e implementa a interface, delegando chamadas para o
   Componente.
4. **Decorators Concretos:** Extendem a classe Decorator para adicionar funcionalidades específicas ao objeto.

---

## Exemplo em Go

Abaixo, um exemplo detalhado implementando o padrão Decorator em Go. Suponha que temos uma interface básica chamada
`Notifier`, e queremos decorá-la com novos comportamentos.

### Código

```go
package main

import "fmt"

// Componente: Interface que define o comportamento básico
type Notifier interface {
	Send(message string)
}

// Componente Concreto: Implementação básica do Componente
type EmailNotifier struct{}

// Método do Componente Concreto
func (e *EmailNotifier) Send(message string) {
	fmt.Println("Enviando e-mail com mensagem:", message)
}

// Decorator: Classe base para os decoradores
type NotifierDecorator struct {
	Notifier Notifier
}

// Método básico que apenas chama o método do Componente decorado
func (d *NotifierDecorator) Send(message string) {
	d.Notifier.Send(message)
}

// Decorator Concreto 1: Adiciona envio de SMS
type SMSNotifier struct {
	NotifierDecorator
}

// Sobrescreve o método Send para adicionar o envio de SMS
func (s *SMSNotifier) Send(message string) {
	// Primeiro, delega o envio inicial
	s.Notifier.Send(message)
	// Adiciona a funcionalidade extra
	fmt.Println("Enviando SMS com mensagem:", message)
}

// Decorator Concreto 2: Adiciona envio de notificações Push
type PushNotifier struct {
	NotifierDecorator
}

// Sobrescreve o método Send para adicionar envio de notificações Push
func (p *PushNotifier) Send(message string) {
	// Primeiro, delega o envio inicial
	p.Notifier.Send(message)
	// Adiciona a funcionalidade extra
	fmt.Println("Enviando notificação push com mensagem:", message)
}

// Função principal para testar o padrão
func main() {
	// Criando um componente básico (EmailNotifier)
	emailNotifier := &EmailNotifier{}

	// Decorando com envio de SMS
	smsNotifier := &SMSNotifier{
		NotifierDecorator: NotifierDecorator{Notifier: emailNotifier},
	}

	// Decorando com envio de notificação Push
	pushNotifier := &PushNotifier{
		NotifierDecorator: NotifierDecorator{Notifier: smsNotifier},
	}

	// Envio final (com email, SMS e push)
	pushNotifier.Send("Olá! Este é um teste de notificação.")
}
```

### Saída do Programa

```plaintext
Enviando e-mail com mensagem: Olá! Este é um teste de notificação.
Enviando SMS com mensagem: Olá! Este é um teste de notificação.
Enviando notificação push com mensagem: Olá! Este é um teste de notificação.
```

---

## Prós e Contras

### Prós

1. **Flexibilidade:** Funcionalidades podem ser adicionadas ou removidas sem alterar o código existente.
2. **Reutilização de Código:** Decorators podem ser combinados e reutilizados em diferentes partes do sistema.
3. **Reduz a Explosão de Herança:** Evita a criação de múltiplas subclasses para combinar diferentes tipos de
   comportamentos.

### Contras

1. **Maior Complexidade:** O código pode se tornar mais difícil de entender devido ao encadeamento de muitos
   decoradores.
2. **Problemas de Depuração:** Traçar o fluxo de execução pode ser complicado ao usar vários decoradores, especialmente
   em sistemas complexos.

---

## Conclusão

O padrão de projeto **Decorator** é um recurso poderoso para adicionar funcionalidades dinâmicas a objetos em tempo de
execução, oferecendo uma solução mais flexível em comparação com herança. Ele é especialmente útil quando o
comportamento precisa ser estendido sem alterar diretamente o código do objeto original. Contudo, deve ser usado com
cuidado para não introduzir complexidade desnecessária ao sistema.

---

## Links de Referência

1. [Documentação Oficial de Go](https://golang.org/doc/)
2. [Padrões de Projeto na Wikipédia](https://pt.wikipedia.org/wiki/Decorator)
3. [Refactoring Guru - Decorator (em inglês)](https://refactoring.guru/design-patterns/decorator)

---
