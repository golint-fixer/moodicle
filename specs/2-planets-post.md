# Planets API

The Planets API is a simple REST api for testing against.

## `POST /api/dot/planets`

* API_KEY: "1234"

Post a new planet in the system.

### Example request

```
{
    "name":"Matt",
    "hoursInDay":1}
    "moons":[
        {"name":"one"},
        {"name":"two"}
        ]
}
```


===

### Example response

* Status: `201`
* Content-Type: `application/json`
* Data.id: /1/ // unique {id}

## `GET /api/dot/planets/{id}`

* API_KEY: "1234"

Gets a planet in the system.

===

### Example response

* Status: `200`
* Content-Type: `application/json`