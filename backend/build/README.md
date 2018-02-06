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

create and configure environment files (e.g. .env.development)
```
cp examples/.env.dropstack .env.dropstack
cp examples/.env.development .env.development
cp examples/.env.production .env.production
```

## Run Tasks

```
run test   // execute tests
run start  // start backend on local machine
run build  // build backend into bin directory
run deploy // deploy to dropstack
```
