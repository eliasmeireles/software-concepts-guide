# CQRS: Command Query Responsibility Segregation

## O que é CQRS?

**CQRS (Command Query Responsibility Segregation)** é um padrão arquitetural que separa as responsabilidades de leitura
e escrita em um sistema. Em vez de usar um único modelo de dados e camadas de serviço para manipular e consultar dados,
o CQRS sugere ter dois modelos distintos:

- Um modelo para **comandos** (escrita): Responsável por modificar o estado do sistema.
- Um modelo para **consultas** (leitura): Responsável por retornar dados sem alterar o estado do sistema.

Esse padrão permite otimizar cada lado (leitura e escrita) de forma independente, melhorando desempenho, escalabilidade
e manutenção do sistema.

---

## Benefícios do CQRS

1. **Escalabilidade**:
   - Permite escalar separadamente os sistemas de leitura e escrita. Por exemplo, se o sistema recebe muito mais
     leituras do que escrituras, os recursos podem ser otimizados para atender à demanda de consultas.

2. **Simplificação de Lógicas Complexas**:
   - Ao separar leitura e escrita, as lógicas associadas a cada operação tornam-se mais simples e diretas.

3. **Flexibilidade no Modelo de Dados**:
   - Os modelos de dados podem ser projetados de forma diferente para leitura e escrita, permitindo ajustes que atendam
     melhor às necessidades específicas de cada operação.

4. **Suporte a Event Sourcing**:
   - O CQRS é frequentemente usado junto com o Event Sourcing, onde eventos são armazenados como uma fonte de verdade, e
     os estados de leitura são gerados a partir desses eventos.

5. **Isolamento de Conflitos de Concurrency**:
   - Como as operações de leitura e escrita usam caminhos diferentes, é menos provável que ocorram conflitos entre elas.

---

## Desvantagens do CQRS

1. **Complexidade Adicional**:
   - A separação de leitura e escrita pode introduzir complexidade no design do sistema e exigir mais esforço de
     manutenção.

2. **Sincronização de Dados**:
   - Como os modelos de leitura e escrita são separados, pode haver desafios para mantê-los sincronizados em sistemas
     altamente dinâmicos.

3. **Curva de Aprendizado**:
   - Requer que a equipe esteja familiarizada com o padrão e suas implicações para evitar implementações equivocadas.

4. **Latência em Sistemas Eventuais**:
   - Quando combinado com consistência eventual, pode haver um pequeno atraso entre a execução de uma escrita e a
     disponibilidade do dado para leitura.

---

## Componentes do CQRS

1. **Modelo de Comandos**:
   - Responsável pelas operações que alteram o estado do sistema (create, update, delete).
   - Inclui regras de negócio e validações associadas aos comandos.

2. **Modelo de Consultas**:
   - Otimizado para a obtenção de dados, geralmente usando um modelo de dados diferente do modelo de comandos.
   - Pode ser baseado em bancos de dados otimizados para consultas, como caches ou bases de dados orientadas a
     documentos.

3. **Eventos**:
   - Em implementações com Event Sourcing, eventos representam as mudanças no sistema e são usados para atualizar os
     modelos de leitura.

---

## Cenários de Uso

O CQRS é mais adequado para:

1. **Sistemas com Altas Taxas de Consulta**:
   - Aplicativos onde as leituras superam amplamente as operações de escrita, como dashboards e relatórios.

2. **Domínios Complexos**:
   - Sistemas onde as lógicas de leitura e escrita são muito diferentes ou possuem complexidade elevada.

3. **Escalabilidade e Desempenho**:
   - Necessidade de escalar ou otimizar separadamente os fluxos de leitura e escrita.

4. **Integração com Event Sourcing**:
   - Quando utilizado com Event Sourcing, o CQRS permite reconstruir o estado de leitura a partir dos eventos
     armazenados.

---

## Considerações Finais

O CQRS é um padrão poderoso que traz vantagens significativas em sistemas complexos e de alta escalabilidade. Contudo, a
implementação deve ser cuidadosamente planejada para evitar complexidade desnecessária em sistemas simples ou de baixa
escala. Avalie as necessidades do projeto antes de adotar esse padrão e considere as possíveis integrações com outros
conceitos, como Event Sourcing e Consistência Eventual.

