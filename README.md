# Go Study 


### go environment setting for Mac
```
echo 'export GOPATH=$HOME/Code/go' >> ~/.bash_profile
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bash_profile
source ~/.bash_profile

mkdir -p ~/Code/go/src/github.com/csk6124
cd !$
```

OXS use .bash_profile, linux use .bashrc

you have setup your environment setup you can install some tools
```
go get -u golang.org/x/tools/cmd/goimports
go get -u golang.org/x/tools/cmd/vet
go get -u golang.org/x/tools/cmd/oracle
go get -u golang.org/x/tools/cmd/godoc
```

### sublime install for golang 
* GoSublime
Preferences -> Settings -> More -> Syntax Specific -> User
```
{

	// you may set specific environment variables here
	// e.g "env": { "PATH": "$HOME/go/bin:$PATH" }
	// in values, $PATH and ${PATH} are replaced with
	// the corresponding environment(PATH) variable, if it exists.
	"env": {"GOPATH": "$HOME/Code/go", "PATH": "$GOPATH/bin:$PATH" },

  "fmt_cmd": ["goimports"],

	// enable comp-lint, this will effectively disable the live linter
	"comp_lint_enabled": true,

	// list of commands to run
	"comp_lint_commands": [
	    // run `golint` on all files in the package
	    // "shell":true is required in order to run the command through your shell (to expand `*.go`)
	    // also see: the documentation for the `shell` setting in the default settings file ctrl+dot,ctrl+4
	    {"cmd": ["golint *.go"], "shell": true},

	    // run go vet on the package
	    {"cmd": ["go", "vet"]},

	    // run `go install` on the package. GOBIN is set,
	    // so `main` packages shouldn't result in the installation of a binary
	    {"cmd": ["go", "install"]}
	],

	"on_save": [
	    // run comp-lint when you save,
	    // naturally, you can also bind this command `gs_comp_lint`
	    // to a key binding if you want
	    {"cmd": "gs_comp_lint"}
	]
}
```

* GoOracle
* GitGutter
Preferences -> Settings -> More -> Syntax Specific -> User
```
{
  "spell_check": true
}

```

# Go How to write go code  https://golang.org/doc/code.html
workspace
* src
* pkg
* bin

```
bin/
	hello
pkg/
	linux_amd64/
		github.com/csk6124/..
src/
	github.com/csk6124/juneGoHome/
		.git
		/hello
			hello.go
		/stringutil
			reverse.go
			reverse_test.go
```


# How to build and install go package
```
cd $GOPATH
go install github.com/csk6124/juneGoHome/hello
bin/hello or hello
```

### util build
```
go build github.com/csk6124/juneGoHome/stringutil
```


# Go syntest

### Go type
private는 소문자
public은 대문자로 한다.

private 예제)
type message struct {
    Name string
    Body string
    Time int64
}

public 예제)
type Message struct {
    Name string
    Body string
    Time int64
}




