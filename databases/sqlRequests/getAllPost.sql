SELECT p.id,
    p.user_id,
    p.group_id,
    p.image,
    p.content,
    p.type,
    p.privacy,
    p.created_at
FROM posts AS p
    LEFT JOIN post_permissions AS pp ON (
        p.id = pp.post_id
        AND pp.user_id = ?
    )
    LEFT JOIN (
        SELECT follower_id,
            following_id
        FROM followers
        WHERE follower_id = ?
    ) AS f ON p.user_id = following_id
WHERE type = 'all'
    AND(
        p.privacy = 'public'
        OR (
            p.privacy = 'private'
            AND f.follower_id IS NOT NULL
        )
        OR (
            p.privacy = 'almost_private'
            AND pp.user_id IS NOT NULL
        )
    )
ORDER BY createdAt DESC