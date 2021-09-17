# Attribution
Authentication code based on Laura Vuorenoja's [kaneli](https://github.com/lauravuo/kaneli) application, and corresponding [tutorial on dev.to](https://dev.to/lauravuo/how-to-oauth-from-the-command-line-47j0)

Also, thanks to Isaac Johnson for crucial bit of info on Clickup's API in his [Clickup-Github tutorial](https://www.freshbrewed.science/clickup-and-gitabl-part-2/index.html)

## Setup

1. Install [golang](https://golang.org/)
1. Create new application in Clickup > Settings > Integrations > Clickup API. Define `http://localhost:4321` as the callback URL.
1. Copy the app client id and secret from the Clickup UI. 
1. Copy the team id from a task URL in the ClickUp web app. It's usually a 7-digit numeral, found here: ```https://app.clickup.com/<TEAMID>/rest/of/url```
1. Define following env variables:

```bash
export CLICKUP_CLIENT_ID=xxx
export CLICKUP_CLIENT_SECRET=xxx
export CLICKUP_TEAM_ID=1234567
```

## Usage

1. In the Clickup app: create a new task (or use an existing task.) Copy the TaskID for the task (e.g. `#123456`).
1. Run the app by giving the task id (remove the # at the front of the ID) as command line argument:

    ```bash
    go run . 123456
    ```
    
1. The app will open browser for Clickup authentication. Log in and give app permission to retrieve data from the corresponding workspace.
1. Task details will be downloaded to the current directory in JSON format.


## Build .exe for Windows

1. Download source, with golang installed:

```GOOS=windows go build ./...```
 