# Mailbox App

## Overview

This repository contains the source code for the mailbox-app project, which is divided into backend and frontend components. The backend is implemented in Golang and includes an API and CLI functionalities. The frontend is developed using Vue 3 for building a Single Page Application (SPA).

### Folder Structure
```bash
mailbox-app/
├── README.md # Main project documentation file
├── backend/
│ ├── api/ # API files for Golang backend
│ ├── cmd/ # Command files for Golang backend
│ ├── internal/ # Common layer between API and CLI
│ └── .env.dist # Environment variables distribution file for backend
├── frontend/
│ ├── src/ # Source files for Vue 3 SPA frontend
│ └── .env.dist # Environment variables distribution file for frontend
```

## Setup

### Pre-requisites:

* Docker
* docker-compose

#### Docker

Install Docker for your platform.

* Mac: https://store.docker.com/editions/community/docker-ce-desktop-mac
* Windows: https://store.docker.com/editions/community/docker-ce-desktop-windows
* Linux: Please see your distributions package management system

#### docker-compose

Install docker-compose for your platform.

* Mac: Included with Docker
* Windows: Included with Docker
* Linux: Please see your distributions package management system

### Installation

First we need to clone the project.

```bash
cd ~/code # or whatever directory you want

git clone https://github.com/MiguelTzab/mailbox-app.git
cd mailbox-app
```

Next we need to copy the `docker-compose.yml.dist` file and update it for our system.

```bash
cp backend/.env.dist backend/.env
cp frontend/.env.dist frontend/.env

cp docker-compose.override.yml.dist docker-compose.override.yml
```

> **Note:** _Personalize the `docker-compose.override.yml` for your system. If you are using Mac, it is recommended to change `- .:/application` to `- .:/application:cached` for the best performance. However, the default file will work out of the box._

Now that we have the application configured, we need to install our dependencies. Before doing that though we need the docker images we use.

```bash
docker-compose pull
```

And finally it's time to start up our containers:

```bash
docker-compose up -d
```

To access the site, visit one of the following URL:

* http://localhost:8004

### Importing
For use the importing CLI (indexer), please [see](./backend/CLI.md).