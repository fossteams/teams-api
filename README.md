# teams-api

An unofficial wrapper for the unofficial Microsoft Teams API

## Description

This library (still work in progress) was created with the goal
of helping the creation of alternative (and open source) 
Microsoft Teams clients. Currently the library itself can only
handle a couple of endpoints and is by far not ready yet.  
My goal is to complete this library as soon as possible and start
the creation of an alternative client based on it.

## Usage

For now, you can only run the tests or use the library on your own.
For this you'll need a Microsoft Teams Token.

### Obtaining a token

For the moment, the token must be retrieved by using the Developer Console
and by looking at the requests headers when teams.microsoft.com is open.
Please, note that there are different tokens for different services:
decode your JWT with [step](https://github.com/smallstep/cli) 
(`step crypto jwt inspect --insecure`) and check the `aud` field to
verify you're using the right token for the right service.
CSA tokens will have the following `aud`:
```json
{
    "header": {},
    "payload": {
        "aud": "https://chatsvcagg.teams.microsoft.com"
    }
}
```

An useful starting point on how to programmatically get the token later
on might be [this note](https://github.com/seancabahug/UHDiscordBot/blob/636bd329c085203648b65944a71a68656371385e/notes.txt).

## Setting the environment variable
```bash
read -s -r MS_TEAMS_TOKEN
# paste your token and press [Enter]
export MS_TEAMS_TOKEN
```

## Running a test
```bash
# CSA (Chat Svc Agg)
go test -v github.com/fossteams/teams-api/api/csa
```