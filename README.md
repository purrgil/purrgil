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
Purrgil is a simple wrapper for **Docker stuff** that help you to manage a **mult-container enviroment**. After read a little about Docker Orchestrators and projects that use Docker to manage your enviroments I noticed that some people have a single repository with a `docker-compose.yml` and an `init.sh` to make all project work aright and use **orchestrators** only to deploy your application.

The ideia of Purrgil is help you construct this enviroment, help you to provide across your stages and the principal: Give you more flexibility and power to use docker in development enviroment!

## Motivation
In final of 2016 we start migrate some stuff of our enviroment to Docker in enterprise, talk n read about that we found some problems and nice ways to resolve that. At some talks we discuss about create some scripts to make our work easily while we develop and deploy our services. So I create Purrgil to make those scripts and help people to create your enviroments faster!

## Install
```
curl https://raw.githubusercontent.com/purrgil/purrgil/master/install.sh | sh
```

## Commands
Purrgil have some API, you can read about that using `purrgil --help`

### Init
This command create a folder with project name and put a `docker-compose.yml`, `purrgil.yml` and `.gitignore` inside them.
```
purrgil init <projectName>
```

### Install
Install read `purrgil.yml` to understant your `packages` (dependencies) and download that inside your folder.
```
purrgil install
```
### Add
When you **add** a `package` in purrgil you need pass the identity of your package, in this case a **github** sign or a **dockerhub** image. Some example `guidiego/purrgil` or `node:6`, by default packages are setting with **github** as provider, but you can use `--dockerhub` to demonstrate that this service are a image. You can know more about using `purrgil add --help`

```
purrgil add <pkgIndentity> [<flag>]
```
### Remove
Remove command in simple way do the reverse of ADD, remove the signs from `purrgil.yml`, `docker-compose.yml`, `.gitignore` and the project folder from project root.
```
purrgil remove <packageName>
```

### Packages
This command list the **name <- identity** from packages, and can be filter by flags, more about `purrgil packages --help`
```
purrgil packages
```
### Deploy
> Todo...

### Up / Down
**Up** command activer docker-compose up passing some aditional parameter to garanted a secure instance of your containers, after that you can drop a project up with purrgil with **down** command :)

```
# Start the applicaiton
purrgil up

# Stop the application
purrgil down
```

## Contribution
The purpose of this project was explained in first 2 sections, this repository are here to maintain the ideia and improve that! Contribution Guide and Code of Conduct are in todo list.
