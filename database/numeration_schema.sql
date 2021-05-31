CREATE SCHEMA numeration;

CREATE TABLE numeration.steps
(
  id Serial PRIMARY KEY,
  day Date NOT NULL,
  steps_number Bigint NOT NULL
);