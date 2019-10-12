-- ----------------------------
-- Table structure for build
-- ----------------------------
DROP TABLE IF EXISTS "ntci"."build";
CREATE TABLE "ntci"."build" (
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "id" int4 NOT NULL,
  "branch" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "git" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "timestamp" date NOT NULL,
  "status" int4 NOT NULL DEFAULT 0
)
;
ALTER TABLE "ntci"."build" OWNER TO "ntci";

-- ----------------------------
-- Primary Key structure for table build
-- ----------------------------
ALTER TABLE "ntci"."build" ADD CONSTRAINT "build_pkey" PRIMARY KEY ("name", "id");
