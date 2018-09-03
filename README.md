# Jexia Star Trek Test
Translate a name written in English to Klingon and find out its species

## Pre Requirements
- Internet Connection with access to GitHub and stapi.co
- 350 Mb free space for GoLang and application
- Any O.S. compatible with GoLang

## Environment Setup
To run this program, you will need first to install:
```
Git
Go
```

After that you can download it from GitHub as following:
```
git clone https://github.com/RoValentim/startrek.git /tmp/RoValentim
```

And set GoLang Application Path to this new directory
```
export GOPATH=/tmp/RoValentim
```

## Running the tests
Once installed, you can run unit test for the entire application entering in src directory and running:
```
go test -v ...
```

But this take a while, since this will run a lot of GoLang tests, so it would be better to get in each directory and run manually as following:
```
go test -v
```

## Running the application
To run the application, get inside main folder and run main.go with a name as parameter
```
cd /tmp/RoValentim/src/main
go run main.go Uhuru
```

This should give this output:
```
0XF8E5 0XF8D6 0XF8E5 0XF8E1 0XF8D0
Human
```

If you want to see more information about what's going on, you can set a log level higher than the default, level 1, until level 4, exporting an O.S. variable named LOG_LEVEL
```
export LOG_LEVEL=4
```

## Building an Executable
To build an executable for the application, get inside main folder and compile it:
```
cd /tmp/RoValentim/src/main
go build main
```

## Application Design Overview
This application are using multi tier architecture concepts with DRY, so each function do only one job and are reused as needed.
Following the tree folder with some comments:
```
/src
   /bin  (compiled application should stay in this folder)
   /main
      /main.go (the main application)
      /stt.log (in Windows environment, for POSIX this log should be in /var/log/stt.log)
   /services
      /character (folder with character specie font and tests)
      /translate (folder with name translation font and tests)
   /shared
      /defines  (defined variables)
      /logger   (log function with writing in both file and stdout)
      /messages (defined messages for the application)
      /urls     (URL functions, such as those used to consume REST)
  /README.md  (this file)
```
