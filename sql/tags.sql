CREATE TABLE tags (
                      id VARCHAR(32) PRIMARY KEY COMMENT '标签 ID',
                      name VARCHAR(100) NOT NULL UNIQUE,
                      created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                      updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
