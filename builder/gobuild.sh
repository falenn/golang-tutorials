#!/bin/bash

# This is more for demonstration of the env vars set when installing Go.

# For convenience, if this builder.sh is on the path aready, when invoked, it
# will build the CURRENT_LOCATION (current directory), assuming there is a .go
# file in the current directory.  It will then try to run the resulting bineary
# of the same name.

# SETUP
# Place onto your path the builder.sh:
# gob -> ~/dev/go/src/golang-tutorials/builder/gobuild.sh
# I created a simlink in my ~/bin which is already on the path.

GO_PATH=/usr/local/go
GOPATH=$HOME/dev/go
GOBIN=$HOME/dev/go/bin
export GOPATH
export GOBIN

# Put Go on the path as well as the bin dir on the path
PATH=$PATH:$GO_PATH/bin:$GOBIN

# Grab current directory
CURRENT_LOCATION=${PWD##*/}

clean() {
  echo "Cleaning $GOBIN...."
  rm -rf $GOBIN/*.go
}

installFiles() {
  # install all of the files in this directory and below
  #echo "Trying to install $1"
  if [ -d "$1" ]
  then
    #echo "Found directory $1"
    for f in $1/*
    do
      installFiles $f
    done
  else
    #echo "Installing $1"
    #execute the file to test
    xpath=${1%/*}
    xbase=${1##*/}
    xfext=${xbase##*.}
    xpref=${xbase%.*}
    #echo; echo  path=${xpath}; echo pref=${xpref}; echo ext=${xfext}
    if [ ${xfext} == "go" ]
    then
      go install $1
      if [ $? -eq 0 ]
      then
        echo "Running ${xpref}..............................."
        $GOBIN/${xpref}
      else
        echo "Error installing $1"
      fi
    fi
  fi
}

clean
installFiles $PWD

# Now, run the code - this binary should be found on the path, Go having
# deposited it into the GOBIN location
#$GOBIN/$CURRENT_LOCATION
