BEGIN;
/*COPY users FROM '/var/sampledata/users.csv' DELIMITER ',' CSV HEADER;*/
INSERT INTO users(user_id,name,user_name,email,password,created_on) VALUES (1,'Seth McGee','Charles','hej@zi.th','IN','2016-06-22 19:10:25-07') ON CONFLICT DO NOTHING;
INSERT INTO users(user_id,name,user_name,email,password,created_on) VALUES (2,'Roy Phillips','Stanley','nohocu@zo.tr','NV','2016-06-22 19:10:25-07') ON CONFLICT DO NOTHING;
INSERT INTO users(user_id,name,user_name,email,password,created_on) VALUES (3,'Joshua Buchanan','Todd','ap@ziuv.ax','OK','2016-06-22 19:10:25-07') ON CONFLICT DO NOTHING;
INSERT INTO users(user_id,name,user_name,email,password,created_on) VALUES (4,'Harriett McDonald','Jay','bimpik@josovik.ie','SC','2016-06-22 19:10:25-07') ON CONFLICT DO NOTHING;
INSERT INTO users(user_id,name,user_name,email,password,created_on) VALUES (5,'Franklin Patrick','Hunter','ow@civul.lv','NY','2016-06-22 19:10:25-07') ON CONFLICT DO NOTHING;

/*COPY cuisines FROM '/var/sampledata/cuisines.csv' DELIMITER ',' CSV HEADER;*/
INSERT INTO cuisines(cuisine_id,cuisine_name) VALUES (1,'American') ON CONFLICT DO NOTHING;
INSERT INTO cuisines(cuisine_id,cuisine_name) VALUES (2,'French') ON CONFLICT DO NOTHING;
INSERT INTO cuisines(cuisine_id,cuisine_name) VALUES (3,'Japanese') ON CONFLICT DO NOTHING;
INSERT INTO cuisines(cuisine_id,cuisine_name) VALUES (4,'Spanish') ON CONFLICT DO NOTHING;

/*COPY image_store FROM '/var/sampledata/imagestore.csv' DELIMITER ',' CSV HEADER;*/
INSERT INTO image_store(image_id,image_data) VALUES (1,6020091711) ON CONFLICT DO NOTHING;
INSERT INTO image_store(image_id,image_data) VALUES (2,4726389271) ON CONFLICT DO NOTHING;
INSERT INTO image_store(image_id,image_data) VALUES (3,4937009651) ON CONFLICT DO NOTHING;
INSERT INTO image_store(image_id,image_data) VALUES (4,3139135474) ON CONFLICT DO NOTHING;
INSERT INTO image_store(image_id,image_data) VALUES (5,4235235523) ON CONFLICT DO NOTHING;
INSERT INTO image_store(image_id,image_data) VALUES (6,5235253253) ON CONFLICT DO NOTHING;
INSERT INTO image_store(image_id,image_data) VALUES (7,764575656) ON CONFLICT DO NOTHING;
INSERT INTO image_store(image_id,image_data) VALUES (8,567657643) ON CONFLICT DO NOTHING;
INSERT INTO image_store(image_id,image_data) VALUES (9,7463486657) ON CONFLICT DO NOTHING;
INSERT INTO image_store(image_id,image_data) VALUES (10,8659875764) ON CONFLICT DO NOTHING;
INSERT INTO image_store(image_id,image_data) VALUES (11,3763547664) ON CONFLICT DO NOTHING;
INSERT INTO image_store(image_id,image_data) VALUES (12,346734677) ON CONFLICT DO NOTHING;
INSERT INTO image_store(image_id,image_data) VALUES (13,474574757) ON CONFLICT DO NOTHING;
INSERT INTO image_store(image_id,image_data) VALUES (14,3474563343) ON CONFLICT DO NOTHING;
INSERT INTO image_store(image_id,image_data) VALUES (15,3647374554) ON CONFLICT DO NOTHING;
INSERT INTO image_store(image_id,image_data) VALUES (16,4564647475) ON CONFLICT DO NOTHING;
INSERT INTO image_store(image_id,image_data) VALUES (17,4575473547) ON CONFLICT DO NOTHING;

/*COPY recipes FROM '/var/sampledata/recipes.csv' DELIMITER ',' CSV HEADER;*/
INSERT INTO recipes(recipe_id,recipe_name,r_time,num_servings,difficulty,image_id,cuisine_id) VALUES (1,'Hamburger',30,3,'medium',1,1) ON CONFLICT DO NOTHING;
INSERT INTO recipes(recipe_id,recipe_name,r_time,num_servings,difficulty,image_id,cuisine_id) VALUES (2,'Canele',60,4,'hard',2,2) ON CONFLICT DO NOTHING;
INSERT INTO recipes(recipe_id,recipe_name,r_time,num_servings,difficulty,image_id,cuisine_id) VALUES (3,'Sushi',20,7,'easy',3,3) ON CONFLICT DO NOTHING;
INSERT INTO recipes(recipe_id,recipe_name,r_time,num_servings,difficulty,image_id,cuisine_id) VALUES (4,'Tamale',20,2,'easy',4,4) ON CONFLICT DO NOTHING;

/*COPY ingredients FROM '/var/sampledata/ingredients.csv' DELIMITER ',' CSV HEADER;*/
INSERT INTO ingredients(ingredient_id,ingredient_name,image_id) VALUES (1,'Fish',5) ON CONFLICT DO NOTHING;
INSERT INTO ingredients(ingredient_id,ingredient_name,image_id) VALUES (2,'Beef',6) ON CONFLICT DO NOTHING;
INSERT INTO ingredients(ingredient_id,ingredient_name,image_id) VALUES (3,'Cheese',7) ON CONFLICT DO NOTHING;
INSERT INTO ingredients(ingredient_id,ingredient_name,image_id) VALUES (4,'Sugar',8) ON CONFLICT DO NOTHING;
INSERT INTO ingredients(ingredient_id,ingredient_name,image_id) VALUES (5,'Flower',9) ON CONFLICT DO NOTHING;
INSERT INTO ingredients(ingredient_id,ingredient_name,image_id) VALUES (6,'Butter',10) ON CONFLICT DO NOTHING;
INSERT INTO ingredients(ingredient_id,ingredient_name,image_id) VALUES (7,'Rice',11) ON CONFLICT DO NOTHING;
INSERT INTO ingredients(ingredient_id,ingredient_name,image_id) VALUES (8,'Bread',12) ON CONFLICT DO NOTHING;
INSERT INTO ingredients(ingredient_id,ingredient_name,image_id) VALUES (9,'Lettuce',13) ON CONFLICT DO NOTHING;
INSERT INTO ingredients(ingredient_id,ingredient_name,image_id) VALUES (10,'Tomato',14) ON CONFLICT DO NOTHING;
INSERT INTO ingredients(ingredient_id,ingredient_name,image_id) VALUES (11,'Corn',15) ON CONFLICT DO NOTHING;
INSERT INTO ingredients(ingredient_id,ingredient_name,image_id) VALUES (12,'Seaweed',16) ON CONFLICT DO NOTHING;
INSERT INTO ingredients(ingredient_id,ingredient_name,image_id) VALUES (13,'Chicken',17) ON CONFLICT DO NOTHING;

/*COPY recipe_ingredients FROM '/var/sampledata/recipeingredients.csv' DELIMITER ',' CSV HEADER;*/
INSERT INTO recipe_ingredients(ingredient_id,recipe_id,amount) VALUES (2,1,7) ON CONFLICT DO NOTHING;
INSERT INTO recipe_ingredients(ingredient_id,recipe_id,amount) VALUES (4,2,7) ON CONFLICT DO NOTHING;
INSERT INTO recipe_ingredients(ingredient_id,recipe_id,amount) VALUES (1,3,2) ON CONFLICT DO NOTHING;
INSERT INTO recipe_ingredients(ingredient_id,recipe_id,amount) VALUES (11,4,7) ON CONFLICT DO NOTHING;
INSERT INTO recipe_ingredients(ingredient_id,recipe_id,amount) VALUES (3,1,6) ON CONFLICT DO NOTHING;
INSERT INTO recipe_ingredients(ingredient_id,recipe_id,amount) VALUES (5,2,7) ON CONFLICT DO NOTHING;
INSERT INTO recipe_ingredients(ingredient_id,recipe_id,amount) VALUES (7,3,9) ON CONFLICT DO NOTHING;
INSERT INTO recipe_ingredients(ingredient_id,recipe_id,amount) VALUES (13,4,0) ON CONFLICT DO NOTHING;
INSERT INTO recipe_ingredients(ingredient_id,recipe_id,amount) VALUES (8,1,4) ON CONFLICT DO NOTHING;
INSERT INTO recipe_ingredients(ingredient_id,recipe_id,amount) VALUES (6,2,0) ON CONFLICT DO NOTHING;
INSERT INTO recipe_ingredients(ingredient_id,recipe_id,amount) VALUES (12,3,7) ON CONFLICT DO NOTHING;
INSERT INTO recipe_ingredients(ingredient_id,recipe_id,amount) VALUES (9,1,7) ON CONFLICT DO NOTHING;
INSERT INTO recipe_ingredients(ingredient_id,recipe_id,amount) VALUES (10,1,8) ON CONFLICT DO NOTHING;

/*COPY saved_recipes FROM '/var/sampledata/savedrecipes.csv' DELIMITER ',' CSV HEADER;*/
INSERT INTO saved_recipes(user_id,recipe_id) VALUES (1,1) ON CONFLICT DO NOTHING;
INSERT INTO saved_recipes(user_id,recipe_id) VALUES (2,2) ON CONFLICT DO NOTHING;
INSERT INTO saved_recipes(user_id,recipe_id) VALUES (3,3) ON CONFLICT DO NOTHING;
INSERT INTO saved_recipes(user_id,recipe_id) VALUES (4,4) ON CONFLICT DO NOTHING;
INSERT INTO saved_recipes(user_id,recipe_id) VALUES (5,1) ON CONFLICT DO NOTHING;
INSERT INTO saved_recipes(user_id,recipe_id) VALUES (1,3) ON CONFLICT DO NOTHING;
INSERT INTO saved_recipes(user_id,recipe_id) VALUES (2,4) ON CONFLICT DO NOTHING;
INSERT INTO saved_recipes(user_id,recipe_id) VALUES (3,2) ON CONFLICT DO NOTHING;

COMMIT;