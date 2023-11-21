create table commit_stats (
    lines_added int, 
    lines_subtracted int,
    name text,
    file_path text,
    commit_hash text,
    date text  -- should use datetime
    );

drop table commit_stats (
    
)