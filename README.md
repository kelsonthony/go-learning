pacote para validar email

go get github.com/badoux/checkmail

go mod tidy para remover as deps que não estão sendo usadas 

go é fortemente tipada

iniciar um mod go mod init 

adicionar um pacote externo

go get github.com/urfave/cli

	//Concorrência é a capacidade de um programa lidar com múltiplas tarefas ao mesmo tempo, mas não necessariamente executá-las simultaneamente. O sistema operacional pode alternar entre as tarefas, dando a impressão de que estão sendo executadas ao mesmo tempo.
	//Paralelismo é a capacidade de um programa executar múltiplas tarefas simultaneamente, utilizando múltiplos processadores ou núcleos. Isso é possível quando as tarefas são independentes e podem ser executadas em paralelo.

	//Goroutines são funções ou métodos que são executados concorrentemente com outras goroutines. Elas são leves e eficientes, permitindo que você crie milhares ou até milhões de goroutines sem esgotar os recursos do sistema.

	//Para criar uma goroutine, basta usar a palavra-chave "go" seguida da chamada da função ou método que você deseja executar concorrentemente. Por exemplo: