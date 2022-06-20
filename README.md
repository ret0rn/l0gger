# L0gger 
<!-- [![Build Status](https://app.travis-ci.com/ret0rn/l0gger.svg?branch=master)](https://app.travis-ci.com/ret0rn/l0gger) -->

Wrapper for quick and easy work with the [logrus](https://github.com/sirupsen/logrus) package in golang


[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ret0rn/l0gger)](https://github.com/ret0rn/l0gger/blob/master/go.mod)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/ret0rn/l0gger?style=plastic)](https://github.com/ret0rn/l0gger/releases)
[![GitHub](https://img.shields.io/github/license/ret0rn/l0gger?style=plastic)](https://github.com/ret0rn/l0gger/blob/master/LICENSE)

[![GitHub Repo stars](https://img.shields.io/github/stars/ret0rn/l0gger?style=social)](https://github.com/ret0rn/l0gger/stargazers)


 ## Install 

```
go get github.com/ret0rn/l0gger
```


 ## Formaters 

 - ```JsonFormater```
 - ```TextFormater```

 ## Log Levels
- ```DebugLvl```
- ```InfoLvl```
- ```WarningLvl```
- ```ErrorLvl```
- ```FatalLvl ```


## Log Out

- File
- Сonsole
- Others ```io.Writer```'s

## Usage

```golang
var log, err = l0gger.New([1], [2], [3]) 
```
1. **Log Out** 

	- type ```string``` - log to file
	- ```nil``` or ```os.Stdout``` - log to console 
	- type ```io.Writer``` - others io.Writers
2. **Formaters**
	- ```JsonFormater``` - JSON Formater
	
		```{"file":"test.go:18","level":"info","msg":"hello","time":"2022-06-19T22:14:02+03:00"}```

	- ```TextFormater``` - Text Formater

		```time="2022-06-19T22:15:01+03:00" level=info msg=hello file="test.go:18"```
3. **Level**
	- ```DebugLvl```
	- ```InfoLvl```
	- ```WarningLvl```
	- ```ErrorLvl```
	- ```FatalLvl ```


### Log to Console 
```go
import "github.com/ret0rn/l0gger"

func main() {

	var log, err = l0gger.New(nil, l0gger.JsonFormater, l0gger.DebugLvl)

	log.Info("Info test")
}
```

### Log to File
```go
import "github.com/ret0rn/l0gger"

func main() {

	var log, err = l0gger.New("logger.log", l0gger.JsonFormater, l0gger.DebugLvl)

	log.Info("Info test")
}
```
