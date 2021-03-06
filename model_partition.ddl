CREATE TABLE CUSTOMER ( C_CUSTKEY     INTEGER PRIMARY KEY,
                             C_NAME        VARCHAR(25) NOT NULL,
                             C_ADDRESS     VARCHAR(40) NOT NULL,
                             C_NATIONKEY   INTEGER NOT NULL,
                             C_PHONE       CHAR(15) NOT NULL,
                             C_ACCTBAL     DECIMAL(15,2)   NOT NULL,
                             C_MKTSEGMENT  CHAR(10) NOT NULL,
                             C_COMMENT     VARCHAR(117) NOT NULL) PARTITION BY HASH (C_CUSTKEY);

CREATE TABLE CUSTOMER_P0
    PARTITION OF CUSTOMER FOR VALUES WITH (MODULUS 10, REMAINDER 0);
CREATE TABLE CUSTOMER_P1
    PARTITION OF CUSTOMER FOR VALUES WITH (MODULUS 10, REMAINDER 1);
CREATE TABLE CUSTOMER_P2
    PARTITION OF CUSTOMER FOR VALUES WITH (MODULUS 10, REMAINDER 2);
CREATE TABLE CUSTOMER_P3
    PARTITION OF CUSTOMER FOR VALUES WITH (MODULUS 10, REMAINDER 3);
CREATE TABLE CUSTOMER_P4
    PARTITION OF CUSTOMER FOR VALUES WITH (MODULUS 10, REMAINDER 4);
CREATE TABLE CUSTOMER_P5
    PARTITION OF CUSTOMER FOR VALUES WITH (MODULUS 10, REMAINDER 5);
CREATE TABLE CUSTOMER_P6
    PARTITION OF CUSTOMER FOR VALUES WITH (MODULUS 10, REMAINDER 6);
CREATE TABLE CUSTOMER_P7
    PARTITION OF CUSTOMER FOR VALUES WITH (MODULUS 10, REMAINDER 7);
CREATE TABLE CUSTOMER_P8
    PARTITION OF CUSTOMER FOR VALUES WITH (MODULUS 10, REMAINDER 8);
CREATE TABLE CUSTOMER_P9
    PARTITION OF CUSTOMER FOR VALUES WITH (MODULUS 10, REMAINDER 9);
