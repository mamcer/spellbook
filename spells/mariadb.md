# mariadb

after installation

```bash
sudo mariadb -u root 
```

queries

```sql
SHOW DATABASES;
CREATE database cookbook;
CREATE OR REPLACE USER cookbook@localhost IDENTIFIED BY 'cookbook';
GRANT ALL ON cookbook.* TO cookbook@localhost IDENTIFIED BY 'cookbook';
FLUSH PRIVILEGES;
```

next logins

```bash
mariadb -u root -p
```

## error remote connection refused ((2003, "Can't connect to MySQL server on '192.168.100.100' ([Errno 111] Connection refused)"))

edit 

```bash
sudo vim /etc/mysql/mariadb.conf.d/50-server.cnf
```
change `bind-address` from `127.0.0.1` to:

```
bind-address = 0.0.0.0
```

