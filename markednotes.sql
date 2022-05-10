CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "folders" (
  "id" SERIAL PRIMARY KEY,
  "parent_id" int,
  "user_id" int NOT NULL,
  "name" varchar NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "notes" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "folder_id" int,
  "name" varchar NOT NULL,
  "body" text NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "folders" ("id");

CREATE INDEX ON "folders" ("user_id");

CREATE INDEX ON "folders" ("parent_id");

CREATE INDEX ON "notes" ("id");

CREATE INDEX ON "notes" ("user_id");

CREATE INDEX ON "notes" ("folder_id");

ALTER TABLE "folders" ADD FOREIGN KEY ("parent_id") REFERENCES "folders" ("id") ON DELETE CASCADE;

ALTER TABLE "folders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "notes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "notes" ADD FOREIGN KEY ("folder_id") REFERENCES "folders" ("id") ON DELETE CASCADE;
