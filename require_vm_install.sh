## install docker 
sudo apt update 
sudo apt-get install docker-ce docker-ce-cli containerd.io
## 

## create a cluster k3s
## 
curl -sfL https://get.k3s.io | sh -

##install Go 
#rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.6.linux-amd64.tar.gz
#export PATH=$PATH:/usr/local/go/bin
#go version
