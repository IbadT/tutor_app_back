CREATE TABLE courses (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    student_id UUID REFERENCES users(id) ON DELETE CASCADE,
    tutor_id UUID REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    progress INT NOT NULL DEFAULT 0,
    total_lessons INT NOT NULL DEFAULT 0,
    completed_lessons INT NOT NULL DEFAULT 0,
    students_count INT NOT NULL DEFAULT 0,
    rating FLOAT NOT NULL DEFAULT 0,
    category_id UUID,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Add foreign key constraints after all tables are created
ALTER TABLE courses ADD CONSTRAINT fk_courses_category 
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL;
ALTER TABLE courses ADD CONSTRAINT fk_courses_student 
    FOREIGN KEY (student_id) REFERENCES users(id) ON DELETE CASCADE;
ALTER TABLE courses ADD CONSTRAINT fk_courses_tutor 
    FOREIGN KEY (tutor_id) REFERENCES users(id) ON DELETE CASCADE;