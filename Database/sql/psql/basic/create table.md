```sql
-------------------- [creating enum tipe]--------------------------
CREATE TYPE STATUS AS ENUM ('conformed', 'dispatched', 'on the way', 'delivered');
-- -------------------- [creating table]--------------------------


CREATE TABLE IF NOT EXISTS orders(
  id SERIAL PRIMARY KEY,
  product_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  status STATUS DEFAULT 'conformed',
  delivered_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW(),
  amount INTEGER
);
-- -------------------- [insert values ]--------------------------
INSERT INTO
  products(id,name,price)
VALUES
  (1, 'vivo',500),
  (2, 'iphone',1000),
  (3, 'redme',600),
  (4, 'realme',400),
  (5, 'one+',500);
INSERT INTO
  users
VALUES
  (1, 'Hemath'),
  (2, 'Justin'),
  (3, 'Hello'),
  (4, 'Anu');
INSERT INTO
  orders
values
  (1, 1, 1),
  (2, 2, 2),
  (3, 3, 3),
  (4, 1, 1),
  (5, 1, 2),
  (6, 1, 4);

```
