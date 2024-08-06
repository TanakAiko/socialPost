SELECT id,
    post_id,
    user_id,
    content,
    image,
    created_at,
FROM comments
WHERE post_id = ?
ORDER BY created_at DESC