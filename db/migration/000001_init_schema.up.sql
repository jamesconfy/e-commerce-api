CREATE EXTENSION IF NOT EXISTS dblink;

DO $$
BEGIN
   IF EXISTS (SELECT FROM pg_database WHERE datname = 'e_commerce_api') THEN
      RAISE NOTICE 'Database already exists';
   ELSE
      PERFORM dblink_exec('dbname=' || current_database(), 'CREATE DATABASE e_commerce_api');
   END IF;
END $$;