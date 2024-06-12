CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE IF NOT EXISTS "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "full_name" VARCHAR(255) NOT NULL,
    "nick_name" VARCHAR,
    "photo" TEXT,
    "birthday" DATE,
    "location"  GEOMETRY(POINT, 4326) NOT NULL, 
    "created_at" TIMESTAMPTZ DEFAULT NOW(), -- TIMESTAMP ni o'zidan ham foydalansam bo'lar edi. Lekin time zonda muammo bo'lmasligi uchun bundan foydalandim.
    "deleted_at" TIMESTAMPTZ,
    "updated_at" TIMESTAMPTZ DEFAULT NOW(),
    "created_by" VARCHAR NOT NULL,
    "updated_by" VARCHAR,
    "deleted_by" VARCHAR
);