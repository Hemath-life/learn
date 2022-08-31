## Whats Views ?

In a database, a view is the result set of a stored query on the data,

using views can update data and delete data

database users can query just as they would in a persistent database collection object.

This pre-established query command is kept in the database dictionary.

## benefits of view ?

- simplifies query
- reduce the impact of changes
- `restrict access` to the data
- views is abstraction of the table
- can we update or drop the view
- using view we can make changes on the attracted tables like `CRUD operations`

## What if you need to changes database structure ?

if u changes have change all query u wrote so if we use view u can simply changes the view section it self

syntax:
View like create variable and storing the data after we want we can access like a table

```SQL
-- create view
CREATE VIEW amount_by_brand AS
 SELECT
     p.brand,
     sum(p.price) AS amount
 FROM products p JOIN orders o ON p.id  = o.product_id
 GROUP BY p.brand;

-- execute the view
SELECT * FROM amount_by_brand;

-- dropt the view
DROP VIEW amount_by_brand;
```

## Access control of the data or restrict access

`with check option` prevents form delete or update in view.

```SQL
CREATE OR REPLACE VIEW amount_by_brand AS
SELECT
	p.price,
    o.amount AS sold_amount,
    (o.amount - p.price) AS profit
FROM products p INNER JOIN orders o ON o.product_id=p.id
WHERE  (o.amount - p.price) > 0
WITH CHECK OPTION;


-- delete the data form view where price = 400
delete from amount_by_brand where price = 400;
-- with check option prevents form delete or update
```
