普通索引
索引能够显著提高查询的速度，尤其是在大型表中进行搜索时。通过使用索引，MySQL 可以直接定位到满足查询条件的数据行，而无需逐行扫描整个表。

创建索引
使用 CREATE INDEX 语句可以创建普通索引。

普通索引是最常见的索引类型，用于加速对表中数据的查询。

CREATE INDEX 的语法：

CREATE INDEX index_name
ON table_name (column1 [ASC|DESC], column2 [ASC|DESC], ...);
CREATE INDEX: 用于创建普通索引的关键字。
index_name: 指定要创建的索引的名称。索引名称在表中必须是唯一的。
table_name: 指定要在哪个表上创建索引。
(column1, column2, ...): 指定要索引的表列名。你可以指定一个或多个列作为索引的组合。这些列的数据类型通常是数值、文本或日期。
ASC和DESC（可选）: 用于指定索引的排序顺序。默认情况下，索引以升序（ASC）排序。
以下实例假设我们有一个名为 students 的表，包含 id、name 和 age 列，我们将在 name 列上创建一个普通索引。

CREATE INDEX idx_name ON students (name);
上述语句将在 students 表的 name 列上创建一个名为 idx_name 的普通索引，这将有助于提高通过姓名进行搜索的查询性能。

需要注意的是，如果表中的数据量较大，索引的创建可能会花费一些时间，但一旦创建完成，查询性能将会显著提高。

修改表结构(添加索引)
我们可以使用 ALTER TABLE 命令可以在已有的表中创建索引。

ALTER TABLE 允许你修改表的结构，包括添加、修改或删除索引。

ALTER TABLE 创建索引的语法：

ALTER TABLE table_name
ADD INDEX index_name (column1 [ASC|DESC], column2 [ASC|DESC], ...);
ALTER TABLE: 用于修改表结构的关键字。
table_name: 指定要修改的表的名称。
ADD INDEX: 添加索引的子句。ADD INDEX用于创建普通索引。
index_name: 指定要创建的索引的名称。索引名称在表中必须是唯一的。
(column1, column2, ...): 指定要索引的表列名。你可以指定一个或多个列作为索引的组合。这些列的数据类型通常是数值、文本或日期。
ASC和DESC（可选）: 用于指定索引的排序顺序。默认情况下，索引以升序（ASC）排序。
下面是一个实例，我们将在已存在的名为 employees 的表上创建一个普通索引:

ALTER TABLE employees
ADD INDEX idx_age (age);
上述语句将在 employees 表的 age 列上创建一个名为 idx_age 的普通索引。

创建表的时候直接指定
我们可以在创建表的时候，你可以在 CREATE TABLE 语句中直接指定索引，以创建表和索引的组合。

CREATE TABLE table_name (
  column1 data_type,
  column2 data_type,
  ...,
  INDEX index_name (column1 [ASC|DESC], column2 [ASC|DESC], ...)
);
CREATE TABLE: 用于创建新表的关键字。
table_name: 指定要创建的表的名称。
(column1, column2, ...): 定义表的列名和数据类型。你可以指定一个或多个列作为索引的组合。这些列的数据类型通常是数值、文本或日期。
INDEX: 用于创建普通索引的关键字。
index_name: 指定要创建的索引的名称。索引名称在表中必须是唯一的。
(column1, column2, ...): 指定要索引的表列名。你可以指定一个或多个列作为索引的组合。这些列的数据类型通常是数值、文本或日期。
ASC和DESC（可选）: 用于指定索引的排序顺序。默认情况下，索引以升序（ASC）排序。
下面是一个实例，我们要创建一个名为 students 的表，并在 age 列上创建一个普通索引。

CREATE TABLE students (
  id INT PRIMARY KEY,
  name VARCHAR(50),
  age INT,
  INDEX idx_age (age)
);
在上述实例中，我们在 students 表的 age 列上创建了一个名为 idx_age 的普通索引。

删除索引的语法
我们可以使用 DROP INDEX 语句来删除索引。

DROP INDEX 的语法：

DROP INDEX index_name ON table_name;
DROP INDEX: 用于删除索引的关键字。
index_name: 指定要删除的索引的名称。
ON table_name: 指定要在哪个表上删除索引。
使用 ALTER TABLE 语句删除索引的语法如下：

ALTER TABLE table_name
DROP INDEX index_name;
ALTER TABLE: 用于修改表结构的关键字。
table_name: 指定要修改的表的名称。
DROP INDEX: 用于删除索引的子句。
index_name: 指定要删除的索引的名称。
以下实例假设我们有一个名为 employees 的表，并在 age 列上有一个名为 idx_age 的索引，现在我们要删除这个索引：

DROP INDEX idx_age ON employees;
或使用 ALTER TABLE 语句：

ALTER TABLE employees
DROP INDEX idx_age;
这两个命令都会从 employees 表中删除名为 idx_age 的索引。

如果该索引不存在，执行命令时会产生错误。因此，在删除索引之前最好确认该索引是否存在，或者使用错误处理机制来处理可能的错误情况。

唯一索引
在 MySQL 中，你可以使用 CREATE UNIQUE INDEX 语句来创建唯一索引。

唯一索引确保索引中的值是唯一的，不允许有重复值。

创建索引
CREATE UNIQUE INDEX index_name
ON table_name (column1 [ASC|DESC], column2 [ASC|DESC], ...);
CREATE UNIQUE INDEX: 用于创建唯一索引的关键字组合。
index_name: 指定要创建的唯一索引的名称。索引名称在表中必须是唯一的。
table_name: 指定要在哪个表上创建唯一索引。
(column1, column2, ...): 指定要索引的表列名。你可以指定一个或多个列作为索引的组合。这些列的数据类型通常是数值、文本或日期。
ASC和DESC（可选）: 用于指定索引的排序顺序。默认情况下，索引以升序（ASC）排序。
以下是一个创建唯一索引的实例： 假设我们有一个名为 employees的 表，包含 id 和 email 列，现在我们想在email列上创建一个唯一索引，以确保每个员工的电子邮件地址都是唯一的。

CREATE UNIQUE INDEX idx_email ON employees (email);以上实例将在 employees 表的 email 列上创建一个名为 idx_email 的唯一索引。
修改表结构
我们可以使用 ALTER TABLE 命令来创建唯一索引。

ALTER TABLE命令允许你修改已经存在的表结构，包括添加新的索引。

ALTER table mytable ADD UNIQUE [indexName] (columnName(length))
ALTER TABLE: 用于修改表结构的关键字。
table_name: 指定要修改的表的名称。
ADD CONSTRAINT: 这是用于添加约束（包括唯一索引）的关键字。
index_name: 指定要创建的唯一索引的名称。约束名称在表中必须是唯一的。
UNIQUE (column1, column2, ...): 指定要索引的表列名。你可以指定一个或多个列作为索引的组合。这些列的数据类型通常是数值、文本或日期。
ASC和DESC（可选）: 用于指定索引的排序顺序。默认情况下，索引以升序（ASC）排序。
以下是一个使用 ALTER TABLE 命令创建唯一索引的实例：假设我们有一个名为 employees 的表，包含 id 和 email 列，现在我们想在 email 列上创建一个唯一索引，以确保每个员工的电子邮件地址都是唯一的。

ALTER TABLE employees
ADD CONSTRAINT idx_email UNIQUE (email);
以上实例将在 employees 表的 email 列上创建一个名为 idx_email 的唯一索引。

请注意，如果表中已经有重复的 email 值，那么添加唯一索引将会失败。在创建唯一索引之前，你可能需要确保表中的 email 列没有重复的值。

创建表的时候直接指定
我们也可以在创建表的同时，你可以在 CREATE TABLE 语句中使用 UNIQUE 关键字来创建唯一索引。

这将在表创建时同时定义唯一索引约束。

CREATE TABLE 语句中创建唯一索引的语法：

CREATE TABLE table_name (
  column1 data_type,
  column2 data_type,
  ...,
  CONSTRAINT index_name UNIQUE (column1 [ASC|DESC], column2 [ASC|DESC], ...)
);
CREATE TABLE: 用于创建新表的关键字。
table_name: 指定要创建的表的名称。
(column1, column2, ...): 定义表的列名和数据类型。你可以指定一个或多个列作为索引的组合。这些列的数据类型通常是数值、文本或日期。
CONSTRAINT: 用于添加约束的关键字。
index_name: 指定要创建的唯一索引的名称。约束名称在表中必须是唯一的。
UNIQUE (column1, column2, ...): 指定要索引的表列名。
以下是一个在创建表时创建唯一索引的实例：假设我们要创建一个名为 employees 的表，其中包含 id、name 和 email 列，我们希望 email 列的值是唯一的，因此我们要在创建表时定义唯一索引。

CREATE TABLE employees (
  id INT PRIMARY KEY,
  name VARCHAR(50),
  email VARCHAR(100) UNIQUE
);
在这个例子中，email 列被定义为唯一索引，因为在它的后面加上了 UNIQUE 关键字。

请注意，使用 UNIQUE 关键字后，索引名称将自动生成，你也可以根据需要指定索引名称。

使用ALTER 命令添加和删除索引
有四种方式来添加数据表的索引：

ALTER TABLE tbl_name ADD PRIMARY KEY (column_list): 该语句添加一个主键，这意味着索引值必须是唯一的，且不能为NULL。
ALTER TABLE tbl_name ADD UNIQUE index_name (column_list): 这条语句创建索引的值必须是唯一的（除了NULL外，NULL可能会出现多次）。
ALTER TABLE tbl_name ADD INDEX index_name (column_list): 添加普通索引，索引值可出现多次。
ALTER TABLE tbl_name ADD FULLTEXT index_name (column_list):该语句指定了索引为 FULLTEXT ，用于全文索引。
以下实例为在表中添加索引。

mysql> ALTER TABLE testalter_tbl ADD INDEX (c);
你还可以在 ALTER 命令中使用 DROP 子句来删除索引。尝试以下实例删除索引:

mysql> ALTER TABLE testalter_tbl DROP INDEX c;
使用 ALTER 命令添加和删除主键
主键作用于列上（可以一个列或多个列联合主键），添加主键索引时，你需要确保该主键默认不为空（NOT NULL）。实例如下：

mysql> ALTER TABLE testalter_tbl MODIFY i INT NOT NULL;
mysql> ALTER TABLE testalter_tbl ADD PRIMARY KEY (i);
你也可以使用 ALTER 命令删除主键：

mysql> ALTER TABLE testalter_tbl DROP PRIMARY KEY;
删除主键时只需指定PRIMARY KEY，但在删除索引时，你必须知道索引名。

显示索引信息
你可以使用 SHOW INDEX 命令来列出表中的相关的索引信息。

可以通过添加 \G 来格式化输出信息。

SHOW INDEX 语句：:

mysql> SHOW INDEX FROM table_name\G
........
SHOW INDEX: 用于显示索引信息的关键字。
FROM table_name: 指定要查看索引信息的表的名称。
\G: 格式化输出信息。
执行上述命令后，将会显示指定表中所有索引的详细信息，包括索引名称（Key_name）、索引列（Column_name）、是否是唯一索引（Non_unique）、排序方式（Collation）、索引的基数（Cardinality）等。