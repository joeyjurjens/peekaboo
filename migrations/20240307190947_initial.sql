-- +goose Up
-- +goose StatementBegin
CREATE TABLE projects (
    id INTEGER NOT NULL PRIMARY KEY,
    project_name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE urls (
    id INTEGER NOT NULL PRIMARY KEY,
    url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    project_id INT NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id)
);

CREATE TABLE screenshots (
    id INTEGER NOT NULL PRIMARY KEY,
    type TEXT NOT NULL CHECK (type IN ('old', 'new', 'diff')),
    filepath TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    url_id INT NOT NULL,
    FOREIGN KEY (url_id) REFERENCES urls(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE screenshots;
DROP TABLE urls;
DROP TABLE projects;
-- +goose StatementEnd
