-- name: login.current
select id from users where (username = $1 or email = $1) and password = $2;

-- name: login.legacy
select id from account where (username = $1 or email = $1) and password = $2;

-- name: user.by_id
select id, username, email, created_at from users where id = $1;
