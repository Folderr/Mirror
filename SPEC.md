# The Folderr Mirror Specification

This specification (known as `spec` from now on) will declare how every version of Mirror works, expected returns, usages, and any other relevant information to expected setup

## What version of Mirrors does this spec cover

Currently, `v0` (development)/`v1` (will be released alongside Folderr V2)

## Compatibility matrix

| Mirror Spec Version | Folderr Version | FoldCLI   |
| ------------------- | --------------- | --------- |
| `v0`/`v1`           | `v2.x`          | `v0`/`v1` |

## Endpoints

### Images, Videos, Files, Links

#### GET /image/:id, /i/:id

Allows you to view an image someone uploaded
When in a browser this will show the image

**Response `200: OK`**

```
Content-Type: image/*
```

#### GET /video/:id, /v/:id

Allows you to view a video someone uploaded
When in a browser this will show a video player

**Response `200: OK`**

```
Content-Type: video/*
```

#### GET /file/:id, /f/:id

Sends a file to you, downloads.

**Response `200: OK`**

```
Content-Type: *
```

#### GET /link/:id, /l/:id

redirects you to the url uploaded under ID "id"

**Response `308: Permanent Redirect`**

### Health Check

This just proves to Folderr that there is a Mirror instance online here
Folderr times out after ~5-10 seconds of checking

**Response `200: OK`**

```json
{
	"message": "Mirror Operational",
	"code": 200
}
```

## Verifying you own a domain (user, for your own Mirror)

When used in `user` mode, Mirror will require you to uploaded a proof of your domain as DNS TXT entry

This allows Folderr to verify you own the domain.

Example Entry

```sh
Name: _mirror-user-verification
Content: <some randomized text>
TTL: 300
Type: TXT
Weight: 0
```

Note: Mirror will *not* generate this, Folderr will.

## Verifying you own a domain (service, for an instance's Mirror)

When used in `service` mode, Mirror requires you to upload a proof of your domain as DNS TXT entry

```sh
Name: _mirror-service-verification
Content: <some randomized text>
TTL: 300
Type: TXT
Weight: 0
```
