To define a primary key that auto increments in `PostgreSQL` you create the table row using the `SERIAL` type with the `PRIMARY KEY` constraint, like this:

```sql
CREATE TABLE cars (
  id    SERIAL PRIMARY KEY,
  brand VARCHAR(30) NOT NULL,
  model VARCHAR(30) NOT NULL,
  year  CHAR(4) NOT NULL
);
```

In MySQL / MariaDB this is equivalent to

```sql
CREATE TABLE cars (
  id    INT(11) NOT NULL AUTO_INCREMENT,
  brand VARCHAR(30) NOT NULL,
  model VARCHAR(30) NOT NULL,
  year CHAR(30) NOT NULL,
  PRIMARY KEY (`id`)
);
```
