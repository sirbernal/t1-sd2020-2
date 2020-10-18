## Laboratorio N°1 Sistemas distribuidos Grupo GCR - G
##### Integrantes:

| Nombre  |   Rol|
| ------------ | ------------ |
| Cristian Bernal R.  | 201773026-9   |
|  Raúl Álvarez C. |  201773010-2 |



#### Como correr el código en cada maquina
Dependiendo de la máquina, se deberá ejecutar el comando make correspondiente:

|  Comando make | Máquina   |
| ------------ | ------------ |
|  runcliente | 10.10.28.81 |
|   runlogistica | 10.10.28.82  |
| runcamion |   10.10.28.83 |
|  runfinanciero | 10.10.28.84  |


#### Funcionamiento del Código
Se usó el lenguaje Go junto con RabbitMQ y gRPC para los envios de datos entre las maquinas. La logica de cada código esta comentada. Se detalla que en la maquina de financiero (10.10.28.84) esta corriendo el server Rabbitmq.
