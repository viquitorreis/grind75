Nesse problema, precisamos ordenar por **START CRESCENTE**, porém, se os starts forem iguais, precisamos ordenar por **END DECRESCENTE**.

- Exemplo SEM ORDENAR POR END (quando starts começam iguais):
```
Sorted apenas por start: [[1,2],[1,4]]

maxEnd = 0

[1,2]: end=2 <= 0? NÃO → maxEnd = 2
[1,4]: end=4 <= 2? NÃO → maxEnd = 4

Resultado: 2 intervalos sobrevivem
```

Isso está ERRADO, pois [1,2] está coberto por [1,4], mas processou [1,2] primeiro e foi configurado maxEnd=2. Quando chegou em [1,4] ao comparar 4 <= 2 (falso) achou que [1,4] não estava coberto.


- Cenário ordenando POR end (quando starts começam iguais):

```
Sorted: [[1,4],[1,2]]  ← [1,4] vem primeiro agora

maxEnd = 0

[1,4]: end=4 <= 0? NÃO → maxEnd = 4
[1,2]: end=2 <= 4? SIM → coberto!

Resultado: 1 intervalo sobrevive ([1,4])
```