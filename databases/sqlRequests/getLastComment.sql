SELECT id,
    post_id,
    user_id,
    content,
    image,
    created_at,
FROM comments
ORDER BY id DESC LIMIT 1