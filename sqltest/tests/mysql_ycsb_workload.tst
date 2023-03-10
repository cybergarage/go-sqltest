- Copyright (C) 2020 Satoshi Konno. All rights reserved.
-
- Licensed under the Apache License, Version 2.0 (the "License");
- you may not use this file except in compliance with the License.
- You may obtain a copy of the License at
-
-  http:-www.apache.org/licenses/LICENSE-2.0
-
- Unless required by applicable law or agreed to in writing, software
- distributed under the License is distributed on an "AS IS" BASIS,
- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
- See the License for the specific language governing permissions and
- limitations under the License.

CREATE TABLE usertable (
	YCSB_KEY VARCHAR(255) PRIMARY KEY,
	FIELD0 TEXT, FIELD1 TEXT,
	FIELD2 TEXT, FIELD3 TEXT,
	FIELD4 TEXT, FIELD5 TEXT,
	FIELD6 TEXT, FIELD7 TEXT,
	FIELD8 TEXT, FIELD9 TEXT)
{
}
INSERT INTO usertable VALUES (
    YCSB_KEY = 'key001', 
    FIELD0 = 'field000',
    FIELD1 = 'field001',
    FIELD2 = 'field002',
    FIELD3 = 'field003',
    FIELD4 = 'field004',
    FIELD5 = 'field005',
    FIELD6 = 'field006',
    FIELD7 = 'field007',
    FIELD8 = 'field008',
    FIELD9 = 'field009')
{
}
SELECT * FROM usertable WHERE YCSB_KEY = 'key001';
{
	"rows" : 
    [
            {
                "YCSB_KEY" : "key001",
                "FIELD0" : "field000",
                "FIELD1" : "field001",
                "FIELD2" : "field002",
                "FIELD3" : "field003",
                "FIELD4" : "field004",
                "FIELD5" : "field005",
                "FIELD6" : "field006",
                "FIELD7" : "field007",
                "FIELD8" : "field008",
                "FIELD9" : "field009"
            }
    ]
}
UPDATE usertable SET FIELD0 = 'field100' WHERE YCSB_KEY = 'key001'
{
}
SELECT * FROM usertable WHERE YCSB_KEY = 'key001';
{
	"rows" : 
    [
            {
                "YCSB_KEY" : "key001",
                "FIELD0" : "field100",
                "FIELD1" : "field001",
                "FIELD2" : "field002",
                "FIELD3" : "field003",
                "FIELD4" : "field004",
                "FIELD5" : "field005",
                "FIELD6" : "field006",
                "FIELD7" : "field007",
                "FIELD8" : "field008",
                "FIELD9" : "field009"
            }
    ]
}
