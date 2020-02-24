# kvant-node

## Start KVANT node

Kvant is a high-performance blockchain. We recommend that you run it only with limited access for trusted administrators. Therefore, the launch of the node is from the root and to the root of the system.

```bash
wget -O kvant_installer.sh  https://raw.githubusercontent.com/kvantdev-test/kvant-node/master/installer.sh && \
chmod +x kvant_installer.sh && \
./kvant_installer.sh
```


## Addons installations

```bash
wget https://github.com/google/leveldb/archive/v1.20.tar.gz && \
  tar -zxvf v1.20.tar.gz && \
  apt-get -f install build-essential make &&\
  cd leveldb-1.20/ && \
  make && \
  sudo cp -r out-static/lib* out-shared/lib* /usr/local/lib/ && \
  cd include/ && \
  sudo cp -r leveldb /usr/local/include/ && \
  sudo ldconfig && \
  rm -f ../v1.20.tar.gz
```


## Compiling from source

### Install GOLANG

```bash
wget -O go.tgz https://dl.google.com/go/go1.13.8.linux-amd64.tar.gz && \
tar xzf go.tgz && \
chown -R root:root ./go && \
mv go /usr/local && \
echo export GOPATH=/d01/go >> ~/.profile && \
echo 'export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin'  >> ~/.profile && \
source ~/.profile && go version
```

### Configuring OS
```bash
apt-get install build-essential
```

### Download and Preparing source code

```bash
wget -O kvant_compiller.sh  https://raw.githubusercontent.com/kvantdev-test/kvant-node/master/compiller.sh && \
chmod +x kvant_compiller.sh && \
./kvant_compiller.sh
```

