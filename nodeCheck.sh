#!/bin/bash
######################################################
#
#  Shell Script to Check the Health of the Node
#
######################################################

TNT_ADDR="$1"

if [ -z "$2" ]
then
	$HOME/go/src/github.com/SomniaStellarum/TNTNodeChecker/bin/TNTNodeChecker --tnt=$TNT_ADDR
else
	HOST=$2
	$HOME/go/src/github.com/SomniaStellarum/TNTNodeChecker/bin/TNTNodeChecker --tnt=$TNT_ADDR --host=$HOST
fi
