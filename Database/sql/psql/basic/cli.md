basic commands ?

- \h help for sql commands
- \? help for psql commands
- \g exit or terminate with semicolon to execute query
- \q to quit
- \l list all of the database
- \d describe the tables
- \d `table_name` describe the table
- \i file_path to run the sql files

# `always use sing quotes. `

## export as csv File ?

- \copy (SELECT \* FROM person LEFT JOIN car ON car.id = person.car_id) TO 'D:\Hemath\Bitcot\resultls.csv' DELIMITER ',' CSV HEADER;

### restart the sequence ?

```sql
ALTER SEQUENCE person_id_seq RESTART WITH 10;
SELECT nextval('person_id_seq'::regclass);
ALTER SEQUENCE person_id_seq RESTART WITH 9;
SELECT nextval('person_id_seq'::regclass);

```
