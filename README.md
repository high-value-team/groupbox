# groupbox
Ad-hoc online collaboration on information collections

![box](documentation/images-groupbox/02_box.png)

Read more about Groupbox: [Details](documentation/groupbox.md)

live demo: https://groupbox.sloppy.zone

## up and running

To quickly just spin up the application, please run the following command:

```
docker-compose up
```

then later visit http://localhost:8080



## frontend
```
cd frontend/build
yarn install
run // will display available 'run' tasks, e.g run docker:build
```

Read more about the build system: [read more ...](development/frontend/build/README.md)

Read more about Frontend Testing (Playground-UI): [read more ...](documentation/cosmos.md)

## backend
```
cd backend/build
yarn install
run // will display available 'run' tasks, e.g run docker:build
```

[read more about Backend build system](development/backend/build/README.md)
