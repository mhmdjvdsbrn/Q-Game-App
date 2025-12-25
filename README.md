# Q-Game-App



# Databases
``migrations command``
```bash
go install github.com/rubenv/sql-migrate/...@latest
 
sql-migrate status -env="development" -config=repository/mysql/dbconfig.yml

sql-migrate up -env="development" -config=repository/mysql/dbconfig.yml 
 
sql-migrate down -env="development" -config=repository/mysql/dbconfig.yml -limit=1
```