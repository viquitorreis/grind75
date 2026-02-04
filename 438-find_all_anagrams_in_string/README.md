# Find All Anagrams in String

O problema quer que encontremos todas as substrings de tamanho fixo `len(p)`em `s`que são anagramas de `p`.

Anagrams tem as mesmas letras com a mesma quantidade. Ex: `abc` e `cba`. Ou seja, encontrar um anagrama significa que uma substring tem exatamente as mesmas letras que `p` na mesma quantidade. Não importa a ordem.

Esse problema, podemos resolver com sliding window:

    - Uma janela de tamanho fixo desliza pela string e verificamos alguma propriedade a cada posição. Aqui a propriedade é a *frequência* de letras em p*.

A solução usa um **contador** `count` que rastreia quanto falta ou sobra de cada letra. A lógica é:

- Valores negativos em `count` significam "ainda precisa de mais letras para formar anagrama".
- Valores positivos em `count` significam "tenho mais dessa letra do que preciso".
- Valores zero em `count` significam "tenho exatamente a quantidade certa"

---

## 1. Setup inicial

Aqui começámos marcando o que precisamos da string `p`:

```go
cnt := [26]int{}
for _, ch := range p {
    cnt[ch-'a']--
}
```

Para `p = "abc"`, depois desse loop teremos:

```
cnt['a'] = -1  (preciso de 1 'a')
cnt['b'] = -1  (preciso de 1 'b')
cnt['c'] = -1  (preciso de 1 'c')
cnt[todas as outras] = 0
```

Esses valores negativos representam uma espécie de "dívida" que precisamos quitar, zerar. Conforme processamos a string `s`, vamos "pagando" essa dívida adicionando letras na janela.

---

## 2. Por que o ponteiro right se move com o loop em s?

O ponteiro right representa o lado direito da janela que desliza (sliding window).

Ele sempre avança com o loop em s, pois estamos processando a string `s` da esquerda para a direita, expanding a janela conforme vamos.

```go
for right, ch := range s {
    x := ch - 'a'
    cnt[x]++
    // ...
}
```

A cada iteração, estamos: "adicionando essa nova letra 'ch' na janela de count (diminuindo a dívida).

Se a dívida ficar zerada, temos a quantidade necessária para aquele caractere.

## Porque quando count[x] > 0 precisamos mover para left?

O left determina onde um anagrama válido se inicia.

Quando `cnt[x] > 0`, significa que temos MAIS da letra `x` do que precisamos. Isso invalida a janela atual como um possível anagrama... Ou seja, não tem a quantidade de letras necessárias na janela atual para formar um anagrama, portanto precisamos mover left para frente, que é o que determina onde o anagrama se inicia.

Exemplo `s = "cbba"` e `p = "abc"`:

Começamos com cnt['a']=-1, cnt['b']=-1, cnt['c']=-1

```
right=0, ch='c': cnt['c']++ -> cnt['c']=0
  janela = "c"
  Tudo ok, não tem nada positivo

right=1, ch='b': cnt['b']++ -> cnt['b']=0
  janela = "cb"
  Tudo ok ainda

right=2, ch='b': cnt['b']++ -> cnt['b']=1  <- PROBLEMA!
  janela = "cbb"
  Temos 2 'b' mas precisamos de apenas 1
```

Quando isso acontece, não adianta mais expandir a janela pela direita. Precisamos **encolher pela esquerda até a janela ficar válida de novo.** É aqui que entra o loop interno:

```go
for cnt[x] > 0 {
    cnt[s[left]-'a']--
    left++
}
```

Esse loop vai:

1. Retornar a dívida (para que tentemos pagar ela em outra janela)
2. Mover a janela para frente, para tentarmos com outra combinação.

Continuando o exemplo:

```
cnt['b'] = 1 (temos 1 'b' a mais)

Iteração 1 do loop interno:
  Remove s[left] = s[0] = 'c' da janela
  cnt['c']--         -> cnt['c'] = -1 (agora falta 'c')
  left++             -> left=1
  janela = "bb"
  cnt['b'] ainda é 1, continua o loop

Iteração 2 do loop interno:
  Remove s[left] = s[1] = 'b' da janela
  cnt['b']--            -> cnt['b'] = 0 (agora tá certo)
  left++                -> left=2
  janela = "b"
  cnt['b'] agora é 0, sai do loop
```

## Por que right - left + 1 == pLen prova que é anagrama?

Depois de todo o processo que fizemos antes, temos algumas garantias:

1. Nenhum 'cnt[x]' é positivo (foi garantido pelo loop interno)
2. A janela tem tamanho exato de `len(p)` (quando `right - left + 1 == pLen`)

Se essas duas condições são verdadeiras, então a janela é um anagrama.

Pois:

- A cada letra que adicionamos na janela, aumentamos o contador. O loop interno garante que nunca deixamos nenhum contador ficar positivo. Então no momento em que a jenala tem o tamanho exato de `p` e nenhum contador está positivo, a única possibilidade é que todos os contadores estejam em zero.
- Se todos estão em zero, significa que pagamos todas as dívida exatamente. Logo, a janela tem exatamente as mesmas letras em `p`.

