CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "email" varchar NOT NULL,
    "password" varchar NOT NULL,
    "email_verification_code" varchar,
    "email_verification_sent_at" varchar,
    "forgot_password_code" varchar,
    "forgot_password_sent_at" varchar,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "created_by" bigint NOT NULL,
    "updated_by" bigint,
    "deleted_by" bigint
);

CREATE TABLE "products" (
    "id" bigserial PRIMARY KEY,
    "title" varchar NOT NULL,
    "description" varchar,
    "seller_name" varchar,
    "photo_url" varchar, 
    "price" numeric,
    "category_name" varchar,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "created_by" bigint NOT NULL,
    "updated_by" bigint,
    "deleted_by" bigint
);

CREATE INDEX ON "products" ("title");

COMMENT ON COLUMN "users"."email_verification_code" IS 'code when user register new account';

COMMENT ON COLUMN "users"."forgot_password_code" IS 'code when user forgot password';