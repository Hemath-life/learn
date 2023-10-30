[Mongodb setup in linux](#mongodb-setup)

[Mongodb status in Linux](#check-mongodb-status)

### Mongodb setup

```bash
sudo apt install dirmngr gnupg apt-transport-https ca-certificates software-properties-common

sudo add-apt-repository 'deb [arch=amd64] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/4.4 multiverse'

wget -qO - https://www.mongodb.org/static/pgp/server-4.4.asc | sudo apt-key add -

echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/4.4 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.4.list

sudo apt install mongodb-org

mongod

mongos

sudo apt-get update

sudo apt-get install -y mongodb-org

```

### Check mongodb status

```bash
# ------------ systemd (systemctl) ------------
sudo systemctl start mongod
sudo systemctl status mongod
sudo systemctl stop mongodb

```
