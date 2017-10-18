# Groupbox Frontend

## Build instructions

If necessary, install yarn

    brew install yarn

Then

    cd src/frontend
    yarn build

Finally generate the go source code to embed the app

    cd ..
    go generate
