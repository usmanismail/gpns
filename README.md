Go Push Notification Service (GPNS)
===
This project has been moved too [https://github.com/gopns/gopns](https://github.com/gopns/gopns)

GPNS provides scalable and efficient a mass push notificaiton service built on top of Amazon's Simple Notification Service. The project provides REST end points for registering Device IDs with arbitrary metadata. These can then be used to send a push notification to all or a segment of the user base based on the metadata. This project is written in the Go Programming language and is designed to be resource efficient. In addition it uses Amazon Simple Queue Service in order to distribute work load out to multiple instances if you have them available. 

### Build from Source
* Install [Go 1.2](http://golang.org/doc/install#install)
* Setup GOPATH
* Add GOPATH/bin to PATH
* go get gpns
* go install gpns
* run gpns

```bash
mkdir ~/go
export GOPATH=~/go
export PATH=$PATH:$GOPATH/bin
go get github.com/usmanismail/gpns
go install github.com/usmanismail/gpns
gpns
```

Please see the [Wiki](https://github.com/usmanismail/gpns/wiki) for more information and documentation

### Contributors 
* [Usman Ismail](http://techtraits.com/usman.html)
