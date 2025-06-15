INSERT INTO bookcabin.aircrafts
(id, code, created_at, updated_at)
VALUES(1,'738', NOW(), NOW());

-- 

INSERT INTO bookcabin.cabins
(id, aircraft_id, deck, first_row, last_row, created_at, updated_at)
VALUES(0, 1, 'MAIN', 4, 30, NOW(), NOW());

--

INSERT INTO bookcabin.seat_columns
(id, cabin_id, column_code, created_at, updated_at)
VALUES
(1, 1, 'LEFT_SIDE', NOW(), NOW()),
(2, 1, 'A', NOW(), NOW()),
(3, 1, 'B', NOW(), NOW()),
(4, 1, 'C', NOW(), NOW()),
(5, 1, 'AISLE', NOW(), NOW()),
(6, 1, 'D', NOW(), NOW()),
(7, 1, 'E', NOW(), NOW()),
(8, 1, 'F', NOW(), NOW()),
(9, 1, 'RIGHT_SIDE', NOW(), NOW());

--
-- ROW 0
INSERT INTO bookcabin.seats
(id, seat_row_id, code, storefront_slot_code, refund_indicator, free_of_charge, created_at, updated_at)
VALUES
(1, 1, 'BLANK', 'BLANK', 0, 1, NOW(), NOW()),
(2, 1, 'BULKHEAD', 'BULKHEAD', 0, 1, NOW(), NOW()),
(3, 1, 'BULKHEAD', 'BULKHEAD', 0, 1, NOW(), NOW()),
(4, 1, 'BULKHEAD', 'BULKHEAD', 0, 1, NOW(), NOW()),
(5, 1, 'BLANK', 'BLANK', 0, 1, NOW(), NOW()),
(6, 1, 'BULKHEAD', 'BULKHEAD', 0, 1, NOW(), NOW()),
(7, 1, 'BULKHEAD', 'BULKHEAD', 0, 1, NOW(), NOW()),
(8, 1, 'BULKHEAD', 'BULKHEAD', 0, 1, NOW(), NOW()),
(9, 1, 'BLANK', 'BLANK', 0, 1, NOW(), NOW());


INSERT INTO bookcabin.seat_characteristics
(id, seat_id, characteristic, created_at, updated_at)
VALUES
(1, 1, 'LEFT_SIDE', NOW(), NOW()),
(2, 9, 'RIGHT_SIDE', NOW(), NOW());

INSERT INTO bookcabin.raw_seat_characteristics
(id, seat_id, raw_characteristic, created_at, updated_at)
VALUES
(1, 1, 'LEFT_SIDE', NOW(), NOW()),
(2, 9, 'RIGHT_SIDE', NOW(), NOW());


-- ROW 4
INSERT INTO bookcabin.seats
(id, seat_row_id, code, storefront_slot_code, refund_indicator, free_of_charge, created_at, updated_at)
VALUES
(10, 2, 'BLANK', 'BLANK', 0, 1, NOW(), NOW()),     -- LEFT_SIDE BLANK
(11, 2, '4A', 'SEAT', 'R', 0, NOW(), NOW()),
(12, 2, '4B', 'SEAT', 'R', 0, NOW(), NOW()),
(13, 2, '4C', 'SEAT', 'R', 0, NOW(), NOW()),
(14, 2, 'AISLE', 'AISLE', 0, 1, NOW(), NOW()),
(15, 2, '4D', 'SEAT', 'R', 0, NOW(), NOW()),
(16, 2, '4E', 'SEAT', 'R', 0, NOW(), NOW()),
(17, 2, '4F', 'SEAT', 'R', 0, NOW(), NOW()),
(18, 2, 'BLANK', 'BLANK', 0, 1, NOW(), NOW());     -- RIGHT_SIDE BLANK

INSERT INTO bookcabin.seat_characteristics
(id, seat_id, characteristic, created_at, updated_at)
VALUES
(3, 10, 'LEFT_SIDE', NOW(), NOW()), -- BLANK with LEFT_SIDE
(4, 11, 'CH', NOW(), NOW()),
(5, 11, 'W', NOW(), NOW()),
(6, 12, 'CH', NOW(), NOW()),
(7, 12, '9', NOW(), NOW()),
(8, 13, 'A', NOW(), NOW()),
(9, 13, 'CH', NOW(), NOW()),
(10, 15, 'A', NOW(), NOW()),
(11, 15, 'CH', NOW(), NOW()),
(12, 16, 'CH', NOW(), NOW()),
(13, 16, '9', NOW(), NOW()),
(14, 17, 'CH', NOW(), NOW()),
(15, 17, 'W', NOW(), NOW()),
(16, 18, 'RIGHT_SIDE', NOW(), NOW()); -- BLANK with RIGHT_SIDE

INSERT INTO bookcabin.raw_seat_characteristics
(id, seat_id, raw_characteristic, created_at, updated_at)
VALUES
(3, 10, 'LEFT_SIDE', NOW(), NOW()),

(4, 11, 'K', NOW(), NOW()),
(5, 11, 'W', NOW(), NOW()),
(6, 11, 'LS', NOW(), NOW()),
(7, 11, 'L', NOW(), NOW()),
(8, 11, 'CH', NOW(), NOW()),

(9, 12, 'K', NOW(), NOW()),
(10, 12, 'LS', NOW(), NOW()),
(11, 12, 'L', NOW(), NOW()),
(12, 12, 'CH', NOW(), NOW()),
(13, 12, '9', NOW(), NOW()),

(14, 13, 'A', NOW(), NOW()),
(15, 13, 'K', NOW(), NOW()),
(16, 13, 'LS', NOW(), NOW()),
(17, 13, 'L', NOW(), NOW()),
(18, 13, 'CH', NOW(), NOW()),

(19, 15, 'A', NOW(), NOW()),
(20, 15, 'K', NOW(), NOW()),
(21, 15, 'RS', NOW(), NOW()),
(22, 15, 'L', NOW(), NOW()),
(23, 15, 'CH', NOW(), NOW()),

(24, 16, 'K', NOW(), NOW()),
(25, 16, 'RS', NOW(), NOW()),
(26, 16, 'L', NOW(), NOW()),
(27, 16, 'CH', NOW(), NOW()),
(28, 16, '9', NOW(), NOW()),

(29, 17, 'K', NOW(), NOW()),
(30, 17, 'W', NOW(), NOW()),
(31, 17, 'RS', NOW(), NOW()),
(32, 17, 'L', NOW(), NOW()),
(33, 17, 'CH', NOW(), NOW()),

(34, 18, 'RIGHT_SIDE', NOW(), NOW());

-- ROW 5


