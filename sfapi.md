# Protocol Documentation
<a name="top"/>

## Table of Contents

- [sfapi.proto](#sfapi.proto)
    - [Film](#sfapi.v1.Film)
    - [GetFilmRequest](#sfapi.v1.GetFilmRequest)
    - [GetFilmResponse](#sfapi.v1.GetFilmResponse)
    - [GetPersonRequest](#sfapi.v1.GetPersonRequest)
    - [GetPersonResponse](#sfapi.v1.GetPersonResponse)
    - [GetPlanetRequest](#sfapi.v1.GetPlanetRequest)
    - [GetPlanetResponse](#sfapi.v1.GetPlanetResponse)
    - [GetSpeciesRequest](#sfapi.v1.GetSpeciesRequest)
    - [GetSpeciesResponse](#sfapi.v1.GetSpeciesResponse)
    - [GetStarshipRequest](#sfapi.v1.GetStarshipRequest)
    - [GetStarshipResponse](#sfapi.v1.GetStarshipResponse)
    - [GetVehicleRequest](#sfapi.v1.GetVehicleRequest)
    - [GetVehicleResponse](#sfapi.v1.GetVehicleResponse)
    - [InvalidKey](#sfapi.v1.InvalidKey)
    - [ListFilmsRequest](#sfapi.v1.ListFilmsRequest)
    - [ListFilmsResponse](#sfapi.v1.ListFilmsResponse)
    - [ListPeopleRequest](#sfapi.v1.ListPeopleRequest)
    - [ListPeopleResponse](#sfapi.v1.ListPeopleResponse)
    - [ListPlanetsRequest](#sfapi.v1.ListPlanetsRequest)
    - [ListPlanetsResponse](#sfapi.v1.ListPlanetsResponse)
    - [ListSpeciesRequest](#sfapi.v1.ListSpeciesRequest)
    - [ListSpeciesResponse](#sfapi.v1.ListSpeciesResponse)
    - [ListStarshipActionsRequest](#sfapi.v1.ListStarshipActionsRequest)
    - [ListStarshipsRequest](#sfapi.v1.ListStarshipsRequest)
    - [ListStarshipsResponse](#sfapi.v1.ListStarshipsResponse)
    - [ListVehiclesRequest](#sfapi.v1.ListVehiclesRequest)
    - [ListVehiclesResponse](#sfapi.v1.ListVehiclesResponse)
    - [Person](#sfapi.v1.Person)
    - [Planet](#sfapi.v1.Planet)
    - [Species](#sfapi.v1.Species)
    - [Starship](#sfapi.v1.Starship)
    - [StarshipAction](#sfapi.v1.StarshipAction)
    - [ValidateSpeciesRequest](#sfapi.v1.ValidateSpeciesRequest)
    - [ValidateSpeciesResponse](#sfapi.v1.ValidateSpeciesResponse)
    - [Vehicle](#sfapi.v1.Vehicle)

    - [StarshipAction.Action](#sfapi.v1.StarshipAction.Action)


    - [Starfriends](#sfapi.v1.Starfriends)


- [Scalar Value Types](#scalar-value-types)



<a name="sfapi.proto"/>
<p align="right"><a href="#top">Top</a></p>

## sfapi.proto



<a name="sfapi.v1.Film"/>

### Film



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |
| episode_id | [int32](#int32) |  |  |
| opening_crawl | [string](#string) |  |  |
| director | [string](#string) |  |  |
| producer | [string](#string) |  |  |
| release_date | [string](#string) |  |  |
| character_ids | [string](#string) | repeated | Person.id |
| planet_ids | [string](#string) | repeated | Planet.id |
| starship_ids | [string](#string) | repeated | Starship.id |
| vehicle_ids | [string](#string) | repeated | Vehicle.id |
| species_ids | [string](#string) | repeated | Species.id |
| id | [string](#string) |  |  |






<a name="sfapi.v1.GetFilmRequest"/>

### GetFilmRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="sfapi.v1.GetFilmResponse"/>

### GetFilmResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| film | [Film](#sfapi.v1.Film) |  |  |






<a name="sfapi.v1.GetPersonRequest"/>

### GetPersonRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="sfapi.v1.GetPersonResponse"/>

### GetPersonResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| person | [Person](#sfapi.v1.Person) |  |  |






<a name="sfapi.v1.GetPlanetRequest"/>

### GetPlanetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="sfapi.v1.GetPlanetResponse"/>

### GetPlanetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| planet | [Planet](#sfapi.v1.Planet) |  |  |






<a name="sfapi.v1.GetSpeciesRequest"/>

### GetSpeciesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="sfapi.v1.GetSpeciesResponse"/>

### GetSpeciesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| species | [Species](#sfapi.v1.Species) |  |  |






<a name="sfapi.v1.GetStarshipRequest"/>

### GetStarshipRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="sfapi.v1.GetStarshipResponse"/>

### GetStarshipResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| starship | [Starship](#sfapi.v1.Starship) |  |  |






<a name="sfapi.v1.GetVehicleRequest"/>

### GetVehicleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="sfapi.v1.GetVehicleResponse"/>

### GetVehicleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vehicle | [Vehicle](#sfapi.v1.Vehicle) |  |  |






<a name="sfapi.v1.InvalidKey"/>

### InvalidKey



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| message | [string](#string) |  |  |






<a name="sfapi.v1.ListFilmsRequest"/>

### ListFilmsRequest







<a name="sfapi.v1.ListFilmsResponse"/>

### ListFilmsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| films | [Film](#sfapi.v1.Film) | repeated |  |






<a name="sfapi.v1.ListPeopleRequest"/>

### ListPeopleRequest







<a name="sfapi.v1.ListPeopleResponse"/>

### ListPeopleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| people | [Person](#sfapi.v1.Person) | repeated |  |






<a name="sfapi.v1.ListPlanetsRequest"/>

### ListPlanetsRequest







<a name="sfapi.v1.ListPlanetsResponse"/>

### ListPlanetsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| planets | [Planet](#sfapi.v1.Planet) | repeated |  |






<a name="sfapi.v1.ListSpeciesRequest"/>

### ListSpeciesRequest







<a name="sfapi.v1.ListSpeciesResponse"/>

### ListSpeciesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| species | [Species](#sfapi.v1.Species) | repeated |  |






<a name="sfapi.v1.ListStarshipActionsRequest"/>

### ListStarshipActionsRequest







<a name="sfapi.v1.ListStarshipsRequest"/>

### ListStarshipsRequest







<a name="sfapi.v1.ListStarshipsResponse"/>

### ListStarshipsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| starships | [Starship](#sfapi.v1.Starship) | repeated |  |






<a name="sfapi.v1.ListVehiclesRequest"/>

### ListVehiclesRequest







<a name="sfapi.v1.ListVehiclesResponse"/>

### ListVehiclesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vehicles | [Vehicle](#sfapi.v1.Vehicle) | repeated |  |






<a name="sfapi.v1.Person"/>

### Person



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| height | [int32](#int32) |  |  |
| mass | [float](#float) |  |  |
| hair_color | [string](#string) |  |  |
| skin_color | [string](#string) |  |  |
| eye_color | [string](#string) |  |  |
| birth_year | [string](#string) |  |  |
| gender | [string](#string) |  |  |
| homeworld | [string](#string) |  |  |
| film_ids | [string](#string) | repeated | Film.id |
| species_ids | [string](#string) | repeated | Species.id |
| vehicle_ids | [string](#string) | repeated | Vehicle.id |
| starship_ids | [string](#string) | repeated | Starship.id |
| id | [string](#string) |  |  |






<a name="sfapi.v1.Planet"/>

### Planet



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| rotation_period | [int32](#int32) |  |  |
| orbital_period | [int32](#int32) |  |  |
| diameter | [int32](#int32) |  |  |
| climate | [string](#string) |  |  |
| gravity | [float](#float) |  |  |
| terrain | [string](#string) |  |  |
| surface_water | [float](#float) |  |  |
| population | [int64](#int64) |  |  |
| resident_ids | [string](#string) | repeated | Person.id |
| film_ids | [string](#string) | repeated | Film.id |
| id | [string](#string) |  |  |






<a name="sfapi.v1.Species"/>

### Species



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| classification | [string](#string) |  |  |
| designation | [string](#string) |  |  |
| average_height | [int32](#int32) |  |  |
| skin_colors | [string](#string) |  |  |
| eye_colors | [string](#string) |  |  |
| average_lifespan | [int32](#int32) |  |  |
| homeworld | [string](#string) |  |  |
| language | [string](#string) |  |  |
| people_ids | [string](#string) | repeated | Person.id |
| film_ids | [string](#string) | repeated | Film.id |
| id | [string](#string) |  |  |
| hair_colors | [string](#string) |  |  |






<a name="sfapi.v1.Starship"/>

### Starship



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| model | [string](#string) |  |  |
| manufacturer | [string](#string) |  |  |
| cost_in_credits | [int64](#int64) |  |  |
| length | [float](#float) |  |  |
| crew | [int32](#int32) |  |  |
| passengers | [int32](#int32) |  |  |
| cargo_capacity | [int64](#int64) |  |  |
| consumables | [string](#string) |  |  |
| hyperdrive_rating | [float](#float) |  |  |
| mglt | [int32](#int32) |  |  |
| starship_class | [string](#string) |  |  |
| pilot_ids | [string](#string) | repeated | Person.id |
| film_ids | [string](#string) | repeated | Film.id |
| id | [string](#string) |  |  |
| max_atmosphering_speed | [int64](#int64) |  |  |






<a name="sfapi.v1.StarshipAction"/>

### StarshipAction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| starship | [Starship](#sfapi.v1.Starship) |  |  |
| action | [StarshipAction.Action](#sfapi.v1.StarshipAction.Action) |  |  |






<a name="sfapi.v1.ValidateSpeciesRequest"/>

### ValidateSpeciesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="sfapi.v1.ValidateSpeciesResponse"/>

### ValidateSpeciesResponse







<a name="sfapi.v1.Vehicle"/>

### Vehicle



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| model | [string](#string) |  |  |
| manufacturer | [string](#string) |  |  |
| cost_in_credits | [int32](#int32) |  |  |
| length | [float](#float) |  |  |
| max_atmosphering_speed | [int32](#int32) |  |  |
| crew | [int32](#int32) |  |  |
| passengers | [int32](#int32) |  |  |
| cargo_capacity | [int32](#int32) |  |  |
| consumables | [string](#string) |  |  |
| vehicle_class | [string](#string) |  |  |
| pilot_ids | [string](#string) | repeated | Person.id |
| film_ids | [string](#string) | repeated | Film.id |
| id | [string](#string) |  |  |








<a name="sfapi.v1.StarshipAction.Action"/>

### StarshipAction.Action


| Name | Number | Description |
| ---- | ------ | ----------- |
| TOOKOFF | 0 |  |
| LANDED | 1 |  |
| HYPERDRIVE | 2 |  |
| HIDING_IN_A_MOUTH | 3 |  |







<a name="sfapi.v1.Starfriends"/>

### Starfriends


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetFilm | [GetFilmRequest](#sfapi.v1.GetFilmRequest) | [GetFilmResponse](#sfapi.v1.GetFilmRequest) | Get a single Film |
| ListFilms | [ListFilmsRequest](#sfapi.v1.ListFilmsRequest) | [ListFilmsResponse](#sfapi.v1.ListFilmsRequest) | Get a list of Films |
| GetVehicle | [GetVehicleRequest](#sfapi.v1.GetVehicleRequest) | [GetVehicleResponse](#sfapi.v1.GetVehicleRequest) | Get a single Vehicle |
| ListVehicles | [ListVehiclesRequest](#sfapi.v1.ListVehiclesRequest) | [ListVehiclesResponse](#sfapi.v1.ListVehiclesRequest) | Get a list of Vehicles |
| GetStarship | [GetStarshipRequest](#sfapi.v1.GetStarshipRequest) | [GetStarshipResponse](#sfapi.v1.GetStarshipRequest) | Get a single Starship |
| ListStarships | [ListStarshipsRequest](#sfapi.v1.ListStarshipsRequest) | [ListStarshipsResponse](#sfapi.v1.ListStarshipsRequest) | Get a list of Starships |
| GetSpecies | [GetSpeciesRequest](#sfapi.v1.GetSpeciesRequest) | [GetSpeciesResponse](#sfapi.v1.GetSpeciesRequest) | Get a single Species |
| ListSpecies | [ListSpeciesRequest](#sfapi.v1.ListSpeciesRequest) | [ListSpeciesResponse](#sfapi.v1.ListSpeciesRequest) | Get a list of Species |
| GetPlanet | [GetPlanetRequest](#sfapi.v1.GetPlanetRequest) | [GetPlanetResponse](#sfapi.v1.GetPlanetRequest) | Get a single Planet |
| ListPlanets | [ListPlanetsRequest](#sfapi.v1.ListPlanetsRequest) | [ListPlanetsResponse](#sfapi.v1.ListPlanetsRequest) | Get a list of Planets |
| GetPerson | [GetPersonRequest](#sfapi.v1.GetPersonRequest) | [GetPersonResponse](#sfapi.v1.GetPersonRequest) | Get a single Person |
| ListPeople | [ListPeopleRequest](#sfapi.v1.ListPeopleRequest) | [ListPeopleResponse](#sfapi.v1.ListPeopleRequest) | Get a list of People |
| ListStarshipActions | [ListStarshipActionsRequest](#sfapi.v1.ListStarshipActionsRequest) | [StarshipAction](#sfapi.v1.ListStarshipActionsRequest) | Watch starships doing stuff, in realtime |
| ValidateSpecies | [ValidateSpeciesRequest](#sfapi.v1.ValidateSpeciesRequest) | [ValidateSpeciesResponse](#sfapi.v1.ValidateSpeciesRequest) | is the given species valid? |





## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |
