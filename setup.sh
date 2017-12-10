#!/bin/bash
#Setup script for Xavier

DIR="$(pwd)"
if [[ $DIR == *"/go/src/"* ]]; then
	echo "Error! Cannot execute setup script from this location! Exiting."
	exit
fi

if [ -z %1 ]; then
	echo "Error! No GitHub username was provided"
	exit
else
	GITUSER=$1
fi 

if [ -z "$GOPATH" ]; then
	export $GOPATH=$HOME/go
fi

if [ ! -d "$GOPATH" ]; then
	echo "Go workspace does not exist at $GOPATH, creating a new one."
	mkdir "$GOPATH/src/github.com/$USER/Xavier" && cd "$_"
	git clone https://github.com/awarrier99/Xavier ./
	rm -rf "$DIR"
else
	echo "Go workspace exists at $GOPATH. Adding files to pre-existing workspace."
	cd "$GOPATH"
	if [ ! -d "src" ]; then
		mkdir src
	fi
	cd src

	if [ ! -d "github.com" ]; then
		mkdir github.com
	fi
	cd github.com

	if [ ! -d "$GITUSER" ]; then
		mkdir $GITUSER
	fi
	cd $GITUSER

	if [ -d "Xavier" ]; then
		echo "Directory Xavier already exists!"
		echo "Do you want to overwrite its contents?"
		read res
		if [ $res == 'y' ] || [ $res == 'Y' ]; then
			rm -rf "Xavier"
			mkdir "Xavier"
		else
			echo "Exiting setup script."
			exit
		fi
	else
		mkdir "Xavier"
	fi

	git clone https://github.com/awarrier99/Xavier Xavier/
	rm -rf "$DIR"
fi

echo "Workspace set up successfully!"
cd "$GOPATH/src/github.com/$GITUSER/Xavier"
