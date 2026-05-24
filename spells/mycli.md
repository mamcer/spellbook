# mycli

## install

```bash
sudo apt install mycli
```

## connect to database

```bash
mycli -h [host] -u [user] -D [database] -P [port]
```

## connect and execute query

```bash
mycli -te "select * from [table-name] limit 10" mysql://[user[:password]@][host][:port][/dbname] 
```

## dsn alias

```bash
vim ~/.myclirc
```

Add a line in the alias_dsn section

```bash
[alias_dsn]
example_dsn = mysql://[user[:password]@][host][:port][/dbname]
```

You can list your dsn alias

```bash
mycli --list-dsn
```

And then connect to a dsn alias

```bash
mycli -d [dsn-alias-name]
```

## special commands

List databases

```bash
\l
```

Change database

```bash
\u
```

List tables in the current database

```bash
\dt
```

Show table schema

```bash
\dt+ [table-name]
```

## output formatting

Show results vertically

```
[sql-query] \G
```

Change output format. Run:

```
\T
```

To list all available options (default psql). To change format to csv for example run:

```
\T csv
```

## quick csv output

```bash
mycli --csv -e "SELECT * FROM employees LIMIT 10" mysql://root@localhost:3306/employees
```

## favourites queries


```
\f                  # List all favorite queries
\f <name>           # Invoke a specific favorite query
\fs <name> <query>  # Save a favorite query
\fd <name>          # Delete a favorite query
```

## system

Access shell commands

```bash
system ls
system cat myfile.txt
system clear
```

## import SQL file

```
\. filename
```

## write output to file

Export all the terminal output to a file

```bash
tee [-o] filename
# execute query
notee
```

## export only the next query 

```
\o [-o] filename
```

## script example 

```bash
#!/bin/bash

# ask for db username
echo "username?:"
read username

# ask for db password
echo "password?:"
read -s password

# ask for year
echo "year?:"
read year

# ask for month
echo "month?:"
read month

# ask for country
echo "country currency? (MXN, BRL, COP, CLP):"
read country

echo " "
echo "username: $username"
echo "year: $year"
echo "month: $month"
echo "country: $country"
echo " "

echo "processing purchases..."
mycli --csv -e "SELECT * FROM table WHERE YEAR(date_created) = $year AND MONTH(date_created) = $month AND currency_from = '$country' ORDER BY date_created ASC" mysql://$username:$password@host:port/dbname > "[$country]_purchase.csv"

echo "done"
```