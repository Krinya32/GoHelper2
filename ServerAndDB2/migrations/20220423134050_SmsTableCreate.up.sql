create table sms
(
    id       bigint,
    text     text default 'Test text',
    receiver BIGINT not null,
    sender   bigint not null
);