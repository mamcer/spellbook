# mysql

[naming standards](https://www.sqlstyle.guide/)  
[https://www.sqlstyle.guide/](https://www.sqlstyle.guide/)  

## docker

    docker pull mysql:5.7.30

run

    docker run -p 3306:3306 --name cook-me -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7.30

    docker exec -it cook-me mysql -uroot -p
    create database cook-me;
    
## bare metal

[https://dev.mysql.com/doc/mysql-apt-repo-quick-guide/en/](https://dev.mysql.com/doc/mysql-apt-repo-quick-guide/en/)

```bash
wget https://repo.mysql.com//mysql-apt-config_0.8.34-1_all.deb
sudo dpkg -i mysql-apt-config_0.8.34-1_all.deb
```

> configure

```bash
sudo apt update
sudo apt install mysql-server -y
systemctl status mysql
```

connect 

```bash
sudo mysql -u root -p
```
> default password: root

## create user

only local access

```sql
CREATE USER 'mario'@'localhost' IDENTIFIED BY '[password]';
GRANT ALL PRIVILEGES ON *.* TO 'mario'@'localhost';    
FLUSH PRIVILEGES;
```

local and remote access

```sql
CREATE USER 'mario'@'%' IDENTIFIED BY '[password]';
GRANT ALL PRIVILEGES ON *.* TO 'mario'@'%';
FLUSH PRIVILEGES;
```

> if you by mistake create a localhost user and need remote access you can create the user with remote access and then delete de local user: DROP USER 'mario'@'localhost';

check access

```sql
USE mysql;
SELECT User, Host FROM user WHERE User = '[user-name]';
```

if you still present an error:

```bash
mycli -h 192.168.100.100 -u mario -P 3306
(2003, "Can't connect to MySQL server on '192.168.100.100' ([Errno 111] Connection refused)")
```

edit:

```bash
sudo vim /etc/mysql/mysql.conf.d/mysqld.cnf
```

look for `bind-address = 127.0.0.1`

and change it to `bind-address = 0.0.0.0`

then restart mysql

```bash
sudo systemctl restart mysql
```

## status

```bash
systemctl status mysql
```

## connect

```bash
mycli -h localhost -P 3306 -u root
```

## backup

```bash
docker exec CONTAINER /usr/bin/mysqldump -u root --password=root DATABASE > backup.sql
```

backup with compression of a remote server

```bash
mysqldump -h yourserver.mysql.database.azure.com \
  -u dbuser@yourserver -p \
  --single-transaction --quick \
  --routines --triggers --events \
  database_name \
  | gzip > database_name_$(date +%F).sql.gz
```

## restore

```bash
cat backup.sql | docker exec -i CONTAINER /usr/bin/mysql -u root --password=root DATABASE
```

## cli

```bash
mysql --host=localhost --user=myname --password=password mydb
mysql -h localhost -u myname -ppassword mydb
```

## troubleshooting

```bash
show status where `variable_name` = 'Threads_connected';

show processlist;
```

## size of a database

```sql
SELECT table_schema "DB Name",
            ROUND(SUM(data_length + index_length) / 1024 / 1024, 1) "DB Size in MB" 
FROM information_schema.tables 
GROUP BY table_schema;
```

## version

```sql
SELECT version()
```

## naming conventions (db)

### Do

- Stick to lowercase → avoids quoting issues. Example: shopapp_prod
- Use underscores, not dashes → dashes force backticks. Example: shopapp_stage not shopapp-stage
- Keep names short but clear → prod instead of production.
- Follow one consistent pattern → either project_env or env_project.
- Use only letters, numbers, underscores → portable across tools.
- Document the convention → pin it in your repo’s README or wiki.

### Don’t

- Don’t mix styles (shopAppDev, shopapp_stage, shopapp-Prod).
- Don’t reuse database names across environments.
- Don’t use reserved words (test, order, select) as names.
- Don’t rename databases casually → migrations and configs will break.
- Don’t put secrets or versions in the DB name (shopapp_prod_v2, shopapp_pwd123).

### Example Convention

Pattern: project_env

- shopapp_dev
- shopapp_stage
- shopapp_prod

Optional extension (for multi-region):

- shopapp_prod_us
- shopapp_prod_eu