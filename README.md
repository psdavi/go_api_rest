# go_api_rest

go get -u github.com/gorilla/mux
sudo snap start docker


API REST utilizando o GO. Desde os princípios, como retornar uma requisição com JSON, como fazer o roteamento desses endpoints, como conectar com banco de dados.


Vai ser usado o banco de dados no Docker, vamos conectar a nossa API REST consumindo um banco de dados no Docker. Vamos usar o RM do GO para conseguir conectar com o banco de dados e manipular mais fácil as nossas requisições sem precisar ficar escrevendo SQL a mão.

Vamos criar o CRUD completo, vamos ser capazes de criar recursos, deletar, editar, atualizar. Vamos pensar em como deixamos a nossa API com um nível bom de estrutura, informando qual vai ser o Content-Type que estamos devolvendo, vamos criar middlewares também e vamos integrar a nossa API com o React.


Depois que criamos todo nosso esqueleto, a nossa API como integramos o nosso back-end GO com um front-end React.



Vamos criar uma função chamada Home() que vai representar a nossa página inicial. func Home(w http.ResponseWritter, r *http.Request). Sempre quando vamos criar uma nova função com o GO relacionada a abrir uma página, precisamos receber a requisição e responder utilizamos dois atalhos que são esses: ResponseWriter e o Request. O w para o responsewriter e o r apontando para o request que vamos utilizar.



package main

import (
    "fmt"
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.print(w, "Home Page")
}

func HandleRequest() {
    http.HandleFunc("/", Home)
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
    fmt.Println("Iniciando o servidor Rest com Go")
    HandleResquest()
}





Essa função temos o hábito de chamar no GO de func HandleResquest(). “Handle” em inglês significa “lidar”, HandleResquest é lidar com as requisições, algo desse tipo. Eu vou colocar aqui usando o módulo que já temos que é o: http.HandleFunc(), eu vou falar assim: "existe uma função que tem esse endereço e vai ser atendida por alguém."

[06:58] Sempre que chegar uma requisição no barra "/", ou seja, quando colocarmos o endereço da nossa página sem nada, sem "/compras" ou "/produtos" eu quero que ele caia nessa página específica e quem vai ser responsável por atendê-lo será a nossa função Home(), só isso, http.HandleFunc(“/”, Home).

[07:15] Agora aqui no nosso arquivo “main” eu vou executar esse nosso HandleResquest() executando ele aqui. Vamos relembrar o que fizemos: criamos uma função chamada Home() e falamos que toda vez que chegar uma requisição em Home() você vai responder exibindo essa mensagem "Home Page". Depois criamos alguém para lidar com essas requisições, quando chegar uma requisição em barra quem vai atender será o controle Home.

[07:41] Na nossa função main() falamos assim: "HandleResquest, executa", ele vai ficar "ouvindo" isso daqui. Só que se ele vai estar ouvindo o nosso servidor precisa estar funcionando, vou subir o nosso servidor aqui também. Vou usar a função log.Fatal(), que significa que qualquer problema que eu tenha quando eu subir o meu servidor ele vai nos apontar, vai mostrar para nós.

[08:02] E vou usar o mesmo pacote log.Fatal(http.ListenAndServe()), ou seja, escuta e sirva, escute e suba e servidor e vou passar aqui qual é a porta que eu quero subir. Vou subir no local host na porta (“:8000”, nil) mesmo, vou colocar vírgula aqui porque não temos nenhum handler para configurarmos por isso vou deixar como nil. log.Fatal(http.ListenAndServe(":8000", nil)). Não temos nenhuma configuração extra que precisaremos fazer.

[08:29] Vou salvar esse nosso projeto, vou abrir aqui o nosso terminal utilizando as teclas "Ctrl + J" ou “Command + J” e vamos executar esse código: go run main.go. E ele vai dar uma mensagem: "Iniciando o servidor Rest com GO" e parou. Eu vou abrir uma aba anônima e vou colocar aqui nesse endereço “localhost:8000” e temos aqui a página Home Page.

[08:57] Até aí está tranquilo só que não queremos exibir uma página com um texto específico, queremos que quando chegue uma determinada requisição para um endereço específico queremos responder com dados, de preferência com JSON que é o padrão mais utilizado atualmente. Isso vamos aprender na sequência.







[00:00] Subimos a nossa primeira página Home só que estou preocupado com algo, observe que esse código que fizemos aqui é simples e está com 21 linhas. Vamos relembrar os passos que demos: criamos uma função Home() que quando bater uma requisição para essa função Home() vamos exibir uma mensagem retornando o ResponseWriter e a mensagem será home page.

[00:26] Depois temos alguém que só toma conta das requisições HandleRequest(), quando chegar alguém aqui nesse endpoint específico, nesse endereço específico quem vai atender vai ser o Home.

package main

import (
    "fmt"
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.print(w, "Home Page")
}

func HandleRequest() {
    http.HandleFunc("/", Home)
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
    fmt.PrintIn("Iniciando o servidor Rest com Go")
    HandleResquest()
}COPIAR CÓDIGO
[00:38] Já conseguimos quebrar isso, por que? Podíamos colocar todo esse conteúdo aqui em um arquivo específico, em uma pasta específica porque a medida que o nosso projeto for aumentando temos as divisões certinhas, já sabemos exatamente onde precisamos alterar no nosso código.

[00:54] Imagine que estamos com uma página só exibindo um texto home page e estamos com 21 linhas, à medida que o nosso projeto for aumentar muito, tivermos diversos endpoints, rotas, controllers e depois modelos isso vai ficar muito difícil manter. Por isso vamos modularizar o nosso código.

[01:12] A primeira coisa que vou fazer vai ser criar uma pasta chamada routes. Vou criar aqui clicando nesse ícone aqui, um novo arquivo, um new file que vou chamar de “routes”. Aqui dentro eu vou criar um arquivo chamado “routes.go”, esse arquivo eu vou falar qual é o pacote que ele pertence que é o pacote de rotas.

[01:32] E eu vou manter nesse código de rotas esse nossa HandleResquest(). Vou tirar HandleResquest() da pasta “main.go” e vou colocá-lo aqui na pasta “routes.go” utilizando as teclas "Ctrl + C" e "Ctrl + V" e quando eu salvo ele já faz os imports certinhos para mim.

[01:45] Só que existe algo estranho aqui, ele fala que fiz os imports só que existe uma variável que está declarada e não sei o que ela é, que é essa Home. O que é Home? Ele não consegue enxergar o que está escrito aqui no nosso código “main.go”, que é a nossa função Home(). O que vamos fazer?

[02:02] Esse Home a funcionalidade dele é quando chegar uma requisição aí que lá no nosso arquivo de “routes.go” vai ser o responsável, quem vai controlar essa página vai ser essa nossa função.

[02:15] Podemos colocar isso em uma pasta chamada "controllers". Vou colocar aqui uma pasta chama “controllers” e vou criar um arquivo chamado “controllers.go”. Esse arquivo controllers eu vou colocar justamente esse nosso código Home, utilizo as teclas “Ctrl + X" para tirá-lo daqui e antes de colocá-lo aqui no “controller.go” vamos falar que package controllers, coloca aqui a nossa função, salvo e ela já está certinha.

package controllers

import (
    "fmt"
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Home Page")
}COPIAR CÓDIGO
[02:43] Lá nas rotas eu vou falar que quando chegar uma requisição para barra quem vai atender vai ser alguém que está lá nos controllers e aqui eu tenho várias opções. Como a nossa função chama-se Home(), será http.HandlerFunc(“/”, controllers.Home e quando eu salvo esse projeto ele já traz o import certinho.

package routes

import (
    "log"
    "net/http"

    "github.com/guilhermeonrails/go-rest-api/controllers"
)

func HandleResquest() {
    http.HandleFunc("/", controllers.Home)
    log.Fatal(http.ListenAndServe(":8000", nil))
}COPIAR CÓDIGO
[03:13] Esse import é único, é o nome que vai estar linkado com o projeto que você iniciou lá no GoMod. Vamos supor que você colocou o nome do seu projeto só go-rest-api ele vai trazer só: go-rest-api/controllers. Se você colocou um repositório do GitHub ele vai trazer o nome completo para você também.

[03:33] Colocamos o Home() no “controllers.go”, ele fez os imports necessários do Home(), de todos os modos que precisamos para esse nosso projeto funcionar e ele já consegue identificar. No “main.go” temos um único problema, ele sabe o que é HandleResquest(), vamos falar para ele que o HandleResquest() está no nosso arquivo de rotas, é ele quem vai lidar com todas as nossas requisições.

[03:57] Ao colocar routes.HandleResquest() ele já aparece o que eu quero. Eu salvo e ele já está tudo funcionando. Ele sabe que ele precisa executar essa função HandleResquest() lá do nosso módulo de rotas.

[04:16] Vamos ver se isso está funcionando. Eu vou abrir o nosso terminal e rodar o nosso servidor, vou colocar aqui go run main.go ele vai falar: "iniciando o servidor rest com go", vou atualizar e está funcionando.

[04:33] Olha o que fizemos agora: temos um arquivo main com 13 linhas.

package main

import (
    "fmt"

    "github.com/guilhermeonrails/go-rest-api/routes"
)

func main() {
    fmt.Println("Iniciando o servidor Rest com Go")
    routes.HandleResquest()
}COPIAR CÓDIGO
[04:41] Temos um arquivo só para controllers.

package controllers

import (
    "fmt"
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Home Page")
}COPIAR CÓDIGO
[04:43] E um arquivo para manter só as nossas rotas.

package routes

import (
    "log"
    "net/http"

    "github.com/guilhermeonrails/go-rest-api/controllers"
)

func HandleResquest() {
    http.HandleFunc("/", controllers.Home)
    log.Fatal(http.ListenAndServe(":8000", nil))
}COPIAR CÓDIGO
[04:45] Se surgir uma nova rota colocamos qual será o endpoint dessa rota. Só colocamos aqui o endpoint dessa nova rota e depois vincula com o controller que vamos criar e o nosso código ficará mais fácil para manter e organizar.
















[00:00] Vamos evoluir a nossa aplicação e eu vou criar um estrutura para listar personalidades. Na minha cidade existem várias ruas, cujo o nome das pessoas eu não faço ideia de quem são e eu quero criar uma aplicação que cubra isso daí.

[00:14] Eu quero ficar registrando. Na minha cidade existe uma rua chamada Carmela Dutra, vou até abrir aqui para vocês. Se eu vier aqui no buscador do Google e digitar: “Avenida Carmela Dutra” ele vai mostrar que tem uma avenida Carmela Dutra também em Guarulhos. Pesquisando por “Carmela Dutra Mogi das Cruzes”, vai mostrar que em Mogi das Cruzes tem uma rua Carmela Dutra.

[00:44] Só que eu não faço ideia quem foi Carmela Dutra, vamos pesquisar: “Carmela Dutra”, quando eu pesquiso ele fala assim: "Carmela Dutra foi a primeira dama...". Foi a primeira dama do Brasil, casada com o presidente Eurico Gaspar Dutra.

Carmela Dutra: Carmela Teles Leite Dutra foi a primeira-dama do Brasil, de 31 de janeiro de 1946 até a sua morte, tendo sido a esposa de Eurico Gaspar Dutra, 16º Presidente do Brasil. Era, carinhosamente, chamada de Dona Santinha, pela sua forte religiosidade, fazendo seu marido abrir uma capelinha no Palácio Guanabara.

[01:05] Eu queria mostrar justamente isso, uma aplicação que eu mostre o nome da personalidade e mostre um pouco da história dela. Para isso vamos criar uma estrutura, o modelo de estrutura que queremos para que criemos várias personalidades da nossa cidade. Eu vou usar como exemplo nomes da minha cidade, mas você pode usar os nomes da sua cidade também e assim você descobre um pouco mais da história de onde você mora.

[01:29] A primeira que vou fazer vai ser criar uma nova pasta que eu vou chamar de models, essa pasta terá os modelos que eu vou conversar depois com o banco de dados. Vamos aprender mais para frente como isso funciona. Dentro de models eu vou criar um arquivo que eu vou chamar de “personalidades.go”. Eu preciso dizer qual é o pacote, que é o pacote de models, package models.

[01:57] E vou criar aqui uma estrutura de personalidade, eu posso digitar type Personalidade struct, abro e fecho chaves e começo a digitar a minha personalidade. type Personalidade struct { }. Outra coisa legal também para criarmos estruturas de uma forma bem rápida se você estiver usando o VS Code, é digitar apenas tys dar um "Enter" e você coloca Personalidade.

[02:31] Toda personalidade vai ter um nome e vai ter uma história, já começamos assim. O nome vai ser do tipo string, assim como a história também será uma string. Nome string, Historia string. Aqui vem uma sacada legal, para criarmos uma aplicação que seja uma Api Rest conseguimos já identificar a forma que esse nome será exibido na nossa Api Rest, como vamos serializar esse nosso campo.

[03:00] Uma das formas que temos que colocando o Nome string com json:"nome" entre crases. E vou fazer o mesmo no campo História, História strind json:"historia" também com as crases. Quando eu salvo, ele já indenta para ficar bem certinho.

[03:40] Toda personalidade minha vai ter um nome e uma história. Eu vou criar uma variável, eu vou mocar, falamos "mocar" quando não estamos de fato indo lá no banco de dados, estamos criando esses nomes na mão.

[03:51] Vou criar aqui uma variável chamada var Personalidades []Personalidades e vou falar que essa variável personalidades vai ser do tipo uma array de personalidade. Serão várias personalidades dentro dessa nossa variável. Vamos ter várias personalidade dentro da variável personalidades.

package models

type Personalidade struct {
    Nome     string `json:"nome"`
    Historia string `json:"historia"`
}

var Personalidades []PersonalidadeCOPIAR CÓDIGO
[04:09] Espero que vocês entendam. Vamos lá. Eu vou mocar algumas personalidades, como estamos lidando só com exemplo eu vou criar aqui um exemplo, vou colocar aqui mesmo. Antes de subirmos o nosso servidor, depois que iniciamos o servidor Rest eu vou colocar aqui, por exemplo, lá dos meus models eu quero trazer models.Personalidades =[]models.Personalidade {} e vou mocar alguns nomes aqui.

[04:46] Vou colocar aqui o primeiro, toda personalidade eu sei que ela tem um nome. Quando eu coloco nome dois pontos, vou deixar bem simples Nome 1 e, além do nome a personalidade também tem uma história {Nome: "Nome 1", Historia: "Historia 1"}. Vou colocar mais um aqui para termos duas personalidades como exemplo, {Nome: "Nome 2", Historia: "Historia 2"}.

package main

import (
    "fmt"

    "github.com/guilhermeonrails/go-rest-api/models"
    "github.com/guilhermeonrails/go-rest-api/routes"
)

func main() {
    models.Personalidades = []models.Personalidade{
        {Nome: "Nome 1", Historia: "Historia 1"},
        {Nome: "Nome 2", Historia: "Historia 2"},
    }

    fmt.Println("Iniciando o servidor Rest com Go")
    routes.HandleResquest()
}COPIAR CÓDIGO
[05:14] Temos agora essas duas personalidades e o que eu quero fazer na sequência é exibir essas duas personalidades em uma determinada endpoint, em um determinado endereço. Quando a pessoa digitar “localhost8000/api/personalides” eu quero exibir esses dois nomes aqui no formato JSON. Vamos ver como podemos fazer isso.

[05:32] A primeira coisa que vamos fazer para conseguir exibir essas personalidades é criando um determinado controller que é o que exibe todas as personalidades. Vou criar uma função aqui em “controller.go” que vou chamar de func TodasPersonalidades() {}. A função TodasPersonalidades() também recebe como parâmetro o ResponseWriter e Request. func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {}.

[06:02] E uma das formas que temos para conseguir colocar essa personalidade é assim: eu vou chamar do módulo json e aqui é muito importante, existe uma forma para conseguirmos responder usando o ResponseWriter que é o NewEncoder(), não é NewDecoder(), eu quero encodar uma resposta e eu quero encodar nessa resposta, json.NewEncoder(w).Encode(models.Personalides).

[06:40] Tem que colocar personalidades no plural porque se colocar só personalidade vai ser uma só. Já temos o nosso controller, é só isso que precisamos. Eu vou falar que quero ter uma resposta, quero ter um novo Enconder tipo JSON e quero encondar lá dos meus modelos todas as personalidades. Agora falta só criarmos a nossa rota.

package controllers

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/guilhermeonrails/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(models.Personalidades)
}COPIAR CÓDIGO
[06:59] Para criarmos a nossa rota determinamos o caminho, podemos colocar até esse caminho mesmo api/personalidades. Vamos ver. Vou utilizar as teclas "Shift + Ctrl" para baixo para ele copiar, vou colocar http.HandlerFunc(“/api/personalidades” e quem vai atender essa requisição serão todas as personalidades. http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).

package routes

import (
    "log"
    "net/http"

    "github.com/guilhermeonrails/go-rest-api/controllers"
)

func HandleResquest() {
    http.HandleFunc("/", controllers.Home)
    http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
    log.Fatal(http.ListenAndServe(":8000", nil))
}COPIAR CÓDIGO
[07:26] Salvei, tenho aqui as duas personalidades mocada, vou abrir o meu servidor, utilizo as teclas "Ctrl + C" para ele parar, vou rodar mais uma vez, go run main.go. Iniciando o servidor Rest com Go, quando eu atualizo a página no “localhost:8000” temos o home page. Se eu coloco “localhost:8000/api/personalidades” temos lá o "Nome 1" e o "Nome 2" e as histórias que temos. E ele já está devolvendo um formato JSON para nós.

[07:55] Ficou legal com uma resposta JSON e vamos evoluir esse projeto ainda mais nos próximos vídeos.



[00:00] Conseguimos exibir tanto a personalidade 1 quanto a personalidade 2 com esses exemplos que fizemos com o nome mocado, que não estamos pegando no banco de dados e nem nada só exibindo na tela mesmo. Isso ficou legal, só que vamos evoluir um pouco mais a nossa aplicação.

[00:16] Vamos supor que eu coloque aqui em personalidades barra 1, “localhost:8000/api/personalidades/1”, eu quero exibir só ID 1. Se eu fizer isso dessa forma ele não vai encontrar ele vai abrir novamente a página Home Page porque ele não tem esse caminho ainda.

[00:30] Para conseguirmos manipular esse novo endpoint que queremos criar vamos precisar configurar esses nossos valores que fazemos. Em outras palavras, vamos precisar de roteador que vai conseguir pegar essas informações que passamos para os nossos endpoints, como ID 1 ou ID 2.

[00:47] Pensando nisso vamos usar um pacote chamado “Gorilla Mux” e vai abrir esse primeiro link do GitHub mesmo. Ele é uma rota poderosa de HTTP e URL, é justamente o que precisamos. Vamos instalar o pacote gorilla/mux e ver como conseguimos utilizá-lo.

[01:14] Aqui na documentação eu sugiro que você depois dê uma lida passo a passo, vai discorrendo para entender como funciona essa implementação do gorilla/mux. E vamos ver na prática como conseguimos usá-lo.

[01:29] A primeira coisa, para conseguirmos instalá-lo temos esse comando aqui com a configuração correta. Eu vou clicar nesse ícone para copiar ou posso selecionar tudo aqui também, utilizar as teclas "Ctrl + C" para copiar. Vamos incorporar em nosso projeto o gorilla/mux.

[01:45] Para isso eu vou parar o nosso servidor, limpo a tela utilizando as teclas "Ctrl + L", vou utilizar as teclas "Ctrl + C" e "Ctrl + V". go get -u github.com/gorilla/mux, dou um "Enter" e vai instalar esse pacote. Se formos aqui em nosso “go.mod” repare que ele falou que temos o gorilla/mux aqui também no nosso pacote como dependência desse projeto. Isso ficou ótimo, é justamente o que queríamos.

[02:20] Para conseguirmos utilizar agora o gorilla/mux é muito simples, vamos lá na nossa rota e vou criar uma instância dele que vou chamar de r, r := mux.NewRouter(). Ou seja, ele vai criar uma nova rota, um novo mapeamento de rotas.

[03:00] A única coisa que vou fazer vai ser colocar essa instancia que temos aqui na frente: r.http.HandleFunc(), r.http.HandleFunc() e aqui quando subimos o nosso servidor no lugar de deixar o new eu vou passar a nossa instancia. Aqui não vamos precisar do r.http.HandleFunc, vamos usar apenas o r.HandleFunc() do gorilla/mux.

package routes

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/guilhermeonrails/go-rest-api/controllers"
)

func HandleResquest() {
    r := mux.NewRouter()
    r.HandleFunc("/", controllers.Home)
    r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
    log.Fatal(http.ListenAndServe(":8000", r))
}COPIAR CÓDIGO
[03:30] Vamos executar mais uma vez o nosso projeto e limpar o terminal. go run main.go. Se viermos aqui vamos que esse daqui esse daqui está carregando certinho e se eu coloco lá na nossa pasta no “localhost:8000” abre a página home certinho. Tudo está funcionando como queríamos.

[03:50] No lugar de usar o default, o mux default que vem do global estamos utilizando o mux com NewRouter() do gorilla/mux. Lembra que eu falei, agora conseguimos manipular melhor as nossas rotas e as nossas requisições.

[04:05] Na sequência, veremos como fazemos para conseguirmos criar um endpoint para visualizar o ID passado. Se colocamos o id1 aqui queremos exibir só a personalidade 1 ou só a personalidade 2 com id2 e assim por diante.




[00:00] Nessa aula vamos aprender como podemos usar o Mux para conseguir informações da requisição que estamos fazendo para exibir uma personalidade ou outra.

[00:09] Temos aqui nesse endpoint de “localhost:8000/api/personalidades” já exibimos todas e eu quero colocar “localhost:8000/api/personalidades/1” para exibir apenas a primeira personalidade. Só que temos um problema: com a forma que fizemos não temos um identificador único para as personalidades. Vamos criar.

[00:27] Eu vou lá na minha pasta de modelo em “personalidades.go” e vou criar uma identificação para elas. Vou chamar de Id int e passaremos json:"id" entre crases.

package models

type Personalidade struct {
    Id       int    `json:"id"`
    Nome     string `json:"nome"`
    Historia string `json:"historia"`
}

var Personalidades []PersonalidadeCOPIAR CÓDIGO
[00:53] Lá no nosso “main.go” vou colocar uma vírgula antes do nome aqui, se eu dito id já posso colocar Id: 1, Id: 2. Temos o "Nome 1" com id: 1, o "Nome 2" com id: 1 e a "Historia 2". {Id: 1, Nome: "Nome 1", Historia: "Historia 1"}, {Id: 2, "Nome 2": "Nome 2", Historia: "Historia 2"}.

package main

import (
    "fmt"

    "github.com/guilhermeonrails/go-rest-api/models"
    "github.com/guilhermeonrails/go-rest-api/routes"
)

func main() {
    models.Personalidades = []models.Personalidade{
        {Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
        {Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
    }

    fmt.Println("Iniciando o servidor Rest com Go")
    routes.HandleResquest()
}COPIAR CÓDIGO
[01:14] Vou parar aqui o meu servidor para visualizarmos, agora sim iniciando o servidor. Se eu atualizar temos o id: 1 e o id: 2, até aí tudo bem já temos um passo dado. Eu quero colocar /1 e pegar essa informação do barra 1 aqui para ele comparar se esse 1 que passamos aqui é igual a esse id: 1. Se for, exibe só ele.

[01:34] Ou vamos supor que eu coloque o id: 2, “localhost:8000/api/personalidades/2”, esse id: 2 é igual ao id: 1? Não, não é igual. É igual a esse id: 2? É. Exibe só esse aqui. Para fazermos isso de uma forma que fique legal e de uma forma bem simples eu vou fechar essas informações e vou lá no meu “controllers.go” e vou criar uma função específica para fazer isso.

[02:00] Vou criar uma nova função fun RetornaUmaPersonalidade(). Essa função, assim como as outras que temos, vai ter como parâmetro o ResponseWriter e o Request. func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {}.

[02:18] A primeira coisa que vou fazer vai ser usar o Mux para conseguir pegar esse valor que passamos aqui para ele. Vou chamar vars := mux.Vars(). Para essa função funcionar eu tenho que passar para ela qual é a requisição, a Request que estamos recebendo nesse momento e eu passo que a requisição é essa. vars := mux.Vars(r).

[02:44] Já conseguimos pegar todos os valores. Só que eu quero apenas um valor, eu quero o valor do ID. Vou colocar aqui id := vars["id"] só que eu vou identificar, eu quero só o valor que tenha identificação de ID. Tudo o que está em vars eu quero só o ID, só o pedacinho do ID.

[03:08] Vamos continuar. Vou fazer esse for, esse for será assim: eu não quero nenhuma validação, nenhuma tipo de validação desse nosso for range por enquanto, e vou falar for _, personalidade := range models.Personalidades {}.

[03:29] Isso significa que eu vou pegar uma personalidade de cada vez e vou ficar comparando: esse ID que temos aqui é igual a esse ID. Se for exibe essa pessoa. Vai ficar fazendo isso. Vou comparar se essa if personalidade.Id == id for igual ao Id que temos aqui, que trouxemos lá do nosso parâmetro o de cima, se eles forem iguais eu quero exibi-lo com json.NewEncoder(w).Econde(personalidade) } não é Decoder.

[04:12] Aqui temos uma mensagem de erro, um alerta que fala assim: "você não pode comparar o personalidade.Id com Id porque eles são de tipos diferentes, um é int e o outro é string". Faz sentido, tudo que passamos aqui no nosso endpoint como valor é do tipo texto, é do tipo string. E o que temos aqui na personalidade.Id que criamos lá no nosso modelo é do tipo int, não podemos comparar inteiro com string.

[04:40] Para resolver esse problema de uma forma muito simples com um pedaço de código é pegarmos esse personalidade.Id que temos do tipo int e converter para string só para ele fazer essa comparação. Vou colocar entre parênteses e vou usar a função if strconv.Itoa(personalidade.Id) == id queremos de inteiro para string para async.

package controllers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/guilhermeonrails/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(models.Personalidades)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for _, personalidade := range models.Personalidades {
        if strconv.Itoa(personalidade.Id) == id {
            json.NewEncoder(w).Encode(personalidade)
        }
    }
}COPIAR CÓDIGO
[05:05] Só ao fazer isso ele já não exibe mais aquela mensagem de erro. E a nossa função que retorna uma personalidade fica assim. A última coisa que falta é criarmos esse endpoint, falar que todas as vezes que bater uma requisição com “localhost:8000/api/personalides/” e um determinado ID eu quero que essa função aqui que retorna uma personalidade seja invocada.

[05:28] Para isso vamos lá no nossa arquivo de rotas, “routes.go”, e eu vou copiar e colocar essa linha aqui utilizando as teclas "Ctrl + C" e "Ctrl + V" aqui para baixo e toda as vezes que chegar uma requisição para api/personalidades/ um determinado valor que vou identificar por ID eu não quero que o controle de todas as personalidades seja atendido eu quero só o que retorna uma personalidades. r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).

package routes

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/guilhermeonrails/go-rest-api/controllers"
)

func HandleResquest() {
    r := mux.NewRouter()
    r.HandleFunc("/", controllers.Home)
    r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
    r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade)
    log.Fatal(http.ListenAndServe(":8000", r))
}COPIAR CÓDIGO
[05:52] Salvei isso aqui, vou reiniciar o nosso servidor go run main.go, iniciou o servidor. Vamos ver aqui exibindo todas as personalidades, se eu colocar /1 só temos a personalidade 1 com o "Nome 1" e a "Historia 1". /2, Id: 2, "Nome 2", "Historia 2".

[06:11] Para finalizarmos essa aula e deixá-la ainda mais clara, todas as vezes que queremos pegar uma informação, um determinado recurso que falamos estamos utilizando uma requisição que chamamos de "Get". E tem uma forma de especificarmos que uma requisição "Get" para personalidade vai retornar todas as personalidades. Uma requisição "Get" para personalidades com um determinado ID, ele quer aquele recurso, vamos também utilizar a requisição .Methods("Get").

[06:38] Para isso vamos colocar aqui r.HandleFunc("/api/personalidades", controllers.TodasPersonalidade).Methods("Get"). Vamos fazer a mesma aqui para baixo, vamos usar r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get").

package routes

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/guilhermeonrails/go-rest-api/controllers"
)

func HandleResquest() {
    r := mux.NewRouter()
    r.HandleFunc("/", controllers.Home)
    r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
    r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
    log.Fatal(http.ListenAndServe(":8000", r))
}COPIAR CÓDIGO
[06:53] O "Get" é simples, a ideia é a seguinte: eu quero um recurso, uma determinada informação, qual informação? Eu quero todas as personalidades. Vamos utilizar o método "Get". Isso quem faz para nós é próprio Google também, o nosso próprio navegador faz uma requisição "Get". Ele fala que precisa de todas as personalidades, requisição "Get" e ele exibe todas as personalidades. Eu preciso de uma personalidade que tenha o id: 1, aí ele vai devolver só esse id: 1.

[07:23] Em ambos os casos utilizamos por padrão essa nomenclatura, essa requisição "Get". A ação que queremos fazer é buscar um determinado recurso.