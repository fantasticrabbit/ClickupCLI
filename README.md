## Setup to run from code

1. Install [golang](https://golang.org/)
1. Clone this repo
1. Create new application in Clickup > Settings > Integrations > Clickup API. Define `http://localhost:4321` as the callback URL.
1. Copy the app client id and secret from the Clickup UI and set as env variables. 

```bash
export CLICKUP_CLIENT_ID=xxx
export CLICKUP_CLIENT_SECRET=xxx
```
1. Set your re-direct URL port in the vars section of `main.go`.

## Usage

1. In the Clickup app, copy the TaskID for the task you'd like to retrieve (e.g. `#123456`).
1. Run the app by giving the task id (remove the # at the front of the ID) as command line argument:

    ```bash
    go run . 123456
    ```
    
1. The app will open browser for Clickup authentication. Log in and give app permission to retrieve data from the corresponding workspace.
1. Task details will be downloaded to the current directory in JSON format (for now to StdOut).


## Build .exe for Windows

1. push tag to GitHub to create new release .exe 
   ```
   git tag vx.x.x -m "new tag message" && git push origin vx.x.x
   ```
