CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_name" VARCHAR(15),
  "email" VARCHAR(255) /* 練習用に UNIQUE は指定しない */,
  "password" VARCHAR(32),
  "location" VARCHAR(30),
  /* "deleted_at" TIMESTAMP DEFAULT NULL, */
  "created_at" TIMESTAMP DEFAULT current_timestamp,
  "updated_at" TIMESTAMP DEFAULT current_timestamp
);

INSERT INTO users(user_name, email, password) VALUES ('yamada', 'aaa@aaa.com', 'ps123456');
