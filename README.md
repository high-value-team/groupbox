# groupbox
Ad-hoc online collaboration on information collections

tools
```
brew install dropstack-cli
brew install go --with-cc-common # Installs go with cross compilation support
export GOROOT=/usr/local/Cellar/go/1.9.1/libexec # optional
```


deploy to dropstack
```
make build
make deploy
```

run local
```
make local
./dropstack
```

run local with docker
```
make build
cd dropstack
docker build -t groupbox .
docker run --rm -ti -p 8080:80 groupbox
```