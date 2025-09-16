CREATE TABLE lessons (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    video_url VARCHAR(255) NOT NULL,
    duration VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Add foreign key constraint
ALTER TABLE lessons ADD CONSTRAINT fk_lessons_course 
    FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE;