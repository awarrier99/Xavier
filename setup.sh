#!/bin/bash
#Setup script for Xavier

if [ -z %1]; then
	echo "Error! No GitHub username was provided"
	exit
else
	$USER = $1
fi 

if [ -z "$GOPATH"]; then
	export $GOPATH = $HOME/go
fi

DIR = "$(pwd)"

if [ ! -d "$GOPATH"]; then
	mkdir "$GOPATH/src/github.com/$USER/Xavier" && cd "$_"
	git clone https://github.com/awarrier99/Xavier ./
	rm -rf "$DIR"
else
	cd "$GOPATH"
	if [ ! -d "/src"]; then
		mkdir src
	fi
	cd src

	if [ ! -d "/github.com"]; then
		mkdir github.com
	fi
	cd github.com

	if [ ! -d "/$USER"]; then
		mkdir $USER
	fi
	cd $USER

	if [ -d "Xavier"]; then
		echo "Error! Directory Xavier already exists!"
		exit
	fi
	cd Xavier

	git clone https://github.com/awarrier99/Xavier ./
	rm -rf "$DIR"
fi
