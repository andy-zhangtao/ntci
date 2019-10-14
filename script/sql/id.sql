-- ----------------------------
-- Table structure for id
-- ----------------------------
DROP TABLE IF EXISTS "ntci"."id";
CREATE TABLE "ntci"."id" (
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "id" int4 NOT NULL
)
;
ALTER TABLE "ntci"."id" OWNER TO "ntci";
COMMENT ON COLUMN "ntci"."id"."name" IS 'The build project name without namespace';
COMMENT ON COLUMN "ntci"."id"."id" IS 'The next build id';

-- ----------------------------
-- Primary Key structure for table id
-- ----------------------------
ALTER TABLE "ntci"."id" ADD CONSTRAINT "id_pkey" PRIMARY KEY ("name");
