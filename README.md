# test-golang

# Run docker

## Run docker

## $docker-compose up -d

## If you want to look the logs you can run with this command

## docker-compose up -d && docker-compose logs --tail=0 -f api

# Migrate table

## Run Docker and run command below

## $docker-compose exec api bee migrate -driver="postgres" -conn="postgres://postgres:mysecret@postgres/test-golang-development?sslmode=disable"

## After Migrate you can see postgres on your localhost with open this url

### http://localhost:1234

### system: postgres

### server: postgres

### username: postgres

### password: mysecret

### database: test-golang-development
