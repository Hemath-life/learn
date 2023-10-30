```SQL
create table car (
    id BIGSERIAL NOT NULL PRIMARY KEY,
	make VARCHAR(50) NOT NULL,
	model VARCHAR(50) NOT NULL,
	price NUMERIC(19,2 )NOT NULL
);

insert into car (id, make, model, price) values (1, 'Toyota', 'Celica', 9374.35);
insert into car (id, make, model, price) values (2, 'Mercedes-Benz', 'W123', 3222.00);
insert into car (id, make, model, price) values (3, 'Lexus', 'LS', 1460.29);
insert into car (id, make, model, price) values (4, 'Ford', 'Escort', 3887.78);


-- min max avg , ROUNT
SELECT MAX (price) FROM car;
SELECT MIN (price) FROM car;
SELECT ROUND(AVG (price)) FROM car;


-- GROUP BY
SELECT make, model,MIN(price) FROM car Group by make,model;
SELECT make,MAX(price) FROM car Group by make ;



--- sum
SELECT sum(price) FROM car  ;
SELECT make,sum(price) FROM car Group by make ;



-- - Arith metic operators provided by postgres
SELECT 10-1;
SELECT 10*1;
SELECT 10*1*10;
SELECT 10/3;
SELECT  10%3 ;
SELECT 10^3;
SELECT  10 ! ;



-- Arith metic operators Round ALIAS keayword - AS
SELECT id,make,model,price as orginal_price,ROUND(price*.10 ,2) AS discount ,ROUND(price - (price*.10),2) AS after_discount FROM car; -- 10% percentage



-- COALESCE() -- Defual values
SELECT COALESCE(1) AS number;
SELECT COALESCE(null,1) AS number;
SELECT COALESCE(null,null,null,null,1,10) AS number;
SELECT COALESCE(email,'Not Provided') as eamil FROM person LIMIT 5;



-- NULLIF
SELECT 10/0;
--  division by zero
SELECT NULLIF(10,10); -- this is will be 10 equalt to 10 so no output
SELECT NULLIF(10,1); -- this is will be 10 exce 10 not equl to 1
SELECT 10/NULLIF(0,0);
SELECT COALESCE(10/NULLIF(0,0),0);






```
