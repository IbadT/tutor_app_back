-- Remove foreign key constraint and column
ALTER TABLE courses DROP CONSTRAINT IF EXISTS fk_courses_next_lesson;
ALTER TABLE courses DROP COLUMN IF EXISTS next_lesson_id;
