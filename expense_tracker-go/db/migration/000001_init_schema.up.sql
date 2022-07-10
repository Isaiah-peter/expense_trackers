CREATE TYPE "transaction_statue" AS ENUM (
  'expense',
  'income'
);

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now()) NOT NULL,
  "updated_at" timestamptz DEFAULT (now()) NOT NULL
);

CREATE TABLE "categories" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int  NOT NULL,
  "icon" varchar  NOT NULL,
  "name" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now()) NOT NULL,
  "updated_at" timestamptz DEFAULT (now()) NOT NULL
);

CREATE TABLE "transactions" (
  "id" SERIAL PRIMARY KEY,
  "category_id" int  NOT NULL,
  "user_id" int  NOT NULL,
  "ammout" bigint NOT NULL,
  "notes" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()) NOT NULL,
  "updated_at" timestamptz DEFAULT (now()) NOT NULL,
  "status" transaction_statue NOT NULL
);

CREATE INDEX ON "users" ("name");

CREATE INDEX ON "categories" ("name");

CREATE INDEX ON "transactions" ("category_id");

CREATE INDEX ON "transactions" ("user_id");

CREATE INDEX ON "transactions" ("status");

COMMENT ON COLUMN "transactions"."ammout" IS 'but nagative and positive number';

ALTER TABLE "categories" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");