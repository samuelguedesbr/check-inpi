# Sistema Automatizado de Monitoramento de Marcas – INPI

## 1. Objetivo do projeto

Desenvolver um backend capaz de monitorar automaticamente os processos de registro de marcas no Instituto Nacional da Propriedade Industrial (INPI), identificando novas movimentações, publicações e alterações no andamento dos pedidos.

O sistema deverá consultar semanalmente a Revista da Propriedade Industrial (RPI), publicada pelo INPI geralmente às terças-feiras, baixar o arquivo XML disponibilizado e processar suas informações para identificar se algum dos processos cadastrados recebeu uma nova publicação.

## 2. Funcionamento do sistema

O funcionamento será dividido nas seguintes etapas:

**1. Execução automática**
- Configurar um agendador de tarefas (scheduler) para executar o processo semanalmente, às terças-feiras.
- Verificar se a nova edição da RPI já está disponível no site oficial do INPI.

**2. Download e processamento do XML**
- Localizar e baixar automaticamente o arquivo XML da edição publicada.
- Ler e interpretar os dados do XML.
- Extrair informações como número do processo, nome da marca, titular, código do despacho e descrição da publicação.

**3. Identificação de alterações**
- Comparar os dados extraídos com os processos de marcas cadastrados no banco de dados.
- Identificar se houve alguma nova publicação relacionada a determinado processo.
- Verificar se ocorreu alguma alteração no andamento, como exigência, oposição, deferimento, indeferimento ou outra movimentação publicada pelo INPI.

**4. Atualização do banco de dados**
- Registrar as novas movimentações identificadas.
- Atualizar o histórico e o status de cada processo.
- Evitar registros duplicados quando o mesmo XML for processado novamente.

**5. Notificação por e-mail**
- Quando uma nova movimentação for identificada, enviar automaticamente um e-mail ao responsável pelo processo.
- O e-mail deverá informar:
  - Nome da marca.
  - Número do processo.
  - Descrição da nova publicação.
  - Data da publicação.
  - Novo status ou movimentação identificada.

Caso não exista nenhuma alteração, nenhuma notificação será enviada.

## 3. Requisitos técnicos

- **Linguagem:** Go.
- **Framework:** Gin.
- **Banco de dados:** PostgreSQL.
- **Agendamento:** Scheduler ou Cron Jobs.
- **Processamento:** Download, leitura e interpretação de arquivos XML.
- **Notificações:** Integração com serviço de envio de e-mails.
- **Automação:** Execução independente de intervenção humana.
- **Logs:** Registro das consultas, downloads, processamentos, erros e notificações enviadas.

## 4. Resultado esperado

Ter um serviço backend que funcione continuamente, verificando semanalmente as publicações do INPI e acompanhando automaticamente os processos cadastrados.

O objetivo é eliminar a necessidade de consultar manualmente a Revista da Propriedade Industrial todas as semanas, permitindo que os responsáveis sejam informados rapidamente quando houver alguma movimentação relevante em seus pedidos de registro de marca.

https://revistas.inpi.gov.br/rpi/