# Jap Vocab App

Jap Vocab Testing App to test english words to be written in Japanese.

## Quickstart

1. Set up golang env

2. Set up googlesheets permissions (with `credentials.json` in `/conf`).
First time will require authentication and `token.json` to be created and put in `/conf`.

3. Install beego

4. Run

```bash
bee run
```

OR

```bash
go run .
```

## Deploy

This app is deployed with GAE

1. Install GAE SDK and setup

2. Deploy
```bash
cd <app-dir>
gcloud app deploy
```

## TODO

[] Look into Logout XSRF 