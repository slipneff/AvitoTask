create table public.users
(
    id bigserial
        primary key
);

create table public.user_segments
(
    user_id    bigint,
    segment_id bigint,
    expires_at timestamp with time zone
);
create table public.segments
(
    id         bigserial
        primary key,
    name       text not null,
    percentage bigint
);

create table public.segment_histories
(
    id           bigserial
        primary key,
    created_at   timestamp with time zone,
    updated_at   timestamp with time zone,
    deleted_at   timestamp with time zone,
    user_id      text,
    segment_name text,
    operation    text,
    timestamp    timestamp with time zone
);

create index idx_segment_histories_deleted_at
    on public.segment_histories (deleted_at);

