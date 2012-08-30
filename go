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
sudo -v 

if [[ -e /usr/bin/gcc && ! -e /Applications/Xcode.app && ! -e /Developer ]]; then
  echo "detected commandline tools only must enable"
  sudo xcode-select -switch /usr/bin
elif [[ ! -e /usr/bin/gcc && ! -e /Applications/Xcode.app && ! -e /Developer ]]; then
  echo "no xcode detected please install Xcode first then rerun this"
  exit 1;
else
  echo "xcode detected"
fi

echo "backing up possibly conflicting libraries in /opt/local and /usr/local"
sudo mv /opt/local /opt/local.before_CS10 2>&1 | grep -v "No such"
sudo mv /usr/local /usr/local.before_CS10 2>&1 | grep -v "No such"

echo "installing homebrew from http://mxcl.github.com/homebrew/"
echo "it is an open source package manager, very cool"
ruby <(curl -fsSkL raw.github.com/mxcl/homebrew/go) || error "installing homebrew"

DOCTOR=`brew doctor`

if [[ "$DOCTOR" != *raring\ to\ brew* ]]; then
  echo "sorry! something went wrong. please copy paste any error messages and email your TAs"
  echo "error: $DOCTOR"
  exit 1;
fi

echo "reauthenticating if necessary"
sudo -v 

echo "installing python dependency"
brew install python || error "installing python" 

echo "installing other dependencies"
brew install libpng || error "installing other dependencies" 

echo "setting paths"
PATH=/usr/local/bin:/usr/local/share/python:$PATH
export PATH
echo 'export PATH=/usr/local/bin:/usr/local/share/python:$PATH' >> ~/.profile

echo "installing numpy dependency"
/usr/local/share/python/pip -q install numpy || error "installing numpy"

echo "installing opencv"
# just find any old gfortran we don't really care
GFORTRAN=`find /usr/bin -maxdepth 1 -name 'gfortran*' -print -quit`
if [ "$GFORTRAN" == "" ]; then
  #doesnt actually need it set to anything real
  export FC=/usr/bin/gfortran 
else
  export FC=$GFORTRAN
fi

#setting ldflags to pick up the libpng we installed cause there seems to be a system libpng problem
LDFLAGS=-L/usr/local/lib,/usr/local/opt/libpng/lib CPPFLAGS=-I/usr/local/include,/usr/local/opt/libpng/include brew install opencv || error "installing opencv" 


echo "you are all set to go!"
