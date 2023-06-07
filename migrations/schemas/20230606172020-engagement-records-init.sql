-- +migrate Up
CREATE TABLE engagement_records
(
    id               UUID PRIMARY KEY DEFAULT (uuid()),
    discord_user_id  BIGINT,
    discord_username VARCHAR(40),
    -- Discord's usernames have the form "username#number"
    -- for example: thanhnguyen2187#4183
    -- the username part's maximal length is 32, while the number part's length is 4
    channel_id       BIGINT,
    channel_name     VARCHAR(32),
    category_id      BIGINT,
    category_name    VARCHAR(32),
    message_id       BIGINT,
    message_length   INT NOT NULL     DEFAULT 0,
    message_count    INT NOT NULL     DEFAULT 0,
    reaction_count   INT NOT NULL     DEFAULT 0,
    reaction_emoji   VARCHAR(20),
    record_type      VARCHAR(20),
    -- based on Discord Gateway Events:
    -- - MessageCreate
    -- - MessageDelete
    -- - MessageReactionAdd
    -- - MessageReactionRemove
    -- - MessageReactionRemoveAll
    sent_at          TIMESTAMP(6), -- Discord server's creation date,
    deleted_at       TIMESTAMP(6),
    created_at       TIMESTAMP(6), -- local database's creation date
    updated_at       TIMESTAMP(6)
);
-- +migrate Down
DROP TABLE engagement_records;
