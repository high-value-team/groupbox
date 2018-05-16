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

function install() {
    run(`cd ../src && yarn install`);
}
help(install, 'Install all dependencies in "src" folder');

//
// test
//

function cosmos() {
    run(`cd ../src && yarn cosmos`);
}
help(cosmos, 'Start cosmos server for Playground-UI testing');

//
// local
//

function local_development () {
    const envFile = 'env.development';
    const envObj = loadEnvironment(envFile);
    console.log(`using ${envFile}`);
    run(`cd ../src && yarn start`, {env: envObj});
}
help(local_development, 'Run frontend start scripts using env.development');

function local_production () {
    const envFile = 'env.production';
    const envObj = loadEnvironment(envFile);
    console.log(`using ${envFile}`);
    run(`cd ../src && yarn start`, {env: envObj});
}
help(local_production, 'Run frontend start scripts using env.production');

//
// docker
//

function docker_prepare () {
    const envFile = 'env.frontend';
    console.log(`using ${envFile}`);

    const binDir = `docker.${timestamp()}`;
    const envObj = loadEnvironment(envFile);

    // build frontend with ENV placeholders
    run(`mkdir -p ${binDir}/app`);
    run(`cd ../src && yarn build`, {env: envObj});
    run(`cp -r ../src/build/* ${binDir}/app`);

    // copy files
    run(`cp template.nginx.default.conf ${binDir}/nginx.default.conf`);
    run(`cp template.nginx.Dockerfile ${binDir}/Dockerfile`);
    run(`cp template.nginx.replace.sh ${binDir}/replace.sh`);
    run(`cp template.nginx.run.sh ${binDir}/run.sh`);
}
help(docker_prepare, 'Build project and prepare Dockerfile');

function docker_build () {
    const imageName = 'hvt1/groupbox-frontend';

    docker_prepare();

    const binPath = findNewestDockerFolder();
    if (binPath === undefined) {
        console.log('No bin-folder found. Please execute a "run docker:prepare" job first!');
        return
    }

    run(`docker build --tag ${imageName} ${binPath}`);
}
help(docker_build, 'Build frontend and build docker image');

function docker_publish () {
    const imageName = 'hvt1/groupbox-frontend';
    const envFile = 'env.dockerhub';
    console.log(`using ${envFile}`);
    const envObj = loadEnvironment(envFile);

    runSilent(`docker login --username ${envObj.USERNAME} --password ${envObj.PASSWORD}`, {stdio: 'pipe'});
    run(`docker push ${imageName}`);
}
help(docker_publish, 'Push latest docker build to docker hub');

function docker_start () {
    const localURL = '127.0.0.1:9010';
    const imageName = 'hvt1/groupbox-frontend';
    const containerName = 'groupbox-frontend';
    const envFile = 'env.development';
    console.log(`using ${envFile}`);
    const envObj = loadEnvironment(envFile);
    const envArgs = toDockerEnvironmentArgs(envObj);

    const isRunning = checkIsRunning(containerName);
    if (isRunning) {
        console.log(`The docker container '${containerName}' is already running. Make sure to stop it! ('run docker:stop')`);
        return
    }
    run(`docker run --name=${containerName} --publish "${localURL}:80" ${envArgs} ${imageName}`);
}

help(docker_start, 'Start docker container');

function docker_stop () {
    const containerName = 'groupbox-frontend';
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

function sloppy_delete() {
    const envFile = 'env.sloppy';
    console.log(`using ${envFile}`);
    const envObj = loadEnvironment(envFile);

    run(`sloppy delete groupbox-frontend`, {env: envObj});
}
help(sloppy_delete, 'Delete existing project on sloppy.zone');

function sloppy_deploy() {
    const envFile = 'env.sloppy';
    console.log(`using ${envFile}`);
    const envObj = loadEnvironment(envFile);

    const deployDir = `sloppy.${timestamp()}`;
    run(`mkdir ${deployDir}`);
    run(`cp template.sloppy.yml ${deployDir}/sloppy.yml`);
    run(`sloppy start ${deployDir}/sloppy.yml`, {env: envObj});
}
help(sloppy_deploy, 'Deploy to sloppy.zone');

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


function clean_dropstack() {
    fs.readdirSync('.').forEach(file => {
        if (fs.statSync(file).isDirectory() && file.match(/^dropstack\.[0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{6}$/)) {
            const removeDirectory = `rm -rf ${file}`;
            run(removeDirectory);
        }
    })
}
help(clean_dropstack, 'Remove all "dropstack" folders');

function clean_sloppy() {
    fs.readdirSync('.').forEach(file => {
        if (fs.statSync(file).isDirectory() && file.match(/^sloppy\.[0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{6}$/)) {
            const removeDirectory = `rm -rf ${file}`;
            run(removeDirectory);
        }
    })
}
help(clean_sloppy, 'Remove all "sloppy" folders');

function clean_install() {
    run(`cd ../src && rm -rf node_modules`);
}
help(clean_install, 'Remove installed libraries in "src" folder');

//
// drone
//

function build_for_drone() {
    clean_install();
    install();
    setup();
    docker_prepare();
    move_latest_docker_prepare_to_bin();
}
help(build_for_drone, 'Create bin directory with all artefacts for creating a docker image in the Drone-CI workflow.');

function move_latest_docker_prepare_to_bin() {
    const binPath = findNewestDockerFolder();
    if (binPath === undefined) {
        console.log('No bin-folder found. Please execute a "run docker:prepare" job first!');
        return
    }

    run(`rm -rf ../bin`);
    run(`mv ${binPath} ../bin`);
}

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

function findNewestDockerFolder() {
    let binFolders = [];
    fs.readdirSync('.').forEach(file => {
        if (fs.statSync(file).isDirectory() && file.match(/^docker\.[0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{6}$/)) {
            const absolutePath = `${__dirname}/${file}`;
            binFolders.push(absolutePath);
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
    // console.log();

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
    install,
    cosmos,

    'local': local_development,
    'local:development': local_development,
    'local:production': local_production,

    'docker:build': docker_build,
    'docker:publish': docker_publish,
    'docker:start': docker_start,
    'docker:stop': docker_stop,

    'sloppy:delete': sloppy_delete,
    'sloppy:deploy': sloppy_deploy,

    'clean:docker': clean_docker,
    'clean:sloppy': clean_sloppy,
    'clean:install': clean_install,

    'build_for_drone': build_for_drone,

    // timestamp: () => console.log(timestamp()),
};