# Solução Longest Valid Palindrome

Precisamos descobrir o maior palindromo válido, rearranjando as letras disponíveis.

Num primeiro momento, o problema parece complicado de se resolver, e até mesmo não muito viável em tempo hábil. Porém, se soubermos pontos chaves, podemos resolver muito fácilmente.

## Counting

Para resolver usando essa solução (há outras), precisamos entender o seguinte:

1. Um palíndromo válido tem **NO MÁXIMO** um caractere que aparece uma quantidade ímpar de vezes.
2. Todo o resto de caracteres aparecem uma quantidade par.

Logo:

1. Podemos atravessar a string s
2. Contar o número de ocorrências de cada caractere
3. Armazenar em um array ou hash table / map (count)
4. Atravessar o map / array count, e para cada valor v, dividimos v por 2
    Na divisão pegamos a parte inteira e multiplicamos por 2, e adicionamos na resposta
5. Finalmente, se a resposta for menor do que o tamanho da string s, incrementamos a resposta por um, e retornamos a resposta

## Exemplo:

```s = "aabbcc"```

Cada string aparece 2 vezes:

- a: 2
- b: 2
- c: 2

Um exemplo de palindromo válido nessa palavra seria:

```cabbac```

Nesse sentido:

```
c  a  b  b  a  c
↑              ↑  (primeira = última)
   ↑        ↑     (segunda = penúltima)
      ↑  ↑        (terceira = antepenúltima)
```

Nesse sentido, as letras precisam se **espelhar**.

Logo, para cada letra `x` no começo, precisamos de outra letra `x` no final.

Se temos 4 letras 'a', conseguimos colocar:

- 2 letras 'a' na primeira metade do palíndromo
- 2 letras 'a' na segunda metade do palíndromo (espelhando elas)

Logo:

Se tivermos **5 letras 'a'** conseguimos usar 4 no palindromo para se espelharem.

Exceto uma no meio. Por isso o palindromo pode ter apenas uma letra ímpar, se forem duas, não tem como ter duas no meio.

Portanto:

```
count['a'] = 5
5 / 2 = 2 (parte inteira da divisão)
2 * 2 = 4  <- Exatamente o número que queremos / podemos usar.
```

## Resumo do algoritmo:

- Conta ocorrências de cada caractere (use um map)
- Para cada count, pega a parte par: (count / 2) * 2
- Se sobrou alguma letra ímpar (qualquer uma), adiciona +1 no final
