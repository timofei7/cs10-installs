#!/bin/bash

function error {
  echo "Error occured in $1, please email your TAs with copy/pasted error messages"
  exit 1;
}  

echo "this will install homebrew and opencv"
echo "you will see lots of stuff happening but most of it is safe to ignore"
echo "it may take a little while..."
echo ""
echo "authenticating, you may need to enter your system password"
sudo true 

echo "backing up possibly conflicting libraries in /opt and /usr/local"
sudo mv /opt/local /opt/local.before_CS10 2>&1 | grep -v "No such"
sudo mv /usr/local /usr/local.before_CS10 2>&1 | grep -v "No such"

echo "installing homebrew from http://mxcl.github.com/homebrew/"
echo "it is an open source package manager, very cool"
ruby <(curl -fsSkL raw.github.com/mxcl/homebrew/go) || error "installing homebrew"

DOCTOR=`brew doctor`

if [[ "$DOCTOR" != *raring\ to\ brew* ]]; then
  echo "sorry! something went wrong. please copy paste any error messages and email your TAs"
  echo "error: \n$DOCTOR"
  exit 1;
fi


echo "installing python dependency"
brew install python || error "installing python" 
echo "setting paths"
PATH=/usr/local/bin:/usr/local/share/python:$PATH
export PATH
echo 'export PATH=/usr/local/bin:/usr/local/share/python:$PATH' >> ~/.profile
echo "installing numpy dependency"
pip -q install numpy || error "installing numpy
echo "installing opencv"
brew install opencv || error "installing opencv" 
