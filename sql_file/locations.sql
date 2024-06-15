CREATE VIEW locations AS
SELECT
    r.ip,
    SUM(CASE WHEN s.score_ = 1 THEN 1 ELSE 0 END) / COUNT(*) AS pos,
    SUM(CASE WHEN s.score_ = 0 THEN 1 ELSE 0 END) / COUNT(*) AS neg
from records_t r
join scores_t s on r.id = s.record_id
GROUP BY r.ip
ORDER BY pos DESC;