# Introdução à Injeção de Dependência (Dependency Injection) em Go

A **Injeção de Dependência** (Dependency Injection, ou DI) é um padrão de design usado para criar e gerenciar
dependências em software, promovendo o desacoplamento entre os componentes. Em Go, que é bem conhecido por sua
simplicidade, o uso de DI é geralmente feito de forma manual, sem frameworks complexos, usando *interfaces*,
construtores e inicialização explícita.

Neste documento, explicaremos os fundamentos da DI, apresentaremos um exemplo prático em Go, bem como analisaremos seus
prós, contras e conclusões.

---

## O que é Injeção de Dependência?

A DI é um processo no qual os objetos recebem suas dependências em vez de criá-las diretamente. Ela permite que as
dependências sejam fornecidas externamente (via construtores, setters ou injeção direta). Em vez de um objeto criar ou
gerenciar suas próprias dependências, ele as "pede" e as utiliza.

Isso promove:

- **Desacoplamento:** Um código modular que segue princípios como o SOLID.
- **Testabilidade:** Facilita a substituição de dependências por mocks durante os testes.

---

## Formas de Injeção de Dependência

- **Injeção via Construtor**: Dependências fornecidas ao criar um objeto.
- **Injeção via Setter**: Dependências configuradas após a criação do objeto.
- **Injeção via Interface**: Uso explícito de interfaces para delegar responsabilidades a implementações específicas.

---

## Exemplo Prático em Go

### Problema

Imagine que temos um serviço de envio de e-mails. No entanto, queremos que esse componente seja flexível o suficiente
para permitir testes ou substituições futuras sem modificar o código principal.

### Solução: Usando Dependency Injection

```go
package main

import "fmt"

// EmailService define a interface para serviços de envio de e-mails
type EmailService interface {
	SendEmail(to string, message string) error
}

// SMTPService é uma implementação concreta de EmailService
type SMTPService struct{}

func (s *SMTPService) SendEmail(to string, message string) error {
	// Lógica de envio de e-mail via SMTP
	fmt.Printf("Enviando e-mail para %s: %s\n", to, message)
	return nil
}

// MockEmailService é usado para testes
type MockEmailService struct{}

func (m *MockEmailService) SendEmail(to string, message string) error {
	// Simula envio de e-mail para testes
	fmt.Printf("Simulando envio de e-mail para %s: %s\n", to, message)
	return nil
}

// UserNotifier depende de EmailService para notificar os usuários
type UserNotifier struct {
	emailService EmailService
}

// NewUserNotifier é um construtor que injeta a dependência
func NewUserNotifier(es EmailService) *UserNotifier {
	return &UserNotifier{emailService: es}
}

// Notify envia uma mensagem para um usuário
func (n *UserNotifier) Notify(userEmail string) {
	err := n.emailService.SendEmail(userEmail, "Bem-vindo ao serviço!")
	if err != nil {
		fmt.Println("Erro ao enviar o e-mail:", err)
	}
}

func main() {
	// Usando a implementação real
	smtpService := &SMTPService{}
	notifier := NewUserNotifier(smtpService)
	notifier.Notify("usuario@example.com")

	// Usando a implementação mock para testes
	mockService := &MockEmailService{}
	testNotifier := NewUserNotifier(mockService)
	testNotifier.Notify("teste@example.com")
}
```

### Explicação do Código

1. **Interface EmailService**: Define o comportamento esperado de qualquer serviço de envio de e-mails.
2. **Implementações SMTPService e MockEmailService**: Implementam EmailService para diferentes cenários (produção ou
   teste).
3. **UserNotifier**: Depende de EmailService, mas não sabe qual implementação está sendo usada (SMTP ou Mock).
4. **Construtor NewUserNotifier**: Injeta a dependência concreta em UserNotifier.
5. **Facilidade de teste**: A implementação `MockEmailService` permite testar funcionalidades sem enviar e-mails reais.

---

## Prós e Contras da Injeção de Dependência

### **Prós**

1. **Alta Testabilidade**: Facilita a criação de testes unitários definindo implementações diferentes (ex.: mocks).
2. **Desacoplamento**: Componentes independentes baseados em contratos (interfaces) ao invés de implementações
   concretas.
3. **Flexibilidade**: Substituir implementações concretas no futuro é fácil.
4. **Reaproveitamento**: Promove a reutilização de código com componentes bem separados e configuráveis.

### **Contras**

1. **Complexidade Adicional**: Pode introduzir camadas e overhead em projetos simples.
2. **Aprendizado Inicial**: Exige uma compreensão clara de interfaces e design orientado a contratos.
3. **Gerenciamento Manual**: Em linguagens como Go, a DI é gerenciada de forma manual, sem frameworks.
4. **Over-engineering**: Pode levar a um design desnecessariamente complexo para projetos pequenos.

---

## Conclusão

A injeção de dependência é uma técnica poderosa que promove um design desacoplado, modular e testável. No entanto, o uso
inadequado pode adicionar complexidade desnecessária. Em Go, a simplicidade do idioma permite que a DI manual seja clara
e eficaz, especialmente através de interfaces e construtores.

Para equipes que desenvolvem sistemas grandes ou configuráveis, DI é essencial. Em projetos menores, deve ser usada com
cuidado para evitar "over-engineering".

---

## Links Úteis

- [Injeção de Dependência - Wikipedia](https://pt.wikipedia.org/wiki/Inje%C3%A7%C3%A3o_de_depend%C3%AAncia)
- [Guice vs Spring – Dependency Injection](https://www.baeldung.com/guice-spring-dependency-injection)
- [Testando Códigos em Go](https://pkg.go.dev/testing)
