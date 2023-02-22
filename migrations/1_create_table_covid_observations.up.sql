CREATE TABLE covid_observations(
        id VARCHAR(256) NOT NULL,
        observation_date timestamp NOT NULL, 
        state VARCHAR(256) NOT NULL,
        country VARCHAR(256) NOT NULL,
        confirmed NUMERIC NOT NULL,
        deaths NUMERIC NOT NULL,
        recovered NUMERIC NOT NULL,
        created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp NULL DEFAULT NULL,
        deleted_at timestamp NULL DEFAULT NULL
);

CREATE INDEX covid_observations_id_idx ON covid_observations(id);
