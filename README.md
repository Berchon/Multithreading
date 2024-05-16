# Multithreading
Neste desafio você terá que usar o que aprendemos com `Multithreading` e `APIs` para buscar o resultado mais rápido entre duas `APIs` distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

- https://brasilapi.com.br/api/cep/v1/{cep}
- http://viacep.com.br/ws/{cep}/json/

## Os requisitos para este desafio são:

1. Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

2. O resultado da request deverá ser exibido no `command line` com os dados do endereço, bem como qual API a enviou.

3.  Limitar o tempo de resposta em `1 segundo`. Caso contrário, o erro de `timeout` deve ser exibido.

## Requisitos básicos:
- Golang 1.21+

## Como usar:
Para utilizar essa API, siga os seguintes passos:

### Instalação
- Verifique se Go está instalado em seu sistema:
```BASH
go version
```

- Clone o repositório para sua máquina local:
```BASH
git clone https://exemplo.com/seu-repositorio.git
```

### Configuração: 
- Para garantir que todas as dependências sejam resolvidas, execute:
```BASH
go mod tidy
```

### Utilização**
- A API está rodando na porta 8080 e o endpoint para obter informações de endereço com base em um CEP é:
```BASH
http://localhost:8080/{cep}
```

- Para fazer uma requisição ao endpoint utilizando cURL, você pode executar o seguinte comando:
```BASH
curl http://localhost:8080/01310930
```

## Documentação:
Para acessar a documentação do Swagger da API, siga os passos abaixo:

1. Certifique-se de que a API esteja em execução localmente.
2. Abra um navegador da web e digite o seguinte URL na barra de endereço:
```bash
http://localhost:8080/swagger/index.html
```
3. Pressione Enter para carregar a página da documentação do Swagger.
4. Você verá a interface do Swagger, que permite explorar os endpoints da API, testar as operações disponíveis e visualizar detalhes sobre os parâmetros necessários para cada requisição.
5. Agora você pode navegar pela documentação do Swagger para entender melhor como interagir com a API e seus endpoints.

A documentação do Swagger é uma ferramenta útil para facilitar o desenvolvimento, teste e integração de APIs.

## Pontos de melhorias
Para melhorar a API, aqui estão algumas sugestões de melhorias:

1. Melhoria nos Responses de Erros:
    - Aprimorar os textos de erro retornados pela API para serem mais específicos e informativos. Isso ajudará os desenvolvedores a identificar e solucionar problemas com mais facilidade.

2. Tratamento de Erros:
    - Implementar tratamentos de erro mais robustos para lidar com situações em que uma API externa esteja fora do ar ou se o usuário estiver sem conexão com a internet. Isso garantirá uma melhor experiência para o usuário final.

3. Testes de Unidade e Testes de Integração:
    - Implementar testes de unidade para garantir que cada parte do código funcione corretamente isoladamente.
    - Implementar testes de integração para verificar se os componentes da aplicação funcionam bem juntos.

Devido à escolha de exercitaDevido à escolha de exercitar uma arquitetura mais complexa inicialmente, os testes foram deixados para um segundo momento. No entanto, a inclusão de testes de unidade e integração pode trazer benefícios significativos em termos de qualidade e confiabilidade do software. Essas melhorias garantirão que a API seja mais robusta, confiável e escalável no futuro.