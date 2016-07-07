CREATE TABLE pages
(
  "id" SERIAL NOT NULL,
  "shorturl" text NOT NULL UNIQUE,
  "longurl" "text" NOT NULL,
  "created_at" "timestamp" NOT NULL DEFAULT now(),
  CONSTRAINT "PKey" PRIMARY KEY ("id")
)
