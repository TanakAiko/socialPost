SELECT id,
    user_id,
    image,
    content,
    type,
    privacy,
    created_at
FROM posts
WHERE type = ?
ORDER BY createdAt DESC