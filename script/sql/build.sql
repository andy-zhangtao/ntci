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
  "status" int4 NOT NULL DEFAULT 0,
  "owner" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "sha" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "message" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "ntci"."build" OWNER TO "ntci";

-- ----------------------------
-- Indexes structure for table build
-- ----------------------------
CREATE INDEX "userAndName" ON "ntci"."build" USING btree (
  "owner" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table build
-- ----------------------------
ALTER TABLE "ntci"."build" ADD CONSTRAINT "build_pkey" PRIMARY KEY ("name", "id", "owner");
