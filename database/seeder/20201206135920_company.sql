-- +goose Up
-- SQL in this section is executed when the migration is applied.

INSERT INTO `companies`(name, address, majority, established) VALUES ('PT TokoPDA', 'Jl. Prof. DR. Satrio No.Kav. 11, RT.3/RW.3, Karet Semanggi, Setia Budi, Kota Jakarta Selatan, Daerah Khusus Ibukota Jakarta 12950','ECommerce', '17 April 2010');
INSERT INTO `companies`(name, address, majority, established) VALUES ('PT BukaLPK', 'Jl. Prof. DR. Satrio No.Kav. 15, RT.3/RW.3, Karet Semanggi, Setia Budi, Kota Jakarta Selatan, Daerah Khusus Ibukota Jakarta 12950','ECommerce', '31 Maret 2008');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DELETE FROM `companies`;
