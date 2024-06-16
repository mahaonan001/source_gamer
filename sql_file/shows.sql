CREATE VIEW shows AS
SELECT 
    r.id as record_id,
    r.ip,
    r.chat,
    r.cleaned_comments,
    d.id AS dim_id,
    d.dim_ ,
    k.t_room,
    k.s_room,
    k.burnning_t,
    k.device_logo,
    k.hot__t,
    k.time_cyc,
    k.money_cyc,
    k.gas_cyc,
    k.ele_cyc,
    k.boal_cyc,
    s.analysis,
    s.extracted_texts,
    s.option_word,
    s.score_
FROM records_t r
JOIN keywords_t k ON r.id = k.record_id
JOIN scores_t s ON r.id = s.record_id
JOIN dims_t d ON d.id = s.dim_id;