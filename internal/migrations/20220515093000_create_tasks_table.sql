DROP TABLE IF EXISTS doneclub.tasks;
CREATE TABLE doneclub.tasks
(
    id          INT UNSIGNED     NOT NULL AUTO_INCREMENT,
    user_id     INT UNSIGNED     NOT NULL,
    goal_id     INT UNSIGNED,
    status      TINYINT UNSIGNED NOT NULL,
    title       TEXT             NOT NULL,
    deadline    DATETIME,
    finished_at DATETIME,
    created_at  DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at  DATETIME,

    PRIMARY KEY (id),
    CONSTRAINT fk_tasks_user_id FOREIGN KEY (user_id)
        REFERENCES doneclub.users (id)
        ON DELETE CASCADE,
    CONSTRAINT fk_tasks_goal_id FOREIGN KEY (goal_id)
        REFERENCES doneclub.goals (id)
        ON DELETE SET NULL
)