ALTER TABLE visitor_request
ADD COLUMN meeting_room VARCHAR(100) NULL AFTER visit_date,
ADD COLUMN visit_time TIME NULL AFTER meeting_room,
ADD COLUMN duration_minutes INT NULL AFTER visit_time;