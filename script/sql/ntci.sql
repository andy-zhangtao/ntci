-- ----------------------------
-- Table structure for ntci
-- ----------------------------
DROP TABLE IF EXISTS "ntci"."ntci";
CREATE TABLE "ntci"."ntci" (
  "owner" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "ntci" varchar(2048) COLLATE "pg_catalog"."default" NOT NULL,
  "branch" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "ntci"."ntci" OWNER TO "ntci";

-- ----------------------------
-- Primary Key structure for table ntci
-- ----------------------------
ALTER TABLE "ntci"."ntci" ADD CONSTRAINT "ntci_pkey" PRIMARY KEY ("owner", "name", "ntci", "branch");
