CREATE SEQUENCE IF NOT EXISTS ports_id_seq;

-- Table Definition
CREATE TABLE "public"."ports" (
    "name" text,
    "city" text,
    "country" text,
    "alias" text,
    "regions" text,
    "latitude" numeric,
    "longitude" numeric,
    "province" text,
    "timezone" text,
    "unlocs" text,
    "code" text,
    "id" int8 NOT NULL DEFAULT nextval('ports_id_seq'::regclass),
    "created_at" timestamp NOT NULL UNIQUE,
    "updated_at" timestamp NULL DEFAULT NULL,
    "deleted_at" timestamp NULL DEFAULT NULL,
    PRIMARY KEY ("id")
);
