# qualpraia
Qual praia? Informe direção do vento e swell para descobrir o melhor pico para surfar

## Como testar:
`$ curl "localhost:8000/spots?wind=s&swell=l"`

### Comandos:
#### Build:
`go build -o bin/qualpraia *.go`
#### Build and run:
`go build -o bin/qualpraia && ./bin/qualpraia`

#### Project organization
```
meianuca-go/
  main.go
  routes.go
  models.go
  qualpraia/
    qualpraia.go
  bin/
    meianuca
```