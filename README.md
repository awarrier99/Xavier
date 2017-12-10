# Xavier
This repository hosts code for Xavier, a personal assistant powered by Artificial Intelligence (in development).

## Introduction
The end goal of this project is to create an assistant which replicates functionality similar to Tony Stark's *Jarvis*. While working towards this goal, the team is working on smaller projects such as image labeling. This project uses the following languages and libraries:
  * [The Go Programming Language](https://golang.org)
  * [TensorFlow](https://www.tensorflow.org)
  * [Python](https://www.python.org)

## Setup
### Linux
If you already have a pre-existing Go workspace directory, **make sure that the variable $GOPATH points to it** before following these steps. This can be tested using `echo $GOPATH`
  1. Clone this repository into an empty directory *dirname* using `git clone https://awarrier99/Xavier dirname/`

  2. `cd` into the directory using `cd dirname`

  3. Run the setup script, setup.sh, with your GitHub username as a command-line argument using `./setup.sh username`

This script will initialize a new Go workspace for you if a previous one doesn't exist, located at $HOME/go (usually /home/user/go). For more information on Go programming conventions, refer to the [Go documentation](https://golang.org/doc). 

The setup script will then create a conventional directory structure and copy the repository to $HOME/go/src/github.com/username/Xavier. After running this script, **_dirname_ and its contents will be deleted**. The project will not be functional if the directory is located elsewhere, due to the conventions of Go programming.

Or, if you prefer, simply create the Go workspace manually and then clone the repository into the corresponding folder.

### Mac OSX


### Windows
