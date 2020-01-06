#!/bin/sh

echo "Uninstall old versions"
sudo apt -y remove docker docker-engine docker.io containerd runc

echo "Install using the repository"
sudo apt update

echo "Install packages to allow apt to use a repository over HTTPS"
sudo apt -y install \
  apt-transport-https \
  ca-certificates \
  curl \
  gnupg2 \
  software-properties-common

echo "Add Docker’s official GPG key"
curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -

echo "Verify that you now have the key with the fingerprint"
sudo apt-key fingerprint 0EBFCD88

echo "Use the following command to set up the stable repository"
sudo add-apt-repository \
  "deb [arch=amd64] https://download.docker.com/linux/debian \
   $(lsb_release -cs) \
   stable"

echo "Update the apt package index"
sudo apt update

echo "Install the latest version of Docker Engine"
sudo apt -y install docker-ce docker-ce-cli containerd.io

sudo systemctl enable docker