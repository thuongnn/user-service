CREATE TABLE "user" (
	user_id SERIAL PRIMARY KEY NOT NULL,
    username VARCHAR (255) UNIQUE NOT NULL,
    password VARCHAR (255) UNIQUE NOT NULL,
    salt VARCHAR (255) UNIQUE NOT NULL,
	creation_time timestamp default CURRENT_TIMESTAMP,
    update_time timestamp  default CURRENT_TIMESTAMP,
    deleted boolean DEFAULT false NOT NULL
);