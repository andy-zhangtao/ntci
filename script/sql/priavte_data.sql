-- ----------------------------
-- Table structure for priavte_data
-- ----------------------------
DROP TABLE IF EXISTS "ntci"."priavte_data";
CREATE TABLE "ntci"."priavte_data" (
  "owner" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "key" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "value" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "ntci"."priavte_data" OWNER TO "ntci";

-- ----------------------------
-- Indexes structure for table priavte_data
-- ----------------------------
CREATE INDEX "useAndProject" ON "ntci"."priavte_data" USING btree (
  "owner" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table priavte_data
-- ----------------------------
ALTER TABLE "ntci"."priavte_data" ADD CONSTRAINT "priavte_data_pkey" PRIMARY KEY ("owner", "name", "key");
