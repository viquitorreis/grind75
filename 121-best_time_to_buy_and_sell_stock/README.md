# Best time to Buy and Sell Stock

Quando o código faz ```buy = sell```, está dizendo "Eu acabei de encontrar um preço mais baixo do que meu preço de compra atual. Daqui para frente, qualquer venda que eu considerar, vai ter mais lucro do se comprar nesse novo preço mais baixo ao invés do preço antigo mais alto".

Por isso atualizamos o ponteiro de compra para a venda atual. Quando achamos um valor de venda menor do que o de compra, é CERTEZA que esse é o menor dia de compra, daqui para a frente.