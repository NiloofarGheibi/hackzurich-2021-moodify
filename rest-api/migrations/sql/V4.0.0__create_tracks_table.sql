CREATE TABLE tracks (
                       user_id uuid,
                       track_id character varying(256) NOT NULL,
                       features character varying(256),
                       mood integer,
                       created_at TIMESTAMP without time zone NOT NULL
)
