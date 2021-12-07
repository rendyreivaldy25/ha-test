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
# Project Structure
For the structure of this project, I use clean code architecture which will separate the project into several sections :
* Controller     
This is acted as a gateway to our business logics
* Service     
Where business logics are located
* Model     
Where the data model, interfaces located     
     
Why I choose this architecture, because with this structure, it will make the project easier to create unit test to test the business logic and it will be much easier to scale up. 
# Config Files
In the configs folder, I created several config files. The contents of the files are the same, the difference is only the purpose of the configs. 
* config.json     
Config file for the current running app
* config.json.dev     
Config file for development environment
* config.json.prod     
Config file for production environment
* config.json.sample     
Config file as a sample, if someone new start to contribute to this project     
     
With this method, we can make the config interchangeable in the CI/CD process. So, shifting environment from development to production will be seamless
# How To Build And Run Project Using Docker
```$ docker build -t ha-test .```     
```$ docker run -p 10000:10000 -it ha-test```
# How To Run The Project
```$ go run .```
# How To Run Unit Test
```$ go test ./... -v```
# Questions 
For the additional questions that needs written answer, I already created a file named `QUESTIONS.txt` located in the root of the project contains my answer.