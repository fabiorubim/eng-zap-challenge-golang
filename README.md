# eng-zap-challenge-golang
ZAP challenge

Para executar: Ao efetuar o clone do projeto, acessar o diretório 'eng-zap-challenge-golang' e utilizando sua ferramenta/IDE de preferência rode o arquivo main.go. Neste caso ele foi desenvolvido utilizando a IDE GoLand.

Outro método para executar a aplicação é abrir o prompt de comando no diretório do projeto e executar: go run main.go

OBS1: O projeto foi desenvolvido no Windows, neste caso será necessário instalar a a linaguagem Go e seu pacotes. Link: https://golang.org/doc/install?download=go1.13.windows-amd64.msi

OBS2: O projeto utiliza alguns pacotes que estão hospedados no GitHub(por exemplo o fasthttp), logo é necessário ter o Git instalado e configurado na variável de ambiente PATH.

OBS3: Permita a conexão quando o firewall do Windows solicitar

Após a aplicação iniciar as seguintes mensagens serão exibidas no console/prompt:

downloading data...

reading data...

parsing data...

done!

Após done! é possível acessar os seguintes endpoints:

http://localhost:8000/zap

http://localhost:8000/vivareal
