# Mirror-Server

The Mirror Server &amp; Spec for Folderr's "Mirror" service

Essentially, this is the service allowing users to utilize their own domains for Folderr
This service can be ran as either a user service, or can be ran as a Folderr Instance service.

We will first start creating a User service, and later we will add the ability for this to be a Instance Service

You can [read the spec](./SPEC.md) if you'd like, to see how Mirror is supposed to work.

Defintions:

- User: A Folderr instance User. **Not to be confused with an OS user**
- Service Manager: A person who manages the server that Folderr runs on, and likely its database and reverse proxy as well.

## User Service

What you as a user should expect: You can only serve your own domains

### Requirements

- Own a domain
- Have your own server
- Set up the reverse proxy yourself.

### Running / Usage

```sh
mirror-server
```

## Instance Service

What you can expect: Ability to serve any user's domain.

### NOTICE - Planned Feature - Not Ready

This is planned to be finished after the User Service. These notes are to provide a solid specification for developers, and you, the service manager

### Requirements

- Ability to use a wildcard domain
- Using Caddy as your Mirror Server's proxy
- Your server has the ability to connect to your Folderr Instance's database

### Running / Usage

```sh
mirror-server --service
```

The `service` flag tells the Mirror you're running it as a service for multiple users

## Installation

Unsupported during `v0`. Please build from source instead

## Building from Source

### Requirements

- Some way to download the source code. We'll use git, and git is recommended
- Go version 1.22 or later.

### Building

```sh
git clone https://github.com/Folderr/Mirror-Server.git
cd Mirror-Server
go build .
```

### Placing it in your path (system)

```sh
# This assumes you're in the directory of the source code and have built the source code
sudo cp Mirror-Server /opt/mirror-server
```

### Placing it in your path (user)

```sh
cp Mirror-Server /home/$USER/.local/bin/mirror-server
```
