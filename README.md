![Imgur](http://i.imgur.com/ixH5L0K.png)

- [About](https://github.com/purrgil/purrgil#About)
- [Motivation](https://github.com/purrgil/purrgil#Motivation)
- [Install](https://github.com/purrgil/purrgil#Install)
- [Commands](https://github.com/purrgil/purrgil#Commands)
  - Init
  - Install
  - Add
  - Remove
  - Packages
  - ?Deploy
  - ?Up
- [Contribution](https://github.com/purrgil/purrgil#Contribution)

## About
Purrgil is a simple wrapper for **Docker stuff** that helps manage a **multi-container environment**. After reading about Docker Orchestrators and projects that use Docker to manage their
enviroments I noticed that some people had a single repository with a `docker-compose.yml` file and an `init.sh` file to launch the app and using **orchestrators** only for deploying.

Purrgil's goal is to help you build and manage your environments, providing you flexibility and power in your development environment.

## Motivation
By the end of 2016 we (at work) decided to add Docker to our development stack and while doing it we ran into some issues when managing the environments.
We discussed about creating some scripts to make our work easy while we develop and deploy our services. So then Purrgil was born to help people get their environments up and running faster!

## Installation
```
curl https://raw.githubusercontent.com/purrgil/purrgil/master/install.sh | sh
```

## Commands

### init <project_name>
`init` will create a folder with the project's name and put a `docker-compose.yml`, `purrgil.yml` and `.gitignore` inside it.

```
purrgil init <project_name>
```

### install
`install` will read `purrgil.yml` file to map your `packages` (dependencies) and download them inside your project's folder.

```
purrgil install
```

### add
`add` will do exactly what you expect: add a package. This can be a github's repo or a dockerhub's image. Example: `guidiego/purrgil` (github) or `node:6` (dockerhub).
By the default purrgil will expect a github's repo, if you want to use a dockerhub's image you've to specify the flag `--dockerhub`.
```
purrgil add <package> [flag]
```

### remove
`remove` will do exactly what you expected it would do. It will remove the package's signature from `purrgil.yml`, `docker-compose.yml`, `.gitignore` and the
project folder from the root folder.

```
purrgil remove <package>
```

### packages
`packages` will list all the packages
```
purrgil packages
```
### deploy
> Not implemented yet.

### up
`up` will safely build your containers and start your application

### down
`down` will safely stop your application
```
purrgil down
```

## Contributing
This repository is here to keep the ideia of the project alive and to let other people contribute with the project.
The Contribution Guide and Code of Conduct are both in my TODO list..
