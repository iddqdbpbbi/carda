create table users(
    user_id serial primary key,
    created_at timestamp default now(),
    username varchar(30) not null unique);

CREATE TYPE essential_contact_type AS ENUM ('phone', 'email');

create table contacts(
    contact_id serial primary key,
    user_id int not null,
    created_at timestamp default now(),
    contact_type essential_contact_type,
    contact_value varchar(255) not null unique,
    active boolean not null,
    foreign key (user_id) references users(user_id) on delete cascade
);

create or replace procedure add_user_with_email(
    p_username varchar(30),
    p_email varchar(255)
)
language plpgsql
as $$
begin
    insert into users (username) 
        values (p_username);
    insert into contacts(user_id, contact_type, active, contact_value) 
                values (lastval(), 'email', TRUE, p_email);
    raise notice 'start of procedure';
    exception
        WHEN unique_violation THEN
            RAISE EXCEPTION 'user with username already exists';
            rollback;

        when others then
            -- Rollback the transaction in case of an error
            raise exception 'unexpected error big pp';
            rollback;
    commit;
end;
$$;