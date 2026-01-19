Importante:

Nesse problema, **ORDENA POR FIM**. Pois temos várias reuniões conflitantes.

Para encaixar o máximo possível, sempre deve escolher a que TERMINA PRIMEIRO, pois deixa mais tempo livre para as próximas.

Input: [[1,100],[11,22],[1,11],[2,12]]

```
Sorted: [[1,11],[2,12], [11,22] ,[1,100]]
         end=11  end=22  end=12  end=100
```

Aqui, vamos armazenar o **ultimo end**. Se tiver overlap, aumentamos a quantidade a remover.

Ficando:

```
lastEnd = 11 (pegamos [1,11])
toRemove = 0

i=1: [2,12]
  2 < 11? SIM → overlap! toRemove = 1
  (mantemos [1,11] porque termina antes)
  lastEnd continua 11

i=2: [11,22]
  11 >= 11? SIM → sem overlap → lastEnd = 22

i=3: [1,100]
  1 < 22? SIM → overlap! toRemove = 2
  lastEnd continua 22

Resultado: 2
```