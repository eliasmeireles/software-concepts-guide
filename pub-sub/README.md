# Publish/Subscribe

O modelo **Pub/Sub (Publish/Subscribe)** é um padrão de comunicação assíncrona e desacoplada usado principalmente em
sistemas distribuídos, que permite que os produtores de dados (publicadores) e os consumidores (assinantes) interajam
sem precisar se conhecer diretamente. Em vez de um modelo de comunicação ponto a ponto, o Pub/Sub segue uma abordagem
de "publicar" mensagens em um canal e "assinar" para receber essas mensagens de forma independente.

### Como Funciona o Modelo Pub/Sub:

1. **Publicador (Publisher)**:
    - O **publicador** é a parte responsável por gerar e enviar mensagens para um tópico. Ele **publica** as mensagens
      no sistema, sem precisar saber quem está recebendo essas mensagens.
    - O publicador não se importa com os consumidores (assinantes); ele simplesmente envia suas mensagens para um tópico
      específico.

2. **Tópico (Topic)**:
    - O **tópico** é um canal de comunicação onde as mensagens são enviadas pelos publicadores.
    - Pode ser visto como um "filtro" para mensagens de um tipo específico. Por exemplo, em um sistema de e-commerce,
      pode haver um tópico para "novos pedidos" e outro para "alterações de status de pedidos".

3. **Assinante (Subscriber)**:
    - O **assinante** é a parte que se inscreve (assina) em um tópico específico para receber mensagens enviadas por
      publicadores nesse tópico.
    - Um assinante não sabe quem são os publicadores, ele só sabe que irá receber as mensagens de um tópico do qual está
      inscrito.

4. **Mensagem**:
    - A **mensagem** é a informação que o publicador envia para o tópico. Ela pode ter diversos formatos, como texto,
      JSON, XML ou até binário.
    - As mensagens são entregues aos assinantes que estão interessados em recebê-las a partir do momento em que se
      inscrevem no tópico.

### Características do Modelo Pub/Sub:

1. **Desacoplamento**:
    - O modelo Pub/Sub promove um **desacoplamento** entre os publicadores e os assinantes. Isso significa que os
      publicadores não precisam saber quem está consumindo as mensagens e vice-versa. A comunicação é feita através de
      um tópico intermediário.
    - Isso facilita a escalabilidade, já que novos assinantes podem ser adicionados ao sistema sem alterar os
      publicadores, e os publicadores podem mudar sua lógica sem afetar os assinantes.

2. **Assinaturas Dinâmicas**:
    - Os assinantes podem se inscrever em tópicos dinamicamente e também cancelar sua inscrição a qualquer momento. Isso
      é útil em sistemas onde as necessidades dos consumidores podem mudar com frequência.

3. **Escalabilidade**:
    - O modelo Pub/Sub é altamente escalável. À medida que o número de publicadores ou assinantes cresce, a arquitetura
      pode ser facilmente expandida sem impactar a comunicação entre eles.

4. **Desempenho Assíncrono**:
    - O modelo Pub/Sub funciona de maneira assíncrona. Ou seja, os publicadores não precisam esperar uma resposta dos
      assinantes, o que melhora o desempenho e permite que os serviços se comuniquem sem bloqueios.
    - Isso permite que sistemas com alta carga de trabalho sejam mais eficientes e escaláveis.

5. **Entrega de Mensagens**:
    - **At least once**: Cada mensagem será entregue pelo menos uma vez aos assinantes. Isso é garantido mesmo que o
      sistema precise reenviar mensagens em caso de falha.
    - **At most once**: A mensagem será entregue no máximo uma vez, sem garantias de reenvio em caso de falhas.
    - **Exactly once**: A mensagem será entregue exatamente uma vez, sem duplicação.

### Exemplo de Fluxo de Pub/Sub:

1. **Publicação de Mensagem**:
    - Um serviço de pagamento de um sistema de e-commerce gera um evento quando o pagamento de um cliente é processado
      com sucesso. O serviço publica uma mensagem no tópico "pagamentos processados".

2. **Assinatura de Tópico**:
    - Um serviço de inventário se inscreve no tópico "pagamentos processados", porque ele precisa atualizar o estoque
      sempre que um pagamento for processado.

3. **Recebendo Mensagem**:
    - Assim que o serviço de pagamento publica a mensagem no tópico "pagamentos processados", o serviço de inventário
      recebe a mensagem de forma assíncrona e atualiza o estoque.

4. **Desinscrição**:
    - Se o serviço de inventário não precisar mais processar mensagens sobre pagamentos, ele pode se desinscrever do
      tópico a qualquer momento.

### Tipos de Implementações de Pub/Sub:

1. **Sistema de Fila (Queue-based Pub/Sub)**:
    - Em algumas implementações de Pub/Sub, um serviço pode enviar mensagens para uma fila, e os assinantes processam as
      mensagens de maneira sequencial.
    - Exemplos: **RabbitMQ**, **Amazon SQS**.

2. **Broadcasting Pub/Sub**:
    - Em sistemas baseados em broadcasting, as mensagens enviadas para o tópico são entregues a todos os assinantes
      simultaneamente.
    - Exemplos: **Apache Kafka**, **Google Cloud Pub/Sub**, **Redis Pub/Sub**.

### Benefícios do Modelo Pub/Sub:

1. **Desacoplamento e Flexibilidade**:
    - O Pub/Sub permite que sistemas sejam desenvolvidos de maneira mais modular, já que os componentes (publicadores e
      assinantes) são independentes entre si.

2. **Escalabilidade**:
    - O sistema pode crescer facilmente adicionando mais assinantes ou publicadores sem alterar a arquitetura existente.

3. **Facilidade de Integração**:
    - Novos componentes podem ser facilmente integrados ao sistema, já que eles só precisam assinar tópicos de
      interesse.

4. **Desempenho Assíncrono**:
    - Como a comunicação é assíncrona, ela não bloqueia o fluxo de trabalho dos sistemas, permitindo maior performance e
      capacidade de processamento paralelo.

### Desvantagens do Modelo Pub/Sub:

1. **Gerenciamento de Estado e Consistência**:
    - Em sistemas de larga escala, pode ser difícil garantir que todos os assinantes recebam todas as mensagens
      corretamente, especialmente quando há falhas temporárias.

2. **Complicação de Monitoramento**:
    - Como a comunicação é assíncrona e distribuída, pode ser difícil monitorar e depurar os fluxos de mensagens entre
      publicadores e assinantes.

3. **Entrega de Mensagens**:
    - Dependendo da implementação, pode haver o risco de duplicação de mensagens ou de mensagens perdidas, a menos que
      sejam usadas garantias específicas como "exatamente uma vez" ou mecanismos de persistência.

### Exemplo de Tecnologias de Pub/Sub:

- **Apache Kafka**: É uma plataforma distribuída para processar fluxos de dados em tempo real. O Kafka oferece uma
  implementação de Pub/Sub altamente escalável e com armazenamento persistente das mensagens.
- **Google Cloud Pub/Sub**: Serviço gerenciado da Google Cloud para comunicação assíncrona entre aplicativos, ideal para
  sistemas distribuídos e grandes volumes de dados.
- **Redis Pub/Sub**: Redis oferece uma implementação simples e eficiente de Pub/Sub, usada principalmente para troca de
  mensagens em tempo real.
- **Amazon SNS (Simple Notification Service)**: Serviço da AWS que facilita o envio de mensagens para múltiplos
  assinantes, podendo ser usado com outros serviços da AWS.

### Conclusão:

O modelo Pub/Sub é uma arquitetura poderosa para sistemas distribuídos, facilitando a comunicação assíncrona e
desacoplada entre diferentes componentes. Ele é amplamente utilizado em sistemas de microserviços, big data,
processamento de eventos em tempo real e em situações onde a escalabilidade e a flexibilidade são essenciais. Ao
implementar Pub/Sub, as equipes podem garantir que seus sistemas sejam mais ágeis, escaláveis e fáceis de manter,
especialmente em ambientes complexos e distribuídos.