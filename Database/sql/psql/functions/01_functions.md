### Sql functions ?

Functions in SQL

Functions similar like procedures but the function return **single value** not like procedures

Cant return sets or multiple values,rows and columns

return single value

syntax:

```sql
CREATE FUNCTION 'function_name'(
    parameter_name data_type
)

RETURNS INTEGER -- this specifics type fo the value it will return , any other data types in psql
-- ATTRIBUTES
DETERMINISTIC -- make sure what the input data type based on that it will return same data type

READS SQL DATA -- saying this function going to have select statement to read some data

MODIFIES SQL DATA -- saying this function going to do inset,update,delete in rows

BEGIN
-- boyd of the function
RETURN 1;
END

```

example:

```sql
CREATE OR REPLACE FUNCTION totalRecords ()
RETURNS integer AS
$$
DECLARE
	total integer;
BEGIN
   SELECT count(*) into total FROM users;
   RETURN total;
END;
$$
LANGUAGE plpgsql;


select totalRecords();


```
