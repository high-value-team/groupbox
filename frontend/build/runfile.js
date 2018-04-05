const { help, run } = require('runjs');
const dotenv = require('dotenv');
const dotenvExpand = require('dotenv-expand');
const fs = require('fs');

//
// tasks
//

function setup() {
    const envFiles = ['env.development', 'env.production', 'env.dropstack'];
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

function start_development () {
    const envFile = 'env.development';
    const envObj = loadEnvironment(envFile);
    console.log(`using ${envFile}`);
    run(`cd ../src && yarn start`, {env: envObj});
}
help(start_development, 'Run frontend start scripts using env.development');

function start_production () {
    const envFile = 'env.production';
    const envObj = loadEnvironment(envFile);
    console.log(`using ${envFile}`);
    run(`cd ../src && yarn start`, {env: envObj});
}
help(start_production, 'Run frontend start scripts using env.production');

function build_production () {
    const envFile = 'env.production';
    console.log(`using ${envFile}`);
    _build(envFile);
}
help(build_production, 'Run frontend build scripts using env.production');

function build_development () {
    const envFile = 'env.development';
    console.log(`using ${envFile}`);
    _build(envFile);
}
help(build_development, 'Run frontend build scripts using env.development');

function _build (envPath) {
    const envObj = loadEnvironment(envPath);
    run(`cd ../src && yarn build`, {env: envObj});
    run(`cp -r ../src/build bin.${timestamp()}`);
}

function build_clean() {
    fs.readdirSync('.').forEach(file => {
        if (fs.statSync(file).isDirectory() && file.match(/^bin\.[0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{6}$/)) {
            const removeDirectory = `rm -rf ${file}`;
            run(removeDirectory);
        }
    });

}
help(build_clean, 'Remove all "bin" folders');

function deploy () {
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

    start: start_development,
    'start:development': start_development,
    'start:production': start_production,

    deploy,
    'deploy:clean': deploy_clean,

    'build': build_production,
    'build:production': build_production,
    'build:development': build_development,
    'build:clean': build_clean,
}