-- +goose Up

CREATE EXTENSION IF NOT EXISTS vector;

CREATE TABLE repos (
                       id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       installation_id BIGINT NOT NULL,
                       full_name       TEXT NOT NULL UNIQUE,
                       created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE developers (
                            id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                            github_login  TEXT NOT NULL UNIQUE,
                            email         TEXT,
                            created_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE pull_requests (
                               id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                               repo_id     UUID NOT NULL REFERENCES repos(id),
                               pr_number   INT NOT NULL,
                               author_id   UUID REFERENCES developers(id),
                               head_sha    TEXT NOT NULL,
                               opened_at   TIMESTAMPTZ NOT NULL,
                               merged_at   TIMESTAMPTZ,
                               UNIQUE (repo_id, pr_number)
);

CREATE TABLE files (
                       id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       repo_id   UUID NOT NULL REFERENCES repos(id),
                       path      TEXT NOT NULL,
                       language  TEXT NOT NULL CHECK (language IN ('python','java','go','javascript','typescript')),
    UNIQUE (repo_id, path)
);

CREATE TABLE functions (
                           id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                           file_id      UUID NOT NULL REFERENCES files(id),
                           name         TEXT NOT NULL,
                           node_type    TEXT NOT NULL CHECK (node_type IN ('function','method','class')),
                           parent_name  TEXT,
                           UNIQUE (file_id, name, parent_name)
);

CREATE TABLE patterns (
                          id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                          repo_id           UUID NOT NULL REFERENCES repos(id),
                          category          TEXT NOT NULL,
                          description       TEXT NOT NULL,
                          first_seen_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
                          recurrence_count  INT NOT NULL DEFAULT 1,
                          embedding         vector(384)
);

CREATE TABLE review_comments (
                                 id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                 pr_id         UUID NOT NULL REFERENCES pull_requests(id),
                                 file_id       UUID REFERENCES files(id),
                                 function_id   UUID REFERENCES functions(id),
                                 severity      TEXT NOT NULL CHECK (severity IN ('style','performance','bug','security')),
                                 category      TEXT NOT NULL,
                                 message       TEXT NOT NULL,
                                 reason        TEXT NOT NULL,
                                 confidence    REAL NOT NULL,
                                 line_number   INT NOT NULL,
                                 pattern_id    UUID REFERENCES patterns(id),
                                 embedding     vector(384),
                                 created_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE bugs (
                      id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                      introduced_pr_id  UUID NOT NULL REFERENCES pull_requests(id),
                      confirmed_by_dev  UUID REFERENCES developers(id),
                      confirmed_at      TIMESTAMPTZ,
                      description       TEXT
);

CREATE TABLE graph_edges (
                             id           BIGSERIAL PRIMARY KEY,
                             src_type     TEXT NOT NULL,
                             src_id       UUID NOT NULL,
                             edge_type    TEXT NOT NULL,
                             dst_type     TEXT NOT NULL,
                             dst_id       UUID NOT NULL,
                             occurred_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
                             metadata     JSONB
);

CREATE INDEX idx_graph_edges_src ON graph_edges (src_type, src_id, edge_type, occurred_at);
CREATE INDEX idx_graph_edges_dst ON graph_edges (dst_type, dst_id, edge_type, occurred_at);

-- +goose Down

DROP TABLE IF EXISTS graph_edges;
DROP TABLE IF EXISTS bugs;
DROP TABLE IF EXISTS review_comments;
DROP TABLE IF EXISTS patterns;
DROP TABLE IF EXISTS functions;
DROP TABLE IF EXISTS files;
DROP TABLE IF EXISTS pull_requests;
DROP TABLE IF EXISTS developers;
DROP TABLE IF EXISTS repos;