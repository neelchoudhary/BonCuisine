CREATE TABLE Books (
  UserId INT64,
  Name    STRING(40),
  Author  STRING(40)
) PRIMARY KEY(UserId, Name), 
INTERLEAVE IN PARENT Users ON DELETE CASCADE
