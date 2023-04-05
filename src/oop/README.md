# Orientação a objetos em Go
> Structs, não objetos

## Go é orientado a objetos
- Mas de forma diferente, com outras características.

## Métodos
```
func (obj *tipoObj) nome (params tipoParams) (retornos) {
  corpo da função
}
```

- O obj se comporta como o ponteiro "this" em C++
- Para chamar o método: ```objeto.nomeDaFuncao(params)```

## Encapsulamento
- É feito em nível de pacote. Cada pacote empaota um tipo e seus métodos
- Letra maiúscula é exportada, letra minúscula não.

## Herança
- Não há implementação em GoLang

## Polimorfismo
- Não há implementação em GoLang
- Há o conteito de interfaces.

### Interfaces
- Definição de um ou mais conjuntos de métodos comuns a um ou mais tipos.
- Permite uma certa criação de polimorfismo.

## Funções genéricas
- Uso de interfaces
- Dentro da interface, você declara todos os tipos que sua função pode aceitar

