# GoTravel

![alt text](https://raw.githubusercontent.com/Cleopatric/go-travel/main/icon.png?raw=true)

Service to search IATA codes for cities, airports and airlines.
Service supported two languages - English / Russian

## Endpoints

To get IATA codes, use the following links:

- Get list of all cities
    ```http
    GET http://0.0.0.0:8081/cities
    ```
- Get list of all cities with search
    ```http
    GET http://0.0.0.0:8081/cities?search=Minsk
    ```
- Get list of airlines
    ```http
    GET http://0.0.0.0:8081/airlines
    ```
- Get list of airports
    ```http
    GET http://0.0.0.0:8081/airports
    ```
- Get city by city name
    ```http
    GET http://0.0.0.0:8081/get_city/{city_name}
    ```
- Get airline by airline name
    ```http
    GET http://0.0.0.0:8081/get_airline/{airline}
    ```
- Get airports by city IATA code
    ```http
    GET http://0.0.0.0:8081/get_airport/{city_code}
    ```

