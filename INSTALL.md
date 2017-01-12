
# Installation

## Compiled version

*shelltools* is a collection of command line programs run from a shell like Bash (or Powershell in Windows). If you download the repository a compiled version is in the dist directory. The compiled binary matching your computer type and operating system can be copied to a bin directory in your PATH.

Compiled versions are available for Mac OS X (amd64 processor), Linux (amd64), Windows (amd64) and Rapsberry Pi (both ARM6 and ARM7)

### Mac OS X

1. Download **shelltools-binary-release.zip** from [https://github.com/rsdoiel/shelltools/releases/latest](https://github.com/rsdoiel/shelltools/releases/latest)
2. Open a finder window, find and unzip **shelltools-binary-release.zip**
3. Look in the unziped folder and find *dist/macosx-amd64/*
4. Drag (or copy) *findfile*, *finddir*, *mergepath* and *pathparts* to a "bin" directory in your path
5. Open and "Terminal" and run `findfile -h` to confirm you were successful

### Windows

1. Download **shelltools-binary-release.zip** from [https://github.com/rsdoiel/shelltools/releases/latest](https://github.com/rsdoiel/shelltools/releases/latest)
2. Open the file manager find and unzip **shelltools-binary-release.zip**
3. Look in the unziped folder and find *dist/windows-amd64/*
4. Drag (or copy) *findfile.exe*, *finddir.exe*, *mergepath* and *pathparts* to a "bin" directory in your path
5. Open Bash and and run `findfile -h` to confirm you were successful

### Linux

1. Download **shelltools-binary-release.zip** from [https://github.com/rsdoiel/shelltools/releases/latest](https://github.com/rsdoiel/shelltools/releases/latest)
2. Find and unzip **shelltools-binary-release.zip**
3. In the unziped directory and find for *dist/linux-amd64/*
4. Copy *findfile*, *finddir*, *mergepath* and *pathparts* to a "bin" directory (e.g. cp ~/Downloads/shelltools-binary-release/dist/linux-amd64/\* ~/bin/)
5. From the shell prompt run `findfile -h` to confirm you were successful

### Raspberry Pi

If you are using a Raspberry Pi 2 or later use the ARM7 binary, ARM6 is only for the first generaiton Raspberry Pi.

1. Download **shelltools-binary-release.zip** from [https://github.com/rsdoiel/shelltools/releases/latest](https://github.com/rsdoiel/shelltools/releases/latest)
2. Find and unzip **shelltools-binary-release.zip**
3. In the unziped directory and find for *dist/raspberrypi-arm7/*
4. Copy *findfile*, *finddir*, *mergepath* and *pathparts* to a "bin" directory (e.g. cp ~/Downloads/shelltools-binary-release/dist/raspberrypi-arm7/\* ~/bin/)
    + if you are using an original Raspberry Pi you should copy the ARM6 version instead
5. From the shell prompt run `findfile -h` to confirm you were successful


## Compiling from source

If you have go v1.7 or better installed then should be able to "go get" to install all the **shelltools** utilities and
package. You will need the GOBIN environment variable set. In this example I've set it to $HOME/bin.

```bash
    export GOBIN=$HOME/bin
    go get github.com/rsdoiel/shelltools/...
```

Or for Windows 10 Powershell (assumes the Windows versions of Go and Git are previously installed)


```powershell
    $Env:GOPATH = "$HOME"
    go get github.com/rsdoiel/shelltools/...
```

or to install from source on Linux/Unix systems

```bash
    git clone https://github.com/rsdoiel/shelltools src/github.com/rsdoiel/shelltools
    cd src/github.com/rsdoiel/shelltools
    make
    make test
    make install
```

Or for Windows 10 Powershell

```powershell
    $Env:GOBIN = "$HOME\bin"
    git clone https://github.com/rsdoiel/shelltools src/github.com/rsdoiel/shelltools
    cd src\github.com\rsdoiel\shelltools
    go install cmds\csvcols\csvcols.go
    go install cmds\finddir\finddir.go
    go install cmds\findfile\filefile.go
    go install cmds\mergepath\mergepath.go
    go install cmds\pathparts\pathparts.go
    go install cmds\reldate\reldate.go
    go install cmds\timefmt\timefmt.go
    go install cmds\urlparse\urlparse.go
```


