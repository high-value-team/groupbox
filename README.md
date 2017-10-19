# groupbox
Ad-hoc online collaboration on information collections

![travis build status](https://travis-ci.org/high-value-team/groupbox.svg?branch=develop)


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

