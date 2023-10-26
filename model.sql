CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL
);

CREATE TABLE "hospitals" (
  "id" integer PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "mutuelles" (
  "id" integer PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "medical_acts" (
  "id" integer PRIMARY KEY,
  "user_id" integer,
  "hopital_id" integer,
  "mutuelle_id" integer,
  "metadata_1" varchar,
  "metadata_2" varchar,
  "montant_total" float DEFAULT 0,
  "prise_en_charge_hopital" float DEFAULT 0,
  "prise_en_charge_mutuelle" float DEFAULT 0,
  "prise_en_charge_patient" float DEFAULT 0,
  "confirmation_mutuelle" boolean DEFAULT false,
  "confirmation_rdv" boolean DEFAULT false,
  "pourcentage_prise_en_charge" integer,
  "commentaire" text,
  "date_creation" timestamp DEFAULT (now()),
  "date_prevue" timestamp,
  "date_venue" timestamp
);

CREATE TABLE "results" (
  "id" integer PRIMARY KEY,
  "medical_act_id" integer,
  "file_name" varchar NOT NULL,
  "file_data" bytea NOT NULL
);

ALTER TABLE "medical_acts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "medical_acts" ADD FOREIGN KEY ("hopital_id") REFERENCES "hospitals" ("id");

ALTER TABLE "medical_acts" ADD FOREIGN KEY ("mutuelle_id") REFERENCES "mutuelles" ("id");

ALTER TABLE "medical_acts" ADD FOREIGN KEY ("id") REFERENCES "results" ("medical_act_id");
