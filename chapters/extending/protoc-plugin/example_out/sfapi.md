# sfapi.v1

Here you could put any kind of content you like.

Maybe put instructions about cloning your org's primary proto IDL repo, or instructions on getting started with tools, or whatever you like.

## Dependencies

* `google/api/annotations.proto`

## Messages

### Film
|Number|Name|Type|Description|
|----|----|----|----|
|1|title|String||
|2|episode_id|Int32||
|3|opening_crawl|String||
|4|director|String||
|5|producer|String||
|6|release_date|String||
|7|character_ids|String|Person.id|
|8|planet_ids|String|Planet.id|
|9|starship_ids|String|Starship.id|
|10|vehicle_ids|String|Vehicle.id|
|11|species_ids|String|Species.id|
|12|id|String||


### GetFilmRequest
|Number|Name|Type|Description|
|----|----|----|----|
|1|id|String||


### GetFilmResponse
|Number|Name|Type|Description|
|----|----|----|----|
|1|film|[Film](#film)||


### ListFilmsRequest
No fields

### ListFilmsResponse
|Number|Name|Type|Description|
|----|----|----|----|
|1|films|[Film](#film)||


### Vehicle
|Number|Name|Type|Description|
|----|----|----|----|
|1|name|String||
|2|model|String||
|3|manufacturer|String||
|4|cost_in_credits|Int32||
|5|length|Float||
|6|max_atmosphering_speed|Int32||
|7|crew|Int32||
|8|passengers|Int32||
|9|cargo_capacity|Int32||
|10|consumables|String||
|11|vehicle_class|String||
|12|pilot_ids|String|Person.id|
|13|film_ids|String|Film.id|
|14|id|String||


### GetVehicleRequest
|Number|Name|Type|Description|
|----|----|----|----|
|1|id|String||


### GetVehicleResponse
|Number|Name|Type|Description|
|----|----|----|----|
|1|vehicle|[Vehicle](#vehicle)||


### ListVehiclesRequest
No fields

### ListVehiclesResponse
|Number|Name|Type|Description|
|----|----|----|----|
|1|vehicles|[Vehicle](#vehicle)||


### Starship
|Number|Name|Type|Description|
|----|----|----|----|
|1|name|String||
|2|model|String||
|3|manufacturer|String||
|4|cost_in_credits|Int64||
|5|length|Float||
|6|crew|Int32||
|7|passengers|Int32||
|8|cargo_capacity|Int64||
|9|consumables|String||
|10|hyperdrive_rating|Float||
|11|mglt|Int32||
|12|starship_class|String||
|13|pilot_ids|String|Person.id|
|14|film_ids|String|Film.id|
|15|id|String||
|16|max_atmosphering_speed|Int64||


### GetStarshipRequest
|Number|Name|Type|Description|
|----|----|----|----|
|1|id|String||


### GetStarshipResponse
|Number|Name|Type|Description|
|----|----|----|----|
|1|starship|[Starship](#starship)||


### ListStarshipsRequest
No fields

### ListStarshipsResponse
|Number|Name|Type|Description|
|----|----|----|----|
|1|starships|[Starship](#starship)||


### Species
|Number|Name|Type|Description|
|----|----|----|----|
|1|name|String||
|2|classification|String||
|3|designation|String||
|4|average_height|Int32||
|5|skin_colors|String||
|6|eye_colors|String||
|7|average_lifespan|Int32||
|8|homeworld|String||
|9|language|String||
|10|people_ids|String|Person.id|
|11|film_ids|String|Film.id|
|12|id|String||
|13|hair_colors|String||


### GetSpeciesRequest
|Number|Name|Type|Description|
|----|----|----|----|
|1|id|String||


### GetSpeciesResponse
|Number|Name|Type|Description|
|----|----|----|----|
|1|species|[Species](#species)||


### ListSpeciesRequest
No fields

### ListSpeciesResponse
|Number|Name|Type|Description|
|----|----|----|----|
|1|species|[Species](#species)||


### Planet
|Number|Name|Type|Description|
|----|----|----|----|
|1|name|String||
|2|rotation_period|Int32||
|3|orbital_period|Int32||
|4|diameter|Int32||
|5|climate|String||
|6|gravity|Float||
|7|terrain|String||
|8|surface_water|Float||
|9|population|Int64||
|10|resident_ids|String|Person.id|
|11|film_ids|String|Film.id|
|12|id|String||


### GetPlanetRequest
|Number|Name|Type|Description|
|----|----|----|----|
|1|id|String||


### GetPlanetResponse
|Number|Name|Type|Description|
|----|----|----|----|
|1|planet|[Planet](#planet)||


### ListPlanetsRequest
No fields

### ListPlanetsResponse
|Number|Name|Type|Description|
|----|----|----|----|
|1|planets|[Planet](#planet)||


### Person
|Number|Name|Type|Description|
|----|----|----|----|
|1|name|String||
|2|height|Int32||
|3|mass|Float||
|4|hair_color|String||
|5|skin_color|String||
|6|eye_color|String||
|7|birth_year|String||
|8|gender|String||
|9|homeworld|String||
|10|film_ids|String|Film.id|
|11|species_ids|String|Species.id|
|12|vehicle_ids|String|Vehicle.id|
|13|starship_ids|String|Starship.id|
|14|id|String||


### GetPersonRequest
|Number|Name|Type|Description|
|----|----|----|----|
|1|id|String||


### GetPersonResponse
|Number|Name|Type|Description|
|----|----|----|----|
|1|person|[Person](#person)||


### ListPeopleRequest
No fields

### ListPeopleResponse
|Number|Name|Type|Description|
|----|----|----|----|
|1|people|[Person](#person)||


### ListStarshipActionsRequest
No fields

### StarshipAction
|Number|Name|Type|Description|
|----|----|----|----|
|1|starship|[Starship](#starship)||
|2|action|[StarshipAction.Action](#starshipactionaction)||


### InvalidKey
|Number|Name|Type|Description|
|----|----|----|----|
|1|key|String||
|2|message|String||


### ValidateSpeciesRequest
|Number|Name|Type|Description|
|----|----|----|----|
|1|name|String||


### ValidateSpeciesResponse
No fields



## Enums


### StarshipAction.Action

* `TOOKOFF`: 0
* `LANDED`: 1
* `HYPERDRIVE`: 2
* `HIDING_IN_A_MOUTH`: 3



## Sevices

`~` denotes a streaming input/outut.

### Starfriends
|Name|Input|Output|Options|Description|
|----|----|----|----|----|
|GetFilm|[GetFilmRequest](#getfilmrequest)|[GetFilmResponse](#getfilmresponse)|**GET** /sfapi/v1/films/{id}|Get a single Film|
|ListFilms|[ListFilmsRequest](#listfilmsrequest)|[ListFilmsResponse](#listfilmsresponse)|**GET** /sfapi/v1/films|Get a list of Films|
|GetVehicle|[GetVehicleRequest](#getvehiclerequest)|[GetVehicleResponse](#getvehicleresponse)|**GET** /sfapi/v1/vehicles/{id}|Get a single Vehicle|
|ListVehicles|[ListVehiclesRequest](#listvehiclesrequest)|[ListVehiclesResponse](#listvehiclesresponse)|**GET** /sfapi/v1/vehicles|Get a list of Vehicles|
|GetStarship|[GetStarshipRequest](#getstarshiprequest)|[GetStarshipResponse](#getstarshipresponse)|**GET** /sfapi/v1/starships/{id}|Get a single Starship|
|ListStarships|[ListStarshipsRequest](#liststarshipsrequest)|[ListStarshipsResponse](#liststarshipsresponse)|**GET** /sfapi/v1/starships|Get a list of Starships|
|GetSpecies|[GetSpeciesRequest](#getspeciesrequest)|[GetSpeciesResponse](#getspeciesresponse)|**GET** /sfapi/v1/species/{id}|Get a single Species|
|ListSpecies|[ListSpeciesRequest](#listspeciesrequest)|[ListSpeciesResponse](#listspeciesresponse)|**GET** /sfapi/v1/species|Get a list of Species|
|GetPlanet|[GetPlanetRequest](#getplanetrequest)|[GetPlanetResponse](#getplanetresponse)|**GET** /sfapi/v1/planets/{id}|Get a single Planet|
|ListPlanets|[ListPlanetsRequest](#listplanetsrequest)|[ListPlanetsResponse](#listplanetsresponse)|**GET** /sfapi/v1/planets|Get a list of Planets|
|GetPerson|[GetPersonRequest](#getpersonrequest)|[GetPersonResponse](#getpersonresponse)|**GET** /sfapi/v1/people/{id}|Get a single Person|
|ListPeople|[ListPeopleRequest](#listpeoplerequest)|[ListPeopleResponse](#listpeopleresponse)|**GET** /sfapi/v1/people|Get a list of People|
|ListStarshipActions|[ListStarshipActionsRequest](#liststarshipactionsrequest)|[~StarshipAction](#starshipaction)||Watch starships doing stuff, in realtime|
|ValidateSpecies|[ValidateSpeciesRequest](#validatespeciesrequest)|[ValidateSpeciesResponse](#validatespeciesresponse)||is the given species valid?|


