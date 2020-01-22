*For work breakdown, see Notion.*
pipeline test

# BonCuisine <a name="title"></a>
A dynamic recipe application that gives users a variety of highly original and tasty recipes to choose from. Users can bookmark recipes they find interesting and will receive push notifications when new recipes are added.  

## Table of Contents
- [Project Summary](#projectsummary)
- [Project Specifications](#projectspecs)
    - [Non-Technical Overview](#projectspecs_nontech)
    - [Technical Overview](#projectspecs_tech)
- [Frontend](#frontend)
    - [Visual Design](#frontend_design)
    - [Visual Details](#frontend_details)
    - [State Management](#frontend_redux)
- [Backend](#backend)
- [Database](#db)
    - [Normalization](#db_norm)
    - [Integrity Rules](#db_integrity)
    - [Indexing](#db_index)
    - [Index Clustering](#db_cluster)
- [Security](#security)
- [Cloud Services](#cloud)
- [Version Control](#version_control)

## Project Summary: <a name="projectsummary"></a>
A full-stack mobile application made with Flutter, Golang, &amp; Postgresql. This project will serve to give us practice developing and deploying an end-to-end full-stack mobile application. We will utilize the Agile & Scrum Methodology using Notion and Git. This README will serve as a reference for future projects. 


## Project Specifications: <a name="projectspecs"></a>
### Non-Technical Overview <a name="projectspecs_nontech"></a>
A new user should be able to create an account using our app. An existing user should be able to log in. Once logged in, the user will have access to recipes that can be viewed based on category. They should also be able to view their saved recipes. The user can click on a recipe and the app will navigate to the recipe details page, where the user can see more info about the recipe, including how long it takes to make, the number of servings, the category, and the ingredients. Each ingredient has a name, an amount, and an image. The user can also bookmark this recipe on this page, this will add this recipe to this user's saved recipes. The user should receive a notification when a new recipe is added (note: a user cannot add their own recipe). Lastly, a user should be able to log out. 

### Technical Overview <a name="projectspecs_tech"></a>
The project will be broken down into three parts: Frontend, Server, & Database. All of which will be deployed seperately. 
The app should be able to securely handle user authentication and account creation. The webserver should have an API that the client app can access securely - with proper authentication. The API should be RESTful and have a websocket to communicate securely with the client. The frontend should integrate the API to display all recipes, all user saved recipes, and all the ingredients associated with recipes. The frontend should also access the websocket to display notifications when a new recipe is added in the backend. The backend will be connected to the database that models the necessary entities. The backend will be hosted using AWS. 


## Frontend: <a name="frontend"></a>
### Visual Design: <a name="frontend_design"></a>
![BonCuisine App Design](images/BoncuisineAppDesignFull.jpg?raw=true "Title")

### Visual Details: <a name="frontend_details"></a>
The app will consist of three screens: 
- Login/Signup Page
    - Username & Password for login
    - Username, Password, Email for signup
    - Segmented Control Button will switch between logging in view and sign up view. 
    - Logging in / Signing up takes the user to the Recipe Dashboard Page with a hero animation where the name "BonCusine" persists through the screen transisition. 
- Recipe Dashboard Page
    - Button in the top-left to view saved recipes. 
    - Horizontal list view of text buttons for each category (ie. Italian, Asian, Contemporary, etc..) to view recipes for that specific category. 
    - Horizontal list view of recipe images, with the name, time, category, and servings at the bottom, that change every time the user swipes left or right to a new recipe.
    - Clicking on a recipe takes you to Recipe Detail Page with a hero animation where the image persists through the screen transition.
- Recipe Detail Page
    - Back button in the top left to go back to Recipe Dashboard page, with a hero animation again. 
    - Bookmark button in the top right to save the recipe. 
    - Display recipe name, time, servings, and category.
    - Large recipe image in the middle
    - Horizontal list of ingredients at the bottom. Ingredients include name, ingredient image, and amount.
    
### Redux State Management <a name="frontend_redux"></a>
TODO add redux design here. 


## Backend: <a name="backend"></a>
The backend will consist of the Rest API and Websocket API. 
TODO add API design here. 


## Database: <a name="db"></a>
#### Normalization Guidelines: <a name="db_norm"></a>
First Normal Form (1NF): A table is 1NF if every cell contains a single value, not a list of values. 

Second Normal Form (2NF): A table is 2NF, if it is 1NF and every non-key column is fully dependent on the primary key. All non-key columns should be directly associated with the main purpose of the table. 

Third Normal Form (3NF): A table is 3NF, if it is 2NF and the non-key columns are independent of each others. If one non-key column changes directly as a result of another non-key column changing, then it is not indepedent of the others. 

#### Integrity Rules: <a name="db_integrity"></a>
Entity Integrity Rule: The primary key cannot contain NULL.
Referential Integrity Rule: Each foreign key value must be matched to a primary key value in the table referenced. 
- If the value of the key changes in the parent table (e.g., the row updated or deleted), all rows with this foreign key in the child table(s) must be handled accordingly. Cascade the change by deleting the records in the child tables accordingly.

If updating a child table ever requires you to update the parent as well, then we have improper design.
In other words, if updating a table with a foreign key requires you to update the associated record in the foreign key's table, that ain't good. 

#### Column Indexing <a name="db_index"></a>
You could create index on selected column(s) to facilitate data searching and retrieval. An index speeds up data access for SELECT, but may slow down INSERT, UPDATE, and DELETE, as the index needs to be updated each time. Chose indexes wisely. 

#### Clustered Indexing <a name="db_cluster"></a>
TODO

## Security: <a name="security"></a>
TODO add security features here. @RahulToppur 

## Cloud Hosting: <a name="cloud"></a>
TODO add cloud hosting here. 

## Version Control: <a name="version_control"></a>
We will use [GitFlow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow
).
![Git Flow Diagram](images/gitflowdiagram.png?raw=true "Title")
When making a new feature branch, follow the name convention: feature-[your intitials]-[name of feature]
Example: feature-nc-updated-ui

