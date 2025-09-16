-- Add indexes for better performance
CREATE INDEX idx_courses_category_id ON courses(category_id);
CREATE INDEX idx_courses_next_lesson_id ON courses(next_lesson_id);
CREATE INDEX idx_lessons_course_id ON lessons(course_id);
CREATE INDEX idx_categories_name ON categories(name);
