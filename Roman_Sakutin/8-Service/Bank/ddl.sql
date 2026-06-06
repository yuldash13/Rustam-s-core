create table if not exists users
(
    id           bigserial primary key,
    name         varchar(100) default '' not null,
    phone_number varchar(100) default '' not null,
    mail         varchar(100) default '' not null,
    created_at   timestamp    default now()
);

create table if not exists accounts
(
    id         bigserial primary key,
    id_user    bigint,
    balance    bigint not null default 0,
    currency   varchar(100)    default '' not null,
    created_at timestamp       default now()
);

create table if not exists transfers
(
    id              bigserial primary key,
    id_from         bigint,
    id_to           bigint,
    currency        varchar(100) default '' not null,
    value           bigint                  not null default 0,
    operation_state varchar(100) default '' not null,
    created_at      timestamp    default now()
);

--Использовал для обновления id, ну ты как бы и сам понимаешь, знаешь же англ
UPDATE transfers
SET id = 2
WHERE id = 0;

--Использовал чтобы синхронизовать id в таблице
SELECT setval(pg_get_serial_sequence('transfers', 'id'),
              COALESCE((SELECT MAX(id) FROM transfers), 1),
              true);

/*
Users:
id  name     phone_numer   mail                             created_at
1   yuldash  +79522402069  yuldash.nazarov@icloud.com       2025-11-06 14:45:05.056745
2   rustam   +79523821948  hgnmkoujm@icloud.com             2025-11-10 16:08:01.966876
3   nazar    +79313900993  atandratova.shirin80@icloud.com  2025-11-11 14:09:27.494611

Accounts:
id  id_user  balance  currency  created_at
7   3        1000     USD       2025-11-11 14:34:42.902045
8   3        1000     RUB       2025-11-11 14:34:48.156656
9   3        1000     EU        2025-11-11 14:34:52.690874
5   2        1000     RUB       2025-11-10 16:08:21.629962
1   1        1000     RUB       2025-11-10 15:08:24.442769
6   2        1000     USD       2025-11-10 16:08:27.408697
2   1        1000     USD       2025-11-10 16:05:06.642631
4   2        1500     EU        2025-11-10 16:08:13.571565
3   1        1000     EU        2025-11-10 16:05:20.100286

Transfers:
id  id_from  id_to  currency  value  operation_status  created_at
1   1        5      RUB       500    cancel            2025-11-13 16:03:26.671568
4   5        1      RUB       500    success           2025-11-13 16:11:21.167328
2   2        6      USD       500    cancel            2025-11-13 16:04:53.820657
5   6        2      USD       500    success           2025-11-13 16:12:36.444876
3   3        4      EU        500    cancel            2025-11-13 16:09:48.521205
6   4        3      EU        500    success           2025-11-13 16:16:36.583848
*/
