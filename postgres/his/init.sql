CREATE TABLE IF NOT EXISTS patient (
    id SERIAL PRIMARY KEY,
    first_name_th   VARCHAR(100) NOT NULL,
    middle_name_th  VARCHAR(100),
    last_name_th    VARCHAR(100) NOT NULL,
    first_name_en   VARCHAR(100),
    middle_name_en  VARCHAR(100),
    last_name_en    VARCHAR(100),
    date_of_birth   DATE NOT NULL,
    patient_hn      VARCHAR(9)  UNIQUE NOT NULL,
    national_id     VARCHAR(13) UNIQUE NOT NULL,
    passport_id     VARCHAR(20) UNIQUE,
    phone_number    VARCHAR(10),
    email           VARCHAR(150),
    gender          CHAR(1) NOT NULL CHECK (gender IN ('M', 'F')),
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_patient_hn
ON patient (patient_hn);

CREATE INDEX IF NOT EXISTS idx_patient_national_id
ON patient (national_id);

CREATE INDEX IF NOT EXISTS idx_patient_name_th
ON patient (first_name_th, last_name_th);

CREATE INDEX IF NOT EXISTS idx_patient_name_en
ON patient (first_name_en, last_name_en);

CREATE INDEX IF NOT EXISTS idx_patient_phone
ON patient (phone_number);

CREATE INDEX IF NOT EXISTS idx_patient_email
ON patient (email);

CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE INDEX IF NOT EXISTS idx_patient_fname_th_trgm
ON patient USING GIN (first_name_th gin_trgm_ops);

CREATE INDEX IF NOT EXISTS idx_patient_lname_th_trgm
ON patient USING GIN (last_name_th gin_trgm_ops);

CREATE INDEX IF NOT EXISTS idx_patient_gender_dob
ON patient (gender, date_of_birth);


INSERT INTO patient (
    first_name_th, middle_name_th, last_name_th,
    first_name_en, middle_name_en, last_name_en,
    date_of_birth, patient_hn,
    national_id, passport_id,
    phone_number, email, gender
) VALUES
('สมชาย', NULL, 'ใจดี', 'Somchai', NULL, 'Jaidee', '1985-01-15', '000000001', '1101700000011', NULL, '0810000001', 'somchai1@mail.com', 'M'),
('สมหญิง', NULL, 'สุขใจ', 'Somying', NULL, 'Sukjai', '1990-02-20', '000000002', '1101700000012', NULL, '0810000002', 'somying2@mail.com', 'F'),
('วิชัย', NULL, 'บุญมา', 'Wichai', NULL, 'Boonma', '1978-03-10', '000000003', '1101700000013', NULL, '0810000003', 'wichai3@mail.com', 'M'),
('กมล', NULL, 'ศรีสุข', 'Kamon', NULL, 'Srisuk', '1988-04-05', '000000004', '1101700000014', NULL, '0810000004', 'kamon4@mail.com', 'F'),
('อนันต์', NULL, 'ทองดี', 'Anan', NULL, 'Thongdee', '1975-05-18', '000000005', '1101700000015', NULL, '0810000005', 'anan5@mail.com', 'M'),
('พิมพ์', NULL, 'เพ็ญใจ', 'Pim', NULL, 'Penjai', '1992-06-22', '000000006', '1101700000016', NULL, '0810000006', 'pim6@mail.com', 'F'),
('ณัฐ', NULL, 'วัฒนะ', 'Nat', NULL, 'Wattana', '1983-07-30', '000000007', '1101700000017', NULL, '0810000007', 'nat7@mail.com', 'M'),
('ศิริ', NULL, 'แสงทอง', 'Siri', NULL, 'Saengthong', '1995-08-14', '000000008', '1101700000018', NULL, '0810000008', 'siri8@mail.com', 'F'),
('ธนพล', NULL, 'รุ่งเรือง', 'Thanapon', NULL, 'Rungruang', '1981-09-09', '000000009', '1101700000019', NULL, '0810000009', 'thanapon9@mail.com', 'M'),
('อรทัย', NULL, 'บุญช่วย', 'Orathai', NULL, 'Boonchuay', '1989-10-25', '000000010', '1101700000020', NULL, '0810000010', 'orathai10@mail.com', 'F'),

('จักร', NULL, 'อินทร์ทอง', 'Jak', NULL, 'Inthong', '1977-11-11', '000000011', '1101700000021', NULL, '0810000011', 'jak11@mail.com', 'M'),
('จิราภรณ์', NULL, 'แก้วใส', 'Jiraporn', NULL, 'Kaewsai', '1991-12-02', '000000012', '1101700000022', NULL, '0810000012', 'jira12@mail.com', 'F'),
('สุรชัย', NULL, 'นาคดี', 'Surachai', NULL, 'Nakdee', '1984-01-08', '000000013', '1101700000023', NULL, '0810000013', 'surachai13@mail.com', 'M'),
('มณี', NULL, 'จันทร์งาม', 'Manee', NULL, 'Chan-ngam', '1993-02-17', '000000014', '1101700000024', NULL, '0810000014', 'manee14@mail.com', 'F'),
('ปรีชา', NULL, 'สกุลดี', 'Preecha', NULL, 'Sakundee', '1979-03-27', '000000015', '1101700000025', NULL, '0810000015', 'preecha15@mail.com', 'M'),
('รัตนา', NULL, 'พูลสุข', 'Rattana', NULL, 'Poolsuk', '1987-04-19', '000000016', '1101700000026', NULL, '0810000016', 'rattana16@mail.com', 'F'),
('ชัยวัฒน์', NULL, 'สุขสันต์', 'Chaiwat', NULL, 'Suksan', '1982-05-06', '000000017', '1101700000027', NULL, '0810000017', 'chaiwat17@mail.com', 'M'),
('กัญญา', NULL, 'อารีย์', 'Kanya', NULL, 'Aree', '1994-06-28', '000000018', '1101700000028', NULL, '0810000018', 'kanya18@mail.com', 'F'),
('ศักดิ์ชัย', NULL, 'จิตดี', 'Sakchai', NULL, 'Jitdee', '1976-07-13', '000000019', '1101700000029', NULL, '0810000019', 'sakchai19@mail.com', 'M'),
('เบญจา', NULL, 'ศรีทอง', 'Benja', NULL, 'Sritong', '1996-08-21', '000000020', '1101700000030', NULL, '0810000020', 'benja20@mail.com', 'F');
