# Accounts Merge

Esse é um problem de grafos.

## Por quê é um problema de grafos?

Se vemos uma conta com emails A e B, sabemos que A e B são dessa pessoa. se depois vemos outra conta com emails B e C, sabemos que A, B e C são todos da mesma pessoa, mesmo que A e C nunca tenham aparecido juntos diretamente.

Se pensarmos em conexões, O email A está conectado ao email B, pois aparecem juntos. O email B está conectado ao C. Então existe um caminho de A até C, passando por B.

Todos esses e-mails formam um grupo coenctado, um **componente conexo**.

Isso é exatamente o que um grafo representa: entidades que têm conexões entre si.

-> Os emails são o nó do grafo, e as conexões (mesma conta) são as arestas

Olha o input do leetcode:

```
[
    ["John","johnsmith@mail.com","john_newyork@mail.com"],
    ["John","johnsmith@mail.com","john00@mail.com"],
    ["Mary","mary@mail.com"],
    ["John","johnnybravo@mail.com"],
]
```

Podemos ver que, o primeiro e segundo John, tem um primeiro e-mail em comum. E é aqui que podemos fazer a componente conexa, a partir desse primeiro email em comum, fazemos as ligações.

### Visualizando isso

```
Conta 1: John -> [a@mail.com, b@mail.com]
Conta 2: John -> [b@mail.com, c@mail.com]
Conta 3: John -> [d@mail.com]
```

Desenhando as conexões. Quando dois emails aparecem na mesma conta, eu desenho uma linha conectando eles:

```
a d <--> b    (aparecem juntos na conta 1)
b d <--> c    (aparecem juntos na conta 2)
d         (aparece sozinho na conta 3)
```

Ao observar o desenho. Conseguimos ver que existem dois grupos separados.

O grupos {a,b,c} está todo conectado entre si (conseguimos ir de qualquer email para qualquer outro seguindo as conexões). Mas o email `d` está isolado, sozinho

Em teoria dos grafos, esses **grupos separados são chamados de componentes conexas**.

Uma componente conexa é um conjunto de nós onde você consegue ir de qualquer nó para qualquer outro nó seguindo as arestas, mas esse conjunto está desconectado de outros nós do grafo.

### Por que componente conexa vai ser a resposta?

Cada componente conexa **representa uma pessoa única**.

Pensa assim: se você conseguir ir do e-mail A até o e-mail B, seguindo conexões (mesmo que indiretamente), então A e B, pertencem à mesma pessoa, pois cada conexão é uma pessoa em si.

Se não existem um caminho possível de uma pessoa até outra, então pertencem a pessoas diferentes.

### Como o DFS encontra essas componentes

O DFS é uma técnica para explorar um grafo.

Imagina que está em uma sala escura com vários quartos conectados por portas, e você tem uma lanterna. Você começa em um quarto qualquer e vai explorando: entra por uma porta, depois entra por outra porta daquele quarto, e continua indo o mais fundo possível antes de voltar atrás para explorar outros caminhos.

No nosso caso, o DFS funciona:

1. Constrói o grafo
2. Para cada conta, conecta o primeiro email com todos os outros emails da mesma conta
    Então [John, a@mail.com, b@mail.com, c@mail.com], você cria as conexões a <-> b, a <-> c, b <-> a, c <-> a

    Logo, o **primeiro email é um hub central**.
3. Percorre todos os emails. Para cada email que ainda não foi visitado, você começa um DFS a partir dele.
    - O DFS explora todos os emails que estão conectados a esse email inicial, marcando cada um como visitado e adicionando em uma lista.
    - Quando o DFS termina, essa lsita contém todos os emails de uma componente conexa - ou seja, todos os emails de uma pessoa

Exemplo passo a passo:

```
Conta 1: John -> [a, b]
Conta 2: John -> [b, c]  
Conta 3: John -> [d]
```

Primeiro construimos o grafo:

```
a -> [b]
b -> [a, c]
c -> [b]
d -> []
```

Depois de construir o grafo, começamos a procurar as componentes:

- Pegamos o primeiro e-mail não visitado, vamos supor que seja a. Fazemos DFS a partir de a:

1. Marcamos a como visitado, adicionamos na lista atual que fica `[a]`
    a. Olhamos os vizinhos de a, que é `b`, b ainda não foi visitado, fazemos DFS nele.
2. Agora em `b`, marcamos como visitado, adicionamos na lista que fica [a,b].
    a. Olhamos os vizinhos de b, que são [a,c].
        O a já foi visitado, pulamos.
        O c ainda não, fazemos DFS nele

## Código

### Parte 1: Construindo o Grafo

Nessa primeira parte, transformamos a lsita de contas em um granfo onde os **emails são nós** e as conexões representam esses emails aparecendo juntos em alguma conta.

Pense no grafo como um mapa de amizades. Se dois emails aparecem na mesma conta eles são amigos. Precisamos construir esse mapa de amizades para depois descobrir quem conhece quem, mesmo que indiretamente.

Input:

```go
Conta 0: John, a@mail.com, b@mail.com
Conta 1: John, b@mail.com, c@mail.com
```

Código:

```go
graph := make(map[string][]string)
emailToName := make(map[string]string)

for _, acc := range accounts {
    name := acc[0]
    firstEmail := acc[1]
    
    for i := 1; i < len(acc); i++ {
        email := acc[i]
        emailToName[email] = name
        
        graph[firstEmail] = append(graph[firstEmail], email)
        graph[email] = append(graph[email], firstEmail)
    }
}
```



**Estado final:**
```
graph = {
    "a@mail.com": ["a@mail.com", "a@mail.com", "b@mail.com"],
    "b@mail.com": ["a@mail.com", "b@mail.com", "b@mail.com", "c@mail.com"],
    "c@mail.com": ["b@mail.com"]
}

emailToName = {
    "a@mail.com": "John",
    "b@mail.com": "John",
    "c@mail.com": "John"
}