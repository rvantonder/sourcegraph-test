BEGIN;

ALTER TABLE campaign_jobs ALTER COLUMN description DROP NOT NULL;

COMMIT;
