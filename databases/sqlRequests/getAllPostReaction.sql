SELECT pr.id,
    pr.post_id,
    pr.user_id,
    pr.reaction
FROM post_reaction AS pr
    LEFT JOIN posts AS p ON pr.post_id = p.id
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
WHERE p.type = 'all'
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