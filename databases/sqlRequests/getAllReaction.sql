SELECT id,
    comment_id,
    user_id,
    reaction,
FROM comments
WHERE comment_id = ?