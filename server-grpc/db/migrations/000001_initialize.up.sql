BEGIN;
CREATE TABLE image_store (
    image_id SERIAL PRIMARY KEY,
    image_data BYTEA NOT NULL
);

CREATE TABLE cuisines (
    cuisine_id SERIAL PRIMARY KEY,
    cuisine_name VARCHAR (255) UNIQUE NOT NULL
);

CREATE TABLE users (
    user_id VARCHAR (255) PRIMARY KEY,
    name VARCHAR (255) NOT NULL,
    user_name VARCHAR (255) UNIQUE NOT NULL,
    email VARCHAR (355) UNIQUE NOT NULL,
    password TEXT UNIQUE NOT NULL,
    created_on TIMESTAMP NOT NULL
);

CREATE TABLE recipes (
    recipe_id SERIAL,
    recipe_name VARCHAR (255) UNIQUE NOT NULL,
    r_time INTEGER NOT NULL,
    num_servings INTEGER NOT NULL,
    difficulty VARCHAR (255) NOT NULL,
    image_id INTEGER NOT NULL,
    cuisine_id INTEGER NOT NULL,
    PRIMARY KEY (recipe_id),
    CONSTRAINT recipes_image_id_fkey FOREIGN KEY (image_id)
        REFERENCES image_store (image_id) MATCH SIMPLE
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT recipes_cuisine_id_fkey FOREIGN KEY (cuisine_id)
        REFERENCES cuisines (cuisine_id) MATCH SIMPLE
        ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE ingredients (
    ingredient_id SERIAL,
    ingredient_name VARCHAR (255) NOT NULL,
    image_id INTEGER NOT NULL,
    PRIMARY KEY (ingredient_id),
    CONSTRAINT ingredients_image_id_fkey FOREIGN KEY (image_id)
        REFERENCES image_store (image_id) MATCH SIMPLE
        ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE saved_recipes (
    user_id VARCHAR(255) NOT NULL,
    recipe_id INTEGER NOT NULL,
    PRIMARY KEY (user_id, recipe_id),
    CONSTRAINT saved_recipe_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES users (user_id) MATCH SIMPLE
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT saved_recipe_recipe_id_fkey FOREIGN KEY (recipe_id)
        REFERENCES recipes (recipe_id) MATCH SIMPLE
        ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE recipe_ingredients (
    ingredient_id INTEGER NOT NULL,
    recipe_id INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    PRIMARY KEY (ingredient_id, recipe_id ),
    CONSTRAINT saved_recipe_ingredient_id_fkey FOREIGN KEY (ingredient_id)
        REFERENCES ingredients (ingredient_id) MATCH SIMPLE
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT saved_recipe_recipe_id_fkey FOREIGN KEY (recipe_id)
        REFERENCES recipes (recipe_id) MATCH SIMPLE
        ON UPDATE CASCADE ON DELETE CASCADE
);
COMMIT;