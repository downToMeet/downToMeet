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
