CREATE EXTENSION IF NOT EXISTS pgcrypto;
create table if not exists users (
                       id uuid primary key default gen_random_uuid(),
                       status varchar not null ,
                       created_at timestamptz not null default now(),
                       updated_at timestamptz not null default now()
);

create table if not exists user_identities (
                                 id uuid primary key default gen_random_uuid(),
                                 user_id uuid not null references users(id) ,

                                 providers varchar not null ,
                                 provider_subject varchar not null ,
                                 created_at timestamptz not null default now(),
                                 updated_at timestamptz not null default now()
);

create index if not exists user_identities_user_id_idx on user_identities(user_id);

create unique index if not exists  user_identities_provider_subject_idx on user_identities(providers,provider_subject);