BEGIN;
COPY users FROM '/var/sampledata/users.csv' DELIMITER ',' CSV HEADER;
COPY cuisines FROM '/var/sampledata/cuisines.csv' DELIMITER ',' CSV HEADER; 
COPY image_store FROM '/var/sampledata/imagestore.csv' DELIMITER ',' CSV HEADER;
COPY recipes FROM '/var/sampledata/recipes.csv' DELIMITER ',' CSV HEADER; 
COPY ingredients FROM '/var/sampledata/ingredients.csv' DELIMITER ',' CSV HEADER; 
COPY recipe_ingredients FROM '/var/sampledata/recipeingredients.csv' DELIMITER ',' CSV HEADER; 
COPY saved_recipes FROM '/var/sampledata/savedrecipes.csv' DELIMITER ',' CSV HEADER; 
COMMIT;