-- migrations/003_seed_education.sql

-- 1. Seed Streams
INSERT INTO "education_streams" (stream_name_en) VALUES
('General'),
('Madrasah'),
('Technical-Vocational')
ON CONFLICT (stream_name_en) DO NOTHING;

-- 2. Seed Levels
INSERT INTO "education_levels" (stream_id, level_name_en) VALUES
((SELECT id FROM education_streams WHERE stream_name_en = 'General'), 'Primary'),
((SELECT id FROM education_streams WHERE stream_name_en = 'General'), 'Junior Secondary'),
((SELECT id FROM education_streams WHERE stream_name_en = 'General'), 'Secondary'),
((SELECT id FROM education_streams WHERE stream_name_en = 'General'), 'Higher Secondary'),
((SELECT id FROM education_streams WHERE stream_name_en = 'General'), 'Tertiary (Undergraduate)'),
((SELECT id FROM education_streams WHERE stream_name_en = 'General'), 'Tertiary (Graduate)'),
((SELECT id FROM education_streams WHERE stream_name_en = 'General'), 'Tertiary (Doctoral)'),
((SELECT id FROM education_streams WHERE stream_name_en = 'Madrasah'), 'Primary'),
((SELECT id FROM education_streams WHERE stream_name_en = 'Madrasah'), 'Junior Secondary'),
((SELECT id FROM education_streams WHERE stream_name_en = 'Madrasah'), 'Secondary'),
((SELECT id FROM education_streams WHERE stream_name_en = 'Madrasah'), 'Higher Secondary'),
((SELECT id FROM education_streams WHERE stream_name_en = 'Madrasah'), 'Tertiary (Undergraduate)'),
((SELECT id FROM education_streams WHERE stream_name_en = 'Madrasah'), 'Tertiary (Graduate)'),
((SELECT id FROM education_streams WHERE stream_name_en = 'Technical-Vocational'), 'Secondary'),
((SELECT id FROM education_streams WHERE stream_name_en = 'Technical-Vocational'), 'Higher Secondary'),
((SELECT id FROM education_streams WHERE stream_name_en = 'Technical-Vocational'), 'Tertiary (Diploma)'),
((SELECT id FROM education_streams WHERE stream_name_en = 'Technical-Vocational'), 'Tertiary (Undergraduate)'),
((SELECT id FROM education_streams WHERE stream_name_en = 'Technical-Vocational'), 'Tertiary (Graduate)'),
((SELECT id FROM education_streams WHERE stream_name_en = 'Technical-Vocational'), 'Tertiary (Doctoral)')
ON CONFLICT (stream_id, level_name_en) DO NOTHING;

-- 3. Seed Qualifications
INSERT INTO "qualifications" (level_id, qualification_name_en, common_name) VALUES
-- General Stream
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'General') AND level_name_en = 'Primary'), 'Primary Education Certificate', 'PEC'),
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'General') AND level_name_en = 'Junior Secondary'), 'Junior School Certificate', 'JSC'),
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'General') AND level_name_en = 'Secondary'), 'Secondary School Certificate', 'SSC'),
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'General') AND level_name_en = 'Higher Secondary'), 'Higher Secondary Certificate', 'HSC'),
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'General') AND level_name_en = 'Tertiary (Undergraduate)'), 'Bachelor''s (Honours)', 'B.A./B.Sc./B.Com. (Hons)'),
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'General') AND level_name_en = 'Tertiary (Undergraduate)'), 'Bachelor''s (Pass)', 'B.A./B.Sc./B.Com. (Pass)'),
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'General') AND level_name_en = 'Tertiary (Undergraduate)'), 'Bachelor of Medicine, Bachelor of Surgery', 'MBBS'),
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'General') AND level_name_en = 'Tertiary (Graduate)'), 'Master''s Degree', 'M.A./M.Sc./MBA'),
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'General') AND level_name_en = 'Tertiary (Doctoral)'), 'Master of Philosophy', 'M.Phil'),
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'General') AND level_name_en = 'Tertiary (Doctoral)'), 'Doctor of Philosophy', 'PhD'),
-- Madrasah Stream
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'Madrasah') AND level_name_en = 'Primary'), 'Ebtedayee', 'Ebtedayee'),
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'Madrasah') AND level_name_en = 'Junior Secondary'), 'Junior Dakhil Certificate', 'JDC'),
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'Madrasah') AND level_name_en = 'Secondary'), 'Dakhil', 'Dakhil'),
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'Madrasah') AND level_name_en = 'Higher Secondary'), 'Alim', 'Alim'),
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'Madrasah') AND level_name_en = 'Tertiary (Undergraduate)'), 'Fazil', 'Fazil'),
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'Madrasah') AND level_name_en = 'Tertiary (Graduate)'), 'Kamil', 'Kamil'),
-- Technical Stream
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'Technical-Vocational') AND level_name_en = 'Secondary'), 'SSC (Vocational)', 'SSC-Voc'),
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'Technical-Vocational') AND level_name_en = 'Higher Secondary'), 'HSC (Vocational)', 'HSC-Voc'),
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'Technical-Vocational') AND level_name_en = 'Higher Secondary'), 'HSC (Business Management Technology)', 'HSC (BMT)'),
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'Technical-Vocational') AND level_name_en = 'Higher Secondary'), 'Diploma in Commerce', 'Dip. (Comm)'),
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'Technical-Vocational') AND level_name_en = 'Tertiary (Diploma)'), 'Diploma in Engineering', 'Polytechnic Diploma'),
((SELECT id FROM education_levels WHERE stream_id = (SELECT id FROM education_streams WHERE stream_name_en = 'Technical-Vocational') AND level_name_en = 'Tertiary (Undergraduate)'), 'Bachelor of Science in Engineering', 'B.Sc. Engg.')
ON CONFLICT (qualification_name_en) DO NOTHING;

-- 4. Set Equivalencies (This must be run *after* inserting the qualifications)
UPDATE qualifications SET equivalent_to = (SELECT id FROM qualifications WHERE common_name = 'PEC') WHERE common_name = 'Ebtedayee';
UPDATE qualifications SET equivalent_to = (SELECT id FROM qualifications WHERE common_name = 'JSC') WHERE common_name = 'JDC';
UPDATE qualifications SET equivalent_to = (SELECT id FROM qualifications WHERE common_name = 'SSC') WHERE common_name = 'Dakhil' OR common_name = 'SSC-Voc';
UPDATE qualifications SET equivalent_to = (SELECT id FROM qualifications WHERE common_name = 'HSC') WHERE common_name = 'Alim' OR common_name = 'HSC-Voc' OR common_name = 'HSC (BMT)' OR common_name = 'Dip. (Comm)';
UPDATE qualifications SET equivalent_to = (SELECT id FROM qualifications WHERE qualification_name_en = 'Bachelor''s (Pass)') WHERE common_name = 'Fazil';
UPDATE qualifications SET equivalent_to = (SELECT id FROM qualifications WHERE qualification_name_en = 'Master''s Degree') WHERE common_name = 'Kamil';