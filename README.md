## Setup in ClickUp

1. Create new application in Clickup > Settings > Integrations > Clickup API. Define `http://localhost:4321` as the callback URL.
1. Copy the app client id and secret from the Clickup UI and set as env variables. 

```
export CLICKUP_CLIENT_ID=xxx
export CLICKUP_CLIENT_SECRET=xxx
```

## Usage - Get Tasks

1. From the Clickup app, copy the TaskID for the task you'd like to retrieve (e.g. `#123456`).
1. Run the app by passing the -t flag and providing the task id (remove the # at the front of the ID) as command line argument:

    ```
    clickup get -t 123456 <short>
    clickup get --task 123456 <long/explicit>
    ```
    
1. The app will open browser for Clickup authentication. Log in and give app permission to retrieve data from the corresponding workspace.
1. Task details will be output in JSON to StdOut. You can use the -f flag to output to a file clickup_<taskid>.json:
    ```
    clickup get -t 123456 -f <short>
    clickup get -t 123456 --file <long/explicit>
    ```
1. returns clickup_123456.json  
