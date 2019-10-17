USE pet;
CREATE TABLE IF NOT EXISTS pets
(
    id              int        NOT NULL AUTO_INCREMENT,
    categoryId      int        DEFAULT 0,
    name            varchar(50) NOT NULL,
    status          varchar(9) DEFAULT 'available',
    createdDateTime TIMESTAMP,
    updatedDateTime TIMESTAMP,
    deletedDateTime TIMESTAMP,
    PRIMARY KEY petId (`id`)
);

CREATE TABLE IF NOT EXISTS pet_photos
(
    petId           int          NOT NULL,
    url             varchar(255) NOT NULL,
    createdDateTime TIMESTAMP,
    deletedDateTime TIMESTAMP,
    PRIMARY KEY petId_url (`petId`, `url`)
);

CREATE TABLE IF NOT EXISTS pet_tags
(
    petId int NOT NULL,
    tagId int NOT NULL
);

CREATE TABLE IF NOT EXISTS categories
(
    id   int NOT NULL AUTO_INCREMENT,
    name varchar(50),
    PRIMARY KEY categoryId (`id`)
);

CREATE TABLE IF NOT EXISTS tags
(
    id   int NOT NULL AUTO_INCREMENT,
    name varchar(50),
    PRIMARY KEY tagId (`id`)
);

INSERT INTO categories (id, name)
VALUES (1, 'dogs'),
       (2, 'cats'),
       (3, 'birds');
INSERT INTO tags (id, name)
VALUES (1, 'friendly'),
       (2, 'independent'),
       (3, 'easy-care');
INSERT INTO pets (id, name, categoryId, status, createdDateTime)
VALUES (1, 'Nacho', 1, 'available', now()),
       (2, 'Gizmo', 1, 'sold', now()),
       (3, 'Queenie', 2, 'available', now());
INSERT INTO pet_tags (petId, tagId)
VALUES (1, 1),
       (1, 3),
       (2, 1),
       (3, 2);
INSERT INTO pet_photos (petId, url, createdDateTime)
VALUES (1, 'https://s3.public/img1-1.png', now()),
       (1, 'https://s3.public/img1-2.png', now()),
       (2, 'https://s3.public/img2-1.png', now()),
       (3, 'https://s3.public/img3-1.png', now()),
       (3, 'https://s3.public/img3-2.png', now()),
       (1, 'https://s3.public/img1-3.png', now());
