docker run --rm -p 5432:5432 -d\
  --name postgres-test\
  -e POSTGRES_USER=postgres\
  -e POSTGRES_PASSWORD=postgres\
  -e POSTGRES_DB=postgres\
  -v /home/shadow/Public/pg-test:/var/lib/postgresql/data\
  postgres:alpine3.19
