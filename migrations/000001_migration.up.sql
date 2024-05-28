BEGIN;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255),
    phone_number VARCHAR(20),
    birthday DATE,
    verified_status BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS promos (
    id SERIAL PRIMARY KEY,
    promo_code VARCHAR(50),
    name VARCHAR(255),
    amount INT,
    start_date DATE,
    end_date DATE,
    UNIQUE(promo_code)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_promos_code ON promos (promo_code);

CREATE TABLE IF NOT EXISTS promo_eligibilities (
    promo_id INT REFERENCES promos(id),
    user_id INT REFERENCES users(id),
    claimed_at TIMESTAMPTZ,
    PRIMARY KEY(promo_id, user_id)
);


CREATE INDEX IF NOT EXISTS idx_promo_eligibilities_promo_id ON promo_eligibilities (promo_id);
CREATE INDEX IF NOT EXISTS idx_promo_eligibilities_user_id ON promo_eligibilities (user_id);

INSERT INTO users (name, email, phone_number, birthday, verified_status) VALUES
('User1', 'user1@example.com', '1234567890', '2023-05-28', true),
('User2', 'user2@example.com', '1234567891', '2023-05-28', true),
('User3', 'user3@example.com', '1234567892', '2023-05-28', false),
('User4', 'user4@example.com', '1234567893', '2023-05-28', false),

('User5', 'user5@example.com', '1234567894', '2023-05-29', true),
('User6', 'user6@example.com', '1234567895', '2023-05-29', true),
('User7', 'user7@example.com', '1234567896', '2023-05-29', false),
('User8', 'user8@example.com', '1234567897', '2023-05-29', false),

('User9', 'user9@example.com', '1234567898', '2023-05-30', true),
('User10', 'user10@example.com', '1234567899', '2023-05-30', true),
('User11', 'user11@example.com', '1234567890', '2023-05-30', false),
('User12', 'user12@example.com', '1234567891', '2023-05-30', false),

('User13', 'user13@example.com', '1234567892', '2023-05-31', true),
('User14', 'user14@example.com', '1234567893', '2023-05-31', true),
('User15', 'user15@example.com', '1234567894', '2023-05-31', false),
('User16', 'user16@example.com', '1234567895', '2023-05-31', false),

('User17', 'user17@example.com', '1234567896', '2023-06-01', true),
('User18', 'user18@example.com', '1234567897', '2023-06-01', true),
('User19', 'user19@example.com', '1234567898', '2023-06-01', false),
('User20', 'user20@example.com', '1234567899', '2023-06-01', false),

('User21', 'user21@example.com', '1234567890', '2023-06-02', true),
('User22', 'user22@example.com', '1234567891', '2023-06-02', true),
('User23', 'user23@example.com', '1234567892', '2023-06-02', false),
('User24', 'user24@example.com', '1234567893', '2023-06-02', false),

('User25', 'user25@example.com', '1234567894', '2023-06-03', true),
('User26', 'user26@example.com', '1234567895', '2023-06-03', true),
('User27', 'user27@example.com', '1234567896', '2023-06-03', false),
('User28', 'user28@example.com', '1234567897', '2023-06-03', false),

('User29', 'user29@example.com', '1234567898', '2023-06-04', true),
('User30', 'user30@example.com', '1234567899', '2023-06-04', true),
('User31', 'user31@example.com', '1234567890', '2023-06-04', false),
('User32', 'user32@example.com', '1234567891', '2023-06-04', false),

('User33', 'user33@example.com', '1234567892', '2023-06-05', true),
('User34', 'user34@example.com', '1234567893', '2023-06-05', true),
('User35', 'user35@example.com', '1234567894', '2023-06-05', false),
('User36', 'user36@example.com', '1234567895', '2023-06-05', false),

('User37', 'user37@example.com', '1234567896', '2023-06-06', true),
('User38', 'user38@example.com', '1234567897', '2023-06-06', true),
('User39', 'user39@example.com', '1234567898', '2023-06-06', false),
('User40', 'user40@example.com', '1234567899', '2023-06-06', false),

('User41', 'user41@example.com', '1234567890', '2023-06-07', true),
('User42', 'user42@example.com', '1234567891', '2023-06-07', true),
('User43', 'user43@example.com', '1234567892', '2023-06-07', false),
('User44', 'user44@example.com', '1234567893', '2023-06-07', false),

('User45', 'user45@example.com', '1234567894', '2023-06-08', true),
('User46', 'user46@example.com', '1234567895', '2023-06-08', true),
('User47', 'user47@example.com', '1234567896', '2023-06-08', false),
('User48', 'user48@example.com', '1234567897', '2023-06-08', false),

('User49', 'user49@example.com', '1234567898', '2023-06-09', true),
('User50', 'user50@example.com', '1234567899', '2023-06-09', true),
('User51', 'user51@example.com', '1234567890', '2023-06-09', false),
('User52', 'user52@example.com', '1234567891', '2023-06-09', false);

COMMIT;