DO
$do$
BEGIN
   IF EXISTS (
      SELECT
      FROM   pg_catalog.pg_roles
      WHERE  rolname = 'pguser') THEN

      ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO pguser;
END IF;
END
$do$;
