<p align="center">
  <img alt="3D pixelated Kibisis icon with an exclamation mark" src="assets/icon@128x128.png" style="padding-top: 15px" height="128" />
</p>

<h1 align="center">
   Kibisis API
</h1>

<p align="center">
  The Kibisis API is a set of DigitalOcean functions that act as a RESTful API for the Kibisis ecosystem.
</p>

<p align="center">
  <a href="https://github.com/agoralabs-sh/kibisis-api/releases/latest">
    <img alt="GitHub release" src="https://img.shields.io/github/v/release/agoralabs-sh/kibisis-api?&logo=github">
  </a>
  <a href="https://github.com/agoralabs-sh/kibisis-api/releases/latest">
    <img alt="GitHub release date - published at" src="https://img.shields.io/github/release-date/agoralabs-sh/kibisis-api?logo=github">
  </a>
</p>

<p align="center">
  <a href="https://github.com/agoralabs-sh/kibisis-api/blob/main/COPYING">
    <img alt="GitHub license" src="https://img.shields.io/github/license/agoralabs-sh/kibisis-api">
  </a>
</p>

#### Table of contents

* [1. Overview](#-1-overview)
  - [1.1. Project structure](#11-project-structure)
* [2. Development](#-2-development)
  - [2.1. Requirements](#21-requirements)
  - [2.2. Setup `doctl`](#22-setup-doctl)
  - [2.3. Create a personal namespace (optional)](#23-create-a-personal-namespace-optional)
* [3. Appendix](#-3-appendix)
  - [3.1. Useful commands](#31-useful-commands)
* [4. How To Contribute](#-4-how-to-contribute)
* [5. License](#-5-license)

## 🗂️ 1. Overview

### 1.1. Project structure

The project structure is based onthe [DigitalOcean functions][https://docs.digitalocean.com/products/functions/how-to/structure-projects/] project structure.
However, the core directories `lib` and `packages` have specific functionality:

* `packages` - This is where each function resides. Each package/function relates to the API path. For example, `achievements/quest` will correspond to an API path `https://<endpoint>/achievements/quests`
* `lib` - This contains is independent modules that are referenced from the packages.

> ⚠️ **NOTE:** Each path must parse the request header and handle the method, i.e. GET, POST, DELETE e.t.c.

<sup>[Back to top ^][table-of-contents]</sup>

## 🛠️ 2. Development

### 2.1. Requirements

* [`doctl`][doctl]
* [Golang v1.20+][golang]
* [Make][make]

<sup>[Back to top ^][table-of-contents]</sup>

### 2.2. Setup `doctl`

The DigitalOcean CLI client, `doctl`, is used to deploy the function to a remote environment tha can be used to develop.

Follow the instructions outlined in the [documentation][doctl].

> ⚠️ **NOTE:** Once you have setup `doctl` make sure you also install the serverless subcommand using `doctl serverless install`

<sup>[Back to top ^][table-of-contents]</sup>

### 2.3. Create a personal namespace (optional)

> ⚠️ **NOTE:** If you have a personal namespace already setup, you can skip this step.

1. Setup a personal namespace using your name suffixed by the word "`namespace`" ensuring to use kebab-case:

```shell script
doctl serverless namespaces create --label="kieran-namespace" --region="ams3"
```

> ⚠️ **NOTE:** The above example shows setting up a namespace in the Amsterdam region, but it is better if you use a region that is closest to you. Run the command: `doctl serverless namespaces list-regions` to get a list of regions and replace the `--region` value with the desired region form the list.

<sup>[Back to top ^][table-of-contents]</sup>

## 📑 3. Appendix

### 3.1. Useful commands

| Command       | Description                                                                                                        |
|---------------|--------------------------------------------------------------------------------------------------------------------|
| `make`        | Creates the `.env` file and deploys the functions to the remote namespace. Intended for development purposes only. |
| `make clean`  | Removes build files and configurations.                                                                            |
| `make deploy` | Deploys the functions to the remote namespace. Intended for development purposes only.                             |
| `make list`   | Lists the deployed functions in the configured namespace.                                                          |
| `make logs`   | Outputs the activation logs for the deployed functions.                                                            |
| `make setup`  | Creates the `.env` file from an `.env.example` file.                                                               |
| `make test`   | Runs the tests for each function.                                                                                  |
| `make watch`  | Watches for code changes and redeploys functions to namespace.                                                     |

<sup>[Back to top ^][table-of-contents]</sup>

## 👏 4. How To Contribute

Please read the [**Contributing Guide**][contribute] to learn about the development process.

<sup>[Back to top ^][table-of-contents]</sup>

## 📄 5. License

Please refer to the [COPYING][copying] file.

<sup>[Back to top ^][table-of-contents]</sup>

<!-- Links -->
[contribute]: ./CONTRIBUTING.md
[copying]: ./COPYING
[doctl]: https://docs.digitalocean.com/reference/doctl/how-to/install/
[golang]: https://go.dev/doc/install
[make]: https://www.gnu.org/software/make/
[table-of-contents]: #table-of-contents
