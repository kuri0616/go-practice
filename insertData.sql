insert into articles (title, contents, username, nice, created_at)
values ('Hello', 'Hello, world!', 'Alice', 0, now());

insert into articles (title, contents, username, nice, created_at)
values ('Goodbye', 'Goodbye, world!', 'Bob', 0, now());

insert into comments (article_id, message, created_at)
values (1, 'Hello, Alice!', now());

insert into comments (article_id, message, created_at)
values (1, 'Hello, Bob!', now());
