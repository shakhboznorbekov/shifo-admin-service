create table specialties
(
    id         uuid primary key not null ,
    name       varchar not null,
    created_at timestamp default now(),
    deleted_at timestamp,
    updated_at timestamp,
    updated_by uuid references users(id),
    created_by uuid references users(id),
    deleted_by uuid references users(id)
);

