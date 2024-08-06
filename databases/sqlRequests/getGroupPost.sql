SELECT id,
    user_id,
    group_id,
    image,
    content,
    type,
    privacy,
    created_at
FROM posts
WHERE (
        type = 'group'
        AND group_id = ?
    )