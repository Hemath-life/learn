### Postgres setup in Linux

#### install postgres

```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```

#### lunch postgres

```bash
sudo -i -u postgres
psql

'\q' # exit

# ------------ (Run psql ) ------------
sudo service postgresql start
sudo -u postgres psql
exclude=postgresql*


# ------------ (Check psql running) ------------
pgrep -u postgres -fa -- -D
pg_isready

```
