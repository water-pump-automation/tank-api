# water-tank-api
API for water tank's data management.

The API is divided into `Internal` and `External` endpoints.

The `External` ones are read-only access to tank's data, or a group of tanks.
The returned attributes are:

- `Name`
- `Group`
- `MaximumCapacity`
- `TankState`
  - `Empty`
  - `Filling`
  - `Full`
- `CurrentWaterLevel`
- `LastFullTime`

The `Internal` exposes all the `External` ones, plus the hability to register a new tank
and update any water level.

In order to provide a new water state, it must be passed an `access_token`,
working as a password to prevent other API users to do malicious or multiple modifications.

## Deploy options

### Docker

Docker compose can be used to build and run both internal and external API, as the following:
``` bash
docker-compose up
```

But, each API can also be deploy running build separately:

- `Internal`

``` bash
docker-compose build water-tank-api-internal-v1
docker run lo-han/water-tank-api-internal-v1 -p 8081:8080
```

- `External`

``` bash
docker-compose build water-tank-api-external-v1
docker run lo-han/water-tank-api-external-v1 -p 8082:8080
```

## Infrastructure options

### Presenter

- [HTTP](infra/web/routes/routes_internal.go)

### Logs

- [Stdout](infra/stdout/stdout.go)

### Database

- [Mock](infra/database/mock/database_mock.go)

## Endpoints

The endpoints' Postman collection can be downloaded at [Water-tank-api [v1].postman_collection.json](docs/postman_requests/Water-tank-api%20[v1].postman_collection.json).

### /v1/water-tank/:name [GET]

#### Response codes

- `Ok (200)`
- `Not Found (404)`

#### Request header

- `group`

#### Response example
``` json
{
    "code": "WATERTANK_200",
    "content": {
        "current_water_level": "90.00L",
        "datetime": "2023-09-25T19:34:39.775746328Z",
        "group": "GROUP_3",
        "maximum_capacity": "120.00L",
        "name": "TANK_6",
        "tank_state": "FILLING"
    }
}
```

### /v1/water-tank/group/:group [GET]

#### Response codes

- `Ok (200)`
- `Bad Request (400)`
- `Not Found (404)`

#### Response example
``` json
{
    "code": "WATERTANK_200",
    "content": {
        "datetime": "2023-09-25T19:36:20.721605065Z",
        "tanks": [
            {
                "current_water_level": "0.00L",
                "group": "GROUP_1",
                "maximum_capacity": "100.00L",
                "name": "TANK_1",
                "tank_state": "EMPTY"
            },
            {
                "current_water_level": "50.00L",
                "group": "GROUP_1",
                "maximum_capacity": "80.00L",
                "name": "TANK_2",
                "tank_state": "FILLING"
            },
            {
                "current_water_level": "120.00L",
                "group": "GROUP_1",
                "maximum_capacity": "120.00L",
                "name": "TANK_3",
                "tank_state": "FULL"
            }
        ]
    }
}
```


### /v1/water-tank/ [POST]

#### Response codes

- `Ok (200)`
- `Bad Request (400)`
- `Unprocessable Entity (422)`

#### Request body example
``` json
{
    "name": "TANK_7",
    "group": "GROUP_2",
    "maximum_capacity": 45
}
```

#### Response example
``` json
{
    "access_token": "<ACCESS_TOKEN>"
}
```

### /v1/water-tank/:name [PATCH]

#### Response codes

- `No Content (204)`
- `Bad Request (400)`
- `Not Found (404)`
- `Unprocessable Entity (422)`

#### Request header

- `access_token`
- `group`

#### Request body example
``` json
{
    "water_level": 10
}
```
