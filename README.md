## Setup in ClickUp

1. Create new application in Clickup > Settings > Integrations > Clickup API. Define `http://localhost:4321` as the callback URL.
1. Copy the app client id and secret from the Clickup UI and set as env variables. 

```bash
export CLICKUP_CLIENT_ID=xxx
export CLICKUP_CLIENT_SECRET=xxx
```

## Windows Server Usage

1. From the Clickup app, copy the TaskID for the task you'd like to retrieve (e.g. `#123456`).
1. Run the app by giving the task id (remove the # at the front of the ID) as command line argument:

    ```bash
    goFetchClickUp.exe 123456
    ```
    
1. The app will open browser for Clickup authentication. Log in and give app permission to retrieve data from the corresponding workspace.
1. Task details will be output in JSON to StdOut.


## Build .exe for Windows

1. push tag to GitHub to create new release .exe 
   ```
   git tag vx.x.x -m "new tag message" && git push origin vx.x.x
   ```
