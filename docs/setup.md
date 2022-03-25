[Home](../Readme.md)
# Setup information

- Install go
- Install git
- Install VScode
  - For VSCode install go extension, markdown viewer

<hr/>

#### Some commands

go version <br/>
echo $GOPATH <br/>
go env <br/>

<hr/>

#### *Put the GOPATH/bin in your system PATH*
<hr/>

## [Install go tools](https://pkg.go.dev/golang.org/x/tools#section-readme) <br/>
  `go install golang.org/x/tools/...@latest`

<hr/>

#### A sample Debug configuration for VSCode
````
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [        
        
    {
        "name": "explorerapp",
        "type": "go",
        "request": "launch",
        "mode": "auto",
        "program": "${workspaceFolder}/main.go"
        
    }
    ]
}
````

