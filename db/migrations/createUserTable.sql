/* Note: I may not use this at all. I might use something like Firebase for sign ups instead */
create table "users" (
  "id" bigserial primary key,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
);