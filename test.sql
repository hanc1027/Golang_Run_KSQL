-- 顯示所有的Queries
SHOW QUERIES;

-- 顯示所有的Tables
SHOW TALBES;

-- 顯示所有的Topics
SHOW TOPICS;

-- 創建test_table
CREATE TABLE test_table (phoneID VARCHAR) WITH (KAFKA_TOPIC='test_topic', VALUE_FORMAT='JSON', KEY='phoneID');

-- 刪除test_table
DROP table test_table;
