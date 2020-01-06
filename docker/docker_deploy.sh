#!/bin/sh

# Uninstall old versions

sudo apt-get remove docker docker-engine docker.io containerd runc

# Install using the repository

sudo apt-get update

# Install packages to allow apt to use a repository over HTTPS

sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg2 \
    software-properties-common

# Add Docker’s official GPG key

curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -

# Verify that you now have the key with the fingerprint

sudo apt-key fingerprint 0EBFCD88

# Use the following command to set up the stable repository

sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/debian \
   $(lsb_release -cs) \
   stable"

# Update the apt package index

sudo apt-get update

# Install the latest version of Docker Engine

sudo apt-get install docker-ce docker-ce-cli containerd.io