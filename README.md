<<<<<<< HEAD
# Traffic Generator Golang version
## Description of the work

The code read urls from text file and open the page or download html file of each url according to the case chosen. 
it can open pages and download files using http request. 
it takes screenshot of each opened page by the url using headless browser and save it in file.png according to the type wrtten in the the text either it's browser or request.
it generate log file contain the runtime , response time , url , status code for each url in the text file. 
each url link has timeout inserted in config.toml file.
it also can upload any chosen file.

---

## Sequence of running the code
-generating two log files one for http request and another one for headless browrser

-open pages

-download files

-upload

--- 

# Installation 

.**Go intsall**

 A sample installation of go version 1.14:
 **1.for Linux**
 
 Prepare the apt packages
 
 `` sudo apt update                    
  ``
  
 ``sudo apt upgrade
 ``
 
 Link of download of version 1.14 (https://dl.google.com/go/go1.14.linux-amd64.tar.gz)

Use the command

``
 wget https://dl.google.com/go/go1.14.linux-amd64.tar.gz
 ``
 
 Untar in /usr/local

``
 tar -C /usr/local -xzf go1.14.linux-amd64.tar.gz                   
 ``
 
 Add /usr/local/go/bin to the PATH environment variable:
 
 ``
 export PATH=$PATH:/usr/local/go/bin             
``

**2.for windows**

Link of download of vesion 1.15 (https://golang.org/dl/go1.15.1.windows-amd64.msi)

---

To test the code, use the go run command

``go run main.go
``

=======
# GIT_TEST
exploring git and github
learning basics of GIt
## learning basics of html 
making a simple website to learn more 
<<<<<<< HEAD
* tutorial 
* learning
* html

`` sudo apt update ``



>>>>>>> 1d67ac5a884c0209632d2ba1118dfe2a03be303c
