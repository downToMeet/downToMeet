# Client

## Getting Started

System Requirements:

- [Node.js 14+](https://nodejs.org/en/)
- [Yarn 1.x](https://classic.yarnpkg.com/en/)

## Running the Project

Type the following lines in your command line from the project's root directory:

```sh
$ cd client
$ yarn
$ yarn start
```

## Generating Test Coverage

In order to generate test coverage in html format, run the following command
from inside the `client` directory:

```
$ yarn test --watchAll --coveragePathIgnorePatterns=/src/lib --coveragePathIgnorePatterns=node_modules --coverage --coverageReporters=html
```

This will generate the test coverage, excluding `node_modules` and the utility
functions in `/src/lib`. Open `client/coverage/index.html` in your web browser
to view the test coverage. To view the coverage in the command line instead of as
HTML, remove `--coverageReporters=html` from the command.

## Notes

- Be sure to run `yarn` every time there are changes to `package.json`. Usually
  you’ll want to run `yarn` in the following scenarios:
  - after pulling from main
  - after merging main into your branch
  - after switching branches (that may have different dependencies)
- This project uses ESLint to ensure code style compliance. ESLint is
  automatically run when you try to make a Git commit, though this can be
  overridden in exigent circumstances with `--no-verify`. To run ESLint
  manually, do `yarn lint`.
