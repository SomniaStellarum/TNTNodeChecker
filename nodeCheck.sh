#!/bin/bash
######################################################
#
#  Shell Script to Check the Health of the Node
#
######################################################

HOME="$1"
TNT_ADDR="$2"

if [ -z "$3" ]
then
	$HOME/go/src/github.com/SomniaStellarum/TNTNodeChecker/bin/TNTNodeChecker --tnt=$TNT_ADDR
else
	HOST=$3
	$HOME/go/src/github.com/SomniaStellarum/TNTNodeChecker/bin/TNTNodeChecker --tnt=$TNT_ADDR --host=$HOST
fi
