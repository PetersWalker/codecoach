-- up
create table commit_stats (
    lines_added int  NOT NULL, 
    lines_subtracted int  NOT NULL,
    name text  NOT NULL,
    file_path text,
    commit_hash text NOT NULL,
    date text  NOT NULL,
    PRIMARY KEY (commit_hash, file_path)
);

-- down
drop table commit_stats ;