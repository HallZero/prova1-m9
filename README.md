# prova1-m9

Rode o broker local usando o mosquito:
```bash
cd config
mosquitto -c mosquitto.cong
```

Rode o publisher:
```bash
cd publisher
go run .
```

Rode o subscriber:
```bash
cd subscriber
go run .
```

## Testes

Para cada diretório individual, abra e rode o comando:
```bash
go test -v
```
