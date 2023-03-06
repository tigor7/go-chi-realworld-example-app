CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    email varchar(255) NOT NULL UNIQUE, 
    username varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    image varchar(255),
    bio TEXT
);

CREATE TABLE IF NOT EXISTS tags (
    id UUID PRIMARY KEY,
    name varchar(255)

);

CREATE TABLE IF NOT EXISTS articles (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    slug TEXT NOT NULL,
    description TEXT NOT NULL, 
    body TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    author_id UUID NOT NULL,
    FOREIGN KEY (author_id) REFERENCES users(id)
);

CREATE FUNCTION article_updated() RETURNS TRIGGER
    AS 
'
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
'
LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON articles
FOR EACH ROW
EXECUTE PROCEDURE article_updated();

CREATE TABLE IF NOT EXISTS articles_tags (
    article_id UUID NOT NULL,
    tag_id UUID NOT NULL,
    FOREIGN KEY (article_id) REFERENCES articles(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id)
);

CREATE TABLE IF NOT EXISTS favorites (
    user_id UUID NOT NULL,
    article_id UUID NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (article_id) REFERENCES articles(id)
);

CREATE TABLE IF NOT EXISTS follows (
    user_id UUID NOT NULL,
    followed_id UUID NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (followed_id) REFERENCES users(id)
);