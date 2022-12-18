# test-golang

# Run docker

## Run docker

```
 $docker-compose up -d
```

## If you want to look the logs you can run with this command

```
docker-compose up -d && docker-compose logs --tail=0 -f api
```

# Migrate table

## Run Docker and run command below

```
$docker-compose exec api bee migrate -driver="postgres" -conn="postgres://postgres:mysecret@postgres/test-golang-development?sslmode=disable"
```

## After Migrate you can see postgres on your localhost with open this url

http://localhost:1234

1. system: postgres
2. server: postgres
3. username: postgres
4. password: mysecret
5. database: test-golang-development

![Alt text](/design/test-golang1.png)
