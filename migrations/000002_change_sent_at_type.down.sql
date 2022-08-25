ALTER TABLE "users" ADD COLUMN "new_email_verification_sent_at" varchar;
UPDATE "users" SET "new_email_verification_sent_at" = "email_verification_sent_at"::varchar;
ALTER TABLE "users" DROP COLUMN "email_verification_sent_at";
ALTER TABLE "users" RENAME COLUMN "new_email_verification_sent_at" TO "email_verification_sent_at";

ALTER TABLE "users" ADD COLUMN "new_forgot_password_sent_at" varchar;
UPDATE "users" SET "new_forgot_password_sent_at" = "forgot_password_sent_at"::varchar;
ALTER TABLE "users" DROP COLUMN "forgot_password_sent_at";
ALTER TABLE "users" RENAME COLUMN "new_forgot_password_sent_at" TO "forgot_password_sent_at";