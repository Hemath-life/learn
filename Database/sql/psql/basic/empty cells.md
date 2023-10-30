## SQL, Handling empty cells

How to handle null data in a SQL database

```SQL
CREATE TABLE people (
  age INT,
  name CHAR(20)
);
```

SQL freely accepts empty values as records:

```SQL
INSERT INTO people VALUES (null, null);
```

This might be a problem, because now we have a row with null values:

```sql
age  |  name
-----+--------
  37 | Flavio
   8 | Roger
     |
```

To solve this, we can declare constrains on our table rows. `NOT NULL` prevents null values:

```sql
CREATE TABLE people (
  age INT NOT NULL,
  name CHAR(20) NOT NULL
);

```

if we try to execute this query again:

```sql
INSERT INTO people VALUES (null, null);
```

We’d get an error, like this:

```sql
ERROR:  null value in column "age" violates not-null constraint
DETAIL:  Failing row contains (null, null).
```

`Note that an empty string is a valid non-null value.`
