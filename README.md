# Secrets creator

Converta variaveis para base64 com um comando.

# Instruções

## Build

```sh
go build -ldflags='-s -w -extldflags "-static"' -o "create-secret"
```

## Uso

- Use o comando `create-secret` para abrir o editor.
- Com o editor aberto, cole suas secrets no seguinte formato:    
    ```yaml
    CHAVE1: valor1
    CHAVE2: valor2
    CHAVE3: 90 
    ```
- Feche o editor.
- Suas secrets vão aparecer no terminal prontas para serem copiadas.
