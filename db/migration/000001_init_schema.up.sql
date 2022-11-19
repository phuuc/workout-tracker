CREATE TABLE IF NOT EXISTS "users" (
  "id" serial PRIMARY KEY,
  "email" varchar(255) unique,
  "hashed_passwd" char(60),
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE IF NOT EXISTS "excercise" (
  "id" serial PRIMARY KEY,
  "name" varchar(50) not null,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE IF NOT EXISTS "workout" (
  "id" serial PRIMARY KEY,
  "reps" int not null,
  "sets" int not null,
  "user_id" int not null,
  "excercise_id" int not null,
  "created_at" timestamp,
  "updated_at" timestamp
);

ALTER TABLE "workout" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "workout" ADD FOREIGN KEY ("excercise_id") REFERENCES "excercise" ("id");
