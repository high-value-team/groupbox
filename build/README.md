# How to build groupbox
## Den build-Rechner präparieren
###Environment

1. `build/template.environment` kopieren nach `build/environment`. Die Datei `environment` ist bei gitignore ausgeschlossen.
2. Environment-Vars in `environment` passend füllen.
3. In `/build` aufrufen: `source environment`, um die Env-Vars auch wirklich zu setzen.

### Dropstack Client (OSX)
`brew install dropstack-cli`

### Golang (OSX)
```
# Installs go with cross compilation support
brew install go --with-cc-common 

# optional (check GOROOT with `set` if necessary)
export GOROOT=/usr/local/Cellar/go/1.9.1/libexec 
```

## Build und Deployment
### Lokal ohne Container

### Lokal im Container

### Dropstack







deploy to dropstack
```
make build
make deploy
```

run local
```
source build/environment # export environment variables
make local
./dropstack --mongodb-url=$MONGODB_URL
```

run local with docker
```
make build
cd dropstack
docker build -t groupbox .
docker run --rm -ti -p 8080:80 groupbox
```

