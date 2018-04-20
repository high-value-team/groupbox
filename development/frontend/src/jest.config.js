module.exports = {
  "rootDir": "src",
  "coverageReporters": [
    "html",
    "json",
    "lcov",
    "text"
  ],
  "setupFiles": [
    "<rootDir>/../scripts/polyfills.js",
    "raf/polyfill",
    "<rootDir>/../jest.setup.js"
  ],
  "testPathIgnorePatterns": [
    "/node_modules/"
  ]
};
