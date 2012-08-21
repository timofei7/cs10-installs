# CS10 OS X Software Installs#






##1) Install Developer Tools + Command Line Tools:##

###on Lion (10.7.x) or MountainLion (10.8.x):###



1. Download and install Xcode via the App Store: [xcode](http://itunes.apple.com/us/app/xcode)
2. Open the Xcode app (/Applications/Xcode.app)
3. Go to Preferences:
![image](https://raw.github.com/timofei7/cs10-installs/master/images//goto_preferences.png)
4. Make sure Command Line Tools are installed:
![image](https://raw.github.com/timofei7/cs10-installs/master/images/preferences_window.png)


###on Snow Leopard (ideally you should upgrade to 10.7+ but this will work):###


1. Download and Install:  [xcode_4.2_for_snow_leopard.dmg](http://www.cs.dartmouth.edu/~tim/xcode_4.2_for_snow_leopard.dmg)



##2) Install Homebrew (and some dependencies)##
0. Download and install [Xquartz](http://xquartz.macosforge.org/) (optional on 10.6 or 10.7 but required on 10.8)
1. open the builtin Terminal app in /Applications (or if you like customization use [iTerm2](http://www.iterm2.com/))
2. backup older/conflicting libraries:
	1. copy/paste into Terminal: `sudo mv /opt /opt.before_CS10` 
	2. copy/paste into Terminal: `sudo mv /usr/local /usr/local.before_CS10`
3. Install [Homebrew](http://mxcl.github.com/homebrew/) which is a handy opensource sofware manager by just copy/pasting the following into Terminal:  `ruby <(curl -fsSkL raw.github.com/mxcl/homebrew/go)`
4. Make sure it installed correctly: `brew doctor`
5. It should say: "Your system is raring to brew" if not email any errors you encountered to your TAs

##3) Install OpenCV and dependencies##

1. Run: `export PATH=/usr/local/bin:/usr/local/share/python:$PATH`
2. Run: `brew install python`
3. Run: `pip -q install numpy`
4. Run: `brew install opencv`
3. This last step will take a while but notice that you can use `brew` to install all sorts of cool opensource software really easily!


##4) Install Eclipse and configure project for JavaCV##

1. Download and install [Eclipse Classic](http://www.eclipse.org/downloads/) (a java interactive developement environment)
2. Download [JavaCV](http://javacv.googlecode.com/files/javacv-0.2-bin.zip), unzip, put it into some folder that you can remember.
3. Start eclipse and start a new Java Project
4. Navigate to Project > Properties > Java Build Path > Libraries and click "Add External JARs...".
4. Locate the JAR files, select them, and click OK.



##notes##
1. An alternative to steps 2 and 3 is an automated install. Just paste `bash <(curl -fsSkL raw.github.com/timofei7/cs10-installs/master/go)` into Terminal and it will attempt to run all the steps for you
2. An alternative to step 1 would be to only install command line tools rather than all of xcode. You may want Xcode if you ever want to write iOS apps, otherwise you can just download and install these:
	* on Lion:  [command line tools Lion](http://www.cs.dartmouth.edu/~tim/command_line_tools_for_xcode_os_x_lion_aug_2012.dmg)
	* on Mountain Lion: [command line tools Mountain Lion](http://www.cs.dartmouth.edu/~tim/command_line_tools_for_xcode_os_x_mountain_lion_aug_2012.dmg)
	* on Snow Leopard just install Xcode per step 1














