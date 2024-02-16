# clean-go-expert

1-Criar o docker com o comando 
  docker compose up
2-Iniciar as migracoes com o comando 
  make createmigration create -ext=sql -dir=sql/migrations -seq init
3-Rodar as migracoes com o comando
  make migrate
4-Rodar o projeto
  Acessar a pasta cmd
  acessar a pasta ordersystem usar o comando
     go run main.go wire_gen.go
5-Chamadas webserver
  Na pasta api
    Criar uma nova order
      Usar o arquivo create_order.http 
    Listar as orders
      Usar o arquivo list_order.http
6-Chamadas GRPC
  abrir o Evans com o comando
     evans -r repl
  abrir o package com o comando
    package pb
  abrir o service com o comando
    service OrderService
  abrir as chamadas
    criar order com o comando 
      call CreateOrder
    listar as orders com o comando
      call ListOrders
7-Chamadas Grahql
  Para acessar o playground
    abrir o navegador e aceesar
      localhost:8080
  Para criar uma nova order
    mutation createOrder {
      createOrder(
        input:{
          id: "h"
          Price: 5
          Tax: 15.7
        }
      ) {
        id
        Price
        Tax
        FinalPrice
      }
    }
  Para listar as orders
    query ListOrders {
      ListOrders{
        id
        Price
        Tax
        FinalPrice
      }
    }

