# FORK btcgo v1.1

Abra o cmd e rode: go mod tidy

dps entre na pasta divisor de intervalos e com base no arquivo Puzzles.json, escolha a wallet ou as wallets q deseja buscar
voce pode adicionar quantos wallets e intervalos quiser para buscar ao mesmo tempo.
abra o main.go:

ex: para buscar as carteiras 66, 67, 68

Metodo 1 para buscar mais de uma carteira

// Input range
	inputRange := Range{
		Min:    "0x20000000000000000", // range minimo
		Max:    "0xfffffffffffffffff", // range maximo
		Status: "13zb1hQbWVsc2S7ZTZnP2G4undNNpdh5so, 1BY8GQbnueYofwSuFAT3USAhGjPrkxDdW9, 1MVDYgVaSN6iKKEsbzRUAYFrYJadLYZvvZ", // wallets a ser buscada

// Number of intervals 
	numIntervals := 200 // Numero da divisao de intervalos (quanto maior, maior a aleatoriedade) MINIMO 2 INTERVALOS

 Metodo 2 para buscar mais de uma carteira 
 
 Criar um arquivo ranges.json nesse modelo
 
{
  "ranges":  [
            { "min": "0x20000000000000000", "max": "0x3ffffffffffffffff", "status": "13zb1hQbWVsc2S7ZTZnP2G4undNNpdh5so" },
        { "min": "0x40000000000000000", "max": "0x7ffffffffffffffff", "status": "1BY8GQbnueYofwSuFAT3USAhGjPrkxDdW9" },
        { "min": "0x80000000000000000", "max": "0xfffffffffffffffff", "status": "1MVDYgVaSN6iKKEsbzRUAYFrYJadLYZvvZ" }
    }
  ]
}

Metodo 1 para buscar APENAS uma carteira
EX: carteria 66
abra o main.go em "divisor de intervalos"

// Input range
	inputRange := Range{
		Min:    "0x20000000000000000",
		Max:    "0x3ffffffffffffffff",
		Status: "13zb1hQbWVsc2S7ZTZnP2G4undNNpdh5so",
	}

// Number of intervals
	numIntervals := 200 (minimo 2 intervalos)




Método de busca:
 Busca randômica dupla
 random 1, embaralhar os intervalos de ranges.json e saltar em X segundos um em um
 random 2, dentro desses intervalos é gerado um bloco aleatorio cm X chaves q serao testadas

Funcionamento principal
 imagine um range de 1 a 100 (intevalo atual que está lendo), o algoritmo escolhe randomicamente um numero (exemplo 15), e gera um "bloco" do tamanho especificado pelo usuário a partir do numero 15. vamos supor que esse bloco seja de tamanho 10, ele irá ler do 15 ao 25 sequencialmente. Caso a chave n seja encontrada, outro bloco é gerado e o processo se repetirá.

obs: o range de cada bloco lido será armazenado para que n seja sobreposto. Recomendo para o puzzle 66 um bloco de tamanho 10 milhões.

Funcionalidade carregar progresso:
caso seja interrompida a busca, a opção "2. Continuar busca anterior" ira carregar os blocos armazenados para garantir que eles n sejam lidos novamente.
obs: caso carregue um progresso, UTILIZE O MESMO TAMANHO DE BLOCO PARA N HAVER CONFLITOS

Quando uma chave for encontrada a busca irá parar e a chave será mandada para o arquivo found_keys

FORK: BTCGO - Investidor Internacional v0.1
Exemplo de configuraçao para o puzzle 68
___________________________________________
__________BTCGO - Mod BY: Inex_____________
__________________v1.1_____________________
Wallets a serem buscadas:
1MVDYgVaSN6iKKEsbzRUAYFrYJadLYZvvZ
200 intervalos carregados
Digite o intervalo de tempo em segundos para saltar para um novo intervalo: 30
Digite o número de threads: 4
Digite o tamanho do bloco (Recomendado 10M para wallet 66): 10000000
1. Iniciar nova busca
2. Continuar busca anterior
1
