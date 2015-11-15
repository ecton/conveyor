CREATE TABLE infrastructures (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL UNIQUE
);

CREATE TABLE servers (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  infrastructure_id INTEGER NOT NULL REFERENCES infrastructures(id)
);

CREATE TABLE workers (
    id BIGSERIAL PRIMARY KEY,
    server_id INTEGER NOT NULL REFERENCES servers(id),
    pid INTEGER NOT NULL,
    attributes JSONB NOT NULL DEFAULT '{}'::JSONB,
    first_checked_in TIMESTAMP NOT NULL DEFAULT NOW(),
    last_checked_in TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TYPE JOB_STATUS as ENUM('queued', 'executing', 'completed', 'failed');
CREATE TABLE jobs (
  id BIGSERIAL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  ended_at TIMESTAMP,
  status JOB_STATUS NOT NULL DEFAULT 'queued',
  status_message TEXT,
  worker_id BIGINT REFERENCES workers(id),
  arguments JSONB,
  result JSONB
);

-- Automatic partitioning jobs by day to make it easy to clean up old jobs
CREATE FUNCTION partition_and_insert_job()
RETURNS trigger AS $BODY$
  DECLARE
    partition TEXT;
    start_date DATE;
    end_date DATE;
  BEGIN
    partition := 'jobs_' || TO_CHAR(NEW.created_at, 'YYYY_MM_DD');
    BEGIN
      EXECUTE 'INSERT INTO ' || partition || ' VALUES ($1.*)' USING NEW;
      RETURN NULL;
    EXCEPTION WHEN undefined_table THEN
      start_date := NEW.created_at::DATE;
      end_date := start_date + 1;
      BEGIN
        EXECUTE 'CREATE TABLE ' || partition || '(
          CHECK (created_at >= ' || QUOTE_LITERAL(start_date) || ' AND created_at < ' || QUOTE_LITERAL(end_date)')
        ) INHERITS (jobs)';
      EXCEPTION WHEN duplicate_table THEN
        -- Another insert already created the parition
      END;
      EXECUTE 'INSERT INTO ' || partition || ' VALUES ($1.*)' USING NEW;
      RETURN NULL;
    END;
  END;
$BODY$
LANGUAGE plpgsql;

CREATE TRIGGER jobs_partition_and_insert_job
BEFORE INSERT ON jobs
FOR EACH ROW EXECUTE PROCEDURE partition_and_insert_job();
