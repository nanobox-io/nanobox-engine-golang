# Go

This is a Go engine used to launch Go apps on [Nanobox](http://nanobox.io).

## Usage
To use the Go engine, specify `golang` as your `engine` in your boxfile.yml.

```yaml
code.build:
  engine: golang
```

## Build Process
When [running a build](https://docs.nanboox.io/cli/build/), this engine compiles code by doing the following:

```
> go get
> go build
```

*These commands can be modified using the [fetch](#fetch) and [build](#build) config options*

## Configuration Options
This engine exposes configuration options through the [boxfile.yml](http://docs.nanobox.io/app-config/boxfile/), a yaml config file used to provision and configure your app's infrastructure when using Nanobox. This engine makes the following options available.

#### Overview of Boxfile Configuration Options
```yaml
code.build:
  config:
    # Go Settings
    runtime: go-1.6
    package: 'github.com/username/code'
    fetch: 'go get'
    build: 'go build'

    # Node.js Settings
    nodejs_runtime: nodejs-4.4
```

##### Quick Links
[Go Settings](#go-settings)  
[Node.js Settings](#nodejs-settings)

---

### Go Settings
The following setting allows you to define your Go runtime environment.

---

#### runtime
Specifies which Golang runtime to use. The following runtimes are available:

go-1.4
go-1.5
go-1.6 *(default)*

```yaml
code.build:
  config:
    runtime: go-1.6
```

---

#### package *(required)*
Specifies the path to the directory in which your code is stored. This can be a local or remote directory.

```yaml
code.build:
  config:
    package: 'github.com/username/code'
```

---

#### fetch
Defines the command to run to load dependencies in the build process.

```yaml
code.build:
  config:
    fetch: 'go get'
```

---

#### build
Defines the command to run to compile your code in the build process.

```yaml
code.build:
  config:
    build: 'go build'
```

---

### Node.js Settings
Many applications utilize Javascript tools in some way. This engine allows you to specify which Node.js runtime you'd like to use.

---

#### nodejs_runtime
Specifies which Node.js runtime and version to use. You can view the available Node.js runtimes in the [Node.js engine documentation](https://github.com/nanobox-io/nanobox-engine-nodejs#runtime).

```yaml
code.build:
  config:
    nodejs_runtime: nodejs-4.4
```

---

## Help & Support
This is a Go engine provided by [Nanobox](http://nanobox.io). If you need help with this engine, you can reach out to us in the [#nanobox IRC channel](http://webchat.freenode.net/?channels=nanobox). If you are running into an issue with the engine, feel free to [create a new issue on this project](https://github.com/nanobox-io/nanobox-engine-golang/issues/new).
