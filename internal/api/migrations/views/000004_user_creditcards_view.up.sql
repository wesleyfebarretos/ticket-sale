CREATE VIEW user_creditcards AS
SELECT
    c.uuid,
    c."name",
    c."number",
    c.expiration,
    c.user_id,
    c.created_at,
    json_build_object(
        'id', cf.id,
        'name', cf."name",
        'description', cf.description,
        'regex', cf.regex
    ) AS "creditcardFlag",
    json_build_object(
        'id', ct.id,
        'name', ct."name"
    ) AS "creditcardType"
FROM
    fin.creditcard c
JOIN
    fin.creditcard_flag cf ON cf.id = c.creditcard_flag_id 
JOIN 
    fin.creditcard_type ct ON ct.id = c.creditcard_type_id
WHERE
    c.is_deleted IS FALSE
ORDER BY c.id ASC;

