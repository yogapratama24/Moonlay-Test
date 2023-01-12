-- -------------------------------------------------------------
-- TablePlus 5.1.2(472)
--
-- https://tableplus.com/
--
-- Database: moonlay_test
-- Generation Time: 2023-01-12 23:24:46.1970
-- -------------------------------------------------------------


-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS lists_id_seq;

-- Table Definition
CREATE TABLE "public"."lists" (
    "id" int8 NOT NULL DEFAULT nextval('lists_id_seq'::regclass),
    "title" varchar(100),
    "description" text,
    "file" varchar(300),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS sub_lists_id_seq;

-- Table Definition
CREATE TABLE "public"."sub_lists" (
    "id" int8 NOT NULL DEFAULT nextval('sub_lists_id_seq'::regclass),
    "title" varchar(100),
    "description" text,
    "file" varchar(300),
    "list_id" int8,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

INSERT INTO "public"."lists" ("id", "title", "description", "file", "created_at", "updated_at") VALUES
(1, 'list 1', 'list 1', 'http://127.0.0.1:8000/api/v1/public/upload/list/51b8ef8e27aee7d0a3d27c63919dd9c9(1).jpg', '2023-01-11 10:26:18.747223+07', '2023-01-11 10:26:18.747223+07'),
(2, 'list 2', 'list 2', 'http://127.0.0.1:8000/api/v1/public/upload/list/cobaqr.png', '2023-01-11 10:26:42.769823+07', '2023-01-11 10:26:42.769823+07');

INSERT INTO "public"."sub_lists" ("id", "title", "description", "file", "list_id", "created_at", "updated_at") VALUES
(1, 'sub list 1', 'sub list 1', 'http://127.0.0.1:8000/api/v1/public/upload/sublist/51b8ef8e27aee7d0a3d27c63919dd9c9(1).jpg', 1, '2023-01-11 10:27:17.154586+07', '2023-01-11 10:27:17.154586+07'),
(2, 'sub list 2', 'sub list 2', 'http://127.0.0.1:8000/api/v1/public/upload/sublist/51b8ef8e27aee7d0a3d27c63919dd9c9(1).jpg', 1, '2023-01-11 10:27:30.338474+07', '2023-01-11 10:27:30.338474+07');

ALTER TABLE "public"."sub_lists" ADD FOREIGN KEY ("list_id") REFERENCES "public"."lists"("id") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "public"."sub_lists" ADD FOREIGN KEY ("list_id") REFERENCES "public"."lists"("id") ON DELETE CASCADE ON UPDATE CASCADE;
