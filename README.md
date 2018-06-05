# Xavier
This repository hosts code for Xavier, a personal assistant powered by Artificial Intelligence (in development).

## Introduction
The end goal of this project is to create an assistant which replicates functionality similar to Tony Stark's *Jarvis*. While working towards this goal, the team is working on smaller projects such as image labeling. This project uses the following languages and libraries:
  * [Node.js](https://nodejs.org/en/)
  * [Go](https://golang.org)
  * [TensorFlow](https://www.tensorflow.org)
  * [Python](https://www.python.org)

## Setup
### Linux


### Mac OSX
Download and install the necessary distribution files (which can be found on the Go website linked above). If you install your files manually and choose to install them in a location different from /usr/local/go (the default), you must set the environment variable $GOROOT to point to it.\
Next, create your go workspace where all your various Go projects will be stored. Your workspace folder must be named 'go' (without quotes). The default location for this is in your $HOME (i.e. /Users/ashvin) directory. However, if you prefer, you can create this workspace in a different location, but you must set the environment variable $GOPATH to point to this directory.\
Finally, create the necessary directory for this codebase. Go conventions suggest that the path of a project must reflect its source control path. As such, create a directory inside your Go workspace called 'src'. Within this directory, create a folder called 'github.com'. Within that directory, create another named 'awarrier99'. Within this directory, clone this repository into a folder named 'Xavier'.\
The full path of the project folder should be $GOPATH/src/github.com/awarrier99/Xavier.\
A series of simple Terminal commands is listed below in order to easily set up your workspace and project folder.

### Windows
