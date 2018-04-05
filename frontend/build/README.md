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
run start                           - Run frontend start scripts using env.development
run start:development               - Run frontend start scripts using env.development
run start:production                - Run frontend start scripts using env.production
run deploy                          - Create deploy folder and deploy to Dropstack
run deploy:clean                    - Remove all "deploy" folders
run build                           - Run frontend build scripts using env.production
run build:production                - Run frontend build scripts using env.production
run build:development               - Run frontend build scripts using env.development
run build:clean                     - Remove all "bin" folders
```

Execute `run` to list all available tasks
