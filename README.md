A small go app which uses the google maps API to get a list of routes from an origin address to a destination address, showing the tim to reach the destination in usual traffic conditions and the expected time based on current traffic.

## Configuration

route.go requires a google maps API key to function. This key is used via the config.GTWConfiguration struct.

main.go requires a configuration file to load the google maps API key from. the file must be a json formatted file with the following structure:

```
{
    "apiKey": "<API-KEY-GOES-HERE>"
}
```

To generate an API key [follow this link](https://developers.google.com/maps/documentation/javascript/get-api-key).

## Flags

If this is being run via main.go then the following flags must be specified

| Flag        | Description                                                         |
| ----------- | ------------------------------------------------------------------- |
| config      | the path to the config.json file containing the google maps API key |
| origin      | The address for the starting point of the journey                   |
| destination | The address for the destination of the journey                      |
