-- Add next_lesson_id column to courses table
ALTER TABLE courses ADD COLUMN next_lesson_id UUID;

-- Add foreign key constraint for next_lesson_id
ALTER TABLE courses ADD CONSTRAINT fk_courses_next_lesson 
    FOREIGN KEY (next_lesson_id) REFERENCES lessons(id) ON DELETE SET NULL;
