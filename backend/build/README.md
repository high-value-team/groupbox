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

```
run setup                           - Create environment files, e.g. env.production. Please edit files with useful values!
run test                            - Run backend test scripts
run build                           - Run backend build scripts
run build:clean                     - Remove all "bin" folders
run start                           - Run backend start scripts using env.development
run start:development               - Run backend start scripts using env.development
run start:production                - Run backend start scripts using env.production
run deploy                          - Create deploy folder and deploy to Dropstack
run deploy:clean                    - Remove all "deploy" folders
```

Execute `run` to list all available tasks
