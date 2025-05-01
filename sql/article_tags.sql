CREATE TABLE article_tags (
                              article_id VARCHAR(32) NOT NULL,
                              tag_id VARCHAR(32) NOT NULL,

                              PRIMARY KEY (article_id, tag_id),
                              CONSTRAINT fk_article FOREIGN KEY (article_id) REFERENCES articles(id) ON DELETE CASCADE,
                              CONSTRAINT fk_tag FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);
