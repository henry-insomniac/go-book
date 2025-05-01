CREATE TABLE articles (
                          id VARCHAR(32) PRIMARY KEY COMMENT '文章 ID',
                          title VARCHAR(255) NOT NULL,
                          content_md LONGTEXT NOT NULL,
                          summary TEXT,
                          cover_image VARCHAR(512),
                          cover_thumb VARCHAR(512),
                          read_time INT DEFAULT 0,
                          published_at DATETIME NOT NULL,
                          author_id VARCHAR(32) NOT NULL,
                          created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                          updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

                          INDEX idx_title(title),
                          INDEX idx_published_at(published_at),
                          CONSTRAINT fk_articles_author FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
);
