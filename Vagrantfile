# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

$bootstrap=<<SCRIPT
apt-get update
apt-get -y install wget
wget -qO- https://get.docker.com/ | sh
gpasswd -a vagrant docker
service docker restart

wget -qO- https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz | tar xvz -C /usr/local
su vagrant -c "echo 'PATH=$PATH:/usr/local/go/bin' >> ~/.profile"
su vagrant -c "echo 'GOPATH=/vagrant/kcompose/Godeps/_workspace' >> ~/.profile"
#su vagrant -c "go get k8s.io/kubernetes/pkg/api"

SCRIPT

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  # Every Vagrant virtual environment requires a box to build off of.

  config.vm.box = "ubuntu/trusty64"

  # Create a private network, which allows host-only access to the machine
  # using a specific IP.

  config.vm.network "private_network", ip: "192.168.33.10"

  config.vm.provider "virtualbox" do |vb|
     vb.customize ["modifyvm", :id, "--memory", "2048"]
  end

  config.vm.provision :shell, inline: $bootstrap 

end
