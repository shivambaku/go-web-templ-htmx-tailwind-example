-- Create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "username" character varying(255) NOT NULL, "hashed_password" character varying(255) NOT NULL, "created_at" timestamp NOT NULL, "updated_at" timestamp NOT NULL, PRIMARY KEY ("id"));
