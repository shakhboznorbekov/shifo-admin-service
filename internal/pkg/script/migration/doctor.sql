create table doctors
(
    id              uuid primary key not null ,
    first_name      varchar,
    last_name       varchar,
    specialty_id    uuid references specialties(id),
    file_link       varchar,
    work_experience text,
    workplace_id    uuid references workplaces(id),
    work_price      varchar,
    start_work      timestamp,
    end_work        timestamp,
    created_at      timestamp default now(),
    deleted_at      timestamp,
    updated_at      timestamp,
    updated_by uuid references users(id),
    created_by uuid references users(id),
    deleted_by uuid references users(id)
);

