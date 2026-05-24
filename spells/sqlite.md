# SQL Lite

[https://antonz.org/sqlite-is-not-a-toy-database/](https://antonz.org/sqlite-is-not-a-toy-database/)

cron based backups 

[https://litestream.io/alternatives/cron/](https://litestream.io/alternatives/cron/)

[https://gosamples.dev/sqlite-intro/](https://gosamples.dev/sqlite-intro/)

```bash
sqlite3
.open nmovies.db
.tables
.schema [table_name]
select * from nmovie;
.quit
```

create table

```sql
CREATE TABLE IF NOT EXISTS nmovie (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, directory TEXT, synopsis TEXT, year INTEGER, imdb TEXT, poster TEXT);
```

delete all

```sql
DELETE FROM nmovie;
```

count all

```sql
SELECT COUNT(*) FROM nmovie;
```

select from all tables

```sql
SELECT * FROM nmovie;
```

dump all the data from tables 

```bash
sqlite3 nmovies.db .dump > nmovies.sql
```

## data types

```
NULL. The value is a NULL value.
INTEGER. The value is a signed integer, stored in 0, 1, 2, 3, 4, 6, or 8 bytes depending on the magnitude of the value.
REAL. The value is a floating point value, stored as an 8-byte IEEE floating point number.
TEXT. The value is a text string, stored using the database encoding (UTF-8, UTF-16BE or UTF-16LE).
BLOB. The value is a blob of data, stored exactly as it was input.
```

```sql
select strftime('%s','now');
select datetime(1684603958,'unixepoch');
```

```sql
strftime('%s','now')
1684604140
datetime(1684603958,'unixepoch')
2023-05-20 17:32:38
``` 