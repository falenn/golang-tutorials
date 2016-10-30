! About the tutorial
I'm learning GoLang!  Yay!

I've followed the tutorials posted on https://tour.golang.org/basics/ in order to get up to speed.  I've added additional comments as I go along.

The builder.sh is my own creation.



!! SETUP

checkout into ~/dev/go/src such that looks like ~/dev/go/src/golang-tutorials


!! builder/gobuild.sh

gobuild.sh helps build go binaries.

SETUP
First, place gobuild.sh on your path - for instance,

If ~/bin is on your path, you can place a symlink to gobuild.sh in ~/bin.  Now, you can use gobuild.sh from anywhere within the golang-tutorials project.

ln -s ~/dev/go/src/golang-tutorials/builder/gobuild.sh ~/bin/gob


USAGE

to build lesson1.go,

cd ~?/dev/go/src/golang-tutorials/lesson1/;gob

This will cause lesson1.go to compile and execute.  The binary will write to the GOBIN location as configured in gobuild.sh: ~/dev/go/bin
