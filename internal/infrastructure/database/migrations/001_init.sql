-- Initial migration
-- migrations/001_init.sql

-- Part 2: Administrative Geography
CREATE TABLE "divisions" (
  "id" bigserial PRIMARY KEY,
  "name_en" varchar(100) NOT NULL UNIQUE,
  "name_bn" varchar(100)
);

CREATE TABLE "districts" (
  "id" bigserial PRIMARY KEY,
  "division_id" bigint NOT NULL REFERENCES "divisions"("id") ON DELETE CASCADE,
  "name_en" varchar(100) NOT NULL UNIQUE,
  "name_bn" varchar(100)
);

CREATE TABLE "upazilas" (
  "id" bigserial PRIMARY KEY,
  "district_id" bigint NOT NULL REFERENCES "districts"("id") ON DELETE CASCADE,
  "name_en" varchar(100) NOT NULL,
  "name_bn" varchar(100),
  -- Add a unique constraint on (district_id, name_en) to prevent duplicate upazilas in a district
  UNIQUE("district_id", "name_en")
);

-- Part 3: Education System
CREATE TABLE "education_streams" (
  "id" bigserial PRIMARY KEY,
  "stream_name_en" varchar(100) NOT NULL UNIQUE -- e.g., 'General', 'Madrasah', 'Technical'
);

CREATE TABLE "education_levels" (
  "id" bigserial PRIMARY KEY,
  "stream_id" bigint NOT NULL REFERENCES "education_streams"("id") ON DELETE CASCADE,
  "level_name_en" varchar(100) NOT NULL, -- e.g., 'Secondary', 'Higher Secondary', 'Tertiary'
  UNIQUE("stream_id", "level_name_en")
);

CREATE TABLE "qualifications" (
  "id" bigserial PRIMARY KEY,
  "level_id" bigint NOT NULL REFERENCES "education_levels"("id") ON DELETE CASCADE,
  "qualification_name_en" varchar(100) NOT NULL UNIQUE, -- e.g., 'SSC', 'HSC', 'Dakhil', 'Diploma in Engineering'
  "common_name" varchar(100), -- e.g., 'Polytechnic'
  "equivalent_to" bigint REFERENCES "qualifications"("id") ON DELETE SET NULL -- Self-referencing to link equivalents (e.g., Dakhil -> SSC)
);

-- Part 4: Other Static Data for Biodata
CREATE TABLE "religions" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(50) UNIQUE NOT NULL
);

CREATE TABLE "professions" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(100) UNIQUE NOT NULL
);

CREATE TABLE "marital_statuses" (
  "id" bigserial PRIMARY KEY,
  "status" varchar(50) UNIQUE NOT NULL
);

CREATE TABLE "skin_tones" (
  "id" bigserial PRIMARY KEY,
  "tone" varchar(50) UNIQUE NOT NULL
);

CREATE TABLE "blood_groups" (
  "id" bigserial PRIMARY KEY,
  "group_name" varchar(5) UNIQUE NOT NULL -- e.g., 'A+', 'O-'
);

CREATE TABLE "countries" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(100) UNIQUE NOT NULL,
  "iso2_code" varchar(2) UNIQUE NOT NULL
);

-- This is the main table for user data
CREATE TABLE "biodata" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL UNIQUE, -- Link to your user auth service
  "full_name" varchar(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  
  -- Relational links to static data
  "current_district_id" bigint REFERENCES "districts"("id"),
  "home_district_id" bigint REFERENCES "districts"("id"),
  "education_qualification_id" bigint REFERENCES "qualifications"("id"),
  "religion_id" bigint REFERENCES "religions"("id"),
  "profession_id" bigint REFERENCES "professions"("id"),
  "marital_status_id" bigint REFERENCES "marital_statuses"("id"),
  "skin_tone_id" bigint REFERENCES "skin_tones"("id"),
  "blood_group_id" bigint REFERENCES "blood_groups"("id"),
  "current_country_id" bigint REFERENCES "countries"("id"),
  
  -- Other core biodata fields
  "height_cm" smallint,
  "weight_kg" smallint,
  "date_of_birth" date,
  "about_me" text,
  "about_family" text,
  "education_details" text, -- For specific institution names, etc.
  "is_disabled" boolean DEFAULT false,
  "disability_details" text
);