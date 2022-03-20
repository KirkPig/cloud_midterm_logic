CREATE TABLE IF NOT EXISTS "messages"
(
    "uuid"              TEXT NOT NULL,
    "author"            TEXT NOT NULL,
    "message"           TEXT NOT NULL,
    "likes"             TEXT NOT NULL,
    "lastUpdateAuthor"  TEXT NOT NULL,
    "lastUpdateMessage" TEXT NOT NULL,
    "lastUpdateLikes"   TEXT NOT NULL,
    "isDeleted"         TEXT NOT NULL,
    PRIMARY KEY ("uuid")
)
