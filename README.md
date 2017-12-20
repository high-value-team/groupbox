# groupbox
Ad-hoc online collaboration on information collections

![travis build status](https://travis-ci.org/high-value-team/groupbox.svg?branch=develop)


tools
```
npm install -g dropstack-cli
brew install go --with-cc-common # Installs go with cross compilation support
export GOROOT=/usr/local/Cellar/go/1.9.1/libexec # optional
```


deploy to dropstack
```
source build/environment # export environment variables
make docker
make deploy
```

run local
```
source build/environment # export environment variables
make build
./dropstack
```

run local with docker
```
source build/environment # export environment variables
make docker
make run
```

