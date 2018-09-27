# Build Scripts

## Prerequisites

* node v9.5.0
* npm v5.6.0
* yarn v0.23.4

```
brew install node@9
npm install -g npm@5.6.0
npm install -g yarn@0.23.4
npm install -g runjs@4.3.0
```

install build dependencies
```
yarn install
```

## Let's get started

getting started (initial setup)
```
run setup             # prepare config files + edit these files (e.g. env.development, env.sloppy, etc)
run install           # install node dependencies
run local:development # run locally with development configuration
run local:production  # run locally with production configuration
```

run docker container on local machine
```
run docker:build # build docker image
run docker:start # start docker container
run docker:stop  # stop docker container
```

deploy app to sloppy
```
run docker:build   # build docker image
run docker:publish # publish docker image to dockerhub
run sloppy:delete  # existing projects need to be deleted first
run sloppy:deploy  # deploy to sloppy
```

## All Run Tasks

To display all available 'run tasks' just execute `run` on the commandline

`run [taskname]`

e.g. `run start`

Available tasks:
```
setup                           - Create environment files, e.g. env.production. Please edit files with useful values!
install                         - Install all dependencies in "src" folder
cosmos                          - Start cosmos server for Playground-UI testing
local                           - Run frontend start scripts using env.development
local:development               - Run frontend start scripts using env.development
local:production                - Run frontend start scripts using env.production
docker:build                    - Build frontend and build docker image
docker:publish                  - Push latest docker build to docker hub
docker:start                    - Start docker container
docker:stop                     - Stop docker container
sloppy:delete                   - Delete existing project on sloppy.zone
sloppy:deploy                   - Deploy to sloppy.zone
clean:docker                    - Remove all "docker" folders
clean:sloppy                    - Remove all "sloppy" folders
clean:install                   - Remove installed libraries in "src" folder
```

