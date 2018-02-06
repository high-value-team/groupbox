const { help, run } = require('runjs');
const dotenv = require('dotenv');
const dotenvExpand = require('dotenv-expand');
const fs = require('fs');

//
// tasks
//

function test () {
    let dirList = run(`go list ../src/... | grep -v vendor | grep -v playground`, {stdio: 'pipe'})
        .split('\n')
        .filter(dir => dir.length > 0); // remove empty
    const envObj = loadEnvironment('development');
    run(`go test -v ${dirList.join(' ')}`, {env: envObj});
    // TODO check return code
}
help(test, 'Run backend test scripts');

function start () {
    build();
    const envObj = loadEnvironment('development');
    run(`../bin/groupbox-backend`, {env: envObj})
}
help(start, 'Run backend start scripts');

function build () {
    const gitTag = run(`git describe --always --tags --dirty="*"`, {stdio: 'pipe'}).trim();
    run(`go build -ldflags "-X main.VersionNumber=${gitTag}" -o ../bin/groupbox-backend ../src`);
}
help(build, 'Run backend build scripts');

// TODO add flag to keep deploy directory, otherwise delete this directory
function deploy () {
    run(`rm -rf ../deploy`);
    run(`mkdir -p ../deploy`);

    // create .dropstack.json
    const dropstackEnv = loadEnvironment('dropstack');
    var file = fs.readFileSync('templates/.dropstack.json.template', 'utf8'); // read .dropstack.json.template
    var parsedFile = interpolate(file, dropstackEnv); // interpolate with envObj
    fs.writeFileSync('../deploy/.dropstack.json', parsedFile); // write .dropstack.json
    // run(`cat ../deploy/.dropstack.json`);

    // create Dockerfile
    const productionEnv = loadEnvironment('production');
    var dockerfile = fs.readFileSync('templates/Dockerfile.template', 'utf8');
    var parsedDockerfile = interpolate(dockerfile, productionEnv);
    fs.writeFileSync('../deploy/Dockerfile', parsedDockerfile);
    // run(`cat ../deploy/Dockerfile`);

    // build executable
    const gitTag = run(`git describe --always --tags --dirty="*"`, {stdio: 'pipe'}).trim();
    run(`GOOS=linux GOARCH=amd64 go build -ldflags "-X main.VersionNumber=${gitTag}" -o ../deploy/groupbox ../src`);

    // deploy
    console.log(`productionEnv:${JSON.stringify(productionEnv, null, 2)}`);
    const cmd = `cd ../deploy && dropstack deploy --compress --verbose --alias ${dropstackEnv.DROPSTACK_ALIAS}.cloud.dropstack.run --token ${dropstackEnv.DROPSTACK_TOKEN}`;
    run(cmd);
    // console.log(`cmd:${cmd}`);
}
help(deploy, 'Run backend deploy scripts');

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

// TODO pass full file name, avoid building filename string
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
    test,
    build,
    start,
    deploy,
}