![Imgur](https://purrgil.github.io/static/images/full_logo.svg)

Purrgil is a simple wrapper for **Docker stuff** that helps manage a **multi-container environment**. After reading about Docker Orchestrators and projects that use Docker to manage their
enviroments I noticed that some people had a single repository with a `docker-compose.yml` file and an `init.sh` file to launch the app and using **orchestrators** only for deploying.

Purrgil's goal is to help you build and manage your environments, providing you flexibility and power in your development environment.

## Install
For developers that want to contribute, you can install with `go` manage:
```
go get github.com/purrgil/purrgil
```

For soft users on macOS, Linux, or OpenBSD run the following:

```
curl https://raw.githubusercontent.com/apex/apex/master/install.sh | sh
```

Note that you may need to run the sudo version below, or alternatively chown /usr/local:

```
curl https://raw.githubusercontent.com/apex/apex/master/install.sh | sudo sh
```

## Documentation
We have a little documentation here: https://purrgil.github.io/docs/, if you want to contribute in this we are using gitbook, and the source is [here](). Any sugestions you can open a issue on the site repo!

## Code of Conduct

We has adopted a Code of Conduct that we expect project participants to adhere to. Please read [the full text](https://github.com/purrgil/purrgil/tree/master/.github/CODE_OF_CONDUCT.md) so that you can understand what actions will and will not be tolerated.

## Contribuition Guide
If you want to contribute, first read our guide! We based on [react]() one: [Purrgil Guide](https://github.com/purrgil/purrgil/tree/master/.github/CONTRIBUTING.md)

## License

By contributing to Purrgil, you agree that your contributions will be licensed under its MIT license.


