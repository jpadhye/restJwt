# restJwt
Learning project to Implement JWT based authorization using GO in course https://www.udemy.com/course/build-jwt-authenticated-restful-apis-with-golang/

## Pre-requisities:
    - Create free ElephantSQL account and update the url in .env file
    - Create the following table
    ```
        create table users (
        id serial primary key,
        email text not null unique,
        password text not null
        );
    ```


