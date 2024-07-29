# water-tank-api
This API is divided into `Internal` (read-only) and `External` (write-only) endpoints.

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

## Endpoints

### /v1/water-tank/

#### [POST] [Internal]

##### Response codes

- `Created (201)`
- `Bad Request (400)`
- `Unprocessable Entity (422)`

##### Request body example
``` json
{
    "tank_name": "TANK_7",
    "group": "GROUP_2",
    "maximum_capacity": 45
}
```

##### Success response example
``` json
{
    "current_water_level": "0.00L",
    "datetime": "2023-09-25T19:34:39.775746328Z",
    "group": "GROUP_2",
    "maximum_capacity": "45.00L",
    "name": "TANK_7",
    "tank_state": "EMPTY"
}
```

### /v1/water-tank/tank/{tank_name}

#### [PATCH] [Internal]

##### Response codes

- `No Content (204)`
- `Bad Request (400)`
- `Not Found (404)`
- `Unprocessable Entity (422)`

##### Request header

- `group`: Tank's group name

##### Request body example
``` json
{
    "water_level": 10
}
```

#### [GET] [External]

##### Response codes

- `Ok (200)`
- `Bad Request (400)`
- `Not Found (404)`

##### Request header

- `group`: Tank's group name

##### Success response example
``` json
{
    "current_water_level": "90.00L",
    "datetime": "2023-09-25T19:34:39.775746328Z",
    "group": "GROUP_3",
    "maximum_capacity": "120.00L",
    "name": "TANK_6",
    "tank_state": "FILLING"
}
```

### /v1/water-tank/group/{group}

#### [GET] [External]

##### Response codes

- `Ok (200)`
- `Bad Request (400)`
- `Not Found (404)`

##### Success response example
``` json
{
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
```
