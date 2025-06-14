# BookCabin Assesment

```
docker run --name mysql-bookcabin \
  -e MYSQL_ROOT_PASSWORD=12345678 \
  -e MYSQL_DATABASE=bookcabin \
  -e MYSQL_USER=bookcabin \
  -e MYSQL_PASSWORD=12345678 \
  -p 3306:3306 \
  -d mysql:8.0

```