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

CREATE TABLE test (
    k VARCHAR(255) PRIMARY KEY,
    v int
);
{
}
INSERT INTO test (k, v) VALUES ('foo', 0);
{
}
SELECT * FROM test WHERE k = 'foo';
{
	"rows" : 
    [
            {
                "k" : "foo",
                "v" : 0
            }
    ]
}
INSERT INTO test (k, v) VALUES ('bar', 1);
{
}
SELECT * FROM test;
{
	"rows" : 
    [
            {
                "k" : "foo",
                "v" : 0
            },
            {
                "k" : "bar",
                "v" : 1
            }
    ]
}
SELECT * FROM test WHERE k = 'foo';
{
	"rows" : 
    [
            {
                "k" : "foo",
                "v" : 0
            }
    ]
}
UPDATE test SET v = 100 WHERE k = 'bar';
{
}
SELECT * FROM test;
{
	"rows" : 
    [
            {
                "k" : "foo",
                "v" : 0
            },
            {
                "k" : "bar",
                "v" : 100
            }
    ]
}
DELETE FROM test WHERE k = 'foo';
{
}
SELECT * FROM test WHERE k = 'foo';
{
}
SELECT * FROM test;
{
	"rows" : 
    [
            {
                "k" : "bar",
                "v" : 100
            }
    ]
}
