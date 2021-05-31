CREATE SCHEMA numeration;

CREATE TABLE numeration.steps
(
  id Serial PRIMARY KEY,
  day Timestamp NOT NULL,
  steps_number Bigint NOT NULL
);