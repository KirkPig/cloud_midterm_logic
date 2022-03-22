CREATE TABLE IF NOT EXISTS "messages"
(
    "uuid"              VARCHAR(100) NOT NULL,
    "author"            TEXT NOT NULL,
    "message"           TEXT NOT NULL,
    "likes"             TEXT NOT NULL,
    "last_update_author"  TIMESTAMP NOT NULL,
    "last_update_message" TIMESTAMP NOT NULL,
    "last_update_likes"   TIMESTAMP NOT NULL,
    "is_deleted"         BOOLEAN NOT NULL,
    "last_update_delete"  TIMESTAMP NOT NULL,
    PRIMARY KEY ("uuid")
)
