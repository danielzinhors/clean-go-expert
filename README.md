# clean-go-expert

1-Criar o docker com o comando 
  docker compose up
2-Iniciar as migracoes com o comando 
  make createmigration create -ext=sql -dir=sql/migrations -seq init
3-Rodar as migracoes com o comando
  make migrate