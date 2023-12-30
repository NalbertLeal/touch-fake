# Touch Fake

![GitHub](https://img.shields.io/github/license/NalbertLeal/touch-fake)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/NalbertLeal/touch-fake)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/NalbertLeal/touch-fake)

## Descrição

_touch-fake_ é uma ferramenta escrita em Go que permite criar facilmente arquivos falsos com um tamanho específico em bytes. Isso pode ser útil para testes de desempenho, simulações e outros casos de uso onde é necessário um arquivo de determinado tamanho, mas sem a necessidade de conteúdo real. O nome _touch-fake_ foi pensado tomando como base o comando _touch_ utilizado no linux para criar arquivos.

## Como Usar

### Pré-requisitos

Certifique-se de ter o Go instalado em sua máquina. Para mais informações sobre como instalar o Go, visite [https://golang.org/doc/install](https://golang.org/doc/install).

### Instalação

```bash
go get -u github.com/NalbertLeal/touch-fake
```

### Uso Básico

Isso criará um arquivo chamado `arquivo_falso.txt` com um tamanho de 1024 bytes:

```bash
touch-fake arquivo_falso.txt 1024 
```

Já isso cria um arquivo chamado `arquivo_falso.txt` com um tamanho de 4096 bytes e sobre-escreve o arquivo (caso ele já exista):

```bash
touch-fake arquivo_falso.txt 1024 -overwrite
```

### Opções

- `-overwrite`: Sobrescrever o arquivo de saída se ele já existir (opcional).

## Contribuição

Se você quiser contribuir com melhorias, correções de bugs ou novos recursos, sinta-se à vontade para abrir um problema ou enviar uma solicitação pull. Sua contribuição é bem-vinda!

## Licença

Este projeto é licenciado sob a [GNU LESSER GENERAL PUBLIC LICENSE](https://choosealicense.com/licenses/lgpl-3.0/).

---