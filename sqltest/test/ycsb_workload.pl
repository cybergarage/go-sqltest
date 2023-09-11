#!/usr/bin/perl
# Copyright (C) 2022 The go-sqltest Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

use strict;
use warnings;

print<<HEADER;
- Copyright (C) 2020 The go-sqltest Authors. All rights reserved.
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
HEADER
my $key_cnt = 10;
my $field_cnt = 10;
# INSERT queries
for my $key_no (0 .. $key_cnt - 1) {
  my $key = "user${key_no}";
  print "INSERT INTO usertable (YCSB_KEY";
  for my $field_no (0 .. $field_cnt - 1) {
    print ", FIELD${field_no}";
  }
  print ") VALUES ('${key}'";
  for my $field_no (0 .. $field_cnt - 1) {
    my $field = "value00${field_no}";
    print ", '${field}'";
  }
  print ");\n";
  print "{\n";
  print "}\n";
}
# UPDATE queries
