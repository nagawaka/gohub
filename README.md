# Running
Crie uma docker machine com docker-machine create -d "virtualbox" $NOME_DA_MAQUINA

docker-machine start $NOME_DA_MAQUINA
docker-machine env $NOME_DA_MAQUINA
eval $(docker-machine env $NOME_DA_MAQUINA)
docker-compose up

# Testando api
Pega ip com
docker-machine ip $NOME_DA_MAQUINA