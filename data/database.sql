PRAGMA foreign_keys = ON;

-- =========================================================
-- TABLE users
-- Admin panel users
-- =========================================================
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    full_name TEXT,
    email TEXT,

    -- super_admin
    -- admin
    -- operator
    -- viewer
    role TEXT DEFAULT 'admin',

    is_active INTEGER DEFAULT 1,

    last_login_at DATETIME,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_username
ON users(username);

-- =========================================================
-- TABLE api_consumers
-- API consumers / applications
-- =========================================================
CREATE TABLE api_consumers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    consumer_name TEXT NOT NULL UNIQUE,

    description TEXT,

    contact_email TEXT,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- =========================================================
-- TABLE consumer_credentials
-- Consumer credentials
-- =========================================================
CREATE TABLE consumer_credentials (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    consumer_id INTEGER NOT NULL,

    -- basic
    -- api_key
    -- jwt
    auth_type TEXT NOT NULL,

    username TEXT,

    password_hash TEXT,

    api_key TEXT,

    jwt_secret TEXT,

    expired_at DATETIME,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (consumer_id)
        REFERENCES api_consumers(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE hosts
-- Root host grouping
-- =========================================================
CREATE TABLE hosts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    host_name TEXT NOT NULL UNIQUE,

    description TEXT,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- =========================================================
-- TABLE virtual_hosts
-- Virtual domains
-- =========================================================
CREATE TABLE virtual_hosts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    host_id INTEGER NOT NULL,

    vhost_name TEXT NOT NULL UNIQUE,

    -- round_robin
    -- weighted_round_robin
    -- least_conn
    -- ip_hash
    -- random
    -- failover
    lb_algorithm TEXT DEFAULT 'round_robin',

    sticky_session INTEGER DEFAULT 0,

    -- active-active
    -- active-passive
    failover_mode TEXT DEFAULT 'active-active',

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (host_id)
        REFERENCES hosts(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE upstream_servers
-- Backend servers
-- =========================================================
CREATE TABLE upstream_servers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_host_id INTEGER NOT NULL,

    target_host TEXT NOT NULL,

    target_port INTEGER NOT NULL,

    protocol TEXT DEFAULT 'http',

    priority INTEGER DEFAULT 1,

    weight INTEGER DEFAULT 1,

    is_backup INTEGER DEFAULT 0,

    is_active INTEGER DEFAULT 1,

    health_check_enabled INTEGER DEFAULT 1,

    health_check_path TEXT DEFAULT '/health',

    health_check_interval_seconds INTEGER DEFAULT 10,

    health_check_timeout_seconds INTEGER DEFAULT 3,

    max_fails INTEGER DEFAULT 3,

    fail_timeout_seconds INTEGER DEFAULT 30,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_host_id)
        REFERENCES virtual_hosts(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE virtual_directories
-- API routes
-- =========================================================
CREATE TABLE virtual_directories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_host_id INTEGER NOT NULL,

    source_path TEXT NOT NULL,

    target_path TEXT NOT NULL,

    -- exact
    -- prefix
    -- wildcard
    -- regex
    -- parameter
    match_type TEXT DEFAULT 'prefix',

    strip_prefix INTEGER DEFAULT 1,

    preserve_host_header INTEGER DEFAULT 0,

    -- none
    -- basic
    -- jwt
    -- external
    -- api_key
    auth_type TEXT DEFAULT 'none',

    is_active INTEGER DEFAULT 1,

    proxy_timeout_seconds INTEGER DEFAULT 30,

    retry_count INTEGER DEFAULT 0,

    retry_delay_ms INTEGER DEFAULT 100,

    max_request_size_mb INTEGER DEFAULT 10,

    websocket_enabled INTEGER DEFAULT 0,

    cache_enabled INTEGER DEFAULT 0,

    cache_ttl_seconds INTEGER DEFAULT 60,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_host_id)
        REFERENCES virtual_hosts(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE virtual_directory_methods
-- Allowed HTTP methods
-- =========================================================
CREATE TABLE virtual_directory_methods (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    http_method TEXT NOT NULL,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE route_consumer_access
-- Route access control
-- =========================================================
CREATE TABLE route_consumer_access (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    consumer_id INTEGER NOT NULL,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE,

    FOREIGN KEY (consumer_id)
        REFERENCES api_consumers(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE external_auth
-- External auth config
-- =========================================================
CREATE TABLE external_auth (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    auth_url TEXT NOT NULL,

    http_method TEXT DEFAULT 'POST',

    request_timeout_seconds INTEGER DEFAULT 5,

    send_headers INTEGER DEFAULT 1,

    send_body INTEGER DEFAULT 0,

    success_key TEXT DEFAULT 'status',

    success_value TEXT DEFAULT 'true',

    message_key TEXT DEFAULT 'message',

    token_key TEXT,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE jwt_configs
-- JWT config
-- =========================================================
CREATE TABLE jwt_configs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    algorithm TEXT DEFAULT 'HS256',

    jwt_secret TEXT NOT NULL,

    issuer TEXT,

    audience TEXT,

    expired_in_seconds INTEGER DEFAULT 3600,

    clock_skew_seconds INTEGER DEFAULT 30,

    require_exp INTEGER DEFAULT 1,

    require_iat INTEGER DEFAULT 1,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE jwt_tokens
-- JWT blacklist/session
-- =========================================================
CREATE TABLE jwt_tokens (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    consumer_id INTEGER,

    token TEXT NOT NULL,

    issued_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    expired_at DATETIME,

    is_revoked INTEGER DEFAULT 0,

    ip_address TEXT,

    user_agent TEXT,

    FOREIGN KEY (consumer_id)
        REFERENCES api_consumers(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE rate_limits
-- Rate limiting
-- =========================================================
CREATE TABLE rate_limits (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    limit_by TEXT DEFAULT 'ip',

    requests_per_minute INTEGER DEFAULT 60,

    burst INTEGER DEFAULT 10,

    block_duration_seconds INTEGER DEFAULT 60,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE cors_configs
-- CORS settings
-- =========================================================
CREATE TABLE cors_configs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    allowed_origins TEXT DEFAULT '*',

    allowed_methods TEXT DEFAULT 'GET,POST,PUT,DELETE,OPTIONS',

    allowed_headers TEXT DEFAULT '*',

    exposed_headers TEXT,

    allow_credentials INTEGER DEFAULT 0,

    max_age_seconds INTEGER DEFAULT 3600,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE circuit_breakers
-- Circuit breaker protection
-- =========================================================
CREATE TABLE circuit_breakers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    enabled INTEGER DEFAULT 1,

    failure_threshold INTEGER DEFAULT 5,

    recovery_timeout_seconds INTEGER DEFAULT 30,

    half_open_max_requests INTEGER DEFAULT 3,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE request_header_rules
-- Request header manipulation
-- =========================================================
CREATE TABLE request_header_rules (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    header_name TEXT NOT NULL,

    operation TEXT NOT NULL,

    value_source TEXT DEFAULT 'static',

    header_value TEXT,

    source_header TEXT,

    variable_name TEXT,

    execution_order INTEGER DEFAULT 1,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE response_header_rules
-- Response header manipulation
-- =========================================================
CREATE TABLE response_header_rules (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    header_name TEXT NOT NULL,

    operation TEXT NOT NULL,

    header_value TEXT,

    execution_order INTEGER DEFAULT 1,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE query_rewrites
-- Query parameter rewrite
-- =========================================================
CREATE TABLE query_rewrites (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    param_name TEXT NOT NULL,

    param_value TEXT,

    operation TEXT DEFAULT 'set',

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE acme_accounts
-- Let's Encrypt ACME accounts
-- =========================================================
CREATE TABLE acme_accounts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    email TEXT NOT NULL,

    provider_url TEXT NOT NULL,

    account_key_path TEXT NOT NULL,

    is_default INTEGER DEFAULT 0,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- =========================================================
-- TABLE ssl_certificates
-- SSL certificate storage
-- =========================================================
CREATE TABLE ssl_certificates (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    acme_account_id INTEGER,

    provider TEXT DEFAULT 'lets_encrypt',

    challenge_type TEXT DEFAULT 'http01',

    certificate_path TEXT NOT NULL,

    private_key_path TEXT NOT NULL,

    chain_path TEXT,

    auto_renew INTEGER DEFAULT 1,

    renew_before_days INTEGER DEFAULT 30,

    last_renew_at DATETIME,

    expired_at DATETIME,

    renew_status TEXT DEFAULT 'pending',

    last_error TEXT,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (acme_account_id)
        REFERENCES acme_accounts(id)
        ON DELETE SET NULL
);

-- =========================================================
-- TABLE certificate_domains
-- SSL certificate domains
-- =========================================================
CREATE TABLE certificate_domains (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    ssl_certificate_id INTEGER NOT NULL,

    domain_name TEXT NOT NULL,

    is_wildcard INTEGER DEFAULT 0,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (ssl_certificate_id)
        REFERENCES ssl_certificates(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE ssl_certificate_bindings
-- SSL bindings
-- =========================================================
CREATE TABLE ssl_certificate_bindings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    ssl_certificate_id INTEGER NOT NULL,

    -- host
    -- virtual_host
    -- global
    binding_type TEXT NOT NULL,

    host_id INTEGER,

    virtual_host_id INTEGER,

    is_default INTEGER DEFAULT 0,

    priority INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (ssl_certificate_id)
        REFERENCES ssl_certificates(id)
        ON DELETE CASCADE,

    FOREIGN KEY (host_id)
        REFERENCES hosts(id)
        ON DELETE CASCADE,

    FOREIGN KEY (virtual_host_id)
        REFERENCES virtual_hosts(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE tls_options
-- TLS settings
-- =========================================================
CREATE TABLE tls_options (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    binding_type TEXT NOT NULL,

    host_id INTEGER,

    virtual_host_id INTEGER,

    min_tls_version TEXT DEFAULT '1.2',

    http2_enabled INTEGER DEFAULT 1,

    hsts_enabled INTEGER DEFAULT 1,

    hsts_max_age INTEGER DEFAULT 31536000,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (host_id)
        REFERENCES hosts(id)
        ON DELETE CASCADE,

    FOREIGN KEY (virtual_host_id)
        REFERENCES virtual_hosts(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE ip_whitelists
-- Allowed IPs
-- =========================================================
CREATE TABLE ip_whitelists (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    ip_address TEXT NOT NULL,

    description TEXT,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE ip_blacklists
-- Blocked IPs
-- =========================================================
CREATE TABLE ip_blacklists (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    ip_address TEXT NOT NULL,

    reason TEXT,

    expired_at DATETIME,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE service_discovery
-- Backend discovery
-- =========================================================
CREATE TABLE service_discovery (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_host_id INTEGER NOT NULL,

    provider TEXT NOT NULL,

    endpoint_url TEXT NOT NULL,

    refresh_interval_seconds INTEGER DEFAULT 30,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_host_id)
        REFERENCES virtual_hosts(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE config_versions
-- Config versioning
-- =========================================================
CREATE TABLE config_versions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    config_name TEXT NOT NULL,

    version_number INTEGER NOT NULL,

    changed_by TEXT,

    notes TEXT,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- =========================================================
-- TABLE api_keys
-- Dedicated API keys
-- =========================================================
CREATE TABLE api_keys (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    consumer_id INTEGER NOT NULL,

    api_key TEXT NOT NULL UNIQUE,

    description TEXT,

    expired_at DATETIME,

    rate_limit_override INTEGER,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (consumer_id)
        REFERENCES api_consumers(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE maintenance_windows
-- Maintenance mode
-- =========================================================
CREATE TABLE maintenance_windows (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_host_id INTEGER,

    title TEXT,

    start_at DATETIME,

    end_at DATETIME,

    maintenance_response_code INTEGER DEFAULT 503,

    maintenance_message TEXT,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_host_id)
        REFERENCES virtual_hosts(id)
        ON DELETE CASCADE
);
```
```sql
PRAGMA foreign_keys = ON;

-- =========================================================
-- TABLE users
-- Admin panel users
-- =========================================================
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    full_name TEXT,
    email TEXT,

    -- super_admin
    -- admin
    -- operator
    -- viewer
    role TEXT DEFAULT 'admin',

    is_active INTEGER DEFAULT 1,

    last_login_at DATETIME,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_username
ON users(username);

-- =========================================================
-- TABLE api_consumers
-- API consumers / applications
-- =========================================================
CREATE TABLE api_consumers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    consumer_name TEXT NOT NULL UNIQUE,

    description TEXT,

    contact_email TEXT,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- =========================================================
-- TABLE consumer_credentials
-- Consumer credentials
-- =========================================================
CREATE TABLE consumer_credentials (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    consumer_id INTEGER NOT NULL,

    -- basic
    -- api_key
    -- jwt
    auth_type TEXT NOT NULL,

    username TEXT,

    password_hash TEXT,

    api_key TEXT,

    jwt_secret TEXT,

    expired_at DATETIME,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (consumer_id)
        REFERENCES api_consumers(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE hosts
-- Root host grouping
-- =========================================================
CREATE TABLE hosts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    host_name TEXT NOT NULL UNIQUE,

    description TEXT,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- =========================================================
-- TABLE virtual_hosts
-- Virtual domains
-- =========================================================
CREATE TABLE virtual_hosts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    host_id INTEGER NOT NULL,

    vhost_name TEXT NOT NULL UNIQUE,

    -- round_robin
    -- weighted_round_robin
    -- least_conn
    -- ip_hash
    -- random
    -- failover
    lb_algorithm TEXT DEFAULT 'round_robin',

    sticky_session INTEGER DEFAULT 0,

    -- active-active
    -- active-passive
    failover_mode TEXT DEFAULT 'active-active',

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (host_id)
        REFERENCES hosts(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE upstream_servers
-- Backend servers
-- =========================================================
CREATE TABLE upstream_servers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_host_id INTEGER NOT NULL,

    target_host TEXT NOT NULL,

    target_port INTEGER NOT NULL,

    protocol TEXT DEFAULT 'http',

    priority INTEGER DEFAULT 1,

    weight INTEGER DEFAULT 1,

    is_backup INTEGER DEFAULT 0,

    is_active INTEGER DEFAULT 1,

    health_check_enabled INTEGER DEFAULT 1,

    health_check_path TEXT DEFAULT '/health',

    health_check_interval_seconds INTEGER DEFAULT 10,

    health_check_timeout_seconds INTEGER DEFAULT 3,

    max_fails INTEGER DEFAULT 3,

    fail_timeout_seconds INTEGER DEFAULT 30,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_host_id)
        REFERENCES virtual_hosts(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE virtual_directories
-- API routes
-- =========================================================
CREATE TABLE virtual_directories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_host_id INTEGER NOT NULL,

    source_path TEXT NOT NULL,

    target_path TEXT NOT NULL,

    -- exact
    -- prefix
    -- wildcard
    -- regex
    -- parameter
    match_type TEXT DEFAULT 'prefix',

    strip_prefix INTEGER DEFAULT 1,

    preserve_host_header INTEGER DEFAULT 0,

    -- none
    -- basic
    -- jwt
    -- external
    -- api_key
    auth_type TEXT DEFAULT 'none',

    is_active INTEGER DEFAULT 1,

    proxy_timeout_seconds INTEGER DEFAULT 30,

    retry_count INTEGER DEFAULT 0,

    retry_delay_ms INTEGER DEFAULT 100,

    max_request_size_mb INTEGER DEFAULT 10,

    websocket_enabled INTEGER DEFAULT 0,

    cache_enabled INTEGER DEFAULT 0,

    cache_ttl_seconds INTEGER DEFAULT 60,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_host_id)
        REFERENCES virtual_hosts(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE virtual_directory_methods
-- Allowed HTTP methods
-- =========================================================
CREATE TABLE virtual_directory_methods (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    http_method TEXT NOT NULL,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE route_consumer_access
-- Route access control
-- =========================================================
CREATE TABLE route_consumer_access (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    consumer_id INTEGER NOT NULL,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE,

    FOREIGN KEY (consumer_id)
        REFERENCES api_consumers(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE external_auth
-- External auth config
-- =========================================================
CREATE TABLE external_auth (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    auth_url TEXT NOT NULL,

    http_method TEXT DEFAULT 'POST',

    request_timeout_seconds INTEGER DEFAULT 5,

    send_headers INTEGER DEFAULT 1,

    send_body INTEGER DEFAULT 0,

    success_key TEXT DEFAULT 'status',

    success_value TEXT DEFAULT 'true',

    message_key TEXT DEFAULT 'message',

    token_key TEXT,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE jwt_configs
-- JWT config
-- =========================================================
CREATE TABLE jwt_configs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    algorithm TEXT DEFAULT 'HS256',

    jwt_secret TEXT NOT NULL,

    issuer TEXT,

    audience TEXT,

    expired_in_seconds INTEGER DEFAULT 3600,

    clock_skew_seconds INTEGER DEFAULT 30,

    require_exp INTEGER DEFAULT 1,

    require_iat INTEGER DEFAULT 1,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE jwt_tokens
-- JWT blacklist/session
-- =========================================================
CREATE TABLE jwt_tokens (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    consumer_id INTEGER,

    token TEXT NOT NULL,

    issued_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    expired_at DATETIME,

    is_revoked INTEGER DEFAULT 0,

    ip_address TEXT,

    user_agent TEXT,

    FOREIGN KEY (consumer_id)
        REFERENCES api_consumers(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE rate_limits
-- Rate limiting
-- =========================================================
CREATE TABLE rate_limits (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    limit_by TEXT DEFAULT 'ip',

    requests_per_minute INTEGER DEFAULT 60,

    burst INTEGER DEFAULT 10,

    block_duration_seconds INTEGER DEFAULT 60,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE cors_configs
-- CORS settings
-- =========================================================
CREATE TABLE cors_configs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    allowed_origins TEXT DEFAULT '*',

    allowed_methods TEXT DEFAULT 'GET,POST,PUT,DELETE,OPTIONS',

    allowed_headers TEXT DEFAULT '*',

    exposed_headers TEXT,

    allow_credentials INTEGER DEFAULT 0,

    max_age_seconds INTEGER DEFAULT 3600,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE circuit_breakers
-- Circuit breaker protection
-- =========================================================
CREATE TABLE circuit_breakers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    enabled INTEGER DEFAULT 1,

    failure_threshold INTEGER DEFAULT 5,

    recovery_timeout_seconds INTEGER DEFAULT 30,

    half_open_max_requests INTEGER DEFAULT 3,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE request_header_rules
-- Request header manipulation
-- =========================================================
CREATE TABLE request_header_rules (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    header_name TEXT NOT NULL,

    operation TEXT NOT NULL,

    value_source TEXT DEFAULT 'static',

    header_value TEXT,

    source_header TEXT,

    variable_name TEXT,

    execution_order INTEGER DEFAULT 1,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE response_header_rules
-- Response header manipulation
-- =========================================================
CREATE TABLE response_header_rules (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    header_name TEXT NOT NULL,

    operation TEXT NOT NULL,

    header_value TEXT,

    execution_order INTEGER DEFAULT 1,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE query_rewrites
-- Query parameter rewrite
-- =========================================================
CREATE TABLE query_rewrites (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    param_name TEXT NOT NULL,

    param_value TEXT,

    operation TEXT DEFAULT 'set',

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE acme_accounts
-- Let's Encrypt ACME accounts
-- =========================================================
CREATE TABLE acme_accounts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    email TEXT NOT NULL,

    provider_url TEXT NOT NULL,

    account_key_path TEXT NOT NULL,

    is_default INTEGER DEFAULT 0,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- =========================================================
-- TABLE ssl_certificates
-- SSL certificate storage
-- =========================================================
CREATE TABLE ssl_certificates (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    acme_account_id INTEGER,

    provider TEXT DEFAULT 'lets_encrypt',

    challenge_type TEXT DEFAULT 'http01',

    certificate_path TEXT NOT NULL,

    private_key_path TEXT NOT NULL,

    chain_path TEXT,

    auto_renew INTEGER DEFAULT 1,

    renew_before_days INTEGER DEFAULT 30,

    last_renew_at DATETIME,

    expired_at DATETIME,

    renew_status TEXT DEFAULT 'pending',

    last_error TEXT,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (acme_account_id)
        REFERENCES acme_accounts(id)
        ON DELETE SET NULL
);

-- =========================================================
-- TABLE certificate_domains
-- SSL certificate domains
-- =========================================================
CREATE TABLE certificate_domains (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    ssl_certificate_id INTEGER NOT NULL,

    domain_name TEXT NOT NULL,

    is_wildcard INTEGER DEFAULT 0,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (ssl_certificate_id)
        REFERENCES ssl_certificates(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE ssl_certificate_bindings
-- SSL bindings
-- =========================================================
CREATE TABLE ssl_certificate_bindings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    ssl_certificate_id INTEGER NOT NULL,

    -- host
    -- virtual_host
    -- global
    binding_type TEXT NOT NULL,

    host_id INTEGER,

    virtual_host_id INTEGER,

    is_default INTEGER DEFAULT 0,

    priority INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (ssl_certificate_id)
        REFERENCES ssl_certificates(id)
        ON DELETE CASCADE,

    FOREIGN KEY (host_id)
        REFERENCES hosts(id)
        ON DELETE CASCADE,

    FOREIGN KEY (virtual_host_id)
        REFERENCES virtual_hosts(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE tls_options
-- TLS settings
-- =========================================================
CREATE TABLE tls_options (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    binding_type TEXT NOT NULL,

    host_id INTEGER,

    virtual_host_id INTEGER,

    min_tls_version TEXT DEFAULT '1.2',

    http2_enabled INTEGER DEFAULT 1,

    hsts_enabled INTEGER DEFAULT 1,

    hsts_max_age INTEGER DEFAULT 31536000,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (host_id)
        REFERENCES hosts(id)
        ON DELETE CASCADE,

    FOREIGN KEY (virtual_host_id)
        REFERENCES virtual_hosts(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE ip_whitelists
-- Allowed IPs
-- =========================================================
CREATE TABLE ip_whitelists (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    ip_address TEXT NOT NULL,

    description TEXT,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE ip_blacklists
-- Blocked IPs
-- =========================================================
CREATE TABLE ip_blacklists (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_directory_id INTEGER NOT NULL,

    ip_address TEXT NOT NULL,

    reason TEXT,

    expired_at DATETIME,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_directory_id)
        REFERENCES virtual_directories(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE service_discovery
-- Backend discovery
-- =========================================================
CREATE TABLE service_discovery (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_host_id INTEGER NOT NULL,

    provider TEXT NOT NULL,

    endpoint_url TEXT NOT NULL,

    refresh_interval_seconds INTEGER DEFAULT 30,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_host_id)
        REFERENCES virtual_hosts(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE config_versions
-- Config versioning
-- =========================================================
CREATE TABLE config_versions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    config_name TEXT NOT NULL,

    version_number INTEGER NOT NULL,

    changed_by TEXT,

    notes TEXT,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- =========================================================
-- TABLE api_keys
-- Dedicated API keys
-- =========================================================
CREATE TABLE api_keys (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    consumer_id INTEGER NOT NULL,

    api_key TEXT NOT NULL UNIQUE,

    description TEXT,

    expired_at DATETIME,

    rate_limit_override INTEGER,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (consumer_id)
        REFERENCES api_consumers(id)
        ON DELETE CASCADE
);

-- =========================================================
-- TABLE maintenance_windows
-- Maintenance mode
-- =========================================================
CREATE TABLE maintenance_windows (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    virtual_host_id INTEGER,

    title TEXT,

    start_at DATETIME,

    end_at DATETIME,

    maintenance_response_code INTEGER DEFAULT 503,

    maintenance_message TEXT,

    is_active INTEGER DEFAULT 1,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (virtual_host_id)
        REFERENCES virtual_hosts(id)
        ON DELETE CASCADE
);
