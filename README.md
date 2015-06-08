api
---

Package `api` provides the structure `api.Api` which coordinates the [`routes`](https://github.com/elos/api/tree/master/routes), [`middleware`](https://github.com/elos/api/tree/master/middleware) and [`services`](https://github.com/elos/api/tree/master/services) of the Elos API.

### Getting Started

The API comes with a preconfigured [Vagrant](https://www.vagrantup.com) [VirtualBox](https://www.virtualbox.org). Therefore you can develop on the API
locally using your own go workspace, or you can start the vagrant-managed virtual machine and develop in a production-like environment. We use the vagrant box (an Ubuntu 14.04 machine) to build the API for our production environment (which is also Ubuntu 14.04). We recommend you develop using Vagrant.

#### Developing with Vagrant

```bash
    # Need the repo
    git clone https://github.com/elos/api

    # Start the vagrant box
    vagrant up

    # Now we can ssh in
    vagrant ssh

    # Once inside, cd into the preconfigured workspace
    vagrant:$~ cd $GOPATH
```

##### You need
    1. [Vagrant](http://www.vagrantup.com/downloads)
    2. [VirtualBox](https://www.virtualbox.org/wiki/Downloads)

#### Developing Locally

```bash
    go get github.com/elos/api
```

### Run the Tests

```bash
    go test
```
