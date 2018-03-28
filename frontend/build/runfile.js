const { help, run } = require('runjs');
const dotenv = require('dotenv');
const dotenvExpand = require('dotenv-expand');
const fs = require('fs');

//
// tasks
//

function setup() {
    ['env.development', 'env.production', 'env.dropstack'].forEach((filename) => {
        if (!fs.existsSync(filename)) {
            run(`cp examples/${filename} .`);
        }
    });
}
help(start, 'Setup environment files; copy ./example/* to .; Please edit copied files with useful values based.');

function start () {
    const envObj = loadEnvironment('development.env');
    run(`cd ../src && yarn start`, {env: envObj});
}
help(start, 'Run frontend start scripts');

function build_production () {
    console.log('using env.production');
    _build('env.production');
}
help(build_production, 'Run frontend build scripts: production');

function build_development () {
    console.log('using env.development');
    _build('env.development');
}
help(build_development, 'Run frontend build scripts: development');

function _build (envPath) {
    const envObj = loadEnvironment(envPath);
    run(`cd ../src && yarn build`, {env: envObj});
    run(`rm -rf ../bin`);
    run(`cp -r ../src/build ../bin`);
}

function deploy () {
    // create deploy folder: ../deploy
    const productionEnv = loadEnvironment('production.env');
    run(`cd ../src && yarn build`, {env: productionEnv});
    run(`rm -rf ../deploy`);
    run(`cp -r ../src/build ../deploy`);

    // parse dropstack environment variables
    const dropstackEnv = loadEnvironment('dropstack.env');

    // build .dropstack.json file
    var file = fs.readFileSync('template.dropstack.json', 'utf8')
    var parsedFile = interpolate(file, dropstackEnv);
    fs.writeFileSync('../deploy/.dropstack.json', parsedFile);
    // run(`cat ../deploy/.dropstack.json`);

    // deploy to dropstack
    const cmd = `cd ../deploy && dropstack deploy --compress --verbose --alias ${dropstackEnv.DROPSTACK_ALIAS}.cloud.dropstack.run --token ${dropstackEnv.DROPSTACK_TOKEN}`;
    run(cmd);
}
help(deploy, 'Run frontend deployment scripts');


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

//
// exports
//

module.exports = {
    setup,
    start,
    deploy,

    // examples:
    // run build
    // run build:production
    // run build:development
    'build': build_production,
    'build:production': build_production,
    'build:development': build_development,
}