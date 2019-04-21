create table sensor
(
    id             serial           not null
        constraint sensor_pk
            primary key,
    name           varchar(50)      not null,
    serial_no      varchar(50)      not null,
    unit_type      varchar(50)      not null,
    max_safe_value double precision not null,
    min_safe_value double precision not null
);