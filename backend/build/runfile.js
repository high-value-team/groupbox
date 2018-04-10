const { help, run } = require('runjs');
const dotenv = require('dotenv');
const dotenvExpand = require('dotenv-expand');
const fs = require('fs');
const path = require("path");

//
// tasks
//

function setup() {
    const envFiles = ['env.development', 'env.production', 'env.dropstack', 'env.test.unit', 'env.test.mongo', 'env.test.smtp'];
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
            run(`cp examples/${filename} .`);
            console.log(`Please edit this file!`)
        }
    });
}
help(setup, 'Create environment files, e.g. env.production. Please edit files with useful values!');

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

function trimPrefix(str, prefix) {
    if (str.startsWith(prefix)) {
        return str.slice(prefix.length)
    } else {
        return str
    }
}

function build () {
    const gitTag = run(`git describe --always --tags --dirty="*"`, {stdio: 'pipe'}).trim();
    run(`go build -ldflags "-X main.VersionNumber=${gitTag}" -o bin.${timestamp()}/groupbox-backend ../src`);
}
help(build, 'Run backend build scripts');

function build_clean() {
    fs.readdirSync('.').forEach(file => {
        if (fs.statSync(file).isDirectory() && file.match(/^bin\.[0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{6}$/)) {
            const removeDirectory = `rm -rf ${file}`;
            run(removeDirectory);
        }
    });

}
help(build_clean, 'Remove all "bin" folders');

function start_development () {
    const envFile = 'env.development';
    console.log(`using ${envFile}`);
    _start(envFile);
}
help(start_development, 'Run backend start scripts using env.development');

function start_production () {
    const envFile = 'env.production';
    console.log(`using ${envFile}`);
    _start(envFile);
}
help(start_production, 'Run backend start scripts using env.production');

function _start (envFile) {
    const envObj = loadEnvironment(envFile);

    build();

    const binPath = findNewestBinFolder();
    if (binPath === undefined) {
        console.log('No bin-folder found. Please execute a "build" job first!');
        return
    }

    run(`${binPath}/groupbox-backend`, {env: envObj});
}

function deploy () {
    const envFile = 'env.production';
    console.log(`using ${envFile}`);

    const binPath = findNewestBinFolder();
    if (binPath === undefined) {
        console.log('No bin-folder found. Please execute a "build" job first!');
        return
    }

    const deployPath = `deploy.${timestamp()}`;
    const createDeployFolder = `cp -r ${binPath} ${deployPath}`;
    run(createDeployFolder);

    // create and write .dropstack.json in deploy folder
    const dropstackEnv = loadEnvironment('env.dropstack');
    var file = fs.readFileSync('template.dropstack.json', 'utf8')
    var parsedFile = interpolate(file, dropstackEnv);
    fs.writeFileSync(`${deployPath}/.dropstack.json`, parsedFile);

    // create Dockerfile
    const productionEnv = loadEnvironment(envFile);
    var dockerfile = fs.readFileSync('template.Dockerfile', 'utf8');
    var parsedDockerfile = interpolate(dockerfile, productionEnv);
    fs.writeFileSync(`${deployPath}/Dockerfile`, parsedDockerfile);
    // run(`cat `${deployPath}/Dockerfile``);

    console.log(`envrionment:\n${JSON.stringify(productionEnv, null, 2)}`);
    const deployToDropstack = `cd ${deployPath} && dropstack deploy --compress --verbose --alias ${dropstackEnv.DROPSTACK_ALIAS}.cloud.dropstack.run --token ${dropstackEnv.DROPSTACK_TOKEN}`;
    run(deployToDropstack);
}
help(deploy, 'Create deploy folder and deploy to Dropstack');

function deploy_clean() {
    fs.readdirSync('.').forEach(file => {
        if (fs.statSync(file).isDirectory() && file.match(/^deploy\.[0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{6}$/)) {
            const removeDirectory = `rm -rf ${file}`;
            run(removeDirectory);
        }
    })
}
help(deploy_clean, 'Remove all "deploy" folders');

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

function findNewestBinFolder() {
    let binFolders = [];
    fs.readdirSync('.').forEach(file => {
        if (fs.statSync(file).isDirectory() && file.match(/^bin\.[0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{6}$/)) {
            binFolders.push(file);
        }
    });
    let sorted = binFolders.sort();
    return sorted[sorted.length-1];
}

function timestamp() {
    const pad = (n) => String("00" + n).slice(-2);
    const date = new Date();
    return `${date.getFullYear()}-${pad(date.getMonth()+1)}-${pad(date.getDay())}-${pad(date.getHours())}${pad(date.getMinutes())}${pad(date.getSeconds())}`
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

    build,
    'build:clean': build_clean,

    start: start_development,
    'start:development': start_development,
    'start:production': start_production,

    deploy,
    'deploy:clean': deploy_clean,
}