/* 
	Pacote Boom implementa funcionalidades para o suporte a criação de um serviço tanto REST API quanto gRPC.

	...

	GrpcServer = Cria um servidor recebendo como parametro apenas uma variavel do tipo *API, onde o mesmo
	utiliza o host de acordo com o que tem no yaml.

	...

	GrpcClient = Cria um client que envia solicitações para um servidor gRPC, recebendo apenas 3 parametros
	*API, credential para garantir a segurança do transport entre cliente e servidor, interface clientGrpc implementando
	uma função que recebe grpc.ClientConnInterface, ....
*/
package boom