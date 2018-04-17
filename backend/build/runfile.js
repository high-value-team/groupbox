const { help, run } = require('runjs');
const dotenv = require('dotenv');
const dotenvExpand = require('dotenv-expand');
const fs = require('fs');
const { SilentLogger } = require('runjs/lib/common');

//
// decorators
//

function runSilent(command, options = {}) {
    run(command, options, new SilentLogger());
}

//
// setup
//

function setup() {
    let envFiles = [];
    fs.readdirSync('setup').forEach(file => {
        if (file.startsWith('env.')) {
            envFiles.push(file);
        }
    });

    let allEnvFilesExist = true;
    envFiles.forEach((filename) => {
        if (!fs.existsSync(filename)) {
            allEnvFilesExist = false;
        }
    });

    if (allEnvFilesExist) {
        console.log("All environment files already exist.")
    }

    envFiles.forEach((filename) => {
        if (!fs.existsSync(filename)) {
            run(`cp setup/${filename} .`);
            console.log(`Please edit this file!`)
        }
    });
}
help(setup, 'Create environment files, e.g. env.production. Please edit files with useful values!');

//
// test
//

function test_unit () {
    const envFile = 'env.test.unit';
    console.log(`using ${envFile}`);

    // perpare environment variables
    let envObj = loadEnvironment(envFile);
    envObj.GOROOT = process.env.GOROOT;
    envObj.GOPATH = process.env.GOPATH;

    // prepare directory list
    let dirList = run(`go list ../src/... | grep -v vendor | grep -v playground`, {stdio: 'pipe'})
        .split('\n')
        .filter(dir => dir.length > 0); // remove empty
    const prefix = dirList[0].replace(/src$/,"");
    dirList = dirList.map((dir) => dir.replace(prefix, '../'));

    run(`go test -v --tags=unit ${dirList.join(' ')}`, {env: envObj});
}
help(test_unit, 'Run backend unit tests');

function test_mongo () {
    const envFile = 'env.test.mongo';
    console.log(`using ${envFile}`);

    // perpare environment variables
    let envObj = loadEnvironment(envFile);
    envObj.GOROOT = process.env.GOROOT;
    envObj.GOPATH = process.env.GOPATH;

    // prepare directory list
    let dirList = run(`go list ../src/... | grep -v vendor | grep -v playground`, {stdio: 'pipe'})
        .split('\n')
        .filter(dir => dir.length > 0); // remove empty
    const prefix = dirList[0].replace(/src$/,"");
    dirList = dirList.map((dir) => dir.replace(prefix, '../'));

    run(`go test -v -tags=mongo ${dirList.join(' ')}`, {env: envObj});
}
help(test_mongo, 'Run backend mongo tests');

function test_smtp () {
    const envFile = 'env.test.smtp';
    console.log(`using ${envFile}`);

    // perpare environment variables
    let envObj = loadEnvironment(envFile);
    envObj.GOROOT = process.env.GOROOT;
    envObj.GOPATH = process.env.GOPATH;

    // prepare directory list
    let dirList = run(`go list ../src/... | grep -v vendor | grep -v playground`, {stdio: 'pipe'})
        .split('\n')
        .filter(dir => dir.length > 0); // remove empty
    const prefix = dirList[0].replace(/src$/,"");
    dirList = dirList.map((dir) => dir.replace(prefix, '../'));

    run(`go test -v -tags=smtp ${dirList.join(' ')}`, {env: envObj});
}
help(test_smtp, 'Run backend smtp tests');

//
// local
//

function local_development () {
    const envFile = 'env.development';
    console.log(`using ${envFile}`);
    const envObj = loadEnvironment(envFile);
    const binDir = `local.${timestamp()}`;

    // build go executable
    const gitTag = run(`git describe --always --tags --dirty="*"`, {stdio: 'pipe'}).trim();
    run(`go build -ldflags "-X main.VersionNumber=${gitTag}" -o ${binDir}/groupbox-backend ../src`);

    // start go executable
    run(`${binDir}/groupbox-backend`, {env: envObj});
}
help(local_development, 'Build and start go-executable using env.development');

function local_production () {
    const envFile = 'env.production';
    console.log(`using ${envFile}`);
    const envObj = loadEnvironment(envFile);
    const binDir = `local.${timestamp()}`;

    // build go executable
    const gitTag = run(`git describe --always --tags --dirty="*"`, {stdio: 'pipe'}).trim();
    run(`go build -ldflags "-X main.VersionNumber=${gitTag}" -o ${binDir}/groupbox-backend ../src`);

    // start go executable
    run(`${binDir}/groupbox-backend`, {env: envObj});
}
help(local_production, 'Build and start go-executable using env.production');

//
// docker
//

function docker_build () {
    const binDir = `docker.${timestamp()}`;
    const imageName = 'hvt1/groupbox-backend';

    // build go executable
    const gitTag = run(`git describe --always --tags --dirty="*"`, {stdio: 'pipe'}).trim();
    const buildCommand = `go build -ldflags "-X main.VersionNumber=${gitTag}" -o ${binDir}/groupbox-backend ../src`;
    const envObj = { GOOS: "linux", GOARCH: "amd64", GOROOT: process.env.GOROOT, GOPATH: process.env.GOPATH};
    run(buildCommand, {env:envObj});

    // create Dockerfile
    var dockerfile = fs.readFileSync('template.docker.Dockerfile', 'utf8');
    fs.writeFileSync(`${binDir}/Dockerfile`, dockerfile);

    // build docker image
    run(`docker build --tag ${imageName} ${binDir}`);
}
help(docker_build, 'Build go executable with docker build flags and build docker image');

function docker_start () {
    const localURL = `127.0.0.1:8091`;
    const imageName = 'hvt1/groupbox-backend';
    const containerName = 'groupbox-backend';
    const envFile = 'env.development';
    console.log(`using ${envFile}`);
    const envObj = loadEnvironment(envFile);
    const envArgs = toDockerEnvironmentArgs(envObj);

    const isRunning = checkIsRunning(containerName);
    if (isRunning) {
        console.log(`The docker container '${containerName}' is already running. Make sure to stop it! ('run docker:stop')`);
        return
    }
    run(`docker run --name=${containerName} --publish "${localURL}:8080" ${envArgs} ${imageName}`);
}

help(docker_start, 'Start docker container');

function docker_stop () {
    const containerName = 'groupbox-backend';
    console.log(`checking if '${containerName}'-container is running`);
    let containerID = run(`docker ps --filter "name=${containerName}" --format "{{.ID}}"`, {stdio: 'pipe'});
    if (containerID.length !== 0) {
        run(`docker kill ${containerName}`);
    }
    console.log();

    console.log(`checking if '${containerName}'-container exists`);
    containerID = run(`docker ps --all --filter "name=${containerName}" --format "{{.ID}}"`, {stdio: 'pipe'});
    if (containerID.length !== 0) {
        run(`docker rm ${containerName}`);
    }
    console.log();
}
help(docker_stop, 'Stop docker container');

//
// sloppy
//

function sloppy_publish () {
    const imageName = 'hvt1/groupbox-backend';
    const envFile = 'env.dockerhub';
    console.log(`using ${envFile}`);
    const envObj = loadEnvironment(envFile);

    runSilent(`docker login --username ${envObj.USERNAME} --password ${envObj.PASSWORD}`, {stdio: 'pipe'});
    run(`docker push ${imageName}`);
}
help(sloppy_publish, 'Push latest docker build to docker hub');

function sloppy_delete() {
    const envFile = 'env.sloppy';
    console.log(`using ${envFile}`);
    const envObj = loadEnvironment(envFile);

    run(`sloppy delete groupbox-backend`, {env: envObj});
}
help(sloppy_delete, 'Delete existing project on sloppy.zone');

function sloppy_deploy() {
    const envFile = 'env.sloppy';
    console.log(`using ${envFile}`);
    const sloppyEnv = loadEnvironment(envFile);

    const deployDir = `sloppy.${timestamp()}`;
    run(`mkdir ${deployDir}`);

    const productionEnv = loadEnvironment('env.production');
    var sloppyTemplate= fs.readFileSync('template.sloppy.yml', 'utf8');
    var sloppyFileContent = interpolate(sloppyTemplate, productionEnv);
    fs.writeFileSync(`${deployDir}/sloppy.yml`, sloppyFileContent);

    run(`sloppy start ${deployDir}/sloppy.yml`, {env: sloppyEnv});
}
help(sloppy_deploy, 'Deploy to sloppy.zone');

//
// dropstack
//

function dropstack_build() {
    const binDir = `dropstack.${timestamp()}`;

    // build go executable
    const gitTag = run(`git describe --always --tags --dirty="*"`, {stdio: 'pipe'}).trim();
    const buildCommand = `go build -ldflags "-X main.VersionNumber=${gitTag}" -o ${binDir}/groupbox-backend ../src`;
    const envObj = { GOOS: "linux", GOARCH: "amd64", GOROOT: process.env.GOROOT, GOPATH: process.env.GOPATH};
    run(buildCommand, {env:envObj});

    // create .dropstack.json
    const dropstackEnv = loadEnvironment('env.dropstack');
    var file = fs.readFileSync('template.dropstack.json', 'utf8')
    var parsedFile = interpolate(file, dropstackEnv);
    fs.writeFileSync(`${binDir}/.dropstack.json`, parsedFile);

    // create Dockerfile
    const productionEnv = loadEnvironment('env.production');
    var dropstackTemplate = fs.readFileSync('template.dropstack.Dockerfile', 'utf8');
    var dropstackDockerfile = interpolate(dropstackTemplate, productionEnv);
    fs.writeFileSync(`${binDir}/Dockerfile`, dropstackDockerfile);
}
help(dropstack_build, 'Create Dropstack folder');

function dropstack_deploy() {
    const binPath = findNewestDropstacklFolder();
    if (binPath === undefined) {
        console.log('No bin-folder found. Please execute a "build" job first!');
        return
    }

    const dropstackEnv = loadEnvironment('env.dropstack');
    const deployToDropstack = `cd ${binPath} && dropstack deploy --compress --verbose --alias ${dropstackEnv.DROPSTACK_ALIAS}.cloud.dropstack.run --token ${dropstackEnv.DROPSTACK_TOKEN}`;
    run(deployToDropstack);
}
help(dropstack_deploy, 'Deploy to Dropstack');

//
// clean
//

function clean_docker() {
    fs.readdirSync('.').forEach(file => {
        if (fs.statSync(file).isDirectory() && file.match(/^docker\.[0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{6}$/)) {
            const removeDirectory = `rm -rf ${file}`;
            run(removeDirectory);
        }
    });

}
help(clean_docker, 'Remove all "docker" folders');

function clean_sloppy() {
    fs.readdirSync('.').forEach(file => {
        if (fs.statSync(file).isDirectory() && file.match(/^sloppy\.[0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{6}$/)) {
            const removeDirectory = `rm -rf ${file}`;
            run(removeDirectory);
        }
    })
}
help(clean_sloppy, 'Remove all "sloppy" folders');

function clean_local() {
    fs.readdirSync('.').forEach(file => {
        if (fs.statSync(file).isDirectory() && file.match(/^local\.[0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{6}$/)) {
            const removeDirectory = `rm -rf ${file}`;
            run(removeDirectory);
        }
    })
}
help(clean_local, 'Remove all "sloppy" folders');

//
// helper
//

function interpolate(text, env) {
    var matches = text.match(/\$([a-zA-Z0-9_]+)|\${([a-zA-Z0-9_]+)}/g) || [];

    matches.forEach(function (match) {
        var key = match.replace(/\$|{|}/g, '');
        var value = env[key];
        text = text.replace(match, value);
    });

    return text;
}

function loadEnvironment(envPath) {
    var env = dotenv.config({path: envPath});
    // console.log(`loadEnvironment:${JSON.stringify(env, null, 2)}`);

    // override 'system environment variables'
    for (var k in env.parsed) {
        process.env[k] = env.parsed[k];
    }

    env = dotenvExpand(env); // will auto merge with 'system environment variables'
    if (env.error) {
        throw env.error
    }
    // console.log(`loadEnvironment dotenvExpand:${JSON.stringify(env, null, 2)}`);

    return env.parsed;
}

function findNewestDropstacklFolder() {
    let binFolders = [];
    fs.readdirSync('.').forEach(file => {
        if (fs.statSync(file).isDirectory() && file.match(/^dropstack\.[0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{6}$/)) {
            binFolders.push(file);
        }
    });
    let sorted = binFolders.sort();
    return sorted[sorted.length-1];
}

function timestamp() {
    const pad = (n) => String("00" + n).slice(-2);
    const date = new Date();
    return `${date.getFullYear()}-${pad(date.getMonth()+1)}-${pad(date.getDate())}-${pad(date.getHours())}${pad(date.getMinutes())}${pad(date.getSeconds())}`
}

function checkIsRunning(containerName) {
    let containerID = run(`docker ps --filter "name=${containerName}" --format "{{.ID}}"`, {stdio: 'pipe'});
    if (containerID.length !== 0) {
        return true;
    }

    containerID = run(`docker ps --all --filter "name=${containerName}" --format "{{.ID}}"`, {stdio: 'pipe'});
    if (containerID.length !== 0) {
        return true;
    }

    return false
}

function toDockerEnvironmentArgs(envObj) {
    var args = '';
    var keys = Object.keys(envObj);
    keys.forEach((key) => {
        const value = envObj[key];
        args = args.concat(`--env "${key}=${value}" `);
    });
    return args;
}

//
// exports
//

module.exports = {
    setup,

    test: test_unit,
    'test:unit': test_unit,
    'test:mongo': test_mongo,
    'test:smtp': test_smtp,

    'local': local_development,
    'local:development': local_development,
    'local:production': local_production,

    'docker:build': docker_build,
    'docker:start': docker_start,
    'docker:stop': docker_stop,

    'sloppy:publish': sloppy_publish,
    'sloppy:delete': sloppy_delete,
    'sloppy:deploy': sloppy_deploy,

    'dropstack:build': dropstack_build,
    'dropstack:deploy': dropstack_deploy,

    'clean:local': clean_local,
    'clean:docker': clean_docker,
    'clean:sloppy': clean_sloppy,
};