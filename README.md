# groupbox
Ad-hoc online collaboration on information collections


run local
```
cd build
go build -o groupbox ../src
docker build -t groupbox .
docker run --rm -ti -p 80:8080 groupbox
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
