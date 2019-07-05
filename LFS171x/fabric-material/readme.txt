Membros do grupo:
Bruna Paracat        - 1611896
Pedro Sousa Meireles - 1510962

Esse trabalho implementa uma solução em hyperledger fabric para um prontuário médico descentralizado.

Nele, é possível criar pacientes e médicos. O paciente pode autorizar um médico a visualizar seus dados e fazer upload de exames.
A qualquer momento o paciente pode retirar o acesso de qualquer médico.

Requerimentos Técnicos:
    Go
    NodeJs (versão 8)
    Docker

Execuão:
Para executar o programa, execute ./startFabric.sh, npm install, node registerAdmin.js, node registerUser.js e node server.js.
Abra o site em localhost:8000.