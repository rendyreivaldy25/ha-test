# Drone Corp 2218 
Space exploration is underway and mostly done by private companies. You joined Engineering department of one of the private government contractors, Atlas Corporation. In that year and age, everything is automated, so survey and data gathering is done by drones. Unfortunately, although drones are perfectly capable of gathering the data, they have issues with locating databank to upload gathered data. You, as the most promising recruit of Atlas Corp, were tasked to develop a drone navigation service (DNS):
* each observed sector of the galaxy has unique numeric SectorID assigned to it
* each sector will have at least one DNS deployed
* each sector has different number of drones deployed at any given moment
* it’s future, but not that far, so drones will still use JSON REST API
DNS should be designed with following constraints in mind
* SectorID can be considered constant (it won’t change during runtime)
* DNS request example:
```javascript
{
    // x, y, z are coords, vel is velocity
    // (values sent as strings, but must be treated as floating point number)
    "x": "123.12",
    "y": "456.56",
    "z": "789.89",
    "vel": "20.0",
}
```
* Complex math needed to serve a response:
```loc = x*SectorID + y*SectorID + z*SectorID + vel```
* DNS response example:
```javascript
{
    "loc": 1389.57,
}
```
# How To Run The Project
```go run .```
# How To Run Unit Test
```go test . -v```
# Questions 
For the added questions that needs written answer, I already created it in file named `QUESTIONS.txt` located in the root of the project.