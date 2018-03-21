const { help, run } = require('runjs');
const dotenv = require('dotenv');
const dotenvExpand = require('dotenv-expand');
const fs = require('fs');

//
// tasks
//

function start () {
    const envObj = loadEnvironment('development');
    run(`cd ../src && yarn start`, {env: envObj});
}
help(start, 'Run frontend start scripts');

// TODO
// build:production
// build:development
// (default: build:production)
function build () {
    const envObj = loadEnvironment('production');
    run(`cd ../src && yarn build`, {env: envObj});
    run(`rm -rf ../bin`);
    run(`cp -r ../src/build ../bin`);
}
help(build, 'Run frontend build scripts');

// initial dropstack deploy:
// export DROPSTACK_TOKEN=xxx
// export DROPSTACK_ALIAS=groupbox
// dropstack deploy --compress --verbose --alias $DROPSTACK_ALIAS.cloud.dropstack.run --token $DROPSTACK_TOKEN
// => replace DROPSTACK_TOKEN, DROPSTACK_ALIAS, DROPSTACK_NAME in '.env.dropstack' file!
function deploy () {
    // TODO create tmp.dropstack.deploy folder and work from there (cp ../bin ../tmp.dropstack.deploy && write .dropstack.json)

    // create deploy folder
    const productionEnv = loadEnvironment('production');
    run(`cd ../src && yarn build`, {env: productionEnv});
    run(`rm -rf ../deploy`);
    run(`cp -r ../src/build ../deploy`);

    // parse dropstack environment variables
    const dropstackEnv = loadEnvironment('dropstack');

    // build .dropstack.json file
    var file = fs.readFileSync('templates/.dropstack.json.template', 'utf8'); // read .dropstack.json.template
    var parsedFile = interpolate(file, dropstackEnv); // interpolate with envObj
    fs.writeFileSync('../deploy/.dropstack.json', parsedFile); // write .dropstack.json
    // run(`cat ../deploy/.dropstack.json`);

    // deploy to dropstack
    const cmd = `cd ../deploy && dropstack deploy --compress --verbose --alias ${dropstackEnv.DROPSTACK_ALIAS}.cloud.dropstack.run --token ${dropstackEnv.DROPSTACK_TOKEN}`;
    run(cmd);
    // console.log(`cmd:${cmd}`);
}
help(build, 'Run frontend deployment scripts');

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

function loadEnvironment(envName) {
    var env = dotenv.config({path: `.env.${envName}`});
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
    build,
    start,
    deploy,
}