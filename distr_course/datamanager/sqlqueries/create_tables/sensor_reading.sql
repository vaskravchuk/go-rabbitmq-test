create table sensor_reading
(
	id serial not null
		constraint sensor_reading_pk
			primary key,
	value double precision not null,
	sensor_id serial not null
		constraint sensor_reading_sensor_id_fk
			references sensor,
	taken_on timestamp not null
);