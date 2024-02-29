# Relatório

Este relatório apresenta uma análise detalhada dos fundamentos de segurança da informação aplicados a ambientes de Internet das Coisas (IoT). Exploraremos conceitos essenciais, como a proteção de dados em trânsito e em armazenamento, bem como a autenticação e autorização da comunicação. Além disso, discutiremos a relevância e a eficácia de cada uma dessas estratégias de proteção, destacando sua importância para garantir a segurança e integridade das informações em ambientes IoT.

Pilares do CIA Triad:

Integridade
Disponibilidade
Confidencialidade

**Pergunta:** 

O que acontece se você utilizar o mesmo ClientID em outra máquina ou sessão do browser? Algum pilar do CIA Triad é violado com isso?

**Resposta:** 

Quando a segunda sessão do browser é conectado no ClientID a primeira sessão vai ser desconectada automaticamente. Isso pode resultar em uma violação do pilar de Integridade. O pilar de Integridade refere-se à garantia de que os dados permaneçam íntegros, precisos e confiáveis ao longo do tempo e durante todo o processo de manipulação. Quando o mesmo ClientID é usado em diferentes máquinas ou sessões do navegador, pode haver uma falta de integridade nos dados associados a esse identificador. Isso ocorre porque não há garantia de que as alterações feitas em uma máquina ou sessão sejam refletidas de forma consistente em todas as outras instâncias.

**Pergunta:** 

Com os parâmetros de ```resources```, algum pilar do CIA Triad pode ser facilmente violado?

**Resposta:** 

Os parâmetros de recursos definidos estão relacionados à alocação de recursos de computação para o contêiner MQTT5, como CPU e memória. Embora esses parâmetros possam influenciar o desempenho e a disponibilidade do serviço, não estão diretamente ligados aos pilares. Por exemplo, se as limitações dos parâmetros recurso foram forte o suficientes podem afetar os três pilares da cia triade.

**Pergunta:**

Sem autenticação (repare que a variável ```allow_anonymous``` está como ```true```), como a parte de confidencialidade pode ser violada?

**Resposta:** 

Sem a autenticação adequada pode acontecer algumas coisas, por exemplo, o acesso não autorizado, intercepção de dados, roubo de indentidade e vazamento de dados sensíveis. o acesso não autorizado significa que sem autenticação, qualquer pessoa pode acessar os dados, mesmo aquelas que não têm permissão, ou seja, as informações confidenciais podem ser acessadas por qualquer pessoa, incluindo indivíduos mal-intencionados. A intercepção de dados significa que sem autenticação, os dados transmitidos podem ser interceptados por terceiros. Isso é especialmente problemático em comunicações não criptografadas, onde os dados podem ser facilmente lidos por qualquer pessoa que esteja interceptando a comunicação. Em relação ao roubo de identidade, sem autenticação significa que não há maneira de verificar a identidade dos usuários. Isso abre a porta para roubo de identidade, onde um invasor pode se passar por outra pessoa para acessar informações confidenciais. Em relação ao vazamento de dados sensíveis, a falta de autenticação permite que qualquer pessoa acesse os dados, incluindo dados sensíveis e privados. Isso aumenta significativamente o risco de vazamento de informações confidenciais.

**Pergunta:**

Tente simular uma violação do pilar de Confidencialidade.

**Resposta:** 

O video no link abaixo demonstra uma conexão bem sucedida com um broker remoto usando um web client logo depois outra coneão é feito nesse mesmo broker remoto com outro web client porém a primeira conexão foi desconectado. Esse tipo de priorização de conexão para o mais recente viola o pilar de confidencialidade porque permite que usuarios não autorizados vejam a informação confidencial.

**Pergunta:**

Tente simular uma violação do pilar de Integridade.

**Resposta:** 

O video no link abaixo demonstra uma conexão bem sucedida com um broker remoto usando um web client logo depois outra coneão é feito nesse mesmo broker remoto com outro web client porém a primeira conexão foi desconectado. Essa priorização de conexão para a mais recente bem-sucedida viola o pilar de integridade porque não tem como saber o que foi feito depois dessa conexão indevida.

**Pergunta:**
Tente simular uma violação do pilar de Disponibilidade.

**Resposta:** 

O video no link abaixo demonstra uma conexão bem sucedida com um broker remoto usando um web client logo depois outra coneão é feito nesse mesmo broker remoto com outro web client porém a primeira conexão foi desconectado. Isso claramente viola o pilar de disponibilidade porque a coneão com o broker remoto vai ser priorizado para a conexão bem-sucedida mas recente que teve.

**Link da evidência dos 3 pilares**

Link da [evidência](https://drive.google.com/file/d/1PwC6ZWbBV2ofbe-GfEPLq9jZ_SVoXWL2/view?usp=sharing) para a simulação da violação dos três pilares.