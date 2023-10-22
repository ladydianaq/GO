# GO
WHAT IS GO?

It is a programming language developed at Google.
DEVELOPED BY: Robert Griesemer

docker
kubernetes
Terraform

It is an open source language, it is a compiled programming language
which facilitates
They are compiled into a single, independent binary file, making them easy to distribute.
uses and sharing
supports cross compilation, allows compilation of windows for linux or linux for windows

It is a multiparadigm language

Concurrency management

It has a large standard library, sharing similarities with Java, Python and others.


## TIOBE

page where we can see what rank go is in

## USE CASES

Systems level applications, Web applications, network level applications and the cloud
command line tools and uses, distributed systems, base implementation
of data. go.dev (check the different packages).

## What do we need?

Documentation in go.dev/doc/tutorial/getting-started editors to work with go go.dev/doc/editors

## Plaground of go

We can share go code, we can locate it at go.dev/play/
All go files have their own example package: the package main (package main)
LIMITATIONS: We cannot work with external packages


Install go in your OS


Copy the link address from the do.dev/dl/ page


```shell
https://go.dev/dl/go1.21.3.linux-amd64.tar.gz
```


Install and download from Terminal
```shell
cd download
sudo apt install wget
```


Discharge


```shell
wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz
```


Uninstall the previous version
```shell
$ sudo rm -rf /usr/local/go
```


Unzip the file
```shell
$ sudo tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz
```


go to download directory: cd /usr/local/ ,ls , cd go/ , ls , cd bin/ , ls

Add the environment variables in $HOME/.profile to the PATH (for the current user)
/etc/profile (for all users)


For current user
```shell
CD
pwd
ls -la
```


find the .profile file, modify the file with:


```shell
nano .profile
```


```shell
nano .profile
```


then in the final part paste the following code:
```shell
export PATH=$PATH:/usr/local/go/bin
```


It is the place where go is installed, to save the file CTRL+o and finally CTRL+x

The changes that were made within our .profile file are reflected after starting
session again on our computer, but we can execute the following code to
execute the changes at home plate.


```shell
source $HOME/.profile
go version
```


## GO environment variables


Create our workspace
```shell
go send
```

Where you point to the go workspace.
search GOPATH="/home/lady-quinto/go"


Create the folder, in this place the packages (pkg), the binaries (bin) and the tools used in the application will be stored, where the go folders (src) will be created.
```shell
mkdir go
cd go/
mkdir src
cd src/
```


## Third party packages


packages that were not developed in go, then we must initialize a module manager


module manager, initialize the module manager, name with which the modules will be managed
```shell
cd go/
cd src/
cd hello world
go mod init hello world
go mod tidy
```


ls, then a go.mod file will be created that allows you to manage the modules and the dependencies of your modules


quote package has famous quotes from hello world, rsc.io/quote repository path of quote package
```shell
go get rsc.io/quote
```


## Problems with your package generation
```shell
Go to Main Menu (4 lines to the right) -> settings -> Go Modules -> check the check on "Enable Go modules integration"
```


command to run and compile GO applications
```shell
go run
```


```shell

```
```shell

```

```shell

```
```shell

```
```shell

```


Install mysql in your OS
```shell
sudo apt update
sudo apt install mysql-server
```


## Step 2: Configure the kubeconfig file


You need to set up the kubeconfig file to access the cluster. If you already have a kubeconfig file, you can skip this step.


```bash
microk8s config > $HOME/.kube/config-local
```


## Step 3: Install and setup lens


```bash
sudo