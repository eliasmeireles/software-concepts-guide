# Event Sourcing

**Event Sourcing** é um padrão arquitetural utilizado em sistemas de software para gerenciar estado e mudanças de
estado,
capturando todas as mutações como eventos sequenciais e imutáveis. Este padrão é particularmente útil em sistemas onde
você precisa garantir integridade histórica, auditabilidade e reconstrução de estado.

## Conceitos Fundamentais

1. **Evento**:
   Um evento é uma representação de algo que aconteceu no sistema, armazenado como um registro imutável. Por exemplo,
   eventos como `UserRegistered`, `ProductAddedToCart` ou `PaymentProcessed`.

2. **Fonte Única da Verdade**:
   No Event Sourcing, os eventos armazenados servem como a fonte única da verdade do estado da aplicação. O estado atual
   do sistema pode ser derivado da sequência de eventos armazenados.

3. **Armazenamento de Eventos**:
   Todos os eventos são persistidos em um repositório chamado **Event Store**, que armazena esses registros de forma
   sequencial e imutável. O Event Store também pode ser usado como histórico para reconstruir ou debugar o comportamento
   do sistema.

4. **Reprodução de Eventos**:
   O estado atual de um sistema é reconstruído ao reproduzir a sequência de eventos registrada. A reprodução de eventos
   garante que o estado final é consistente e suporta fluxos de recuperação de falhas.

## Benefícios do Event Sourcing

- **Auditabilidade**:
  Como cada mudança de estado é registrada como um evento, o histórico completo está disponível para auditorias.

- **Reversibilidade**:
  Eventos capturados permitem desfazer transações ou reverter estados com facilidade.

- **Consistência em Sistemas Distribuídos**:
  Eventos podem ser propagados para diferentes partes de um sistema distribuído, garantindo consistência entre
  componentes e serviços.

- **Evolução de Dados**:
  Você pode criar novos estados ou gerar relatórios diferentes apenas derivando informações de eventos históricos.

## Exemplo de Implementação: Evento e Armazenamento

Aqui está um exemplo implementado em **Go**, ilustrando um sistema básico de Event Sourcing:

```go
package main

import (
	"fmt"
)

// Evento base
type Event interface {
	EventType() string
}

type UserRegistered struct {
	UserID    string
	UserName  string
	Timestamp string
}

func (e UserRegistered) EventType() string {
	return "UserRegistered"
}

// Event Store (Armazena eventos)
type EventStore struct {
	events []Event
}

func (store *EventStore) AddEvent(event Event) {
	store.events = append(store.events, event)
}

func (store *EventStore) GetEvents() []Event {
	return store.events
}

// Reprodução de eventos para reconstruir estado
type User struct {
	UserID   string
	UserName string
}

func ReplayEvents(events []Event) []User {
	var users []User
	for _, event := range events {
		switch e := event.(type) {
		case UserRegistered:
			users = append(users, User{
				UserID:   e.UserID,
				UserName: e.UserName,
			})
		}
	}
	return users
}

func main() {
	// Criando o Event Store
	store := &EventStore{}

	// Registrando eventos
	store.AddEvent(UserRegistered{"1", "Alice", "2023-10-01T10:00:00Z"})
	store.AddEvent(UserRegistered{"2", "Bob", "2023-10-01T11:00:00Z"})

	// Listando eventos
	for _, event := range store.GetEvents() {
		fmt.Printf("Event Type: %s\n", event.EventType())
	}

	// Reconstruindo estado
	users := ReplayEvents(store.GetEvents())
	fmt.Println("Reconstructed Users:")
	for _, user := range users {
		fmt.Printf("UserID: %s, UserName: %s\n", user.UserID, user.UserName)
	}
}
```

## Ciclo no Event Sourcing

1. **Criação do Evento**:
   O evento é criado e descreve uma mudança específica no sistema.

2. **Persistência do Evento**:
   Esse evento é armazenado no `Event Store`.

3. **Publicação do Evento**:
   Outros serviços ou agregados relevantes são notificados sobre a ocorrência deste evento.

4. **Reprodução do Evento**:
   Eventos são carregados do armazenamento para reconstruir o estado atual.

## Comparação com Persistência Tradicional

| Persistência Tradicional                      | Event Sourcing                               |
|-----------------------------------------------|----------------------------------------------|
| Persiste apenas o estado atual.               | Persiste todas as mudanças como eventos.     |
| Difícil de auditar o histórico.               | Histórico completo é nativamente suportado.  |
| Alterar o modelo de dados pode ser arriscado. | Modelo de eventos permite evolução flexível. |

## Casos de Uso

- Sistemas de pagamento.
- Logística e rastreamento.
- Sistemas financeiros com alta necessidade de auditabilidade.
- Software de gestão de histórico, como CRMs e ERPs.

## Desafios

- **Complexidade Adicional**:
  Implementar Event Sourcing pode tornar os sistemas mais complexos devido à necessidade de gerenciar eventos e versões.

- **Lidar com Evolução de Eventos**:
  Considere como atualizar eventos antigos para novos formatos ao longo do tempo.

- **Volume de Dados**:
  O armazenamento de todos os eventos pode gerar um volume significativo de dados.

---

Event Sourcing, apesar da complexidade, é uma abordagem poderosa que fornece auditabilidade, flexibilidade e
consistência para aplicações que exigem controle estrito de mudanças e históricos.