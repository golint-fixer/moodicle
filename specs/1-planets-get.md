# Planets API

The Planets API is a simple REST api for testing against.

## `GET /api/dot/planets`

* API_KEY: "1234"

Gets all planets in the system.

===

### Example response

* Status: `200`
* Content-Type: `application/json`

```
[{"id":"36ea0a63-dfff-11e5-8236-5cc5d46c2078","name":"Matt","hoursInDay":1},{"id":"4543a1fb-dfff-11e5-8236-5cc5d46c2078","name":"Matt","hoursInDay":1,"moons":[{"name":"trevor"}]},{"id":"abb2a886-dfff-11e5-9e58-5cc5d46c2078","name":"Matt","hoursInDay":12,"moons":[{"name":"trevor"}]}]

```

## `GET /api/dot/planets/36ea0a63-dfff-11e5-8236-5cc5d46c2078`

* API_KEY: "1234"

Gets a planet in the system.

===

### Example response

* Status: `200`
* Content-Type: `application/json`

Returns the planets:

```
{"id":"36ea0a63-dfff-11e5-8236-5cc5d46c2078","name":"Matt","hoursInDay":1}

```

## `GET /api/dot/planets/unknown`

* API_KEY: "1234"

Get an unknown planet returns 404.

===

### Example response

* Status: `404`
* Content-Type: `application/json`
