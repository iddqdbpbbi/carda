CREATE TABLE USERS(
    USER_ID SERIAL PRIMARY KEY,
    CREATED_AT TIMESTAMP DEFAULT NOW(),
    USERNAME VARCHAR(30) NOT NULL UNIQUE);

CREATE TYPE ESSENTIAL_CONTACT_TYPE AS ENUM ('PHONE', 'EMAIL');

CREATE TABLE CONTACTS(
    CONTACT_ID SERIAL PRIMARY KEY,
    USER_ID INT NOT NULL,
    CREATED_AT TIMESTAMP DEFAULT NOW(),
    CONTACT_TYPE ESSENTIAL_CONTACT_TYPE,
    CONTACT_VALUE VARCHAR(255) NOT NULL UNIQUE,
    ACTIVE BOOLEAN NOT NULL,
    FOREIGN KEY (USER_ID) REFERENCES USERS(USER_ID) ON DELETE CASCADE
);

CREATE OR REPLACE PROCEDURE ADD_USERS_WITH_EMAIL(
    P_USERNAME VARCHAR(30),
    P_EMAIL VARCHAR(255)
)
LANGUAGE PLPGSQL
AS $$
BEGIN
    INSERT INTO USERS (USERNAME) 
        VALUES (P_USERNAME);
    INSERT INTO CONTACTS(USER_ID, CONTACT_TYPE, ACTIVE, CONTACT_VALUE) 
                VALUES (LASTVAL(), 'EMAIL', TRUE, P_EMAIL);
    EXCEPTION
        WHEN UNIQUE_VIOLATION THEN
            RAISE EXCEPTION 'USER WITH USERNAME ALREADY EXISTS';
            ROLLBACK;

        WHEN OTHERS THEN
            -- Raise "big pp" exception because i need to know if it handle by this when block
            RAISE EXCEPTION 'UNEXPECTED ERROR BIG PP';
            ROLLBACK;
    COMMIT;
END;
$$;