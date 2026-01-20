## Conceito de interseção:

Dois intervalos [a,b] e [c,d] têm interseção se eles se sobrepõem. A interseção é:

- start = max(a, c) (começo mais tarde)
- end = min(b, d) (termina mais cedo)

```
[1,  3]
    [2,    6]
    [2,3]  ← interseção
```

Podemos usar **two pointers** para resolver esse problema. Um ponteiro para cada lista.

1. Ambos pointers começam no inicio dos intervalos
2. Para cada iteração, achamos o start e o end.
    Start = min(ptrA[0], ptrB[0])
    End = max(ptrA[1], ptrB[1])
3. Se start <= end, existe intersecção, colocamos nos intervalos de resultado
4. Avança o ponteiro do intervalo que termina primeiro