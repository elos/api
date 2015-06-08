#!/usr/bin/env bash

# --- Install Mongo {{{

echo "== Installing Mongo =="

# Import the public key used by the package management system.
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 7F0CEB10

# Create a list file for MongoDB.
echo "deb http://repo.mongodb.org/apt/ubuntu "$(lsb_release -sc)"/mongodb-org/3.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-3.0.list

# Reload local package database.
sudo apt-get update

# TODO: pin a specific version of mongo
sudo apt-get install -y mongodb-org

echo "== Mongo Installed =="

# --- }}}

# --- Install Go {{{

echo "== Installing Go =="

# Install requirements (GVM et al)
apt-get update
apt-get install htop curl git mercurial make binutils bison gcc build-essential --fix-missing -y

# Install GVM
sudo -u vagrant HOME=/home/vagrant bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
source /home/vagrant/.gvm/scripts/gvm
source /home/vagrant/.bashrc

# Install go 1.4
echo "== Installing go1.4 =="
sudo -u vagrant GVM_ROOT=/home/vagrant/.gvm /home/vagrant/.gvm/bin/gvm install go1.4
echo "== Installed go1.4 =="

echo "== Using go1.4 As Default=="
sudo -u vagrant GVM_ROOT=/home/vagrant/.gvm /home/vagrant/.gvm/bin/gvm use go1.4

echo "== Go Installed =="

# --- }}}

# --- Configure Go {{{

echo "== Configuring Go =="

# Add Environment Variables
echo "export GOPATH=/home/vagrant/go" >> /home/vagrant/.bashrc

# Source the bashrc to see the environment variables
source /home/vagrant/.bashrc

# Get all the dependencies
cd /home/vagrant/go/src/github.com/elos/api && sudo -u vagrant go get

echo "== Go Configured =="

# --- }}}
