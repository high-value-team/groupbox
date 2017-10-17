# groupbox
Ad-hoc online collaboration on information collections

tools
```
brew install dropstack-cli
brew install go --with-cc-common # Installs go with cross compilation support
export GOROOT=/usr/local/Cellar/go/1.9.1/libexec # optional
```

run local
```
cd build
go build -o groupbox ../src
docker build -t groupbox .
docker run --rm -ti -p 8080:80 groupbox
```

run on dropstack
```
cd groupbox/build
export GROUPBOX_STACKLETNAME=my_cool_stackletname
./build.sh

cd groupbox/dropstack
dropstack login
dropstack deploy
```
