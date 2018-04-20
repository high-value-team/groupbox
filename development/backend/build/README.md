# Build Scripts

## Prerequisites

* node v9.5.0
* npm v5.6.0
* yarn v0.23.4

```
brew install node@9
npm install -g npm@5.6.0
npm install -g yarn@0.23.4
```

install build dependencies
```
yarn install
```

## Run Tasks

`run [taskname]`

e.g. `run start`

Available tasks:
```
setup                           - Create environment files, e.g. env.production. Please edit files with useful values!
test                            - Run backend unit tests
test:unit                       - Run backend unit tests
test:mongo                      - Run backend mongo tests
test:smtp                       - Run backend smtp tests
local                           - Build and start go-executable using env.development
local:development               - Build and start go-executable using env.development
local:production                - Build and start go-executable using env.production
docker:build                    - Build go executable with docker build flags and build docker image
docker:start                    - Start docker container
docker:stop                     - Stop docker container
sloppy:publish                  - Push latest docker build to docker hub
sloppy:delete                   - Delete existing project on sloppy.zone
sloppy:deploy                   - Deploy to sloppy.zone
dropstack:build                 - Create Dropstack folder
dropstack:deploy                - Deploy to Dropstack
clean:local                     - Remove all "sloppy" folders
clean:docker                    - Remove all "docker" folders
clean:sloppy                    - Remove all "sloppy" folders
```

Execute `run` to list all available tasks
