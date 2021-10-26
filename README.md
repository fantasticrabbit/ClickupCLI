## Configuration

1. In the [ClickUp web app](https://app.clickup.com), create a new application in the Clickup Menu > Settings > Integrations > Clickup API. 
1. Define `http://localhost:<port#>` as the callback URL.
1. Copy the app client id and secret, and redirect port number from the Clickup UI and set as environment variables in your local environment: 

```
export CLICKUP_CLIENT_ID=xxx
export CLICKUP_CLIENT_SECRET=xxx
export CLICKUP_REDIRECT_PORT=9999  (this is optional, will default to port 4321, just make sure it matches in the Clickup custom app definition above.)
```

## Usage  

### Authentication
1. The first time the app is used, it will open a browser and ask the user to log in to ClickUp to give the client app permission to exchange data with the corresponding workspace. The client may only be used with one workspace at a time.

### Get Tasks  

1. From the Clickup app, copy the TaskID for the task you'd like to retrieve (e.g. `#123456`).  
1. Execute the app by passing the -t flag and providing the task ID (remove the # at the front of the ID) as command line argument:

    ```
    clickup get -t 123456 <short>
    clickup get --task 123456 <long/explicit>
    ```
    
1. Task details will be output in JSON to StdOut. 
1. You can use the -f flag to output to a file clickup_<taskid>.json:
    ```
    clickup get -t 123456 -f <short>
    clickup get -t 123456 --file <long/explicit>
    ```

### Logout
1. ClickupCLI will save your authentication token in $HOME/.clickup/config.yaml.  To delete the token issue the logout subcommand:
    ```
    clickup logout
    ```
1. This will clear the token, requiring the user to re-authenticate the next time the client app is executed.
