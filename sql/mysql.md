# 查询
 SELECT a.runoob_id, a.runoob_author, b.runoob_count FROM runoob_tbl a, tcount_tbl b WHERE a.runoob_author = b.runoob_author;

 SELECT a.runoob_id, a.runoob_author, b.runoob_count FROM runoob_tbl a INNER JOIN tcount_tbl b ON a.runoob_author = b.runoob_author;
 
 SELECT a.runoob_id, a.runoob_author, b.runoob_count FROM runoob_tbl a LEFT JOIN tcount_tbl b ON a.runoob_author = b.runoob_author;
 SELECT a.runoob_id, a.runoob_author, b.runoob_count FROM runoob_tbl a RIGHT JOIN tcount_tbl b ON a.runoob_author = b.runoob_author;

SELECT column_name, function(column_name)
FROM table_name
WHERE column_name operator value
GROUP BY column_name;

select university,
    avg(question_cnt) as avg_question_cnt
from user_profile
group by university
order by avg_question_cnt;

SELECT device_id,gender,age,university,gpa FROM user_profile
WHERE (university='山东大学'AND gpa > 3.5) OR (university='复旦大学' AND gpa > 3.8)

SELECT device_id,gender,age,gpa FROM user_profile
WHERE university='山东大学' 
UNION ALL
SELECT device_id,gender,age,gpa FROM user_profile
WHERE gender="male"

select qpd.device_id, qpd.question_id, qpd.result
from question_practice_detail as qpd
inner join user_profile as up
on up.device_id=qpd.device_id and up.university='浙江大学'
order by question_id

select device_id, question_id, result
from question_practice_detail
where device_id in (
    select device_id from user_profile
    where university='浙江大学'
)
order by question_id

SELECT
u.university,
qd.difficult_level,
count(q.question_id)/count(distinct(q.device_id)) AS avg_answer_cnt
FROM question_practice_detail AS q
LEFT JOIN user_profile AS u
ON u.device_id=q.device_id
LEFT JOIN question_detail AS qd
ON q.question_id=qd.question_id
GROUP BY u.university, qd.difficult_level;

# 索引
CREATE INDEX 的语法：

CREATE INDEX index_name
ON table_name (column1 [ASC|DESC], column2 [ASC|DESC], ...);

CREATE INDEX idx_name ON students (name);

ALTER TABLE table_name
ADD INDEX index_name (column1 [ASC|DESC], column2 [ASC|DESC], ...);

ALTER TABLE employees
ADD INDEX idx_age (age);

CREATE TABLE table_name (
  column1 data_type,
  column2 data_type,
  ...,
  INDEX index_name (column1 [ASC|DESC], column2 [ASC|DESC], ...)
);

CREATE TABLE students (
  id INT PRIMARY KEY,
  name VARCHAR(50),
  age INT,
  INDEX idx_age (age)
);

唯一索引确保索引中的值是唯一的，不允许有重复值。
创建索引
CREATE UNIQUE INDEX index_name
ON table_name (column1 [ASC|DESC], column2 [ASC|DESC], ...);
CREATE UNIQUE INDEX idx_email ON employees (email);

ALTER table mytable ADD UNIQUE [indexName] (columnName(length))
ALTER TABLE employees
ADD CONSTRAINT idx_email UNIQUE (email);

CREATE TABLE table_name (
  column1 data_type,
  column2 data_type,
  ...,
  CONSTRAINT index_name UNIQUE (column1 [ASC|DESC], column2 [ASC|DESC], ...)
);
CREATE TABLE employees (
  id INT PRIMARY KEY,
  name VARCHAR(50),
  email VARCHAR(100) UNIQUE
);

有四种方式来添加数据表的索引：
ALTER TABLE tbl_name ADD PRIMARY KEY (column_list): 该语句添加一个主键，这意味着索引值必须是唯一的，且不能为NULL。
ALTER TABLE tbl_name ADD UNIQUE index_name (column_list): 这条语句创建索引的值必须是唯一的（除了NULL外，NULL可能会出现多次）。
ALTER TABLE tbl_name ADD INDEX index_name (column_list): 添加普通索引，索引值可出现多次。
ALTER TABLE tbl_name ADD FULLTEXT index_name (column_list):该语句指定了索引为 FULLTEXT ，用于全文索引。