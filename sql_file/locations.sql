CREATE TRIGGER update_location_on_insert
    AFTER INSERT ON Scores_t
    FOR EACH ROW
BEGIN
    DECLARE user_ip VARCHAR(10);

    SELECT Ip INTO user_ip FROM Records_t WHERE id = NEW.record_id;
    IF NEW.Score_ = 1 THEN
        INSERT INTO locations_t (Ip, Pos, Neg, Num)
        VALUES (user_ip, 1, 0, 1)
        ON DUPLICATE KEY UPDATE
                                    Pos = (Pos*Num + 1)/(Num+1),
                                    Neg = Neg*Num/(Num+1),
                                    Num = Num + 1;
    ELSE
        INSERT INTO locations_t (Ip, Pos, Neg, Num)
        VALUES (user_ip, 0, 1,1)
        ON DUPLICATE KEY UPDATE
                                 Pos = Pos*Num/(Num+1),
                                 Neg = (Neg*Num + 1)/(Num+1),
                                 Num = Num + 1;
    END IF;
END;