# qualpraia
Qual praia? Informe direção do vento e swell para descobrir o melhor pico para surfar

## Testing:
`curl "localhost:8000/spots?wind=s&swell=l"`  
`curl "http://meianuca-go.herokuapp.com/spots?wind=n&swell=l"`

### Commands:

#### Build:
`go build -o bin/qualpraia *.go`
#### Build and run:
`go build -o bin/qualpraia && ./bin/qualpraia`
#### Heroku:
Push changes to Heroku:  
`heroku push heroku master`  
Check logs:  
`heroku logs -- tail`  
Open app:  
`heroku open`  
Testing:
`curl "http://meianuca-go.herokuapp.com/spots?wind=n&swell=l"`


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
