CREATE TABLE app_files (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    original_img VARCHAR NOT NULL,
    cropped_img VARCHAR,
    status INT NOT NULL,
    public_id UUID NOT NULL UNIQUE
)