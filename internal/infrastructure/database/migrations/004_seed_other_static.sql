-- migrations/004_seed_other_static.sql

-- Seed Marital Statuses
INSERT INTO "marital_statuses" (status) VALUES
('Unmarried'),
('Married'),
('Divorced'),
('Widowed'),
('Annulled')
ON CONFLICT (status) DO NOTHING;

-- Seed Religions
INSERT INTO "religions" (name) VALUES
('Islam'),
('Hinduism'),
('Christianity'),
('Buddhism'),
('Other')
ON CONFLICT (name) DO NOTHING;

-- Seed Professions (A comprehensive starter list)
INSERT INTO "professions" (name) VALUES
('Doctor'),
('Engineer'),
('Teacher'),
('Government Service'),
('Private Service'),
('Business'),
('Lawyer'),
('Banker'),
('Accountant'),
('Software Developer'),
('IT Professional'),
('Homemaker'),
('Student'),
('Unemployed'),
('Self-employed'),
('Farmer'),
('Police'),
('Defense (Army/Navy/Air)'),
('Journalist'),
('Artist'),
('Other')
ON CONFLICT (name) DO NOTHING;

-- Seed Skin Tones
INSERT INTO "skin_tones" (tone) VALUES
('Fair'),
('Very Fair'),
('Wheatish'),
('Olive'),
('Dark')
ON CONFLICT (tone) DO NOTHING;

-- Seed Blood Groups
INSERT INTO "blood_groups" (group_name) VALUES
('A+'),
('A-'),
('B+'),
('B-'),
('AB+'),
('AB-'),
('O+'),
('O-')
ON CONFLICT (group_name) DO NOTHING;

-- Seed Countries (A small subset, including Bangladesh)
INSERT INTO "countries" (name, iso2_code) VALUES
('Bangladesh', 'BD'),
('United States', 'US'),
('United Kingdom', 'GB'),
('Australia', 'AU'),
('Canada', 'CA'),
('India', 'IN'),
('Saudi Arabia', 'SA'),
('United Arab Emirates', 'AE'),
('Malaysia', 'MY')
ON CONFLICT (name) DO NOTHING;