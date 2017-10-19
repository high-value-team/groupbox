# How to build groupbox
## Den build-Rechner präparieren
###Environment

1. `build/template.environment` kopieren nach `build/environment`. Die Datei `environment` ist bei gitignore ausgeschlossen.
2. Environment-Vars in `environment` passend füllen.

### Dropstack Client (OSX)
```
brew install dropstack-cli
```

### Golang (OSX)
```
# Installs go with cross compilation support
brew install go --with-cc-common 

# optional (check GOROOT with `set` if necessary)
export GOROOT=/usr/local/Cellar/go/1.9.1/libexec 
```

## Build und Deployment
### Lokal ohne Container
```
# in /groupbox
source build/environment

make local
./groupbox
```
Anschließend läuft die Anwendung lokal auf [http://localhost:8080](http://localhost:8080).

### Lokal im Container
```
make build
make run
```
Anschließend läuft die Anwendung lokal auf [http://localhost:9090](http://localhost:9090).

### Dropstack
```
# in /groupbox
source build/environment

make build
make deploy
```
Anschließend läuft die Anwendung in einem Stacklet.