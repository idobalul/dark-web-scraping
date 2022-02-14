# Dark Web Scraping

## Intro
This is an API that scrapes information from paste sites in the dark web (currently scrapes only from stronghold paste), saves the pastes in a DB and give analysis about the data.

## Technologies Used
- Gin
- Colly
- MySQL
- Docker

## How to Use
### Using docker-compose
1. Fork or clone the repo and open it in your IDE of choice.  
2. Open your terminal and run `docker-compose up --build`  
3. Wait until the containers will finish their setup (the Gin server container will restart multiple times until the MySQL container will be ready).  
4. Everything is running and ready to work, the API server is running on port 8080.  

### Running Without Docker
1. Fork or clone the repo and open it in your IDE of choice.  
2. Setup your MySQL locally.
3. For your MySQL DB run the following code in the MySQL workbench:
    ```
    CREATE TABLE IF NOT EXISTS pastes (id INT NOT NULL AUTO_INCREMENT, title VARCHAR(255) NOT NULL, author VARCHAR(255) NOT NULL, content LONGTEXT NOT NULL, date VARCHAR(128) NOT NULL, PRIMARY KEY (id), UNIQUE INDEX date_UNIQUE (date ASC) VISIBLE);
    ```
4. Edit the config.go file in the db folder to have your MySQL credentials.   
5. Now open your terminal and run the following commands:
    ```
    go mod download
    docker run -it -p 8118:8118 -p 9050:9050 -d dperson/torproxy
    go run .
    ```
    The first command download all the necessary dependencies.  
    The second command creates the TOR proxy container.  
    The third command runs the server.  
6. Everything is running and ready to work, the API server is running on port 8080.  

## Endpoints
Right now the API has only one route "/scrape" that handles everything.  
More endpoints maybe added in the future.  

## Versions
- V1.0 - First working version with basic scraping and analysis.  