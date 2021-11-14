## Configuration

1. In the [ClickUp web app](https://app.clickup.com), create a new application in the Clickup Menu > Settings > Integrations > Clickup API. 
1. Define `http://localhost:<port#>` as the callback URL. The CLI defaults to 4321, but you may use any valid port and configure the CLI.
1. Copy the app client id and secret, and redirect port number from the Clickup UI and set as environment variables in your local environment: 

```
export CLICKUP_CLIENT_ID=xxx
export CLICKUP_CLIENT_SECRET=xxx
export CLICKUP_PORT=9999
```

### SET command

1. You can optionally set the local host port with the set command:
    ```
    clickup set --port=1234
    ```
1. If your space is using custom Task ID's, you will need to set the Team ID. Grab this code from your Clickup web app URL, for example, "https://app.clickup.com/1234567/...". 
    ```
    clickup set --team=1234567
    ```
1. Both parameters can be set in the same command, or provided as environment variables with the "CLICKUP_" prefix.

1. You can additionally set the token manually (--token), to use a personal token or for environments without a browser

## Usage  

### Authentication
1. The first time the app connects to the Clickup API, it will open a browser and prompt the user to give the client app permission to exchange data with a workspace. The client may only be used with one workspace at a time.

1. ClickupCLI will save your authentication token in $HOME/.clickup/config.yaml.  To delete the token issue the logout subcommand:
    ```
    clickup logout
    ```
1. This will clear the token, requiring the user to re-authenticate the next time the client app is executed.

### Get Task  

1. From the Clickup app, copy the TaskID for the task you'd like to retrieve (e.g. `#123456`).  
1. Provide the task ID as command line argument (with or without the "#" prefix):

    ```
    clickup get task #123456
    ```
    
1. Task details will be output in JSON to StdOut. 
1. If Clickup space is using custom task IDs, be sure to set the Team ID with either the set command or env variable, and pass the -c flag:
    ```
    clickup get task CUSTOM-1234 -c        <short>
    clickup get task CUSTOM-1234 --custom  <long/explicit>
    ```
1. To include sub-tasks in the JSON output, add the -s/--subtasks flag:
    ```
    clickup get task 123456 -s
    clickup get task 123456 --subtasks
    ```

### Get List
1. From the Clickup app, copy the List ID for the list you'd like to retrieve (e.g. `1234`).  
1. Provide the list ID as command line argument:
    ```
    clickup get list 1234
    ```
1. List details will be output in JSON to StdOut.