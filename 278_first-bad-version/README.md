# Explicação 278

Quando encontramos uma versão ruim (bad version), todas as da frente serão bad version, por herança. Apesar disso, pode ser que tenha alguma anterior que  era bad version.

Para entender esse algoritmo, precisamos entender alguns pontos:

1. Se ao checarmos o mid, a versão não for bad, ou seja **false**, significa que podemos ignorar toda a primeira metade.
2. Ao acharmos uma bad version, pode ser que existe alguma anterior que também era bad version.

Durante todo esse algoritmo, temos uma **invariante** que não muda. A **primeira bad version** está SEMPRE entre o intervalo de left e right.

```
Versões: [1, 2, 3, 4, 5, 6, 7, 8]
          G  G  G  G  G  B  B  B
```

A primeira bad version seria o 6.

- Início:

```
left=1, right=8
Invariante: primeira bad está em [1,8] - (6 está entre 1 e 8)
```

- Iteração 1:

```
mid = 4
isBadVersion(4) = false
```

Se esa é boa, todas anteriores também são. Precisamos buscar na segunda parte.

Logo:

```
left = mid + 1 = 5
```

- Iteração 2:

```
left=5, right=8
Invariante: primeira bad está em [5,8]? - (6 está entre 5 e 8)
```

Ficando:

```
mid = 6
isBadVersion(6) = true
```

Aqui. Encontramos uma bad version. **Pode ser** que tenhamos outras anteriores a essa.

Então precisamos buscar antes, para isso diminuimos a janela do ponteiro extremo a direita.

```
right = mid = 6 
```

- iteração 3:

```
left=5, right=6
Invariante: primeira bad está em [5,6]? Sim.
```

Nesse caso:

```
mid = 5
isBadVersion(5) = false
```

Como é boa, iremos atualizar o ponteiro da esquerda para buscar no lado direito:

```
left = mid + 1 = 6
```

Aqui, quebra o loop.

Retornamo so left, que será o valor correto.
