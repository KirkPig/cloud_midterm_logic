CREATE TABLE IF NOT EXISTS "messages"
(
    "uuid"              TEXT NOT NULL,
    "author"            TEXT NOT NULL,
    "message"           TEXT NOT NULL,
    "likes"             TEXT NOT NULL,
    "lastUpdateAuthor"  TIMESTAMP NOT NULL,
    "lastUpdateMessage" TIMESTAMP NOT NULL,
    "lastUpdateLikes"   TIMESTAMP NOT NULL,
    "isDeleted"         BOOLEAN NOT NULL,
    "lastUpdateDelete"  TIMESTAMP NOT NULL,
    PRIMARY KEY ("uuid")
)
