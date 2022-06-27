DROP TABLE IF EXISTS doneclub.category_task;
CREATE TABLE doneclub.category_task
(
    id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    category_id INT UNSIGNED NOT NULL,
    task_id     INT UNSIGNED NOT NULL,
    created_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),

    CONSTRAINT fk_category_task_category_id FOREIGN KEY (category_id)
        REFERENCES doneclub.categories (id)
        ON DELETE CASCADE,

    CONSTRAINT fk_category_task_task_id FOREIGN KEY (task_id)
        REFERENCES doneclub.tasks (id)
        ON DELETE CASCADE
)