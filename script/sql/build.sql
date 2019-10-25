/*
 Navicat PostgreSQL Data Transfer

 Source Server         : ntci-ntci
 Source Server Type    : PostgreSQL
 Source Server Version : 90602
 Source Host           : 192.168.2.108:5432
 Source Catalog        : ntci
 Source Schema         : ntci

 Target Server Type    : PostgreSQL
 Target Server Version : 90602
 File Encoding         : 65001

 Date: 25/10/2019 10:39:01
*/


-- ----------------------------
-- Table structure for build
-- ----------------------------
DROP TABLE IF EXISTS "ntci"."build";
CREATE TABLE "ntci"."build" (
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "id" int4 NOT NULL,
  "branch" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "git" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "timestamp" timestamp(0) NOT NULL,
  "status" int4 NOT NULL DEFAULT 0,
  "owner" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "sha" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "message" varchar(255) COLLATE "pg_catalog"."default",
  "language" varchar(255) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "langver" varchar(255) COLLATE "pg_catalog"."default" DEFAULT ''::character varying
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
