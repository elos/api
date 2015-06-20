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
SRCROOT="/opt/go"
SRCPATH="/opt/gopath"

# Get the ARCH
ARCH=`uname -m | sed 's|i686|386|' | sed 's|x86_64|amd64|'`

# Install Prereq Packages
sudo apt-get update
sudo apt-get install -y build-essential curl git-core libpcre3-dev mercurial pkg-config zip

# Install Go
cd /tmp
wget -q https://storage.googleapis.com/golang/go1.4.2.linux-${ARCH}.tar.gz
tar -xf go1.4.2.linux-${ARCH}.tar.gz
sudo mv go $SRCROOT
sudo chmod 775 $SRCROOT
sudo chown vagrant:vagrant $SRCROOT

# Setup the GOPATH; even though the shared folder spec gives the working
# directory the right user/group, we need to set it properly on the
# parent path to allow subsequent "go get" commands to work.
sudo mkdir -p $SRCPATH
sudo chown -R vagrant:vagrant $SRCPATH 2>/dev/null || true
# ^^ silencing errors here because we expect this to fail for the shared folder

cat <<EOF >/tmp/gopath.sh
export GOPATH="$SRCPATH"
export GOROOT="$SRCROOT"
export PATH="$SRCROOT/bin:$SRCPATH/bin:\$PATH"
EOF
sudo mv /tmp/gopath.sh /etc/profile.d/gopath.sh
sudo chmod 0755 /etc/profile.d/gopath.sh
source /etc/profile.d/gopath.sh

echo "== Go Installed =="

# --- }}}

# --- Configure Go {{{

echo "== Configuring Go =="

# Source the bashrc to see the environment variables
source /home/vagrant/.bashrc

# Get all the dependencies
cd $SRCPATH/src/github.com/elos/api && go get

echo "== Go Configured =="

# --- }}}

# --- Clear Up Memory {{{

sudo apt-get purge -y chef puppet

# --- }}}
