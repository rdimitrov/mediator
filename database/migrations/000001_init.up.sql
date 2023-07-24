-- Copyright 2023 Stacklok, Inc
--
-- Licensed under the Apache License, Version 2.0 (the "License");
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--      http://www.apache.org/licenses/LICENSE-2.0
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-- See the License for the specific language governing permissions and
-- limitations under the License.

-- organizations table
CREATE TABLE organizations (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    company TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- groups table
CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    organization_id INTEGER NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    description TEXT,
    is_protected BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);


-- roles table
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    organization_id INTEGER NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    group_id INTEGER REFERENCES groups(id) ON DELETE CASCADE,
    name TEXT NOT NULL,    
    is_admin BOOLEAN NOT NULL DEFAULT FALSE,
    is_protected BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    organization_id INTEGER NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    email TEXT,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    needs_password_change BOOLEAN NOT NULL DEFAULT TRUE,
    first_name TEXT,
    last_name TEXT,
    is_protected BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    min_token_issued_time TIMESTAMP
);

-- user/groups
CREATE TABLE user_groups (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    group_id INTEGER NOT NULL REFERENCES groups(id) ON DELETE CASCADE
);

-- user/roles
CREATE TABLE user_roles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role_id INTEGER NOT NULL REFERENCES roles(id) ON DELETE CASCADE
);

-- provider_access_tokens table
CREATE TABLE provider_access_tokens (
    id SERIAL PRIMARY KEY,
    provider TEXT NOT NULL,
    group_id INTEGER NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    encrypted_token TEXT NOT NULL,
    expiration_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- signing_keys table
CREATE TABLE signing_keys (
    id SERIAL PRIMARY KEY,
    group_id INTEGER NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    private_key TEXT NOT NULL,
    public_key TEXT NOT NULL,
    passphrase TEXT NOT NULL,
    key_identifier TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- repositories table
CREATE TABLE repositories (
    id SERIAL PRIMARY KEY,
    provider TEXT NOT NULL,
    group_id INTEGER NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    repo_owner TEXT NOT NULL,
    repo_name TEXT NOT NULL,
    repo_id INTEGER NOT NULL,
    is_private BOOLEAN NOT NULL,
    is_fork BOOLEAN NOT NULL,
    webhook_id INTEGER,
    webhook_url TEXT NOT NULL,
    deploy_url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE session_store (
    id SERIAL PRIMARY KEY,
    provider TEXT NOT NULL,
    grp_id INTEGER,
    port INTEGER,
    session_state TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- policy type
CREATE TABLE policy_types (
    id SERIAL PRIMARY KEY,
    provider TEXT NOT NULL,
    policy_type VARCHAR(50) NOT NULL,
    description TEXT,
    version varchar(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE policies (
    id SERIAL PRIMARY KEY,
    provider TEXT NOT NULL,
    group_id INTEGER NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    policy_type INTEGER NOT NULL REFERENCES policy_types ON DELETE CASCADE,
    policy_definition JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

create type policy_status_types as enum ('success', 'failure');
CREATE TABLE policy_status (
    id SERIAL PRIMARY KEY,
    repository_id INTEGER NOT NULL REFERENCES repositories(id) ON DELETE CASCADE,
    policy_id INTEGER NOT NULL REFERENCES policies(id) ON DELETE CASCADE,
    policy_status policy_status_types NOT NULL,
    last_updated TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE policy_violations (
    id SERIAL PRIMARY KEY,
    repository_id INTEGER NOT NULL REFERENCES repositories(id) ON DELETE CASCADE,
    policy_id INTEGER NOT NULL REFERENCES policies(id) ON DELETE CASCADE,
    metadata JSONB NOT NULL,
    violation JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE rule_type (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    provider TEXT NOT NULL,
    group_id INTEGER NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    rule_type VARCHAR(50) NOT NULL,
    definition JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Unique constraint
ALTER TABLE provider_access_tokens ADD CONSTRAINT unique_group_id UNIQUE (group_id);
ALTER TABLE repositories ADD CONSTRAINT unique_repo_id UNIQUE (repo_id);
ALTER TABLE signing_keys ADD CONSTRAINT unique_key_identifier UNIQUE (key_identifier);

-- Indexes
CREATE UNIQUE INDEX organizations_name_lower_idx ON organizations (LOWER(name));
CREATE UNIQUE INDEX organizations_company_lower_idx ON organizations (LOWER(company));
CREATE INDEX idx_users_organization_id ON users(organization_id);
CREATE INDEX idx_groups_organization_id ON groups(organization_id);
CREATE INDEX idx_roles_group_id ON roles(group_id);
CREATE UNIQUE INDEX roles_organization_id_name_lower_idx ON roles (organization_id, LOWER(name));
CREATE INDEX idx_provider_access_tokens_group_id ON provider_access_tokens(group_id);
CREATE UNIQUE INDEX users_organization_id_email_lower_idx ON users (organization_id, LOWER(email));
CREATE UNIQUE INDEX users_organization_id_username_lower_idx ON users (organization_id, LOWER(username));
CREATE UNIQUE INDEX repositories_repo_id_idx ON repositories(repo_id);
CREATE UNIQUE INDEX policies_group_id_policy_type_idx ON policies(provider, group_id, policy_type);
CREATE UNIQUE INDEX policy_types_idx ON policy_types(provider, policy_type);
CREATE UNIQUE INDEX policy_status_idx ON policy_status(repository_id, policy_id);
CREATE UNIQUE INDEX rule_type_idx ON rule_type(provider, group_id, rule_type);

-- Create default root organization

INSERT INTO organizations (name, company) 
VALUES ('Root Organization', 'Root Company');

INSERT INTO groups (organization_id, name, is_protected)
VALUES (1, 'Root Group', TRUE);

-- superadmin role
INSERT INTO roles (organization_id, name, is_admin, is_protected)
VALUES (1, 'Superadmin Role', TRUE, TRUE);

INSERT INTO users (organization_id, email, username, password, first_name, last_name, is_protected, needs_password_change)
VALUES (1, 'root@localhost', 'root', '$argon2id$v=19$m=16,t=2,p=1$c2VjcmV0aGFzaA$WP4Vqo6QtHBY+n0x99R81Q', 'Root', 'Admin', TRUE, FALSE);   -- password is P4ssw@rd

INSERT INTO user_groups (user_id, group_id) VALUES (1, 1);
INSERT INTO user_roles (user_id, role_id) VALUES (1, 1);

-- policy types
INSERT INTO policy_types (provider, policy_type, description, version) VALUES
('github', 'branch_protection', 'Policy type to enforce branch protection rules on a repo', '1.0.0');
INSERT INTO policy_types (provider, policy_type, description, version) VALUES
('github', 'secret_scanning', 'Policy type to enforce secret scanning a repo', '1.0.0');