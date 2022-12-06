# Introduction 
A dating application

# Functionalities
The dating app have the following functionalities
1. Create random user
2. Create user profile
3. Login via user credentials
4. Refresh the token
5. Get the potential match for the user
6. Swipe the user based on the match skills
7. Filter the user profile by age or gender or both
8. User profiles can filter or listed out as per the swipes stats
9. Once the swap is done, next time the profile is not listed on the potential match

# Environments
There are two environments configured on the project
1. local/Dev
2. Prod

## local/Dev Environment
The local/dev enviroginment can be setup from the .env.local file which is listed on the project no need to create seperatly.

### How to run local/dev project
1. First just have to run the docker file by running the following command:
```bash
docker compose up
or
go run main.go
```

But for this you need to create a docker image for mongo db

### How to run on Prod
1. There is nothing to do, everything is setup on AWS also using serverless just hitting the api endpoints
2. Mongodb is configured on AWS, also mongodb uri is mentioned on the .env.prod

if you want to deploy the project on serverless, first need to make a change in config.service.go and add this line:
```bash
v.SetConfigName(".env.prod")
make deploy_prod
```

## API's collection
Local and Prod api collection.json is added on the project directory just simple import the collection in to postman, it works

## Swagger Documentation
Swagger docs is also added, if you want to explore the api's from swagger docs just simple run this URL:
Local Swagger URL: http://localhost:8080/swagger/index.html#/

### PROD endpoints::
```bash

GET - https://d5mo5y0bc3.execute-api.eu-central-1.amazonaws.com/prod/v1/hello
POST - https://d5mo5y0bc3.execute-api.eu-central-1.amazonaws.com/prod/v1/auth/login
POST - https://d5mo5y0bc3.execute-api.eu-central-1.amazonaws.com/prod/v1/auth/refresh
POST - https://d5mo5y0bc3.execute-api.eu-central-1.amazonaws.com/prod/v1/user/create
POST - https://d5mo5y0bc3.execute-api.eu-central-1.amazonaws.com/prod/v1/profile/create
GET - https://d5mo5y0bc3.execute-api.eu-central-1.amazonaws.com/prod/v1/profile
GET - https://d5mo5y0bc3.execute-api.eu-central-1.amazonaws.com/prod/v1/profile/filter
GET - https://d5mo5y0bc3.execute-api.eu-central-1.amazonaws.com/prod/v1/swipe
```

## Solution Implementation:
1. Write an endpoint to create a random user at /user/create?
   1. User can create randomly by just hitting this endpoint all the data can be generated randomly and stored into mongodb 
2. Write an endpoint to fetch profiles of potential matches at /profiles?
   1. For this designed three endpoints:
      ```bash
         POST - https://d5mo5y0bc3.execute-api.eu-central-1.amazonaws.com/prod/v1/profile/create
         GET - https://d5mo5y0bc3.execute-api.eu-central-1.amazonaws.com/prod/v1/profile
         GET - https://d5mo5y0bc3.execute-api.eu-central-1.amazonaws.com/prod/v1/profile/filter
         ```
      2. Potential match criteria
         1. Create profile is used to create a user profile because on behalf of the profile user can find a potential match, so profile creation is neccessary
         2. Find potential match for the user on behalf of the following criteria:
            1. Profiles must be under the range of user min or max age criteria
            2. Fetch only those profiles who have the location mentioned by the user
            3. If User interested in match with the profiles we can get the all potential match for the specific user
3. Write an endpoint to respond to a profile at /swipe?
   1. Swipe can be done on the following criteria:
      1. For example if user a is new and first time user can swipe any profile we store the user id and the preference of the user on the profile of user b , user a likes
      2. Second if the user b also like the profile of user a, first we check the user id is listed in my swipe dict if yes so we can have a perfect match
      
4. Write an endpoint to authenticate a user at /login?
   1. User can login via username and passowrd will get the user access token and refresh token in response.
5. Extend /profiles and /swipe to be authenticated by a login token?
   1. The token we get in response of the login, same token is used as a header for the profile and swipe endpoints, otherthan that won't be able these endpoints because user is not authroized
6. Extend /profiles to filter results by age and or gender?
   1. In profile functionality, user can hit the below endpoint and get the desired filter profiles on the based of gender or age filter
    ```bash
    GET - https://d5mo5y0bc3.execute-api.eu-central-1.amazonaws.com/prod/v1/profile/filter
    ```
7. Extend /profiles to sort profiles by attractiveness?
   1. On the bases of swipe count, user can get the potential match of the most likely profile who have the most swipes

Tools and Technologies Used:
1. API Backend: Golang - YES
2. JSON for request and response payloads - YES
3. Database:
   1. MongoDB  - YES
4. Source Control Git - YES
5. CloudFormation/serverless - YES
    

### Logs
logs can be store in a file called access.log in logs directory.

### Request and Response Screenshots
![](https://i.postimg.cc/9MsDht2B/Screenshot-2022-12-06-at-4-47-09-PM.png)
![](https://i.postimg.cc/9QHyT2vB/Screenshot-2022-12-06-at-4-47-20-PM.png)
![](https://i.postimg.cc/8CKHL3dP/Screenshot-2022-12-06-at-5-04-01-PM.png)
![](https://i.postimg.cc/KjnBpsTT/Screenshot-2022-12-06-at-5-09-44-PM.png)
![](https://i.postimg.cc/MHxvHZcc/Screenshot-2022-12-06-at-5-10-05-PM.png)
![](https://i.postimg.cc/VNqw6qXR/Screenshot-2022-12-06-at-5-10-20-PM.png)
