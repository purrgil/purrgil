# Contributing Guide

Purrgil born with a simple ideia and many things to do! We will maintain/manage this project the better that we can, this simple docs serve to make your contribuition more easy :D

### Code of Conduct

We has adopted a Code of Conduct that we expect project participants to adhere to. Please read [the full text](https://github.com/purrgil/purrgil/tree/master/.github/CODE_OF_CONDUCT.md) so that you can understand what actions will and will not be tolerated.

### Open Development

All work on Purrgil happens directly on [GitHub](https://github.com/purrgil/purrgil). Both core team members and external contributors send pull requests which go through the same review process.

### Branch Organization

We will do our best to keep the [`master` branch](https://github.com/purrgil/purrgil/tree/master) in good shape, with tests passing at all times. But in order to move fast, we will make API changes that your application might not be compatible with. We recommend that you use [the latest stable version of Purrgil](https://github.com/purrgil/purrgil/releases).

If you send a pull request, please do it against the `master` branch. We maintain stable branches for major versions separately but we don't accept pull requests to them directly. Instead, we cherry-pick non-breaking changes from master to the latest stable major version.

### Semantic Versioning

Purrgil follows [semantic versioning](http://semver.org/). We release patch versions for bugfixes, minor versions for new features, and major versions for any breaking changes. When we make breaking changes, we also introduce deprecation warnings in a minor version so that our users learn about the upcoming changes and migrate their code in advance.

### Bugs

#### Where to Find Known Issues

We are using [GitHub Issues](https://github.com/purrgil/purrgil/issues) for our public bugs. We keep a close eye on this and try to make it clear when we have an internal fix in progress. Before filing a new task, try to make sure your problem doesn't already exist.

#### Discussions
If you want to init a discussion about something releated with Purrgil, create an issue with `[FORUM]` in start of the title! We have some forums opened by core team :)

### Sending a Pull Request

The core team is monitoring for pull requests. We will review your pull request and either merge it, request changes to it, or close it with an explanation.

### Contribution Prerequisites

* You have `docker`, `docker-compose` and `docker-machine` installed.
* You have `editorconfig` installed in your main editor.
* You have go installed `1.5+`.
* You are familiar with `git`.

### Style Guide

We use basic `gofmt` to format our files, but we have a `.editorconfig` to garanted that we maintain code on lint!

### License

By contributing to Purrgil, you agree that your contributions will be licensed under its MIT license.
