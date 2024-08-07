SELECT pr.id,
    pr.post_id,
    pr.user_id,
    pr.reaction
FROM post_reaction AS pr
    LEFT JOIN posts AS p ON pr.post_id = p.id
WHERE (
        p.type = 'group'
        AND p.group_id = ?
    )