/*
 Navicat Premium Dump SQL

 Source Server         : swantara-gate
 Source Server Type    : SQLite
 Source Server Version : 3045000 (3.45.0)
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3045000 (3.45.0)
 File Encoding         : 65001

 Date: 28/05/2026 17:50:28
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for access_logs
-- ----------------------------
DROP TABLE IF EXISTS "access_logs";
CREATE TABLE "access_logs" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "timestamp" DATETIME DEFAULT CURRENT_TIMESTAMP,
  "method" TEXT,
  "host" TEXT,
  "path" TEXT,
  "status_code" INTEGER,
  "duration_ms" INTEGER,
  "client_ip" TEXT,
  "user_agent" TEXT,
  "bytes_sent" INTEGER,
  "vhost_id" INTEGER,
  "vdir_id" INTEGER,
  "upstream_host" TEXT,
  "upstream_port" INTEGER
);

-- ----------------------------
-- Records of access_logs
-- ----------------------------

-- ----------------------------
-- Table structure for acme_accounts
-- ----------------------------
DROP TABLE IF EXISTS "acme_accounts";
CREATE TABLE "acme_accounts" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "email" TEXT NOT NULL,
  "provider_url" TEXT NOT NULL,
  "account_key_path" TEXT NOT NULL,
  "is_default" INTEGER DEFAULT 0,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- ----------------------------
-- Records of acme_accounts
-- ----------------------------

-- ----------------------------
-- Table structure for api_consumers
-- ----------------------------
DROP TABLE IF EXISTS "api_consumers";
CREATE TABLE "api_consumers" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "consumer_name" TEXT NOT NULL,
  "description" TEXT,
  "contact_email" TEXT,
  "is_active" INTEGER DEFAULT 1,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  "updated_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  UNIQUE ("consumer_name" ASC)
);

-- ----------------------------
-- Records of api_consumers
-- ----------------------------

-- ----------------------------
-- Table structure for api_keys
-- ----------------------------
DROP TABLE IF EXISTS "api_keys";
CREATE TABLE "api_keys" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "consumer_id" INTEGER NOT NULL,
  "api_key" TEXT NOT NULL,
  "description" TEXT,
  "expired_at" DATETIME,
  "rate_limit_override" INTEGER,
  "is_active" INTEGER DEFAULT 1,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("consumer_id") REFERENCES "api_consumers" ("id") ON DELETE CASCADE ON UPDATE NO ACTION,
  UNIQUE ("api_key" ASC)
);

-- ----------------------------
-- Records of api_keys
-- ----------------------------

-- ----------------------------
-- Table structure for certificate_domains
-- ----------------------------
DROP TABLE IF EXISTS "certificate_domains";
CREATE TABLE "certificate_domains" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "ssl_certificate_id" INTEGER NOT NULL,
  "domain_name" TEXT NOT NULL,
  "is_wildcard" INTEGER DEFAULT 0,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("ssl_certificate_id") REFERENCES "ssl_certificates" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of certificate_domains
-- ----------------------------

-- ----------------------------
-- Table structure for circuit_breakers
-- ----------------------------
DROP TABLE IF EXISTS "circuit_breakers";
CREATE TABLE "circuit_breakers" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "virtual_directory_id" INTEGER NOT NULL,
  "enabled" INTEGER DEFAULT 1,
  "failure_threshold" INTEGER DEFAULT 5,
  "recovery_timeout_seconds" INTEGER DEFAULT 30,
  "half_open_max_requests" INTEGER DEFAULT 3,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("virtual_directory_id") REFERENCES "virtual_directories" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of circuit_breakers
-- ----------------------------

-- ----------------------------
-- Table structure for config_versions
-- ----------------------------
DROP TABLE IF EXISTS "config_versions";
CREATE TABLE "config_versions" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "config_name" TEXT NOT NULL,
  "version_number" INTEGER NOT NULL,
  "changed_by" TEXT,
  "notes" TEXT,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- ----------------------------
-- Records of config_versions
-- ----------------------------

-- ----------------------------
-- Table structure for consumer_credentials
-- ----------------------------
DROP TABLE IF EXISTS "consumer_credentials";
CREATE TABLE "consumer_credentials" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "consumer_id" INTEGER NOT NULL,
  "auth_type" TEXT NOT NULL,
  "username" TEXT,
  "password_hash" TEXT,
  "api_key" TEXT,
  "jwt_secret" TEXT,
  "expired_at" DATETIME,
  "is_active" INTEGER DEFAULT 1,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("consumer_id") REFERENCES "api_consumers" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of consumer_credentials
-- ----------------------------

-- ----------------------------
-- Table structure for cors_configs
-- ----------------------------
DROP TABLE IF EXISTS "cors_configs";
CREATE TABLE "cors_configs" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "virtual_directory_id" INTEGER NOT NULL,
  "allowed_origins" TEXT DEFAULT '*',
  "allowed_methods" TEXT DEFAULT 'GET,POST,PUT,DELETE,OPTIONS',
  "allowed_headers" TEXT DEFAULT '*',
  "exposed_headers" TEXT,
  "allow_credentials" INTEGER DEFAULT 0,
  "max_age_seconds" INTEGER DEFAULT 3600,
  "is_active" INTEGER DEFAULT 1,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("virtual_directory_id") REFERENCES "virtual_directories" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of cors_configs
-- ----------------------------

-- ----------------------------
-- Table structure for external_auth
-- ----------------------------
DROP TABLE IF EXISTS "external_auth";
CREATE TABLE "external_auth" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "virtual_directory_id" INTEGER NOT NULL,
  "auth_url" TEXT NOT NULL,
  "http_method" TEXT DEFAULT 'POST',
  "request_timeout_seconds" INTEGER DEFAULT 5,
  "send_headers" INTEGER DEFAULT 1,
  "send_body" INTEGER DEFAULT 0,
  "success_key" TEXT DEFAULT 'status',
  "success_value" TEXT DEFAULT 'true',
  "message_key" TEXT DEFAULT 'message',
  "token_key" TEXT,
  "is_active" INTEGER DEFAULT 1,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("virtual_directory_id") REFERENCES "virtual_directories" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of external_auth
-- ----------------------------

-- ----------------------------
-- Table structure for hosts
-- ----------------------------
DROP TABLE IF EXISTS "hosts";
CREATE TABLE "hosts" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "host_name" TEXT NOT NULL,
  "description" TEXT,
  "is_active" INTEGER DEFAULT 1,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  "updated_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  UNIQUE ("host_name" ASC)
);

-- ----------------------------
-- Records of hosts
-- ----------------------------
INSERT INTO "hosts" VALUES (1, 'api.example.local', 'Main API Host', 1, '2026-05-28 10:11:47', '2026-05-28 10:11:47');
INSERT INTO "hosts" VALUES (2, 'api-dua.example.local', 'Main API Host', 1, '2026-05-28 10:11:55', '2026-05-28 10:11:55');

-- ----------------------------
-- Table structure for ip_blacklists
-- ----------------------------
DROP TABLE IF EXISTS "ip_blacklists";
CREATE TABLE "ip_blacklists" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "virtual_directory_id" INTEGER NOT NULL,
  "ip_address" TEXT NOT NULL,
  "reason" TEXT,
  "expired_at" DATETIME,
  "is_active" INTEGER DEFAULT 1,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("virtual_directory_id") REFERENCES "virtual_directories" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of ip_blacklists
-- ----------------------------

-- ----------------------------
-- Table structure for ip_whitelists
-- ----------------------------
DROP TABLE IF EXISTS "ip_whitelists";
CREATE TABLE "ip_whitelists" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "virtual_directory_id" INTEGER NOT NULL,
  "ip_address" TEXT NOT NULL,
  "description" TEXT,
  "is_active" INTEGER DEFAULT 1,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("virtual_directory_id") REFERENCES "virtual_directories" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of ip_whitelists
-- ----------------------------

-- ----------------------------
-- Table structure for jwt_configs
-- ----------------------------
DROP TABLE IF EXISTS "jwt_configs";
CREATE TABLE "jwt_configs" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "virtual_directory_id" INTEGER NOT NULL,
  "algorithm" TEXT DEFAULT 'HS256',
  "jwt_secret" TEXT NOT NULL,
  "issuer" TEXT,
  "audience" TEXT,
  "expired_in_seconds" INTEGER DEFAULT 3600,
  "clock_skew_seconds" INTEGER DEFAULT 30,
  "require_exp" INTEGER DEFAULT 1,
  "require_iat" INTEGER DEFAULT 1,
  "is_active" INTEGER DEFAULT 1,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("virtual_directory_id") REFERENCES "virtual_directories" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of jwt_configs
-- ----------------------------

-- ----------------------------
-- Table structure for jwt_tokens
-- ----------------------------
DROP TABLE IF EXISTS "jwt_tokens";
CREATE TABLE "jwt_tokens" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "consumer_id" INTEGER,
  "token" TEXT NOT NULL,
  "issued_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  "expired_at" DATETIME,
  "is_revoked" INTEGER DEFAULT 0,
  "ip_address" TEXT,
  "user_agent" TEXT,
  FOREIGN KEY ("consumer_id") REFERENCES "api_consumers" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of jwt_tokens
-- ----------------------------

-- ----------------------------
-- Table structure for maintenance_windows
-- ----------------------------
DROP TABLE IF EXISTS "maintenance_windows";
CREATE TABLE "maintenance_windows" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "virtual_host_id" INTEGER,
  "title" TEXT,
  "start_at" DATETIME,
  "end_at" DATETIME,
  "maintenance_response_code" INTEGER DEFAULT 503,
  "maintenance_message" TEXT,
  "is_active" INTEGER DEFAULT 1,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("virtual_host_id") REFERENCES "virtual_hosts" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of maintenance_windows
-- ----------------------------

-- ----------------------------
-- Table structure for query_rewrites
-- ----------------------------
DROP TABLE IF EXISTS "query_rewrites";
CREATE TABLE "query_rewrites" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "virtual_directory_id" INTEGER NOT NULL,
  "param_name" TEXT NOT NULL,
  "param_value" TEXT,
  "operation" TEXT DEFAULT 'set',
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("virtual_directory_id") REFERENCES "virtual_directories" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of query_rewrites
-- ----------------------------

-- ----------------------------
-- Table structure for rate_limits
-- ----------------------------
DROP TABLE IF EXISTS "rate_limits";
CREATE TABLE "rate_limits" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "virtual_directory_id" INTEGER NOT NULL,
  "limit_by" TEXT DEFAULT 'ip',
  "requests_per_minute" INTEGER DEFAULT 60,
  "burst" INTEGER DEFAULT 10,
  "block_duration_seconds" INTEGER DEFAULT 60,
  "is_active" INTEGER DEFAULT 1,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("virtual_directory_id") REFERENCES "virtual_directories" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of rate_limits
-- ----------------------------

-- ----------------------------
-- Table structure for refresh_tokens
-- ----------------------------
DROP TABLE IF EXISTS "refresh_tokens";
CREATE TABLE "refresh_tokens" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "user_id" INTEGER NOT NULL,
  "token" TEXT NOT NULL,
  "expires_at" DATETIME NOT NULL,
  "is_revoked" INTEGER DEFAULT 0,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION,
  UNIQUE ("token" ASC)
);

-- ----------------------------
-- Records of refresh_tokens
-- ----------------------------
INSERT INTO "refresh_tokens" VALUES (1, 2, 'fb5fdf757ecf8a613a7c7272898f1400103b8202cebfe1427acd808674321c0f', '2026-06-04 14:16:12.8755799 +0700 +07 m=+604862.494775301', 1, '2026-05-28 07:16:12');
INSERT INTO "refresh_tokens" VALUES (2, 2, '87aeb1edd402bdcf8832e041337b7fb26b5572f5f85d3d5a9aa1b1ddcd13c4f6', '2026-06-04 14:16:23.5708094 +0700 +07 m=+604873.190004801', 1, '2026-05-28 07:16:23');
INSERT INTO "refresh_tokens" VALUES (3, 2, '85543eab90329ff29907cfa29304654e3c198f9d99d00d75bdb40d5ba1a6304d', '2026-06-04 14:17:22.9891554 +0700 +07 m=+604932.608350801', 1, '2026-05-28 07:17:22');
INSERT INTO "refresh_tokens" VALUES (4, 2, '9ed8057462fed4c272a30e653e6bad5107814a6a45f1f07da973b666585b532d', '2026-06-04 14:17:25.8200707 +0700 +07 m=+604935.439266101', 1, '2026-05-28 07:17:25');
INSERT INTO "refresh_tokens" VALUES (5, 2, '8150a1326dd817107616e7b807adffcec5469f134a57699871831070f7d172a7', '2026-06-04 14:17:27.4387923 +0700 +07 m=+604937.057987701', 1, '2026-05-28 07:17:27');
INSERT INTO "refresh_tokens" VALUES (6, 2, 'eac3c84daf518803aee60f19ea36fa27676089686f444e86cde39235b06679d8', '2026-06-04 14:17:28.4457341 +0700 +07 m=+604938.064929501', 1, '2026-05-28 07:17:28');
INSERT INTO "refresh_tokens" VALUES (7, 2, 'b47389f6c7d3c7430b0f01e5971f2164f87d6c32e886c0d88d36e00ec6ff0e31', '2026-06-04 14:17:29.4100503 +0700 +07 m=+604939.029245701', 1, '2026-05-28 07:17:29');
INSERT INTO "refresh_tokens" VALUES (8, 2, '1bb9baac6b8c7a1724a44a0a51189d422851928ded2d075b57135737d543474c', '2026-06-04 14:17:30.3330198 +0700 +07 m=+604939.952215201', 1, '2026-05-28 07:17:30');
INSERT INTO "refresh_tokens" VALUES (9, 2, '1f4668b84c2bcb2fa1c89f835b4ed2e4f0e593a8926a4eb4afa7b10033c6245a', '2026-06-04 14:32:17.3450272 +0700 +07 m=+604835.016655501', 0, '2026-05-28 07:32:17');
INSERT INTO "refresh_tokens" VALUES (10, 2, 'b0bb40c18cca63784621994a79154ec987903226fd79d65da9dd87f2ac7df2bc', '2026-06-04 14:32:38.4942696 +0700 +07 m=+604856.165897901', 1, '2026-05-28 07:32:38');
INSERT INTO "refresh_tokens" VALUES (11, 2, '9055bd0c40a608bde61325db0f6424a6ec9ea04afcac5599a7cf2e22243788cd', '2026-06-04 14:40:54.8555722 +0700 +07 m=+604999.770039901', 1, '2026-05-28 07:40:54');
INSERT INTO "refresh_tokens" VALUES (12, 2, '122bb994ddda41b7c3447ee56d5662a5f81e468f18a9e33e5cee9a17530ce95e', '2026-06-04 14:41:10.2787611 +0700 +07 m=+605015.193228801', 1, '2026-05-28 07:41:10');
INSERT INTO "refresh_tokens" VALUES (13, 2, 'a7fbdf36f3342ef7723c0032e96ad0ddc3dfa6548e24a74f46aba29d4e3b4da0', '2026-06-04 14:41:11.470979 +0700 +07 m=+605016.385446701', 1, '2026-05-28 07:41:11');
INSERT INTO "refresh_tokens" VALUES (14, 2, '6163f3fc8c04f87ee90fd634ad5424a104e07f9260e302eec000e565acc01007', '2026-06-04 14:41:32.0628135 +0700 +07 m=+605036.977281201', 1, '2026-05-28 07:41:32');
INSERT INTO "refresh_tokens" VALUES (15, 2, '9638ebc58f4aa97e1980903af2e8f634fce8629cc1d5289dd2e9aec778086178', '2026-06-04 14:41:32.8535485 +0700 +07 m=+605037.768016201', 1, '2026-05-28 07:41:32');
INSERT INTO "refresh_tokens" VALUES (16, 2, '124b0cfb6659ebf325e96d77188f052608f0356d5a3a3f5de1f7f323c2e5260c', '2026-06-04 16:30:41.8576588 +0700 +07 m=+604856.144131301', 0, '2026-05-28 09:30:41');
INSERT INTO "refresh_tokens" VALUES (17, 2, '256e926171bccd4cd1fe0a43dae47499d16fe43d0c5d3e0086c0b4ecac928beb', '2026-06-04 17:11:45.748833 +0700 +07 m=+606973.514789301', 0, '2026-05-28 10:11:45');
INSERT INTO "refresh_tokens" VALUES (18, 2, 'd6cce7ba738f75b311253490861b4566a82e4bd84fc515b7fb587755fe5e478a', '2026-06-04 17:42:24.1007189 +0700 +07 m=+604870.531067001', 0, '2026-05-28 10:42:24');

-- ----------------------------
-- Table structure for request_header_rules
-- ----------------------------
DROP TABLE IF EXISTS "request_header_rules";
CREATE TABLE "request_header_rules" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "virtual_directory_id" INTEGER NOT NULL,
  "header_name" TEXT NOT NULL,
  "operation" TEXT NOT NULL,
  "value_source" TEXT DEFAULT 'static',
  "header_value" TEXT,
  "source_header" TEXT,
  "variable_name" TEXT,
  "execution_order" INTEGER DEFAULT 1,
  "is_active" INTEGER DEFAULT 1,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("virtual_directory_id") REFERENCES "virtual_directories" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of request_header_rules
-- ----------------------------

-- ----------------------------
-- Table structure for response_header_rules
-- ----------------------------
DROP TABLE IF EXISTS "response_header_rules";
CREATE TABLE "response_header_rules" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "virtual_directory_id" INTEGER NOT NULL,
  "header_name" TEXT NOT NULL,
  "operation" TEXT NOT NULL,
  "header_value" TEXT,
  "execution_order" INTEGER DEFAULT 1,
  "is_active" INTEGER DEFAULT 1,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("virtual_directory_id") REFERENCES "virtual_directories" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of response_header_rules
-- ----------------------------

-- ----------------------------
-- Table structure for route_consumer_access
-- ----------------------------
DROP TABLE IF EXISTS "route_consumer_access";
CREATE TABLE "route_consumer_access" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "virtual_directory_id" INTEGER NOT NULL,
  "consumer_id" INTEGER NOT NULL,
  "is_active" INTEGER DEFAULT 1,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("virtual_directory_id") REFERENCES "virtual_directories" ("id") ON DELETE CASCADE ON UPDATE NO ACTION,
  FOREIGN KEY ("consumer_id") REFERENCES "api_consumers" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of route_consumer_access
-- ----------------------------

-- ----------------------------
-- Table structure for service_discovery
-- ----------------------------
DROP TABLE IF EXISTS "service_discovery";
CREATE TABLE "service_discovery" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "virtual_host_id" INTEGER NOT NULL,
  "provider" TEXT NOT NULL,
  "endpoint_url" TEXT NOT NULL,
  "refresh_interval_seconds" INTEGER DEFAULT 30,
  "is_active" INTEGER DEFAULT 1,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("virtual_host_id") REFERENCES "virtual_hosts" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of service_discovery
-- ----------------------------

-- ----------------------------
-- Table structure for sqlite_sequence
-- ----------------------------
DROP TABLE IF EXISTS "sqlite_sequence";
CREATE TABLE "sqlite_sequence" (
  "name",
  "seq"
);

-- ----------------------------
-- Records of sqlite_sequence
-- ----------------------------
INSERT INTO "sqlite_sequence" VALUES ('users', 3);
INSERT INTO "sqlite_sequence" VALUES ('refresh_tokens', 18);
INSERT INTO "sqlite_sequence" VALUES ('hosts', 2);
INSERT INTO "sqlite_sequence" VALUES ('virtual_hosts', 2);
INSERT INTO "sqlite_sequence" VALUES ('upstream_servers', 4);
INSERT INTO "sqlite_sequence" VALUES ('virtual_directories', 2);

-- ----------------------------
-- Table structure for ssl_certificate_bindings
-- ----------------------------
DROP TABLE IF EXISTS "ssl_certificate_bindings";
CREATE TABLE "ssl_certificate_bindings" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "ssl_certificate_id" INTEGER NOT NULL,
  "binding_type" TEXT NOT NULL,
  "host_id" INTEGER,
  "virtual_host_id" INTEGER,
  "is_default" INTEGER DEFAULT 0,
  "priority" INTEGER DEFAULT 1,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("ssl_certificate_id") REFERENCES "ssl_certificates" ("id") ON DELETE CASCADE ON UPDATE NO ACTION,
  FOREIGN KEY ("host_id") REFERENCES "hosts" ("id") ON DELETE CASCADE ON UPDATE NO ACTION,
  FOREIGN KEY ("virtual_host_id") REFERENCES "virtual_hosts" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of ssl_certificate_bindings
-- ----------------------------

-- ----------------------------
-- Table structure for ssl_certificates
-- ----------------------------
DROP TABLE IF EXISTS "ssl_certificates";
CREATE TABLE "ssl_certificates" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "acme_account_id" INTEGER,
  "provider" TEXT DEFAULT 'lets_encrypt',
  "challenge_type" TEXT DEFAULT 'http01',
  "certificate_path" TEXT NOT NULL,
  "private_key_path" TEXT NOT NULL,
  "chain_path" TEXT,
  "auto_renew" INTEGER DEFAULT 1,
  "renew_before_days" INTEGER DEFAULT 30,
  "last_renew_at" DATETIME,
  "expired_at" DATETIME,
  "renew_status" TEXT DEFAULT 'pending',
  "last_error" TEXT,
  "is_active" INTEGER DEFAULT 1,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  "updated_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("acme_account_id") REFERENCES "acme_accounts" ("id") ON DELETE SET NULL ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of ssl_certificates
-- ----------------------------

-- ----------------------------
-- Table structure for tls_options
-- ----------------------------
DROP TABLE IF EXISTS "tls_options";
CREATE TABLE "tls_options" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "binding_type" TEXT NOT NULL,
  "host_id" INTEGER,
  "virtual_host_id" INTEGER,
  "min_tls_version" TEXT DEFAULT '1.2',
  "http2_enabled" INTEGER DEFAULT 1,
  "hsts_enabled" INTEGER DEFAULT 1,
  "hsts_max_age" INTEGER DEFAULT 31536000,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("host_id") REFERENCES "hosts" ("id") ON DELETE CASCADE ON UPDATE NO ACTION,
  FOREIGN KEY ("virtual_host_id") REFERENCES "virtual_hosts" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of tls_options
-- ----------------------------

-- ----------------------------
-- Table structure for upstream_servers
-- ----------------------------
DROP TABLE IF EXISTS "upstream_servers";
CREATE TABLE "upstream_servers" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "virtual_host_id" INTEGER NOT NULL,
  "target_host" TEXT NOT NULL,
  "target_port" INTEGER NOT NULL,
  "protocol" TEXT DEFAULT 'http',
  "priority" INTEGER DEFAULT 1,
  "weight" INTEGER DEFAULT 1,
  "is_backup" INTEGER DEFAULT 0,
  "is_active" INTEGER DEFAULT 1,
  "health_check_enabled" INTEGER DEFAULT 1,
  "health_check_path" TEXT DEFAULT '/health',
  "health_check_interval_seconds" INTEGER DEFAULT 10,
  "health_check_timeout_seconds" INTEGER DEFAULT 3,
  "max_fails" INTEGER DEFAULT 3,
  "fail_timeout_seconds" INTEGER DEFAULT 30,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  "updated_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("virtual_host_id") REFERENCES "virtual_hosts" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of upstream_servers
-- ----------------------------
INSERT INTO "upstream_servers" VALUES (1, 1, 'localhost', 7071, 'http', 1, 1, 0, 1, 0, '/health', 10, 3, 3, 30, '2026-05-28 10:36:26', '2026-05-28 10:36:26');
INSERT INTO "upstream_servers" VALUES (2, 1, 'localhost', 7072, 'http', 1, 1, 0, 1, 0, '/health', 10, 3, 3, 30, '2026-05-28 10:36:34', '2026-05-28 10:36:34');
INSERT INTO "upstream_servers" VALUES (3, 2, 'localhost', 7073, 'http', 1, 1, 0, 1, 0, '/health', 10, 3, 3, 30, '2026-05-28 10:36:42', '2026-05-28 10:36:42');
INSERT INTO "upstream_servers" VALUES (4, 2, 'localhost', 7074, 'http', 1, 1, 0, 1, 0, '/health', 10, 3, 3, 30, '2026-05-28 10:36:50', '2026-05-28 10:36:50');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "users";
CREATE TABLE "users" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "username" TEXT NOT NULL,
  "password_hash" TEXT NOT NULL,
  "full_name" TEXT,
  "email" TEXT,
  "role" TEXT DEFAULT 'admin',
  "is_active" INTEGER DEFAULT 1,
  "last_login_at" DATETIME,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  "updated_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  UNIQUE ("username" ASC)
);

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO "users" VALUES (1, 'testcli', '$2a$10$qayue2aZ/3iwuOtnS3ZQ6uPwOv4oZ0yxNGH8Ng289BxKaq4gVhMmi', '', '', 'admin', 1, NULL, '2026-05-28 14:10:50.0340632 +0700 +07 m=+0.048142601', '2026-05-28 14:10:54.8167647 +0700 +07 m=+0.047703201');
INSERT INTO "users" VALUES (2, 'admin', '$2a$10$rjZ0raqP973SCG5bWefneuJb6sbkXfSIiyxbiqAsNX0QMbHa1VACq', 'Super Admin', 'admin@example.com', 'super_admin', 1, '2026-05-28 17:42:24.1053844 +0700 +07 m=+70.535732501', '2026-05-28 14:12:02.8123223 +0700 +07 m=+0.046588701', '2026-05-28 17:42:24.1053844 +0700 +07 m=+70.535732501');

-- ----------------------------
-- Table structure for virtual_directories
-- ----------------------------
DROP TABLE IF EXISTS "virtual_directories";
CREATE TABLE "virtual_directories" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "virtual_host_id" INTEGER NOT NULL,
  "source_path" TEXT NOT NULL,
  "target_path" TEXT NOT NULL,
  "match_type" TEXT DEFAULT 'prefix',
  "strip_prefix" INTEGER DEFAULT 1,
  "preserve_host_header" INTEGER DEFAULT 0,
  "auth_type" TEXT DEFAULT 'none',
  "is_active" INTEGER DEFAULT 1,
  "proxy_timeout_seconds" INTEGER DEFAULT 30,
  "retry_count" INTEGER DEFAULT 0,
  "retry_delay_ms" INTEGER DEFAULT 100,
  "max_request_size_mb" INTEGER DEFAULT 10,
  "websocket_enabled" INTEGER DEFAULT 0,
  "cache_enabled" INTEGER DEFAULT 0,
  "cache_ttl_seconds" INTEGER DEFAULT 60,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  "updated_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("virtual_host_id") REFERENCES "virtual_hosts" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of virtual_directories
-- ----------------------------
INSERT INTO "virtual_directories" VALUES (1, 1, '/', '/', 'prefix', 1, 0, 'none', 1, 30, 2, 100, 10, 0, 0, 0, '2026-05-28 10:40:17', '2026-05-28 10:40:17');
INSERT INTO "virtual_directories" VALUES (2, 1, '/jamet', '/jamet.php', 'prefix', 1, 0, 'none', 1, 30, 2, 100, 10, 0, 0, 0, '2026-05-28 10:41:05', '2026-05-28 10:41:05');

-- ----------------------------
-- Table structure for virtual_directory_methods
-- ----------------------------
DROP TABLE IF EXISTS "virtual_directory_methods";
CREATE TABLE "virtual_directory_methods" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "virtual_directory_id" INTEGER NOT NULL,
  "http_method" TEXT NOT NULL,
  FOREIGN KEY ("virtual_directory_id") REFERENCES "virtual_directories" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- ----------------------------
-- Records of virtual_directory_methods
-- ----------------------------

-- ----------------------------
-- Table structure for virtual_hosts
-- ----------------------------
DROP TABLE IF EXISTS "virtual_hosts";
CREATE TABLE "virtual_hosts" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "host_id" INTEGER NOT NULL,
  "vhost_name" TEXT NOT NULL,
  "lb_algorithm" TEXT DEFAULT 'round_robin',
  "sticky_session" INTEGER DEFAULT 0,
  "failover_mode" TEXT DEFAULT 'active-active',
  "is_active" INTEGER DEFAULT 1,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  "updated_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("host_id") REFERENCES "hosts" ("id") ON DELETE CASCADE ON UPDATE NO ACTION,
  UNIQUE ("vhost_name" ASC)
);

-- ----------------------------
-- Records of virtual_hosts
-- ----------------------------
INSERT INTO "virtual_hosts" VALUES (1, 1, 'server-1-2.local', 'round_robin', 0, 'active-active', 1, '2026-05-28 10:35:07', '2026-05-28 10:35:07');
INSERT INTO "virtual_hosts" VALUES (2, 1, 'server-3-4.local', 'round_robin', 0, 'active-active', 1, '2026-05-28 10:35:13', '2026-05-28 10:35:13');

-- ----------------------------
-- Auto increment value for hosts
-- ----------------------------
UPDATE "sqlite_sequence" SET seq = 2 WHERE name = 'hosts';

-- ----------------------------
-- Auto increment value for refresh_tokens
-- ----------------------------
UPDATE "sqlite_sequence" SET seq = 18 WHERE name = 'refresh_tokens';

-- ----------------------------
-- Auto increment value for upstream_servers
-- ----------------------------
UPDATE "sqlite_sequence" SET seq = 4 WHERE name = 'upstream_servers';

-- ----------------------------
-- Auto increment value for users
-- ----------------------------
UPDATE "sqlite_sequence" SET seq = 3 WHERE name = 'users';

-- ----------------------------
-- Indexes structure for table users
-- ----------------------------
CREATE INDEX "idx_users_username"
ON "users" (
  "username" ASC
);

-- ----------------------------
-- Auto increment value for virtual_directories
-- ----------------------------
UPDATE "sqlite_sequence" SET seq = 2 WHERE name = 'virtual_directories';

-- ----------------------------
-- Auto increment value for virtual_hosts
-- ----------------------------
UPDATE "sqlite_sequence" SET seq = 2 WHERE name = 'virtual_hosts';

PRAGMA foreign_keys = true;
